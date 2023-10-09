package decode

import (
	"github.com/paulmach/orb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_makeMultiPolygon(t *testing.T) {
	type args struct {
		lengths   []uint32
		inCords   []int64
		precision uint32
		dimension uint32
	}
	tests := []struct {
		name string
		args args
		want orb.MultiPolygon
	}{
		{
			name: "simple multi polygon",
			args: args{
				inCords:   []int64{10200, 200, 100, 0, 0, 100, -100, 0},
				lengths:   []uint32{1, 1, 4},
				dimension: 2,
				precision: 2,
			},
			want: []orb.Polygon{
				[]orb.Ring{
					[]orb.Point{{102, 2}, {103, 2}, {103, 3}, {102, 3}, {102, 2}},
				},
			},
		},
		{
			name: "simple multi polygon with hole",
			args: args{
				inCords:   []int64{10000, 0, 100, 0, 0, 100, -100, 0, 10020, 20, 60, 0, 0, 60, -60, 0},
				lengths:   []uint32{1, 2, 4, 4},
				dimension: 2,
				precision: 2,
			},
			want: []orb.Polygon{
				[]orb.Ring{
					[]orb.Point{{100, 0}, {101, 0}, {101, 1}, {100, 1}, {100, 0}},
					[]orb.Point{{100.2, 0.2}, {100.8, 0.2}, {100.8, 0.8}, {100.2, 0.8}, {100.2, 0.2}},
				},
			},
		},
		{
			name: "multi polygon with two separate polygons and one of them has a hole",
			args: args{
				inCords:   []int64{10200, 200, 100, 0, 0, 100, -100, 0, 10000, 0, 100, 0, 0, 100, -100, 0, 10020, 20, 60, 0, 0, 60, -60, 0},
				lengths:   []uint32{2, 1, 4, 2, 4, 4},
				dimension: 2,
				precision: 2,
			},
			want: []orb.Polygon{
				[]orb.Ring{
					[]orb.Point{{102, 2}, {103, 2}, {103, 3}, {102, 3}, {102, 2}},
				},
				[]orb.Ring{
					[]orb.Point{{100, 0}, {101, 0}, {101, 1}, {100, 1}, {100, 0}},
					[]orb.Point{{100.2, 0.2}, {100.8, 0.2}, {100.8, 0.8}, {100.2, 0.8}, {100.2, 0.2}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := makeMultiPolygon(tt.args.lengths, tt.args.inCords, tt.args.precision, tt.args.dimension)
			assert.Equal(t, tt.want, got)
		})
	}
}
