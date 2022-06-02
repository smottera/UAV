//build all the control structs here. then copy to txrx.go once working
package main

import (
	"fmt"

	pb "./protofiles"
	"github.com/golang/protobuf/proto"
)

func main() {

	p := &pb.CameraControlPacket{
		Pan:                  10,
		Tilt:                 20,
		Record:               30,
		Focus:                40,
		Zoom:                 550,
		Flash:                false,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     []byte{},
		XXX_sizecache:        0,
	}
	p1 := &pb.CameraControlPacket{}

	body, _ := proto.Marshal(p)

	_ = proto.Unmarshal(body, p1)

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshaled proto data: ", body)
	fmt.Println("Unmarshaled struct: ", p1)
}
