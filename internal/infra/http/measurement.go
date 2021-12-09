package http

import (
	"errors"
	"log"
	"net/http"

	"github.com/tomchavakis/turf-api/internal/app/measurement"
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

	// Business Logic
	d, err := sh.measurementSvc.GetDistance()
	if err != nil {
		log.Printf("error %v", err)
		e := NewResponseError(errors.New(err.Error()), http.StatusInternalServerError)
		return nil, e
	}
	return NewResponse(d, http.StatusOK), nil
}
