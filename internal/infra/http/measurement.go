package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/tomchavakis/geo-api/internal/app/measurement"
	"github.com/tomchavakis/geojson/geometry"
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

// NearestPointMessage ...
type NearestPointMessage struct {
	ReferencePoint *geometry.Point  `json:"ref,omitempty"`
	Points         []geometry.Point `json:"points"`
	Units          string           `json:"units"`
}

func (sh *MeasurementHandler) distanceRoute(w http.ResponseWriter, r *http.Request) (*Response, error) {
	latA, lonA, err := getLatLon(r, "latA", "lonA")

	if err != nil {
		return nil, NewResponseError(errors.New("invalid input"), http.StatusBadRequest)
	}

	latB, lonB, err := getLatLon(r, "latB", "lonB")

	if err != nil {
		return nil, NewResponseError(errors.New("invalid input"), http.StatusBadRequest)
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
		return nil, NewResponseError(errors.New(err.Error()), http.StatusInternalServerError)
	}

	return NewResponse(d, http.StatusOK), nil
}

func (sh *MeasurementHandler) bearingRoute(w http.ResponseWriter, r *http.Request) (*Response, error) {
	latA, lonA, err := getLatLon(r, "latA", "lonA")

	if err != nil {
		return nil, NewResponseError(errors.New("invalid input"), http.StatusBadRequest)
	}

	latB, lonB, err := getLatLon(r, "latB", "lonB")

	if err != nil {
		return nil, NewResponseError(errors.New("invalid input"), http.StatusBadRequest)
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
		return nil, NewResponseError(errors.New("invalid input"), http.StatusBadRequest)
	}

	latB, lonB, err := getLatLon(r, "latB", "lonB")

	if err != nil {
		return nil, NewResponseError(errors.New("invalid input"), http.StatusBadRequest)
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

	return NewResponse(*midpoint, http.StatusOK), nil
}

func (sh *MeasurementHandler) nearestPointRoute(w http.ResponseWriter, r *http.Request) (*Response, error) {
	if r.Body == nil {
		err := errors.New("invalid Body")
		return nil, NewResponseError(err, http.StatusBadRequest)
	}
	var np NearestPointMessage
	err := json.NewDecoder(r.Body).Decode(&np)
	if err != nil {
		return nil, NewResponseError(errors.New("invalid input"), http.StatusBadRequest)
	}

	if np.ReferencePoint == nil {
		err := errors.New("reference point can't be empty")
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	if len(np.Points) == 0 {
		err := errors.New("points can't be empty")
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	nearestPoint, err := sh.measurementSvc.GetNearestPoint(*np.ReferencePoint, np.Points, np.Units)

	if err != nil {
		return nil, NewResponseError(err, http.StatusBadRequest)
	}

	return NewResponse(nearestPoint, http.StatusOK), nil
}

func (sh *MeasurementHandler) destinationRoute(w http.ResponseWriter, r *http.Request) (*Response, error) {
	lat, lon, err := getLatLon(r, "lat", "lon")

	if err != nil {
		return nil, NewResponseError(errors.New("invalid input"), http.StatusBadRequest)
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

func getLatLon(r *http.Request, lat, lon string) (*float64, *float64, error) {
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
