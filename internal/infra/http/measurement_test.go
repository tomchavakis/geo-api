package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomchavakis/geo-api/internal/common"
	"github.com/tomchavakis/geo-api/test/mock"
	"github.com/tomchavakis/geojson/geometry"
)

func TestDistance(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := map[string]struct {
		mockGetDistance func(x, y geometry.Point) (*float64, error)
		want            *Response
		request         string
		wantErr         bool
		err             error
		args            args
	}{
		"invalid input A": {
			want: nil,
			mockGetDistance: func(x, y geometry.Point) (*float64, error) {
				return common.Float64Ptr(10.0), nil
			},
			request: "/api/v1/distance?latA=a&lonA=34.44&latB=23.44&lonB=34.42",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid input"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"invalid input B": {
			want: nil,
			mockGetDistance: func(x, y geometry.Point) (*float64, error) {
				return common.Float64Ptr(10.0), nil
			},
			request: "/api/v1/distance?latA=23.33&lonA=34.44&latB=b&lonB=34.42",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid input"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"get distance error": {
			want: nil,
			mockGetDistance: func(x, y geometry.Point) (*float64, error) {
				return nil, errors.New("get distance error")
			},
			request: "/api/v1/distance?latA=23.33&lonA=34.44&latB=23.44&lonB=34.42",
			wantErr: true,
			err:     NewResponseError(errors.New("get distance error"), http.StatusInternalServerError),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"happy path": {
			want: NewResponse(common.Float64Ptr(10.0), http.StatusOK),
			mockGetDistance: func(x, y geometry.Point) (*float64, error) {
				return common.Float64Ptr(10.0), nil
			},
			request: "/api/v1/distance?latA=23.33&lonA=34.44&latB=23.44&lonB=34.42",
			wantErr: false,
			err:     nil,
			args: args{
				w: nil,
				r: nil,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.request, nil)
			assert.NoError(t, err)
			tt.args.r = req

			MockSvc := mock.NewMockMeasurementRepository()
			MockSvc.GetDistanceFn = tt.mockGetDistance
			h := NewMeasurementHandler(MockSvc)
			got, err := h.distanceRoute(tt.args.w, tt.args.r)
			if tt.wantErr || err != nil {
				assert.Equal(t, tt.err, err, "distance() error = %v,expected = %v", err, tt.err)
				return
			}
			assert.Equal(t, tt.want, got, "distance() got = %v, want %v", got, tt.want)
		})
	}
}

func TestBearing(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := map[string]struct {
		mockGetBearing func(x, y geometry.Point) (*float64, error)
		want           *Response
		request        string
		wantErr        bool
		err            error
		args           args
	}{
		"invalid input A": {
			want: nil,
			mockGetBearing: func(x, y geometry.Point) (*float64, error) {
				return common.Float64Ptr(10.0), nil
			},
			request: "/api/v1/bearing?latA=a&lonA=34.44&latB=23.44&lonB=34.42",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid input"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"invalid input B": {
			want: nil,
			mockGetBearing: func(x, y geometry.Point) (*float64, error) {
				return common.Float64Ptr(10.0), nil
			},
			request: "/api/v1/bearing?latA=23.33&lonA=34.44&latB=b&lonB=34.42",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid input"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"get bearing error": {
			want: nil,
			mockGetBearing: func(x, y geometry.Point) (*float64, error) {
				return nil, errors.New("get bearing error")
			},
			request: "/api/v1/bearing?latA=23.33&lonA=34.44&latB=23.44&lonB=34.42",
			wantErr: true,
			err:     NewResponseError(errors.New("get bearing error"), http.StatusInternalServerError),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"happy path": {
			want: NewResponse(common.Float64Ptr(10.0), http.StatusOK),
			mockGetBearing: func(x, y geometry.Point) (*float64, error) {
				return common.Float64Ptr(10.0), nil
			},
			request: "/api/v1/bearing?latA=23.33&lonA=34.44&latB=23.44&lonB=34.42",
			wantErr: false,
			err:     nil,
			args: args{
				w: nil,
				r: nil,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.request, nil)
			assert.NoError(t, err)
			tt.args.r = req

			MockSvc := mock.NewMockMeasurementRepository()
			MockSvc.GetBearingFn = tt.mockGetBearing
			h := NewMeasurementHandler(MockSvc)
			got, err := h.bearingRoute(tt.args.w, tt.args.r)
			if tt.wantErr || err != nil {
				assert.Equal(t, tt.err, err, "bearing() error = %v,expected = %v", err, tt.err)
				return
			}
			assert.Equal(t, tt.want, got, "bearing() got = %v, want %v", got, tt.want)
		})
	}
}

