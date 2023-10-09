package geobuf

import (
	"github.com/mscno/go-geobuf/internal/encode"
	"github.com/mscno/go-geobuf/internal/math"
	geoproto "github.com/mscno/go-geobuf/proto"
	"github.com/paulmach/orb/geojson"
)

func Encode(obj interface{}) (*geoproto.Data, error) {
	return EncodeWithOptions(obj, encode.FromAnalysis(obj))
}

type EncodingOption func(o *encode.EncodingConfig)

func WithPrecision(precision uint) EncodingOption {
	return func(o *encode.EncodingConfig) {
		o.Precision = uint(math.DecodePrecision(uint32(precision)))
	}
}

func WithDimension(dimension uint) EncodingOption {
	return func(o *encode.EncodingConfig) {
		o.Dimension = dimension
	}
}

func WithKeyStore(store encode.KeyStore) EncodingOption {
	return func(o *encode.EncodingConfig) {
		o.Keys = store
	}
}

func EncodeWithOptions(obj interface{}, opts ...EncodingOption) (*geoproto.Data, error) {
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

	data := &geoproto.Data{
		Keys:       cfg.Keys.Keys(),
		Dimensions: uint32(cfg.Dimension),
		Precision:  math.EncodePrecision(cfg.Precision),
	}

	switch t := obj.(type) {
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
		data.DataType = &geoproto.Data_Geometry_{
			Geometry: encode.EncodeGeometry(t.Geometry(), cfg),
		}
	}

	return data, nil
}
