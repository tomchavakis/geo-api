package geometry

import "github.com/tomchavakis/geojson"

// Object is the base interface for GeometryObject types
type Object struct {
	Type geojson.OBjectType
}
