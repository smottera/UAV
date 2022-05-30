//build all the control structs here. then copy to txrx.go once working
//upgrade golang version then try bidirectional gRPC
package main

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

type cameraControlPacket struct {
	pan    int32
	tilt   int32
	record int32
	focus  int32
	zoom   float32
	flash  bool
}

func main() {

	p := &cameraControlPacket{
		1,
		2,
		3,
		4,
		5.11,
		false,
	}
	p1 := &cameraControlPacket{}

	/*

		p := &pb.Person{
			Id:    0,
			Name:  "Roger F",
			Email: "rf@example.com",
		}

		p1 := &pb.Person{} */

	body, _ := proto.Marshal(p)

	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshaled proto data: ", body)
	fmt.Println("Unmarshaled struct: ", p1)
}
