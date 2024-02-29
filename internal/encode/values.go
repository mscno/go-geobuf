package encode

import (
	"encoding/json"
	"reflect"

	"github.com/mscno/go-geobuf/geobufpb"
)

func EncodeValue(val interface{}) (*geobufpb.Data_Value, error) {
	v := reflect.ValueOf(val)
	return encodeValue(v, val)
}

func encodeValue(v reflect.Value, val interface{}) (*geobufpb.Data_Value, error) {
	switch v.Kind() {
	case reflect.Bool:
		return encodeBool(v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intval := v.Int()
		if intval < 0 {
			return encodeInt(uint64(v.Int()*-1), false)
		}
		return encodeInt(uint64(v.Int()), true)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return encodeInt(uint64(v.Uint()), true)
	case reflect.Float32, reflect.Float64:
		return encodeDouble(v.Float())
	case reflect.String:
		return encodeString(v.String())
	case reflect.Ptr:
		return encodeValue(v.Elem(), val)
	case reflect.Slice:
		return encodeJSON(val)
	case reflect.Map:
		return encodeJSON(val)
	case reflect.Struct:
		return encodeJSON(val)
	case reflect.Invalid:
		return encodeJSON(val)
	default:
		panic("Unknown type: " + v.Kind().String())
	}
}

func encodeInt(val uint64, positive bool) (*geobufpb.Data_Value, error) {
	if positive {
		return &geobufpb.Data_Value{
			ValueType: &geobufpb.Data_Value_PosIntValue{
				PosIntValue: val,
			},
		}, nil
	}

	return &geobufpb.Data_Value{
		ValueType: &geobufpb.Data_Value_NegIntValue{
			NegIntValue: val,
		},
	}, nil
}

func encodeDouble(val float64) (*geobufpb.Data_Value, error) {
	return &geobufpb.Data_Value{
		ValueType: &geobufpb.Data_Value_DoubleValue{
			DoubleValue: val,
		},
	}, nil
}

func encodeString(val string) (*geobufpb.Data_Value, error) {
	return &geobufpb.Data_Value{
		ValueType: &geobufpb.Data_Value_StringValue{
			StringValue: val,
		},
	}, nil
}

func encodeBool(val bool) (*geobufpb.Data_Value, error) {
	return &geobufpb.Data_Value{
		ValueType: &geobufpb.Data_Value_BoolValue{
			BoolValue: val,
		},
	}, nil
}

func encodeJSON(val interface{}) (*geobufpb.Data_Value, error) {
	encoded, err := json.Marshal(val)
	return &geobufpb.Data_Value{
		ValueType: &geobufpb.Data_Value_JsonValue{
			JsonValue: encoded,
		},
	}, err
}
