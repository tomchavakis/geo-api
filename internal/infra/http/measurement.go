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

func (sh *MeasurementHandler) destinationRoute(w http.ResponseWriter, r *http.Request) (*Response, error) {
	lat, lon, err := getLatLon(r, "lat", "lon")

	if err != nil {
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	p := geometry.Point{
		Lat: *lat,
		Lng: *lon,
	}

	d := r.URL.Query().Get("distance")
	if d == "" {
		return nil, NewResponseError(errors.New("distance can't be empty"), http.StatusBadRequest)
	}
	distance, err := strconv.ParseFloat(d, 64)

	if err != nil {
		return nil, NewResponseError(errors.New("invalid distance"), http.StatusBadRequest)
	}

	b := r.URL.Query().Get("bearing")
	if b == "" {
		return nil, NewResponseError(errors.New("bearing can't be empty"), http.StatusBadRequest)
	}
	bearing, err := strconv.ParseFloat(b, 64)

	if err != nil {
		return nil, NewResponseError(errors.New("invalid bearing"), http.StatusBadRequest)
	}

	units := r.URL.Query().Get("units")

	dp, err := sh.measurementSvc.GetDestination(p, distance, bearing, units)
	if err != nil {
		log.Printf("error %v", err)
		return nil, NewResponseError(errors.New(err.Error()), http.StatusInternalServerError)
	}

	return NewResponse(dp, http.StatusOK), nil
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
