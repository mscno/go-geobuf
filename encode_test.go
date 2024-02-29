package geobuf

import (
	"github.com/mscno/go-geobuf/internal/encode"
	"github.com/paulmach/orb/geojson"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEncodeEmpty(t *testing.T) {
	_, err := Encode(nil)
	require.Error(t, err)
	require.ErrorIs(t, err, ErrNilInput)
}

func TestEncodeBadType(t *testing.T) {
	_, err := Encode("bad type")
	require.Error(t, err)
	require.ErrorIs(t, err, ErrUnsupportedType)
}

func TestEncodeEmptyFeatureCollection(t *testing.T) {
	fc := geojson.NewFeatureCollection()
	data, err := Encode(fc)
	require.NoError(t, err)
	require.NotNil(t, data)
}

func TestEncodeEmptyFeature(t *testing.T) {
	fc := geojson.NewFeature(nil)
	data, err := Encode(fc, WithAllowEmptyGeometry(true))
	require.NoError(t, err)
	require.NotNil(t, data)

	fc2, err := Decode(data)
	require.NoError(t, err)
	require.NotNil(t, fc2)
	fc3, ok := fc2.(*geojson.Feature)
	require.True(t, ok)
	require.NotNil(t, fc3)
	require.Nil(t, fc3.Geometry)
}

var emptyGeometryFeature = `{"type":"Feature","geometry":null}`

func TestEncodeEmptyGeometryFeature(t *testing.T) {
	f, err := geojson.UnmarshalFeature([]byte(emptyGeometryFeature))

	data, err := Encode(f, WithAllowEmptyGeometry(false))
	require.Error(t, err)
	require.ErrorIs(t, err, encode.ErrEmptyGeometry)

	data, err = Encode(f)
	require.Error(t, err)
	require.ErrorIs(t, err, encode.ErrEmptyGeometry)

	data, err = Encode(f, WithAllowEmptyGeometry(true))
	require.NoError(t, err)
	require.NotNil(t, data)

	f2, err := Decode(data)
	require.NoError(t, err)
	require.NotNil(t, f2)
	f3, ok := f2.(*geojson.Feature)
	require.True(t, ok)
	require.NotNil(t, f3)
	require.Nil(t, f3.Geometry)
}
