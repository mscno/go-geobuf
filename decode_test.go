package geobuf_test

import (
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"

	. "github.com/mscno/go-geobuf"
)

func TestDecodePoint(t *testing.T) {
	p := geojson.NewGeometry(orb.Point([]float64{124.123, 234.456}))
	encoded, err := Encode(p)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeMultiPoint(t *testing.T) {
	p := geojson.NewGeometry(orb.MultiPoint([]orb.Point{
		orb.Point([]float64{124.123, 234.456}),
		orb.Point([]float64{345.567, 456.678}),
	}))
	encoded, err := Encode(p)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeLineString(t *testing.T) {
	p := geojson.NewGeometry(orb.LineString([]orb.Point{
		orb.Point([]float64{124.123, 234.456}),
		orb.Point([]float64{345.567, 456.678}),
	}))
	encoded, err := Encode(p)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeMultiLineString(t *testing.T) {
	p := geojson.NewGeometry(orb.MultiLineString([]orb.LineString{
		orb.LineString([]orb.Point{
			orb.Point([]float64{124.123, 234.456}),
			orb.Point([]float64{345.567, 456.678}),
		}),
		orb.LineString([]orb.Point{
			orb.Point([]float64{224.123, 334.456}),
			orb.Point([]float64{445.567, 556.678}),
		}),
	}))
	encoded, err := Encode(p)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodePolygon(t *testing.T) {
	p := geojson.NewGeometry(orb.Polygon([]orb.Ring{
		orb.Ring([]orb.Point{
			orb.Point([]float64{124.123, 234.456}),
			orb.Point([]float64{345.567, 456.678}),
			orb.Point([]float64{124.123, 234.456}),
		}),
		orb.Ring([]orb.Point{
			orb.Point([]float64{224.123, 334.456}),
			orb.Point([]float64{445.567, 556.678}),
			orb.Point([]float64{224.123, 334.456}),
		}),
	}))
	encoded, err := Encode(p)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeMultiPolygon(t *testing.T) {
	p := geojson.NewGeometry(
		orb.MultiPolygon([]orb.Polygon{
			orb.Polygon([]orb.Ring{
				orb.Ring([]orb.Point{
					orb.Point([]float64{124.123, 234.456}),
					orb.Point([]float64{345.567, 456.678}),
					orb.Point([]float64{124.123, 234.456}),
				}),
				orb.Ring([]orb.Point{
					orb.Point([]float64{224.123, 334.456}),
					orb.Point([]float64{445.567, 556.678}),
					orb.Point([]float64{224.123, 334.456}),
				}),
			}),
			orb.Polygon([]orb.Ring{
				orb.Ring([]orb.Point{
					orb.Point([]float64{124.123, 234.456}),
					orb.Point([]float64{345.567, 456.678}),
					orb.Point([]float64{124.123, 234.456}),
				}),
				orb.Ring([]orb.Point{
					orb.Point([]float64{224.123, 334.456}),
					orb.Point([]float64{445.567, 556.678}),
					orb.Point([]float64{224.123, 334.456}),
				}),
			}),
		}))
	encoded, err := Encode(p)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeMultiPolygonEfficient(t *testing.T) {
	p := geojson.NewGeometry(
		orb.MultiPolygon([]orb.Polygon{
			orb.Polygon([]orb.Ring{
				orb.Ring([]orb.Point{
					orb.Point([]float64{124.123, 234.456}),
					orb.Point([]float64{345.567, 456.678}),
					orb.Point([]float64{124.123, 234.456}),
				}),
				orb.Ring([]orb.Point{
					orb.Point([]float64{224.123, 334.456}),
					orb.Point([]float64{445.567, 556.678}),
					orb.Point([]float64{224.123, 334.456}),
				}),
			}),
		}))
	encoded, err := Encode(p)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeFeatureIntId(t *testing.T) {
	p := geojson.NewFeature(orb.Polygon([]orb.Ring{
		orb.Ring([]orb.Point{
			orb.Point([]float64{124.123, 234.456}),
			orb.Point([]float64{345.567, 456.678}),
			orb.Point([]float64{124.123, 234.456}),
		}),
		orb.Ring([]orb.Point{
			orb.Point([]float64{224.123, 334.456}),
			orb.Point([]float64{445.567, 556.678}),
			orb.Point([]float64{224.123, 334.456}),
		}),
	}))
	p.ID = int64(1)
	p.Properties["int"] = uint(4)
	p.Properties["float"] = float64(2.0)
	p.Properties["neg_int"] = -1
	p.Properties["string"] = "string"
	p.Properties["bool"] = true
	encoded, err := Encode(p)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeFeatureStringId(t *testing.T) {
	p := geojson.NewFeature(orb.Polygon([]orb.Ring{
		orb.Ring([]orb.Point{
			orb.Point([]float64{124.123, 234.456}),
			orb.Point([]float64{345.567, 456.678}),
			orb.Point([]float64{124.123, 234.456}),
		}),
		orb.Ring([]orb.Point{
			orb.Point([]float64{224.123, 334.456}),
			orb.Point([]float64{445.567, 556.678}),
			orb.Point([]float64{224.123, 334.456}),
		}),
	}))
	p.ID = "1234"
	p.Properties["int"] = uint(4)
	p.Properties["float"] = float64(2.0)
	p.Properties["neg_int"] = -1
	p.Properties["string"] = "string"
	p.Properties["bool"] = true
	encoded, err := Encode(p)
	require.NoError(t, err)

	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(p, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeFeatureCollection(t *testing.T) {
	p := geojson.NewFeature(orb.Polygon([]orb.Ring{
		orb.Ring([]orb.Point{
			orb.Point([]float64{124.123, 234.456}),
			orb.Point([]float64{345.567, 456.678}),
			orb.Point([]float64{124.123, 234.456}),
		}),
		orb.Ring([]orb.Point{
			orb.Point([]float64{224.123, 334.456}),
			orb.Point([]float64{445.567, 556.678}),
			orb.Point([]float64{224.123, 334.456}),
		}),
	}))
	p.ID = "1234"
	p.Properties["int"] = uint(4)
	p.Properties["float"] = float64(2.0)
	p.Properties["neg_int"] = -1
	p.Properties["string"] = "string"
	p.Properties["bool"] = true

	p2 := geojson.NewFeature(orb.Polygon([]orb.Ring{
		orb.Ring([]orb.Point{
			orb.Point([]float64{224.123, 334.456}),
			orb.Point([]float64{445.567, 556.678}),
			orb.Point([]float64{224.123, 334.456}),
		}),
		orb.Ring([]orb.Point{
			orb.Point([]float64{124.123, 234.456}),
			orb.Point([]float64{345.567, 456.678}),
			orb.Point([]float64{124.123, 234.456}),
		}),
	}))
	p2.ID = "5679"
	p2.Properties["int"] = uint(4)
	p2.Properties["float"] = float64(2.0)
	p2.Properties["neg_int"] = -1
	p2.Properties["string"] = "string"
	p2.Properties["bool"] = true

	collection := geojson.NewFeatureCollection()
	collection.Append(p)
	collection.Append(p2)
	encoded, err := Encode(collection)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	if !reflect.DeepEqual(collection, decoded) {
		t.Errorf("Expected %+v, got %+v", p, decoded)
	}
}

func TestDecodeFeatureMultiPolygonWithCustomPrecision(t *testing.T) {
	// todo add test for bounding box: "bbox":[-83.647031,33.698307,-83.275933,33.9659119]
	var feature_s = `{"id":"1000001","type":"Feature","geometry":{"type":"MultiPolygon","coordinates":[[[[-83.537385,33.9659119],[-83.5084519,33.931233],[-83.4155119,33.918541],[-83.275933,33.847977],[-83.306619,33.811444],[-83.28034,33.7617739],[-83.29145,33.7343149],[-83.406189,33.698307],[-83.479523,33.802265],[-83.505928,33.81776],[-83.533165,33.820923],[-83.647031,33.9061979],[-83.537385,33.9659119]]],[[[-83.537385,33.9659119],[-83.5084519,33.931233],[-83.4155119,33.918541],[-83.275933,33.847977],[-83.306619,33.811444],[-83.28034,33.7617739],[-83.29145,33.7343149],[-83.406189,33.698307],[-83.479523,33.802265],[-83.505928,33.81776],[-83.533165,33.820923],[-83.647031,33.9061979],[-83.537385,33.9659119]]],[[[-83.537385,33.9659119],[-83.5084519,33.931233],[-83.4155119,33.918541],[-83.275933,33.847977],[-83.306619,33.811444],[-83.28034,33.7617739],[-83.29145,33.7343149],[-83.406189,33.698307],[-83.479523,33.802265],[-83.505928,33.81776],[-83.533165,33.820923],[-83.647031,33.9061979],[-83.537385,33.9659119]]]]},"properties":{"AREA":"13219","COLORKEY":"#03E174","area":"13219","index":1109}}`
	var feature, _ = geojson.UnmarshalFeature([]byte(feature_s))

	encoded, err := EncodeWithOptions(feature, WithPrecision(7))
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)

	assert.Equal(t, feature.ID, "1000001")
	assert.Equal(t, feature, decoded)

}

func TestDecodeFeatureMultiPolygonWithHoles(t *testing.T) {
	var feature_s = `{"type":"Feature","properties":{},"geometry":{"type":"MultiPolygon","coordinates":[[[[102,2],[103,2],[103,3],[102,3],[102,2]]],[[[100,0],[101,0],[101,1],[100,1],[100,0]],[[100.2,0.2],[100.8,0.2],[100.8,0.8],[100.2,0.8],[100.2,0.2]]]]}}`
	var feature, _ = geojson.UnmarshalFeature([]byte(feature_s))

	encoded, err := Encode(feature)
	require.NoError(t, err)
	decoded, err := Decode(encoded)
	require.NoError(t, err)
	assert.Equal(t, uint32(1), encoded.Precision)
	assert.Equal(t, feature, decoded)

}
