package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	internal "jx-ui/internal/kube"

	"github.com/gorilla/mux"
	"github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned"
	jenkinsxv1 "github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned/typed/jenkins.io/v1"
	"github.com/jenkins-x/jx-helpers/v3/pkg/kube/naming"
	"github.com/rs/cors"

	// "github.com/jenkins-x-plugins/jx-pipeline/pkg/cloud"
	pipelineapi "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	tknclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/unrolled/render"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}
type Server struct {
	// addr       string
	server     *http.Server
	jxClient   jenkinsxv1.PipelineActivityInterface
	srClient   jenkinsxv1.SourceRepositoryInterface
	tknClient  tknclient.Interface
	kubeClient kubernetes.Interface
	render     *render.Render
}

// ToDo: make it configurable
const JXNS = "jx"

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

// PipelinesHandler function
func (s *Server) PipelinesHandler(w http.ResponseWriter, r *http.Request) {
	pa, err := s.jxClient.
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// Todo: improve error handling!
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	s.render.JSON(w, http.StatusOK, pa.Items) //nolint:errcheck
}

// PipelineHandler function
func (s *Server) PipelineHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]
	branch := vars["branch"]
	build := vars["build"]
	name := naming.ToValidName(owner + "-" + repo + "-" + branch + "-" + build)
	method := r.Method
	if method == "GET" {
		pa, err := s.jxClient.
			Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			// Todo: improve error handling!
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		fmt.Println(pa.Spec)
		fmt.Println(pa.Spec.BuildLogsURL)

		s.render.JSON(w, http.StatusOK, pa) //nolint:errcheck
	} else {
		pa, err := s.jxClient.
			Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			// Todo: improve error handling!
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		prName := pa.Labels["tekton.dev/pipeline"]

		pr, err := s.tknClient.TektonV1beta1().PipelineRuns("jx").Get(context.Background(), prName, metav1.GetOptions{})
		if err != nil {
			panic(err)
		}

		if pr.Status.CompletionTime == nil {
			pr.Spec.Status = pipelineapi.PipelineRunSpecStatusCancelled
		}
		_, err = s.tknClient.TektonV1beta1().PipelineRuns("jx").Update(context.Background(), pr, metav1.UpdateOptions{})
		if err != nil {
			panic(err)
		}
		s.render.JSON(w, http.StatusOK, "pipeline "+name+" stopped") //nolint:errcheck
	}
}

// PipelineHandler function
// func (s *Server) PipelineDeleteHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	name := vars["name"]
// 	err := s.jxClient.
// 		JenkinsV1().
// 		PipelineActivities(JXNS).
// 		Delete(context.Background(), name, metav1.DeleteOptions{})
// 	if err != nil {
// 		// Todo: improve error handling!
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	s.render.JSON(w, http.StatusOK, "pipeline "+ name + " stopped") //nolint:errcheck
// }

// StageLogHandler function
func (s *Server) StageLogHandler(w http.ResponseWriter, r *http.Request) {
	// Todo: This is incredibly inefficient, but works, best to switch to SSE/Websockets
	vars := mux.Vars(r)
	name := vars["name"]

	// ToDo
	// Seems like the only way to get all the pods/tasks for a pipeline run
	// Not all pods are labelled by jx ... need to look!
	podList, err := s.kubeClient.
		CoreV1().
		Pods(JXNS).
		List(context.Background(), metav1.ListOptions{LabelSelector: "tekton.dev/pipelineRun=" + name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if len(podList.Items) < 1 {
		s.render.JSON(w, http.StatusNotFound, http.StatusText(http.StatusNotFound)) //nolint:errcheck
	}

	logSteps := make(map[string][]string)
	var logStep string
	for k := range podList.Items {
		podName := podList.Items[k].GetName()
		taskLabel := podList.Items[k].Labels["tekton.dev/pipelineTask"]
		list := getPodContainers(podName, s.kubeClient)
		for k1 := range list {
			logStep = getStepLogs(podName, list[k1], s.kubeClient)
			logSteps[taskLabel] = append(logSteps[taskLabel], logStep)
		}
	}

	s.render.JSON(w, http.StatusOK, logSteps) //nolint:errcheck
}

func getPodContainers(podName string, client kubernetes.Interface) []string {
	pod, err := client.
		CoreV1().
		Pods(JXNS).
		Get(context.Background(), podName, metav1.GetOptions{})
	if err != nil {
		panic(err)
	}

	containerList := make([]string, 0)

	for k := range pod.Spec.Containers {
		containerList = append(containerList, pod.Spec.Containers[k].Name)
	}

	return containerList
}

func getStepLogs(podName, containerName string, client kubernetes.Interface) string {
	req := client.
		CoreV1().
		Pods(JXNS).
		GetLogs(podName, &v1.PodLogOptions{Container: containerName})
	podLogs, err := req.Stream(context.Background())
	if err != nil {
		panic(err)
	}
	defer podLogs.Close()
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		panic(err)
	}
	str := buf.String()
	return str
}

// RepositoriesHandler function
func (s *Server) RepositoriesHandler(w http.ResponseWriter, r *http.Request) {
	repo, err := s.srClient.
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// Todo: improve error handling!
		panic(err)
	}

	s.render.JSON(w, http.StatusOK, repo.Items) //nolint:errcheck
}

func registerRoutes(router *mux.Router, server *Server) *mux.Router {
	router.HandleFunc("/api/v1/pipelines", server.PipelinesHandler)
	router.HandleFunc("/api/v1/pipelines/{owner}/{repo}/{branch}/{build}", server.PipelineHandler).Methods("GET", "POST")
	router.HandleFunc("/api/v1/stages/{name}/logs", server.StageLogHandler)
	router.HandleFunc("/api/v1/repositories", server.RepositoriesHandler)
	spa := spaHandler{staticPath: "web/build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)
	return router
}

func main() {
	s := &Server{}
	s.render = render.New(render.Options{
		DisableHTTPErrorRendering: true,
	})
	config, err := internal.GetKubeConfig()
	if err != nil {
		log.Fatal("Cannot get kubeconfig")
	}

	jxClient, err := versioned.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	s.jxClient = jxClient.JenkinsV1().PipelineActivities(JXNS)

	s.srClient = jxClient.JenkinsV1().SourceRepositories(JXNS)

	tknClient, err := tknclient.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	s.tknClient = tknClient

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	s.kubeClient = kubeClient

	router := mux.NewRouter()
	router = registerRoutes(router, s)
	// This is only required for dev mode
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	})
	s.server = &http.Server{
		Handler:      c.Handler(router),
		Addr:         "127.0.0.1:8080", // Todo: Make it configurable
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(s.server.ListenAndServe())
}
