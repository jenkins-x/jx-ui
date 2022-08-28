package server

import (
	"net/http"
	"os"
	"path/filepath"

	internal "jx-ui/internal/kube"

	"github.com/gorilla/mux"
	"github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned"
	jenkinsxv1 "github.com/jenkins-x/jx-api/v4/pkg/client/clientset/versioned/typed/jenkins.io/v1"
	"github.com/jenkins-x/jx-helpers/v3/pkg/kube/jxclient"
	"github.com/rs/cors"

	tknclient "github.com/tektoncd/pipeline/pkg/client/clientset/versioned"
	"github.com/unrolled/render"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// Server struct holds the external dependencies requires for running the UI backend
type Server struct {
	addr       string
	server     *http.Server
	jxIface    versioned.Interface
	jxClient   jenkinsxv1.PipelineActivityInterface
	srClient   jenkinsxv1.SourceRepositoryInterface
	tknClient  tknclient.Interface
	kubeClient kubernetes.Interface
	render     *render.Render
	namespace  string
}

type spaHandler struct {
	staticPath string
	indexPath  string
}

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

func (s *Server) Run() error {
	s.render = render.New(render.Options{
		DisableHTTPErrorRendering: true,
	})
	config, err := internal.GetKubeConfig()
	if err != nil {
		return err
	}

	if s.namespace == "" {
		s.namespace = "jx"
	}

	s.addr = os.Getenv("UI_ADDR")
	if s.addr == "" {
		s.addr = "localhost:9200"
	}

	jxClient, err := versioned.NewForConfig(config)
	if err != nil {
		return err
	}
	s.jxClient = jxClient.JenkinsV1().PipelineActivities(s.namespace)

	s.srClient = jxClient.JenkinsV1().SourceRepositories(s.namespace)

	tknClient, err := tknclient.NewForConfig(config)
	if err != nil {
		return err
	}
	s.tknClient = tknClient

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}
	s.kubeClient = kubeClient

	s.jxIface, err = jxclient.LazyCreateJXClient(s.jxIface)
	if err != nil {
		return err
	}

	router := mux.NewRouter()
	router = registerRoutes(router, s)
	// This is only required for dev mode
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "HEAD", "PUT"},
		AllowCredentials: true,
	})
	s.server = &http.Server{
		Handler: c.Handler(router),
		Addr:    s.addr,
	}

	err = s.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
