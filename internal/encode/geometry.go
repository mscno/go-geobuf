package encode

import (
	"errors"
	"github.com/mscno/go-geobuf/geobufpb"
	"github.com/mscno/go-geobuf/internal/math"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

const (
	GeometryPoint           = "Point"
	GeometryMultiPoint      = "MultiPoint"
	GeometryLineString      = "LineString"
	GeometryMultiLineString = "MultiLineString"
	GeometryPolygon         = "Polygon"
	GeometryMultiPolygon    = "MultiPolygon"
)

const GeometryCollection = "GeometryCollection"

var ErrEmptyGeometry = errors.New("empty geometry")
var ErrGeometryCollectionUnsupported = errors.New("geometry collection not supported")
var ErrUnsupportInputType = errors.New("unsupported input type")

func EncodeGeometry(g orb.Geometry, opt *EncodingConfig) (*geobufpb.Data_Geometry, error) {
	if g == nil {
		if !opt.AllowEmptyGeometry {
			return nil, ErrEmptyGeometry
		}
		return &geobufpb.Data_Geometry{
			Type: geobufpb.Data_Geometry_EMPTY,
		}, nil
	}
	switch g.GeoJSONType() {
	case geojson.TypePoint:
		p := g.(orb.Point)
		return &geobufpb.Data_Geometry{
			Type:   geobufpb.Data_Geometry_POINT,
			Coords: translateCoords(opt.Precision, p[:]),
		}, nil
	case geojson.TypeMultiPoint:
		p := g.(orb.MultiPoint)
		return &geobufpb.Data_Geometry{
			Type:   geobufpb.Data_Geometry_MULTIPOINT,
			Coords: translateLine(opt.Precision, opt.Dimension, p, false),
		}, nil
	case geojson.TypeLineString:
		p := g.(orb.LineString)
		return &geobufpb.Data_Geometry{
			Type:   geobufpb.Data_Geometry_LINESTRING,
			Coords: translateLine(opt.Precision, opt.Dimension, p, false),
		}, nil
	case geojson.TypeMultiLineString:
		p := g.(orb.MultiLineString)
		coords, lengths := translateMultiLine(opt.Precision, opt.Dimension, p)
		return &geobufpb.Data_Geometry{
			Type:    geobufpb.Data_Geometry_MULTILINESTRING,
			Coords:  coords,
			Lengths: lengths,
		}, nil
	case geojson.TypePolygon:
		p := []orb.Ring(g.(orb.Polygon))
		coords, lengths := translateMultiRing(opt.Precision, opt.Dimension, p)
		return &geobufpb.Data_Geometry{
			Type:    geobufpb.Data_Geometry_POLYGON,
			Coords:  coords,
			Lengths: lengths,
		}, nil
	case geojson.TypeMultiPolygon:
		p := []orb.Polygon(g.(orb.MultiPolygon))
		coords, lengths := translateMultiPolygon(opt.Precision, opt.Dimension, p)
		return &geobufpb.Data_Geometry{
			Type:    geobufpb.Data_Geometry_MULTIPOLYGON,
			Coords:  coords,
			Lengths: lengths,
		}, nil
	case GeometryCollection:
		// Hack to ensure that the geometry collection is not empty and that it contains at least one point
		// This is because geojson.GeoJSONType() returns "GeometryCollection" for empty geometries
		if x, ok := g.(orb.Collection); ok {
			if len(x) == 0 {
				if !opt.AllowEmptyGeometry {
					return nil, ErrEmptyGeometry
				}
				return &geobufpb.Data_Geometry{
					Type: geobufpb.Data_Geometry_EMPTY,
				}, nil
			}
			return nil, ErrGeometryCollectionUnsupported
		}
		return nil, ErrGeometryCollectionUnsupported
	default:
		return nil, ErrUnsupportInputType
	}
}

func translateMultiLine(e uint, dim uint, lines []orb.LineString) ([]int64, []uint32) {
	lengths := make([]uint32, len(lines))
	coords := []int64{}

	for i, line := range lines {
		lengths[i] = uint32(len(line))
		coords = append(coords, translateLine(e, dim, line, false)...)
	}
	return coords, lengths
}

func translateMultiPolygon(e uint, dim uint, polygons []orb.Polygon) ([]int64, []uint32) {
	lengths := []uint32{uint32(len(polygons))}
	coords := []int64{}
	for _, rings := range polygons {
		lengths = append(lengths, uint32(len(rings)))
		newLine, newLength := translateMultiRing(e, dim, rings)
		lengths = append(lengths, newLength...)
		coords = append(coords, newLine...)
	}
	return coords, lengths
}

func translateMultiRing(e uint, dim uint, lines []orb.Ring) ([]int64, []uint32) {
	lengths := make([]uint32, len(lines))
	coords := []int64{}
	for i, line := range lines {
		lengths[i] = uint32(len(line) - 1)
		newLine := translateLine(e, dim, line, true)
		coords = append(coords, newLine...)
	}
	return coords, lengths
}

/*
Since we're converting floats to ints, we can get additional compression out of
how Google does varint encoding (#1). Smaller numbers can be packed into less bytes,
even when using large primitives (int64). To take advantage of this, we subtract
out the prior coordinate x/y value from the current coordinate x/y value to (hopefully)
reduce the number to a small integer.

For example: (123.123, 234.234), (123.134, 234.236) would be encoded out to
(123123, 234234), (11, 2). The first point takes the full penalty for encoding size,
while the remaining points become much smaller.

A further enhancement comes from the fact that lines that start and end in the same place,
such as with a polygon, we can skip the last point, and place it back when we decode.

1. https://developers.google.com/protocol-buffers/docs/encoding#varints
*/
func translateLine(precision uint, dim uint, points []orb.Point, isClosed bool) []int64 {
	sums := make([]int64, dim)
	ret := make([]int64, len(points)*int(dim))
	for i, point := range points {
		for j, p := range point {
			n := math.IntWithPrecision(p, precision) - sums[j]
			ret[(int(dim)*i)+j] = n
			sums[j] = sums[j] + n
		}
	}
	if isClosed {
		return ret[:(len(ret) - int(dim))]
	}
	return ret
}

// Converts a floating point geojson point to int64 by multiplying it by a factor of 10,
// potentially truncating and rounding
func translateCoords(precision uint, point []float64) []int64 {
	ret := make([]int64, len(point))
	for i, p := range point {
		ret[i] = math.IntWithPrecision(p, precision)
	}
	return ret
}
