package measurement

import "github.com/tomchavakis/turf-go/geojson/geometry"

// Service ...
type Service interface {
	GetDistance(x geometry.Point, y geometry.Point) (*float64, error)
	GetBearing(x geometry.Point, y geometry.Point) (*float64, error)
	GetDestination(x geometry.Point, distance float64, bearing float64, units string) (*geometry.Point, error)
	GetMidPoint(x geometry.Point, y geometry.Point) geometry.Point
}
