# :hammer: [![GoDoc](https://godoc.org/github.com/tomchavakis/geojson?status.svg)](https://godoc.org/github.com/tomchavakis/geojson) [![GitHub license](https://badgen.net/github/license/tomchavakis/geojson)](https://github.com/tomchavakis/geojson/blob/main/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/tomchavakis/geojson)](https://goreportcard.com/report/github.com/tomchavakis/geojson)

# GeoJSON 

GeoJSON is a GoLang library that implements the [GeoJSON](https://geojson.org/) format specification


## Usage


## GeoJSON Objects - Examples

This library implements the following [GeoJSON Objects](https://datatracker.ietf.org/doc/html/rfc7946#section-3):


### Point
```go
point := geometry.NewPoint(lat,lng)
```

```go
point := geometry.Point{
		Lat: 50.0,
		Lng: 30.0,
	}
```
### MultiPoint

```go
multipoint := geometry.NewPoint(lat,lng)
```

```go
multipoint := geometry.Point{
		Lat: 50.0,
		Lng: 30.0,
	}
```

### LineString

```go
lineString := geometry.LineString{
	Coordinates: []geometry.Point{
		{
			Lat: 30.0,
			Lng: 23.0,
		},
		{
			Lat: 32.0,
			Lng: 24.0,
		},
	},

}
```

### MultiLineString

```go
multiLineString := geometry.MultiLineString{
		Coordinates: []geometry.LineString{
			{
				Coordinates: []geometry.Point{
					{
						Lat: 43.3,
						Lng: 24.5,
					},
					{
						Lat: 44.2,
						Lng: 25.0,
					},
				},
			},
			{
				Coordinates: []geometry.Point{
					{
						Lat: 38.3,
						Lng: 40.5,
					},
					{
						Lat: 38.2,
						Lng: 5.0,
					},
				},
			},
		},
	}
```

### Polygon

```go
poly := geometry.Polygon{
		Coordinates: []geometry.LineString{
			{
				Coordinates: []geometry.Point{
					{
						Lat: 36.171278341935434,
						Lng: -86.76624298095703,
					},
					{
						Lat: 36.170862616662134,
						Lng: -86.74238204956055,
					},
					{
						Lat: 36.19607929145354,
						Lng: -86.74100875854492,
					},
					{
						Lat: 36.2014818084173,
						Lng: -86.77362442016602,
					},
					{
						Lat: 36.171278341935434,
						Lng: -86.76624298095703,
					},
				},
			},
		},
```
### MultiPolygon

### Feature

- New Feature

```go	
	geom := geometry.Geometry{
		GeoJSONType: geojson.Polygon,
		Coordinates: [][]float64{
			{
				36.171278341935434,
				-86.76624298095703,
			},
			{
				36.170862616662134,
				-86.74238204956055,
			},
			{
				36.19607929145354,
				-86.74100875854492,
			},
			{
				36.2014818084173,
				-86.77362442016602,
			},
			{
				36.171278341935434,
				-86.76624298095703,
			},
		},
	}

	f, err := feature.New(geom, []float64{}, nil, "geom-1")```

- FromJSON:

```go
	 gjson: "{ \"type\" : \"Feature\",  \"properties\": {},   \"geometry\": { \"type\": \"Point\",  \"coordinates\": [-71, 41]} }",
	if err != nil {
		t.Errorf("LoadJSONFixture error %v", err)
	}

	f, err := feature.FromJSON(gjson)
```

#### Supported Functions:

  - [x] ToPoint
  - [x] ToPolygon
  - [x] ToMultiPolygon
  - [x] ToLineString
  - [x] ToMultiLineString

### FeatureCollection
```go
	gjson := "{ \"type\": \"FeatureCollection\", \"features\": [ { \"type\": \"Feature\", \"properties\": {}, \"geometry\": { \"type\": \"Polygon\", \"coordinates\": [ [ [-2.109375, 47.040182144806664], [4.5703125, 44.59046718130883], [7.03125, 49.15296965617042], [-3.515625, 49.83798245308484], [-2.109375, 47.040182144806664] ] ] } }, { \"type\": \"Feature\", \"properties\": {}, \"geometry\": { \"type\": \"Polygon\", \"coordinates\": [ [ [9.64599609375, 47.70976154266637], [9.4482421875, 47.73932336136857], [8.8330078125, 47.47266286861342], [10.21728515625, 46.604167162931844], [11.755371093749998, 46.81509864599243], [11.865234375, 47.90161354142077], [9.64599609375, 47.70976154266637] ] ] } } ] }"

	collection, err := feature.CollectionFromJSON(gjson)
	if err != nil {
		t.Errorf("CollectionFromJSON error: %v", err)
	}	
```

### Geometry

```go
geometry, err := geometry.FromJSON("{ \"type\": \"MultiPolygon\", \"coordinates\": [ [ [ [-116, -45], [-90, -45], [-90, -56], [-116, -56], [-116, -45] ] ], [ [ [-90.351563, 9.102097], [-77.695312, -3.513421], [-65.039063, 12.21118], [-65.742188, 21.616579], [-84.023437, 24.527135], [-90.351563, 9.102097] ] ] ] }")
```
#### Supported Functions:

  - [x] ToPoint
  - [x] ToPolygon
  - [x] ToMultiPolygon
  - [x] ToLineString
  - [x] ToMultiLineString

### GeometryCollection

```go
gjson := "{ \"TYPE\": \"GeometryCollection\", \"geometries\": [ { \"TYPE\": \"Point\", \"coordinates\": [-7903683.846322424, 5012341.663847514] }, { \"TYPE\": \"LineString\", \"coordinates\": [ [-2269873.991957, 3991847.410438], [-391357.58482, 6026906.856034], [1565430.33928, 1917652.163291] ] } ] }"

gcl, err := geometry.CollectionFromJSON
if err != nil {
	t.Errorf("CollectionFromJSON error: %v", err)
}	
```

## turf-go
You can also use the library [turf-go](https://github.com/tomchavakis/turf-go) which is an implementation of the [turf.js](https://turfjs.org/) library in GoLang.