func TestMidPoint(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := map[string]struct {
		mockGetMidPoint func(x, y geometry.Point) *geometry.Point
		want            *Response
		request         string
		wantErr         bool
		err             error
		args            args
	}{
		"invalid input A": {
			want: nil,
			mockGetMidPoint: func(x, y geometry.Point) *geometry.Point {
				return nil
			},
			request: "/api/v1/midpoint?latA=a&lonA=34.44&latB=23.44&lonB=34.42",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid input"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"invalid input B": {
			want: nil,
			mockGetMidPoint: func(x, y geometry.Point) *geometry.Point {
				return nil
			},
			request: "/api/v1/midpoint?latA=23.33&lonA=34.44&latB=b&lonB=34.42",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid input"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"happy path": {
			want: NewResponse(*geometry.NewPoint(23.38500031791607, 34.430004151010706), http.StatusOK),
			mockGetMidPoint: func(x, y geometry.Point) *geometry.Point {
				return geometry.NewPoint(23.38500031791607, 34.430004151010706)
			},
			request: "/api/v1/midpoint?latA=23.33&lonA=34.44&latB=23.44&lonB=34.42",
			wantErr: false,
			err:     nil,
			args: args{
				w: nil,
				r: nil,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.request, nil)
			assert.NoError(t, err)
			tt.args.r = req

			MockSvc := mock.NewMockMeasurementRepository()
			MockSvc.GetMidPointFn = tt.mockGetMidPoint
			h := NewMeasurementHandler(MockSvc)
			got, err := h.midpointRoute(tt.args.w, tt.args.r)
			if tt.wantErr || err != nil {
				assert.Equal(t, tt.err, err, "midpoint() error = %v,expected = %v", err, tt.err)
				return
			}
			assert.Equal(t, tt.want, got, "midpoint() got = %v, want %v", got, tt.want)
		})
	}
}

