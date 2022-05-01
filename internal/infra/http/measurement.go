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

	latA, lonA, err := getLatLon(r, "latA", "lonA")

	if err != nil {
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	latB, lonB, err := getLatLon(r, "latB", "lonB")

	if err != nil {
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	p1 := geometry.Point{
		Lat: *latA,
		Lng: *lonA,
	}

	p2 := geometry.Point{
		Lat: *latB,
		Lng: *lonB,
	}

	// Business Logic
	d, err := sh.measurementSvc.GetDistance(p1, p2)
	if err != nil {
		log.Printf("error %v", err)
		return nil, NewResponseError(errors.New(err.Error()), http.StatusInternalServerError)
	}

	return NewResponse(d, http.StatusOK), nil
}

func (sh *MeasurementHandler) bearingRoute(w http.ResponseWriter, r *http.Request) (*Response, error) {
	latA, lonA, err := getLatLon(r, "latA", "lonA")

	if err != nil {
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	latB, lonB, err := getLatLon(r, "latB", "lonB")

	if err != nil {
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	p1 := geometry.Point{
		Lat: *latA,
		Lng: *lonA,
	}

	p2 := geometry.Point{
		Lat: *latB,
		Lng: *lonB,
	}

	// Business Logic
	b, err := sh.measurementSvc.GetBearing(p1, p2)
	if err != nil {
		log.Printf("error %v", err)
		return nil, NewResponseError(errors.New(err.Error()), http.StatusInternalServerError)
	}

	return NewResponse(b, http.StatusOK), nil
}

func (sh *MeasurementHandler) midpointRoute(w http.ResponseWriter, r *http.Request) (*Response, error) {
	latA, lonA, err := getLatLon(r, "latA", "lonA")

	if err != nil {
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	latB, lonB, err := getLatLon(r, "latB", "lonB")

	if err != nil {
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	p1 := geometry.Point{
		Lat: *latA,
		Lng: *lonA,
	}

	p2 := geometry.Point{
		Lat: *latB,
		Lng: *lonB,
	}

	midpoint := sh.measurementSvc.GetMidPoint(p1, p2)

	return NewResponse(midpoint, http.StatusOK), nil
}

func getLatLon(r *http.Request, lat string, lon string) (*float64, *float64, error) {
	lat0 := r.URL.Query().Get(lat)
	latA, err := strconv.ParseFloat(lat0, 64)
	if err != nil {
		return nil, nil, errors.New("invalid point")
	}
	lon0 := r.URL.Query().Get(lon)
	lonA, err := strconv.ParseFloat(lon0, 64)
	if err != nil {
		return nil, nil, errors.New("invalid point")
	}

	return &latA, &lonA, nil
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
