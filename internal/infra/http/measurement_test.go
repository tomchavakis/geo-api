package http

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomchavakis/geo-api/internal/common"
	"github.com/tomchavakis/geo-api/test/mock"
	"github.com/tomchavakis/geojson/geometry"
)

func TestDistance(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/distance?latA=23.33&lonA=34.44&latB=23.44&lonB=34.42", nil)
	assert.NoError(t, err)

	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := map[string]struct {
		mockGetDistance func(x, y geometry.Point) (*float64, error)
		want            *Response
		wantErr         bool
		err             error
		args            args
	}{
		"happy path": {
			want: NewResponse(common.Float64Ptr(10.0), http.StatusOK),
			mockGetDistance: func(x, y geometry.Point) (*float64, error) {
				return common.Float64Ptr(10.0), nil
			},
			wantErr: false,
			err:     nil,
			args: args{
				w: nil,
				r: req,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
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
