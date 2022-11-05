package server

import (
	"github.com/gorilla/mux"
)

func registerRoutes(s *Server) *mux.Router {
	s.router.HandleFunc("/healthz", s.HealthzHandler)
	s.router.HandleFunc("/api/v1/pipelines", s.Pipeline.PipelinesHandler)

	s.router.HandleFunc("/api/v1/pipelines/{owner}/{repo}/{branch}/{build}", s.Pipeline.PipelineHandler).Methods("GET", "PUT")
	s.router.HandleFunc("/api/v1/logs/{owner}/{repo}/{branch}/{build}", s.PipelineStreamingLogHandler)
	s.router.HandleFunc("/api/v1/logs_archived/{owner}/{repo}/{branch}/{build}", s.PipelineArchivedLogHandler)
	s.router.HandleFunc("/api/v1/repositories", s.Repositories.RepositoriesHandler)
	spa := spaHandler{staticPath: "web/build", indexPath: "index.html"}
	s.router.PathPrefix("/").Handler(spa)
	return s.router
}
