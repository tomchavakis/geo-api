package mock

import (
	"github.com/tomchavakis/geojson/geometry"
)

// MeasurementRepository defines mock functions for Measurement repository.
type MeasurementRepository struct {
	GetDistanceFn     func(x, y geometry.Point) (*float64, error)
	GetBearingFn      func(x, y geometry.Point) (*float64, error)
	GetDestinationFn  func(x geometry.Point, d, b float64, units string) (*geometry.Point, error)
	GetNearestPointFn func(refPoint geometry.Point, points []geometry.Point, units string) (*geometry.Point, error)
	GetMidPointFn     func(x, y geometry.Point) *geometry.Point
}

// NewMockMeasurementRepository builds a mock Repository.
func NewMockMeasurementRepository() *MeasurementRepository {
	return &MeasurementRepository{}
}

// GetDistance ...
func (r *MeasurementRepository) GetDistance(x, y geometry.Point) (*float64, error) {
	if r.GetDistanceFn != nil {
		return r.GetDistanceFn(x, y)
	}
	return nil, nil
}

// GetBearing ...
func (r *MeasurementRepository) GetBearing(x, y geometry.Point) (*float64, error) {
	if r.GetBearingFn != nil {
		return r.GetBearingFn(x, y)
	}
	return nil, nil
}

// GetDestination ...
func (r *MeasurementRepository) GetDestination(x geometry.Point, d, b float64, units string) (*geometry.Point, error) {
	if r.GetDestinationFn != nil {
		return r.GetDestinationFn(x, d, b, units)
	}
	return nil, nil
}

// GetNearestPoint
func (r *MeasurementRepository) GetNearestPoint(refPoint geometry.Point, points []geometry.Point, units string) (*geometry.Point, error) {
	if r.GetNearestPointFn != nil {
		return r.GetNearestPointFn(refPoint, points, units)
	}
	return nil, nil
}

// GetMidPoint
func (r *MeasurementRepository) GetMidPoint(x, y geometry.Point) *geometry.Point {
	if r.GetMidPointFn != nil {
		return r.GetMidPointFn(x, y)
	}
	return nil
}
