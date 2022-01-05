package measurement

import "github.com/tomchavakis/turf-go/geojson/geometry"

// Service ...
type Service interface {
	GetDistance(x geometry.Point, y geometry.Point) (*float64, error)
	GetBearing(x geometry.Point, y geometry.Point) (*float64, error)
}
