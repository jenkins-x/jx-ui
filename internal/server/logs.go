package server

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jenkins-x-plugins/jx-pipeline/pkg/tektonlog"
	"github.com/jenkins-x/jx-helpers/v3/pkg/kube/naming"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelineStreamingLogHandler func streams the logs for a running pipeline
// It is a pure golang implementation of Server Sent Events (SSE)
func (s *Server) PipelineStreamingLogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	flusher, ok := w.(http.Flusher)
	if !ok {
		fmt.Println("response writer does not implement http flusher")
	}

	ctx := context.Background()
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]
	branch := vars["branch"]
	build := vars["build"]

	paName := fmt.Sprintf("%s-%s-%s-%s",
		naming.ToValidName(owner),
		naming.ToValidName(repo),
		naming.ToValidName(branch),
		build)

	baseName := fmt.Sprintf("%s/%s/%s #%s",
		naming.ToValidName(owner),
		naming.ToValidName(repo),
		naming.ToValidName(branch),
		strings.ToLower(build))

	logger := tektonlog.TektonLogger{
		JXClient:     s.JxIface,
		TektonClient: s.TknClient,
		KubeClient:   s.KubeClient,
		Namespace:    s.Namespace,
	}

	pa, err := s.JxClient.
		Get(context.Background(), paName, metav1.GetOptions{})
	if err != nil {
		// Todo: improve error handling!
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	triggerContext := pa.Spec.Context
	name := fmt.Sprintf("%s %s", baseName, naming.ToValidName(triggerContext))

	filter := tektonlog.BuildPodInfoFilter{
		Owner:      owner,
		Repository: repo,
		Branch:     branch,
		Build:      build,
	}

	_, _, prMap, err := logger.GetTektonPipelinesWithActivePipelineActivity(ctx, &filter)
	if err != nil {
		fmt.Println(err)
	}

	prList := prMap[name]

	for k := range logger.GetRunningBuildLogs(ctx, pa, prList, name) {
		select {
		case <-r.Context().Done():
			return
		default:
			fmt.Fprintf(w, "data: %s\n\n", k.Line)
			flusher.Flush()
		}
	}
}

// PipelineLogHandler returns the logs for a given pipeline
func (s *Server) PipelineArchivedLogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	owner := vars["owner"]
	repo := vars["repo"]
	branch := vars["branch"]
	build := vars["build"]
	name := naming.ToValidName(owner + "-" + repo + "-" + branch + "-" + build)

	logger := tektonlog.TektonLogger{
		JXClient:     s.JxIface,
		TektonClient: s.TknClient,
		KubeClient:   s.KubeClient,
		Namespace:    s.Namespace,
	}

	pa, err := s.JxClient.
		Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		// Todo: improve error handling!
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	logs := []string{}

	// Read for archived logs if builds are not running
	for line := range logger.StreamPipelinePersistentLogs(pa.Spec.BuildLogsURL) {
		logs = append(logs, line.Line)
	}

	s.render.JSON(w, http.StatusOK, logs) //nolint:errcheck
}
