package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RouteBuilder builds the routes
func (h *HTTP) RouteBuilder() {

	h.Router.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(status)
	})

	h.Router.Get("/metrics", promhttp.Handler().ServeHTTP)
	h.Router.Route("/api/v1", func(r chi.Router) {

		h.Router.Get("/api/v1/distance", handle(h.s.measurementRoute))
	})
}
