package measurement

import (
	"github.com/tomchavakis/turf-go/constants"
	"github.com/tomchavakis/turf-go/geojson/geometry"
	m "github.com/tomchavakis/turf-go/measurement"
)

// Repository ...
type Repository struct {
}

// New generates a new repository.
func New() (*Repository, error) {
	return &Repository{}, nil
}

// GetDistance returns the distance of two points
func (r *Repository) GetDistance(x geometry.Point, y geometry.Point) (*float64, error) {

	d, err := m.PointDistance(x, y, constants.UnitMeters)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// GetBearing returns the bearing of two points
func (r *Repository) GetBearing(x geometry.Point, y geometry.Point) (*float64, error) {
	b := m.PointBearing(x, y)

	return &b, nil
}

// GetDestination returns a point located in a specific destance and bearing from a reference point
func (r *Repository) GetDestination(x geometry.Point, d float64, b float64, units string) (*geometry.Point, error) {
	dest, err := m.Destination(x, d, b, units)
	if err != nil {
		return nil, err
	}

	return dest, nil
}

func (r *Repository) GetMidPoint(x geometry.Point, y geometry.Point) geometry.Point {
	mid := m.MidPoint(x, y)

	return mid
}
