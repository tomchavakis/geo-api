package http

import (
	"errors"
	"net/http"
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
