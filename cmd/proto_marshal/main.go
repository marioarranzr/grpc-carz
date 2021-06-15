package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/marioarranzr/carz/pb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	// var loc pb.Location
	loc := pb.Location{
		Lat: -22.24221,
		Lng: 3.32342,
	}

	data, err := proto.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}
	// os.Stdout.Write(data)
	fmt.Println(data)

	jdata, err := json.Marshal(&loc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON size", len(jdata))

	var loc2 pb.Location
	if err := proto.Unmarshal(data, &loc2); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n\n", &loc2)

	s := pb.Status_FREE
	fmt.Printf("free: %s (%d)\n\n", s, s)

	car := &pb.Car{
		Id:       "95",
		Driver:   "McQueen",
		Location: &loc,
		Status:   pb.Status_FREE,
		Updated:  timestamppb.Now(),
	}
	fmt.Println(car)
	t := car.Updated.AsTime()
	fmt.Println(t)

	fmt.Printf("%#v\n\n", car) // prints the protobuf
	fmt.Printf("%v\n\n", car)  // prints the protobuf

	pdata, err := proto.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PROTO size", len(pdata))
	jdata, err = json.Marshal(&car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON size", len(jdata))

	fmt.Println("json =>", string(jdata)) // Will use the MarshalJSON function overided in json.go
	jdata, err = protojson.Marshal(car)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("protojson =>", string(jdata))
}
