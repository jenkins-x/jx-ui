package server

import (
	"context"
	"net/http"
	"sort"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jenkins-x/jx-helpers/v3/pkg/kube/naming"
	pipelineapi "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type pipelinesRes struct {
	GitOwner      string
	GitRepository string
	GitBranch     string
	Build         string
	Status        string
	StartTime     *metav1.Time
	EndTime       *metav1.Time
}

// PipelinesHandler function
func (s *Server) PipelinesHandler(w http.ResponseWriter, r *http.Request) {
	pa, err := s.jxClient.
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// Todo: improve error handling!
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	Pipelines := make([]pipelinesRes, len(pa.Items))

	for k := range pa.Items {
		Pipelines[k].GitOwner = pa.Items[k].Spec.GitOwner
		Pipelines[k].GitRepository = pa.Items[k].Spec.GitRepository
		Pipelines[k].GitBranch = pa.Items[k].Spec.GitBranch
		Pipelines[k].Build = pa.Items[k].Spec.Build
		Pipelines[k].Status = pa.Items[k].Spec.Status.String()
		Pipelines[k].StartTime = pa.Items[k].Spec.StartedTimestamp
		Pipelines[k].EndTime = pa.Items[k].Spec.CompletedTimestamp
	}

	sort.Slice(Pipelines, func(i, j int) bool {
		return Pipelines[j].StartTime.Before(Pipelines[i].StartTime)
	})

	s.render.JSON(w, http.StatusOK, Pipelines) //nolint:errcheck
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

	switch method {
	case "POST":
		s.render.JSON(w, http.StatusMethodNotAllowed, nil) //nolint:errcheck
	case "PUT":
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

		if strings.EqualFold(pa.Spec.Status.String(), "pending") || strings.EqualFold(pa.Spec.Status.String(), "running") {
			pr.Spec.Status = pipelineapi.PipelineRunSpecStatusCancelled
		}
		_, err = s.tknClient.TektonV1beta1().PipelineRuns("jx").Update(context.Background(), pr, metav1.UpdateOptions{})
		if err != nil {
			panic(err)
		}
		s.render.JSON(w, http.StatusOK, "pipeline "+name+" stopped") //nolint:errcheck
	default:
		pa, err := s.jxClient.
			Get(context.Background(), name, metav1.GetOptions{})
		if err != nil {
			// Todo: improve error handling!
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		s.render.JSON(w, http.StatusOK, pa) //nolint:errcheck
	}
}
