package decode

import (
	"bytes"
	"encoding/json"
	"github.com/mscno/go-geobuf/geobufpb"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
)

func DecodeFeature(msg *geobufpb.Data, feature *geobufpb.Data_Feature, precision, dimension uint32) *geojson.Feature {
	geo := feature.Geometry
	decodedGeo := DecodeGeometry(geo, msg.Precision, msg.Dimensions)

	var geoFeature *geojson.Feature
	switch decodedGeo.Type {
	case "FeatureCollection":
		collection := make(orb.Collection, len(decodedGeo.Geometries))
		for i, child := range decodedGeo.Geometries {
			collection[i] = child.Coordinates
		}
		geoFeature = geojson.NewFeature(collection)
	default:
		geoFeature = geojson.NewFeature(decodedGeo.Coordinates)
	}

	for i := 0; i < len(feature.Properties); i = i + 2 {
		keyIdx := feature.Properties[i]
		valIdx := feature.Properties[i+1]
		val := feature.Values[valIdx]
		switch actualVal := val.ValueType.(type) {
		case *geobufpb.Data_Value_BoolValue:
			geoFeature.Properties[msg.Keys[keyIdx]] = actualVal.BoolValue
		case *geobufpb.Data_Value_DoubleValue:
			geoFeature.Properties[msg.Keys[keyIdx]] = actualVal.DoubleValue
		case *geobufpb.Data_Value_StringValue:
			geoFeature.Properties[msg.Keys[keyIdx]] = actualVal.StringValue
		case *geobufpb.Data_Value_PosIntValue:
			geoFeature.Properties[msg.Keys[keyIdx]] = uint(actualVal.PosIntValue)
		case *geobufpb.Data_Value_NegIntValue:
			geoFeature.Properties[msg.Keys[keyIdx]] = int(actualVal.NegIntValue) * -1
		case *geobufpb.Data_Value_JsonValue:
			if bytes.Equal(actualVal.JsonValue, []byte("null")) {
				geoFeature.Properties[msg.Keys[keyIdx]] = nil
				continue
			}
			if actualVal.JsonValue == nil {
				geoFeature.Properties[msg.Keys[keyIdx]] = nil
				continue
			}
			if actualVal.JsonValue[0] == '{' {
				var m map[string]interface{}
				err := json.Unmarshal(actualVal.JsonValue, &m)
				if err != nil {
					panic(err)
				}
				geoFeature.Properties[msg.Keys[keyIdx]] = m
			} else {
				var a []interface{}
				err := json.Unmarshal(actualVal.JsonValue, &a)
				if err != nil {
					panic(err)
				}
				geoFeature.Properties[msg.Keys[keyIdx]] = a
			}
		}
	}
	switch id := feature.IdType.(type) {
	case *geobufpb.Data_Feature_Id:
		if id != nil {
			geoFeature.ID = id.Id
		}
	case *geobufpb.Data_Feature_IntId:
		geoFeature.ID = id.IntId
	}
	return geoFeature
}
