package encode

import (
	"github.com/mscno/go-geobuf/geobufpb"
	"github.com/paulmach/orb/geojson"
)

func EncodeFeatureCollection(collection *geojson.FeatureCollection, opts *EncodingConfig) (*geobufpb.Data_FeatureCollection, error) {
	features := make([]*geobufpb.Data_Feature, len(collection.Features))

	for i, feature := range collection.Features {
		encoded, err := EncodeFeature(feature, opts)
		if err != nil {
			return nil, err
		}
		features[i] = encoded
	}

	return &geobufpb.Data_FeatureCollection{
		Features: features,
	}, nil
}
