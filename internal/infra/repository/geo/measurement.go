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

// GetDistance returns the distance of 2 points
func (r *Repository) GetDistance(x geometry.Point, y geometry.Point) (*float64, error) {

	d, err := m.PointDistance(x, y, constants.UnitMeters)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
