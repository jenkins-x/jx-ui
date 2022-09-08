package server

import (
	"net/http"
)

// HealthzHandler function
func (s *Server) HealthzHandler(w http.ResponseWriter, r *http.Request) {

	s.render.JSON(w, http.StatusOK, http.StatusText(http.StatusOK)) //nolint:errcheck
}
