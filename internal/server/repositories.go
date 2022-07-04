package server

import (
	"context"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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
