package geobuf

import (
	"errors"
	geoproto "github.com/mscno/go-geobuf/geobufpb"
	"github.com/mscno/go-geobuf/internal/encode"
	"github.com/mscno/go-geobuf/internal/math"
	"github.com/paulmach/orb/geojson"
)

var ErrUnsupportedType = errors.New("unsupported type: object is not geojson")
var ErrNilInput = errors.New("invalid input: object is nil")

type EncodingOption func(o *encode.EncodingConfig)

func WithPrecision(precision uint) EncodingOption {
	return func(o *encode.EncodingConfig) {
		o.HardcodedPrecision = true
		o.Precision = uint(math.DecodePrecision(uint32(precision)))
	}
}

func WithDimension(dimension uint) EncodingOption {
	return func(o *encode.EncodingConfig) {
		o.Dimension = dimension
	}
}

func WithKeys(keys []string) EncodingOption {
	return func(o *encode.EncodingConfig) {
		o.Keys = encode.NewKeyStoreWithKeys(keys)
	}
}

func WithAllowEmptyGeometry(allow bool) EncodingOption {
	return func(o *encode.EncodingConfig) {
		o.AllowEmptyGeometry = allow
	}
}

func Encode(obj interface{}, opts ...EncodingOption) (*geoproto.Data, error) {
	cfg := &encode.EncodingConfig{
		Dimension: 2,
		Precision: 1,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	if cfg.Keys == nil {
		cfg.Keys = encode.NewKeyStore()
		encode.AnalyzeKeys(obj, cfg)
	}

	if !cfg.HardcodedPrecision {
		encode.AnalyzePrecision(obj, cfg)
	}

	data := &geoproto.Data{
		Keys:       cfg.Keys.Keys(),
		Dimensions: uint32(cfg.Dimension),
		Precision:  math.EncodePrecision(cfg.Precision),
	}

	switch t := obj.(type) {
	case nil:
		return nil, ErrNilInput
	case *geojson.FeatureCollection:
		collection, err := encode.EncodeFeatureCollection(t, cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &geoproto.Data_FeatureCollection_{
			FeatureCollection: collection,
		}
	case *geojson.Feature:
		feature, err := encode.EncodeFeature(t, cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &geoproto.Data_Feature_{
			Feature: feature,
		}
	case *geojson.Geometry:
		geom, err := encode.EncodeGeometry(t.Geometry(), cfg)
		if err != nil {
			return nil, err
		}
		data.DataType = &geoproto.Data_Geometry_{
			Geometry: geom,
		}
	default:
		return nil, ErrUnsupportedType
	}

	return data, nil
}
