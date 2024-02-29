package encode

import (
	"github.com/mscno/go-geobuf/geobufpb"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"reflect"
	"testing"
)

func TestEncodePoint(t *testing.T) {
	testCases := []struct {
		Precision uint
		Expected  []int64
	}{
		{
			Precision: 1000,
			Expected:  []int64{124123, 234456},
		},
		// Should round up when truncating precision
		{
			Precision: 100,
			Expected:  []int64{12412, 23446},
		},
		// Should round up (.5) when truncating precision
		{
			Precision: 10,
			Expected:  []int64{1241, 2345},
		},
		{
			Precision: 1,
			Expected:  []int64{124, 234},
		},
	}

	p := geojson.NewGeometry(orb.Point([2]float64{124.123, 234.456}))
	for i, test := range testCases {
		expected := &geobufpb.Data_Geometry{
			Type:   geobufpb.Data_Geometry_POINT,
			Coords: test.Expected,
		}
		encoded, err := EncodeGeometry(p.Geometry(), &EncodingConfig{
			Dimension: 2,
			Precision: test.Precision,
		})
		if err != nil {
			t.Errorf("Case [%d]: Unexpected error: %s", i, err)
		}

		if !reflect.DeepEqual(encoded, expected) {
			t.Errorf("Case [%d]: Expected %+v, got %+v", i, expected, encoded)
		}
	}

}

func Test_translateMultiPolygon(t *testing.T) {
	type args struct {
		e        uint
		dim      uint
		polygons []orb.Polygon
	}
	tests := []struct {
		name        string
		args        args
		wantCoords  []int64
		wantLengths []uint32
	}{
		{
			name: "simple multi polygon",
			args: args{
				e:   100,
				dim: 2,
				polygons: []orb.Polygon{
					[]orb.Ring{
						[]orb.Point{{102, 2}, {103, 2}, {103, 3}, {102, 3}, {102, 2}},
					},
				},
			},
			wantCoords:  []int64{10200, 200, 100, 0, 0, 100, -100, 0},
			wantLengths: []uint32{1, 1, 4},
		},
		{
			name: "simple multi polygon with hole",
			args: args{
				e:   100,
				dim: 2,
				polygons: []orb.Polygon{
					[]orb.Ring{
						[]orb.Point{{100, 0}, {101, 0}, {101, 1}, {100, 1}, {100, 0}},
						[]orb.Point{{100.2, 0.2}, {100.8, 0.2}, {100.8, 0.8}, {100.2, 0.8}, {100.2, 0.2}},
					},
				},
			},
			wantCoords:  []int64{10000, 0, 100, 0, 0, 100, -100, 0, 10020, 20, 60, 0, 0, 60, -60, 0},
			wantLengths: []uint32{1, 2, 4, 4},
		},
		{
			name: "multi polygon with two separate polygons and one of them has a hole",
			args: args{
				e:   100,
				dim: 2,
				polygons: []orb.Polygon{
					[]orb.Ring{
						[]orb.Point{{102, 2}, {103, 2}, {103, 3}, {102, 3}, {102, 2}},
					},
					[]orb.Ring{
						[]orb.Point{{100, 0}, {101, 0}, {101, 1}, {100, 1}, {100, 0}},
						[]orb.Point{{100.2, 0.2}, {100.8, 0.2}, {100.8, 0.8}, {100.2, 0.8}, {100.2, 0.2}},
					},
				},
			},
			wantCoords:  []int64{10200, 200, 100, 0, 0, 100, -100, 0, 10000, 0, 100, 0, 0, 100, -100, 0, 10020, 20, 60, 0, 0, 60, -60, 0},
			wantLengths: []uint32{2, 1, 4, 2, 4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coords, lengths := translateMultiPolygon(tt.args.e, tt.args.dim, tt.args.polygons)
			if !reflect.DeepEqual(coords, tt.wantCoords) {
				t.Errorf("translateMultiPolygon() got coords = %v, want %v", coords, tt.wantCoords)
			}
			if !reflect.DeepEqual(lengths, tt.wantLengths) {
				t.Errorf("translateMultiPolygon() got lengths = %v, want %v", lengths, tt.wantLengths)
			}
		})
	}
}
