package measurement

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tomchavakis/geo-api/internal/common"
	"github.com/tomchavakis/geojson/geometry"
)

func TestGetBearing(t *testing.T) {

	type args struct {
		x geometry.Point
		y geometry.Point
	}

	tests := map[string]struct {
		args    args
		want    *float64
		wantErr bool
		err     error
	}{
		"zero bearing xy": {
			args: args{
				x: geometry.Point{
					Lat: 20.0,
					Lng: 44.0,
				},
				y: geometry.Point{
					Lat: 21.0,
					Lng: 44.0,
				},
			},
			want:    common.Float64Ptr(0),
			wantErr: false,
			err:     nil,
		},
		"positive bearing xy": {
			args: args{
				x: geometry.Point{
					Lat: 20.5,
					Lng: 44.5,
				},
				y: geometry.Point{
					Lat: 21.0,
					Lng: 44.0,
				},
			},
			want:    common.Float64Ptr(317.00810984889324),
			wantErr: false,
			err:     nil,
		},
	}

	for name, tt := range tests {
		r := &Repository{}

		t.Run(name, func(t *testing.T) {
			b, err := r.GetBearing(tt.args.x, tt.args.y)
			assert.NotNil(t, b)
			assert.Equal(t, *b, *tt.want)

			if err != nil || tt.wantErr {
				assert.Equal(t, tt.err, err, "GetBearing() error = %q, wantErr %q", err, tt.err)
			}

		})
	}
}
