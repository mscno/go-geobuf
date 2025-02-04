package encode

import (
	"github.com/mscno/go-geobuf/geobufpb"
	"github.com/paulmach/orb/geojson"
)

func EncodeFeature(feature *geojson.Feature, opts *EncodingConfig) (*geobufpb.Data_Feature, error) {
	oldGeo := geojson.NewGeometry(feature.Geometry) // TODO Do we ned this line?
	geo, err := EncodeGeometry(oldGeo.Geometry(), opts)
	if err != nil {
		return nil, err

	}
	f := &geobufpb.Data_Feature{
		Geometry: geo,
	}

	id, err := EncodeIntId(feature.ID)
	if err == nil {
		f.IdType = id
	} else {
		newId, newErr := EncodeId(feature.ID)
		if newErr != nil {
			return nil, newErr
		}
		f.IdType = newId
	}

	properties := make([]uint32, 0, 2*len(feature.Properties))
	values := make([]*geobufpb.Data_Value, 0, len(feature.Properties))
	for key, val := range feature.Properties {
		encoded, err := EncodeValue(val)
		if err != nil {
			return f, err
		}

		idx := opts.Keys.IndexOf(key)
		values = append(values, encoded)
		properties = append(properties, uint32(idx))
		properties = append(properties, uint32(len(values)-1))
	}

	f.Values = values
	f.Properties = properties
	return f, nil
}