func TestDestinationPoint(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := map[string]struct {
		mockGetDestinationPoint func(x geometry.Point, d, b float64, units string) (*geometry.Point, error)
		want                    *Response
		request                 string
		wantErr                 bool
		err                     error
		args                    args
	}{
		"invalid input A": {
			want: nil,
			mockGetDestinationPoint: func(x geometry.Point, d, b float64, units string) (*geometry.Point, error) {
				return nil, nil
			},
			request: "/api/v1/destination?lat=a&lon=34.44&distance=10&bearing=10&units=m",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid input"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"get destination error": {
			want: nil,
			mockGetDestinationPoint: func(x geometry.Point, d, b float64, units string) (*geometry.Point, error) {
				return nil, errors.New("getDestination error")
			},
			request: "/api/v1/destination?lat=23.33&lon=34.44&distance=10&bearing=10&units=m",
			wantErr: true,
			err:     NewResponseError(errors.New("getDestination error"), http.StatusInternalServerError),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"empty distance": {
			want: nil,
			mockGetDestinationPoint: func(x geometry.Point, d, b float64, units string) (*geometry.Point, error) {
				return nil, nil
			},
			request: "/api/v1/destination?lat=23.33&lon=34.44&distance=&bearing=10&units=m",
			wantErr: true,
			err:     NewResponseError(errors.New("distance can't be empty"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"distance error": {
			want: nil,
			mockGetDestinationPoint: func(x geometry.Point, d, b float64, units string) (*geometry.Point, error) {
				return nil, nil
			},
			request: "/api/v1/destination?lat=23.33&lon=34.44&distance=a&bearing=10&units=m",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid distance"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"empty bearing": {
			want: nil,
			mockGetDestinationPoint: func(x geometry.Point, d, b float64, units string) (*geometry.Point, error) {
				return nil, nil
			},
			request: "/api/v1/destination?lat=23.33&lon=34.44&distance=10&bearing=&units=m",
			wantErr: true,
			err:     NewResponseError(errors.New("bearing can't be empty"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"bearing error": {
			want: nil,
			mockGetDestinationPoint: func(x geometry.Point, d, b float64, units string) (*geometry.Point, error) {
				return nil, nil
			},
			request: "/api/v1/destination?lat=23.33&lon=34.44&distance=10&bearing=a&units=m",
			wantErr: true,
			err:     NewResponseError(errors.New("invalid bearing"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"happy path": {
			want: NewResponse(geometry.NewPoint(24.215564151448817, 34.61122551304243), http.StatusOK),
			mockGetDestinationPoint: func(x geometry.Point, d, b float64, units string) (*geometry.Point, error) {
				return geometry.NewPoint(24.215564151448817, 34.61122551304243), nil
			},
			request: "/api/v1/destination?lat=23.33&lon=34.44&distance=10&bearing=10&units=m",
			wantErr: false,
			err:     nil,
			args: args{
				w: nil,
				r: nil,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.request, nil)
			assert.NoError(t, err)
			tt.args.r = req

			MockSvc := mock.NewMockMeasurementRepository()
			MockSvc.GetDestinationFn = tt.mockGetDestinationPoint
			h := NewMeasurementHandler(MockSvc)
			got, err := h.destinationRoute(tt.args.w, tt.args.r)
			if tt.wantErr || err != nil {
				assert.Equal(t, tt.err, err, "destination() error = %v,expected = %v", err, tt.err)
				return
			}
			assert.Equal(t, tt.want, got, "destination() got = %v, want %v", got, tt.want)
		})
	}
}

func TestNearestPoint(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := map[string]struct {
		mockGetNearestPoint func(refPoint geometry.Point, points []geometry.Point, units string) (*geometry.Point, error)
		want                *Response
		request             string
		payload             NearestPointMessage
		wantErr             bool
		err                 error
		args                args
	}{
		"nil points": {
			want: nil,
			mockGetNearestPoint: func(refPoint geometry.Point, points []geometry.Point, units string) (*geometry.Point, error) {
				return nil, nil
			},
			payload: NearestPointMessage{
				ReferencePoint: nil,
				Points:         nil,
				Units:          "",
			},
			request: "/api/v1/nearestpoint",
			wantErr: true,
			err:     NewResponseError(errors.New("reference point can't be empty"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"nil reference point empty": {
			want: nil,
			mockGetNearestPoint: func(refPoint geometry.Point, points []geometry.Point, units string) (*geometry.Point, error) {
				return nil, nil
			},
			payload: NearestPointMessage{
				ReferencePoint: nil,
				Points:         []geometry.Point{},
				Units:          "",
			},
			request: "/api/v1/nearestpoint",
			wantErr: true,
			err:     NewResponseError(errors.New("reference point can't be empty"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"empty points": {
			want: nil,
			mockGetNearestPoint: func(refPoint geometry.Point, points []geometry.Point, units string) (*geometry.Point, error) {
				return nil, nil
			},
			payload: NearestPointMessage{
				ReferencePoint: &geometry.Point{
					Lat: 39.50,
					Lng: -75.33,
				},
				Points: []geometry.Point{},
				Units:  "",
			},
			request: "/api/v1/nearestpoint",
			wantErr: true,
			err:     NewResponseError(errors.New("points can't be empty"), http.StatusBadRequest),
			args: args{
				w: nil,
				r: nil,
			},
		},
		"happy path": {
			want: NewResponse(geometry.NewPoint(39.46, -75.30), http.StatusOK),
			mockGetNearestPoint: func(refPoint geometry.Point, points []geometry.Point, units string) (*geometry.Point, error) {
				return geometry.NewPoint(39.46, -75.30), nil
			},
			payload: NearestPointMessage{
				ReferencePoint: &geometry.Point{
					Lat: 39.50,
					Lng: -75.33,
				},
				Points: []geometry.Point{
					{
						Lat: 39.44,
						Lng: -75.33,
					},
					{
						Lat: 39.45,
						Lng: -75.33,
					},
					{
						Lat: 39.46,
						Lng: -75.31,
					},
					{
						Lat: 39.46,
						Lng: -75.30,
					},
				},
				Units: "",
			},
			request: "/api/v1/nearestpoint",
			wantErr: false,
			err:     nil,
			args: args{
				w: nil,
				r: nil,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			p, err := json.Marshal(tt.payload)
			assert.NoError(t, err, "error marshalling the payload")

			req, err := http.NewRequest("POST", tt.request, strings.NewReader(string(p)))
			assert.NoError(t, err)
			tt.args.r = req

			MockSvc := mock.NewMockMeasurementRepository()
			MockSvc.GetNearestPointFn = tt.mockGetNearestPoint
			h := NewMeasurementHandler(MockSvc)
			got, err := h.nearestPointRoute(tt.args.w, tt.args.r)
			if tt.wantErr || err != nil {
				assert.Equal(t, tt.err, err, "nearestpoint() error = %v,expected = %v", err, tt.err)
				return
			}
			assert.Equal(t, tt.want, got, "nearestpoint() got = %v, want %v", got, tt.want)
		})
	}
}
