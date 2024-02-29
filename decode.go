package geobuf

import (
	"github.com/mscno/go-geobuf/geobufpb"
	"github.com/mscno/go-geobuf/internal/decode"
	"github.com/paulmach/orb/geojson"
)

func Decode(msg *geobufpb.Data) (interface{}, error) {
	switch v := msg.DataType.(type) {
	case *geobufpb.Data_Geometry_:
		geo := v.Geometry
		return decode.DecodeGeometry(geo, msg.Precision, msg.Dimensions), nil
	case *geobufpb.Data_Feature_:
		return decode.DecodeFeature(msg, v.Feature, msg.Precision, msg.Dimensions), nil
	case *geobufpb.Data_FeatureCollection_:
		collection := geojson.NewFeatureCollection()
		for _, feature := range v.FeatureCollection.Features {
			collection.Append(decode.DecodeFeature(msg, feature, msg.Precision, msg.Dimensions))
		}
		return collection, nil
	}
	return struct{}{}, nil
}
