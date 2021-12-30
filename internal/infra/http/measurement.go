package http

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/tomchavakis/geo-api/internal/app/measurement"
	"github.com/tomchavakis/turf-go/geojson/geometry"
)

// MeasurementHandler struct
type MeasurementHandler struct {
	measurementSvc measurement.Service
}

// NewMeasurementHandler handler
func NewMeasurementHandler(msrSvc measurement.Service) *MeasurementHandler {
	mh := &MeasurementHandler{
		measurementSvc: msrSvc,
	}
	return mh
}

func (sh *MeasurementHandler) measurementRoute(w http.ResponseWriter, r *http.Request) (*Response, error) {

	var latA, lonA, latB, lonB float64
	var err error

	lat0 := r.URL.Query().Get("latA")
	latA, err = strconv.ParseFloat(lat0, 64)
	if err != nil {
		return nil, NewResponseError(errors.New("invalid point"), http.StatusBadRequest)
	}
	lon0 := r.URL.Query().Get("lonA")
	lonA, err = strconv.ParseFloat(lon0, 64)
	if err != nil {
		return nil, NewResponseError(errors.New("invalid point"), http.StatusBadRequest)
	}

	lat1 := r.URL.Query().Get("latB")
	latB, err = strconv.ParseFloat(lat1, 64)
	if err != nil {
		return nil, NewResponseError(errors.New("invalid point"), http.StatusBadRequest)
	}

	lon1 := r.URL.Query().Get("lonB")
	lonB, err = strconv.ParseFloat(lon1, 64)
	if err != nil {
		return nil, NewResponseError(errors.New("invalid point"), http.StatusBadRequest)
	}

	p1 := geometry.Point{
		Lat: latA,
		Lng: lonA,
	}

	p2 := geometry.Point{
		Lat: latB,
		Lng: lonB,
	}

	// Business Logic
	d, err := sh.measurementSvc.GetDistance(p1, p2)
	if err != nil {
		log.Printf("error %v", err)
		return nil, NewResponseError(errors.New(err.Error()), http.StatusInternalServerError)
	}

	return NewResponse(d, http.StatusOK), nil
}

// As Path Params:
// if l0 := chi.URLParam(r, "latA"); l0 != "" {
// 	latA, err = strconv.ParseFloat(l0, 64)
// 	if err != nil {
// 		return nil, NewResponseError(errors.New("invalid point"), http.StatusBadRequest)
// 	}
// }

// if l1 := chi.URLParam(r, "lonA"); l1 != "" {
// 	latA, err = strconv.ParseFloat(l1, 64)
// 	if err != nil {
// 		return nil, NewResponseError(errors.New("invalid point"), http.StatusBadRequest)
// 	}
// }

// if l2 := chi.URLParam(r, "latB"); l2 != "" {
// 	latA, err = strconv.ParseFloat(l2, 64)
// 	if err != nil {
// 		return nil, NewResponseError(errors.New("invalid point"), http.StatusBadRequest)
// 	}
// }

// if l3 := chi.URLParam(r, "lonB"); l3 != "" {
// 	latA, err = strconv.ParseFloat(l3, 64)
// 	if err != nil {
// 		return nil, NewResponseError(errors.New("invalid point"), http.StatusBadRequest)
// 	}
// }
