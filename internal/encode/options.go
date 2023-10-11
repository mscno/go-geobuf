package encode

import (
	"github.com/mscno/go-geobuf/internal/math"
	"github.com/paulmach/orb"
	geojson "github.com/paulmach/orb/geojson"
)

type EncodingConfig struct {
	Dimension          uint
	Precision          uint
	HardcodedPrecision bool
	Keys               KeyStore
}

func AnalyzePrecision(obj interface{}, opts *EncodingConfig) {
	switch t := obj.(type) {
	case *geojson.FeatureCollection:
		for _, feature := range t.Features {
			AnalyzePrecision(feature, opts)
		}
	case *geojson.Feature:
		AnalyzePrecision(geojson.NewGeometry(t.Geometry), opts) // TODO bench fix to not create new geometry
	case *geojson.Geometry:
		switch t.Type {
		case GeometryPoint:
			updatePrecision(t.Geometry().(orb.Point), opts)
		case GeometryMultiPoint:
			coords := t.Geometry().(orb.MultiPoint)
			for _, coord := range coords {
				updatePrecision(coord, opts)
			}
		case GeometryLineString:
			coords := t.Geometry().(orb.LineString)
			for _, coord := range coords {
				updatePrecision(coord, opts)
			}
		case GeometryMultiLineString:
			lines := t.Geometry().(orb.MultiLineString)
			for _, line := range lines {
				for _, coord := range line {
					updatePrecision(coord, opts)
				}
			}
		case GeometryPolygon:
			lines := t.Geometry().(orb.Polygon)
			for _, line := range lines {
				for _, coord := range line {
					updatePrecision(coord, opts)
				}
			}
		case GeometryMultiPolygon:
			polygons := t.Geometry().(orb.MultiPolygon)
			for _, rings := range polygons {
				for _, ring := range rings {
					for _, coord := range ring {
						updatePrecision(coord, opts)
					}
				}
			}
		}
	}
}

func updatePrecision(point orb.Point, opt *EncodingConfig) {
	for _, val := range point {
		e := math.GetPrecision(val)
		if e > opt.Precision {
			opt.Precision = e
		}
	}
}

func AnalyzeKeys(obj interface{}, opts *EncodingConfig) {
	switch t := obj.(type) {
	case *geojson.FeatureCollection:
		for _, feature := range t.Features {
			AnalyzeKeys(feature, opts)
		}
	case *geojson.Feature:
		for key, _ := range t.Properties {
			opts.Keys.Add(key)
		}
	}
}
