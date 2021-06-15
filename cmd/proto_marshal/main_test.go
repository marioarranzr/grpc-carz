package main

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/marioarranzr/carz/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BenchmarkJSON(b *testing.B) {
	car := &pb.Car{
		Id:     "95",
		Driver: "McQueen",
		Location: &pb.Location{
			Lat: -2.242,
			Lng: 2.43,
		},
		Status:  pb.Status_FREE,
		Updated: timestamppb.Now(),
	}
	for i := 0; i < b.N; i++ {
		_, err := proto.Marshal(car)
		if err != nil {
			log.Fatal(err)
		}
	}
}
func BenchmarkPROTO(b *testing.B) {
	car := &pb.Car{
		Id:     "95",
		Driver: "McQueen",
		Location: &pb.Location{
			Lat: -2.242,
			Lng: 2.43,
		},
		Status:  pb.Status_FREE,
		Updated: timestamppb.Now(),
	}
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(car)
		if err != nil {
			log.Fatal(err)
		}
	}
}
