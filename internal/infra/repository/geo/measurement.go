package measurement

import (
	"github.com/tomchavakis/turf-go/constants"
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
func (r *Repository) GetDistance() (*float64, error) {

	d, err := m.Distance(20.0, 44.3, 21.0, 45.9, constants.UnitMeters)
	if err != nil {
		return nil, err
	}

	return &d, nil
}
