package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/pkg/errors"

	"github.com/tomchavakis/geo-api/internal/app/measurement"
)

// HTTP ...
type HTTP struct {
	Router *chi.Mux
	s      *MeasurementHandler
}

// New constructs a new HTTP
func New(msrSvc measurement.Service) *HTTP {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(NewPrometheusMiddleware("geo-go"))

	r.Use(render.SetContentType(render.ContentTypeJSON))
	return &HTTP{
		Router: r,
		s:      NewMeasurementHandler(msrSvc),
	}
}

func handle(fn func(w http.ResponseWriter, r *http.Request) (*Response, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, err := fn(w, r)
		if err != nil {
			status := http.StatusInternalServerError
			if resp != nil {
				status = resp.Status
			}
			log.Printf("error %v", err)
			_, _ = RespondError(w, status, err)
			return
		}
		_, _ = Respond(w, resp.Status, resp.Payload)
	}
}

// RespondError sends an error reponse back to the client.
func RespondError(w http.ResponseWriter, code int, err error) (int, error) {
	if webErr, ok := errors.Cause(err).(*Error); ok {
		er := ErrorResponse{
			Error: webErr.Err.Error(),
		}
		return Respond(w, webErr.Status, er)

	}
	// If not, the handler sent any arbitrary error value so use 500.
	er := ErrorResponse{
		Error: http.StatusText(http.StatusInternalServerError),
	}

	return Respond(w, code, er)
}

// Respond converts a Go value to XML and sends it to the client.
func Respond(w http.ResponseWriter, code int, payload interface{}) (int, error) {
	response, err := json.Marshal(payload)

	if err != nil {
		log.Printf("error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return w.Write(nil)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	return w.Write(response)
}

// Header is the http header representation as a map of strings
type Header map[string]string

// Response definition of the sync Response model.
type Response struct {
	Payload interface{}
	Status  int
	Header  Header
}

// NewResponse creates a new Response.
func NewResponse(p interface{}, s int) *Response {
	return &Response{Payload: p, Status: s, Header: make(map[string]string)}
}
