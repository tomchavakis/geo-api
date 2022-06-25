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
		h.Router.Get("/api/v1/bearing", handle(h.s.bearingRoute))
		h.Router.Get("/api/v1/destination", handle(h.s.destinationRoute))
		h.Router.Get("/api/v1/midpoint", handle(h.s.midpointRoute))
		h.Router.Post("/api/v1/nearestpoint", handle(h.s.nearestPointRoute))
	})
}
