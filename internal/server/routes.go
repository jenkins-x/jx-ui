package server

import "github.com/gorilla/mux"

func registerRoutes(router *mux.Router, server *Server) *mux.Router {
	router.HandleFunc("/api/v1/pipelines", server.PipelinesHandler)

	router.HandleFunc("/api/v1/pipelines/{owner}/{repo}/{branch}/{build}", server.PipelineHandler).Methods("GET", "PUT")
	router.HandleFunc("/api/v1/logs/{owner}/{repo}/{branch}/{build}", server.PipelineStreamingLogHandler)
	router.HandleFunc("/api/v1/logs_archived/{owner}/{repo}/{branch}/{build}", server.PipelineArchivedLogHandler)
	router.HandleFunc("/api/v1/repositories", server.RepositoriesHandler)
	spa := spaHandler{staticPath: "web/build", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)
	return router
}
