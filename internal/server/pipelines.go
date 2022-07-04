package server

import (
	"context"
	"net/http"
	"sort"

	"github.com/gorilla/mux"
	"github.com/jenkins-x/jx-helpers/v3/pkg/kube/naming"
	pipelineapi "github.com/tektoncd/pipeline/pkg/apis/pipeline/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PipelinesHandler function
func (s *Server) PipelinesHandler(w http.ResponseWriter, r *http.Request) {
	pa, err := s.jxClient.
		List(context.Background(), metav1.ListOptions{})
	if err != nil {
		// Todo: improve error handling!
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sort.Slice(pa.Items, func(i, j int) bool {
		return pa.Items[j].Spec.StartedTimestamp.Before(pa.Items[i].Spec.StartedTimestamp)
	})

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
