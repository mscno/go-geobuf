package main

import (
	"encoding/hex"
	"github.com/mscno/go-geobuf"
	"github.com/paulmach/orb"
	"github.com/paulmach/orb/geojson"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
)

func main() {
	demoPoint()
	demoPointWithPrecision()
	demoMultiPolygon()

}

func demoMultiPolygon() {
	f, err := os.ReadFile("cmd/feature.geojson")
	if err != nil {
		log.Fatal(err)
	}

	feature, err := geojson.UnmarshalFeature(f)
	if err != nil {
		log.Fatal(err)
	}

	data, err := geobuf.Encode(feature)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := proto.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	payloadBase64 := hex.EncodeToString(payload)

	err = os.WriteFile("cmd/feature.multipolygon.geobuf.base64.txt", []byte(payloadBase64), 0644)
}

func demoPoint() {
	p := orb.Point([2]float64{124.123, 234.456})
	feature := geojson.NewFeature(p)
	data, err := geobuf.Encode(feature)
	if err != nil {
		log.Fatal(err)
	}

	payload, err := proto.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	payloadBase64 := hex.EncodeToString(payload)

	err = os.WriteFile("cmd/feature.point.geobuf.base64.txt", []byte(payloadBase64), 0644)
}

func demoPointWithPrecision() {
	p := orb.Point([2]float64{124.123, 234.456_789})
	feature := geojson.NewFeature(p)
	data, err := geobuf.Encode(feature, geobuf.WithPrecision(6))
	if err != nil {
		log.Fatal(err)
	}

	payload, err := proto.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	payloadBase64 := hex.EncodeToString(payload)

	err = os.WriteFile("cmd/feature.point_with_precision.geobuf.base64.txt", []byte(payloadBase64), 0644)
}
