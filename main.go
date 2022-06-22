/*
___________________________________________________________________________
                                golang daemon
                                                ___________
                                GUI <---------> | TxRx.go |
                                                |         | 1--------> (simple OS) Fully Automatic pilot intelligence,
                   Commandline shell <--------> |         |                    UAV comms established, Systems initialized, Mission path planning,
												|	      |		      path to rigid body dynamics matrices.
 NLP / Human Behavior based control    <------> |         | 	                  (This is the main logic! Just 1 step behind virtual joystick output
                                                | TxRx.go |
      xbox/playstation/thrustmaster  ---------> |         | 2-------> virtual joystick output
                                                |         |                            (get values from fast buffers, output to virtual driver)
            Custom Physical Joystick ---------> | TxRx.go |                            (Use with RC sim)
                                                |         |                            (Must be integrated with a C++ windows driver)
												|	      |
	                                        	|         | 3-------> sbus output
                                                | TxRx.go |           (directly talk to FC) (Must be integrated with Golang uart packages
	                                        	|         |
	                                        	|         | 4-------> Traffic Management memebership
		                                        | TxRx.go |          (depends on a PostgreSQL DB)
                                                |         |
                                                |_________| 5-------> image/video transmission
								      									(frames are captured, compressed and minced before dispatch to cloud)
                                                                      	(images/frames need to be memcached in Redis
___________________________________________________________________________
FROM TxRx.go
Build supporting data structures, primary variables.
Build a function that randomizes these variables for future testing purposes
Build other 3 variants of gRPC services for testing
Test over LAN and then internet (NGINX required?)


*/

// TESTS
// test with multiple server / client instances
// test what happens to stream during connection loss
// txrx.go should be able to heal connection after loss ASAP
//12. test new proto file -> Test Latency and reliability -> test on SoC -> test under different network circumstances
//5. Implement Drone Social Network (drone ID, drone specific data, authorization)
//6. Implement Unmanned Traffic Management function with the help of the above info
//7. Implement necessary backend storage and caching services as needed
//8. Design and build a scheduler?
//9. create stubClient, stubServer, stubRepeater subPackages?
/*
protoc -I ./ protofiles/person.proto --go_out=plugins=grpc:.
protoc -I ./ protofiles/person.proto --go-grpc_out=plugins=grpc:.

*/

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "./protofiles"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port      = ":50051"
	noOfSteps = 3
)

//priority 0
var (
	uniqueID string

	batteryLow   bool = true
	safe2Land    bool = false
	returnToHome bool = false
	disarm       bool = true
	mode         int
	homeLat      float32
	homeLon      float32
	homeAlt      float32
)

//priority 1
type controlPacket struct {
	throttle   int //Channel 0
	rudder     int
	aileron    int
	elevator   int
	motorPower int
	aux1       int
	aux2       int
	aux3       int //Channel 7
}

//priority 2
type missionPacket struct {
	missionID uint
	latitude  float32
	longitude float32
	altitude  float32
	repeat    bool
}

//priority 3
type telemetryPacket struct {
	batteryVoltage float32
	currentDraw    float32
	longitude      float32
	latitude       float32
	altitude       float32
	temperature    float32
	motorRPM       int
}

func droneDummyDataGenerator(iterate int, delay int) {

	for j := 0; j <= iterate; j++ {

		time.Sleep(time.Millisecond * time.Duration(delay))
		//init these variables

		rand1 := float32(rand.Float64())
		timestamp := time.Now()
		safe2Land = false
		disarm = true
		returnToHome = false

		mode = 0
		homeLat = 123.321 * rand1
		homeLon = 987.789 * rand1
		homeAlt = 6900.32 * rand1

		t1 := telemetryPacket{12.3 * rand1, 123.2 * rand1, 123123.3123123 * rand1, 3333.555 * rand1, 23.46, 69.4 * rand1, 9999}
		m1 := missionPacket{123, 32.55 * rand1, 112.34 * rand1, 12.99 * rand1, false}
		c1 := controlPacket{0, 0, 1, 2, 3, 4, 5, 6}

		fmt.Println("Yoo  ", rand1, timestamp, t1, m1, c1, safe2Land, mode, disarm, returnToHome, homeAlt, homeLat, homeLon)
	}
}

// server is used to create uavControlServer.
type server struct {
	pb.UnimplementedUavControlServer
}

func (s *server) GetTelemetry(in *pb.Acknowledged, stream pb.UavControl_GetTelemetryServer) error {
	log.Printf("Got request for mor....")
	log.Printf("a: $%s", in.A)
	// Send streams here
	for i := 0; i < noOfSteps; i++ {
		// Simulating I/O or Computation process using sleep........
		// Usually this will be saving money transfer details in DB or
		// talk to the third party API
		time.Sleep(time.Second * 1)
		// Once task is done, send the successful message back to the client
		if err := stream.Send(&pb.TelemetryPacket{BatteryVoltage: 1.234, CurrentDraw: 3.2, Altitude: 33.92}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}

	log.Printf("Successfully transfered amount $%v ", in.A)
	return nil
}

func initSys() {
	fmt.Println("Hello Early Adopter. Welcome to TxRx.go")
	batteryService()
	missionService()
	sensorCheck()
	//get list of of all peripherals connected
	//check wireless connections
	//check communication with HQ/cloud
	//check if flight controls are working correctly
	//check if data is correctly being logged

	fmt.Println("All Systems Go!")
}

func sensorCheck() int {
	return 0
}
func batteryService() {
	//check health of all cells
	//log data
	//calculate range, lifespan, next service date
	fmt.Println("Battery status ... unknown")
}

func missionService() {
	fmt.Println("Mission scheduler empty ...!")
	//inputs: mission start time, deadline (time), description, distance,
	// path[[gcode]], landing bool?, payload info, attempts, failure?complete?,
	//return address
	//backup return address
	//check mission queue

}

func landingService() {
	fmt.Println("Preparing for landing ...")
	//stop motion, reset attitude
	//plan path for landing
	//begin path to decend
	//check sensors for possibilty of landing
	//continue phase 2 of decent
	//touch down
	//poweroff main motor
	//Disarm
	fmt.Println("Landing complete.")
}

func PathPlanneringProMax() {
	fmt.Println("Path planned.")
	//inputs:starting,ending points,
	//regions to avoid, regions to pass through
	//collion prevention with other UAVs
	//path dependent on time?
	//output: path coordinates, keep track of timing

}

func commsHeartBeat() {
	//check ping/latency with all comms
	//measure signal strength, in proximity UAVs and HQs
	//log everything
	fmt.Println("All comms okay.")

	//check if connection with HQ is okay
	//check if sbus output to microcontroller is okay
	//in case of signal loss: continue with current path? path plan a new Return to home then land
}
func getJoystickData() {
	fmt.Println("NIGGA PWEASE")
}

func testProtoMarshalling() {

	p := &pb.CameraControlPacket{
		Pan:                  0,
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

	then := time.Now()
	body, _ := proto.Marshal(p)

	_ = proto.Unmarshal(body, p1)

	fmt.Println("---------------------------------------------------")

	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshaled proto data: ", body)
	fmt.Println("Unmarshaled struct: ", p1)

	now := time.Now()
	diff := now.Sub(then)
	fmt.Println("Time taken: ", diff)

}
func main() {

	initSys()
	//testProtoMarshalling()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a new GRPC Server
	s := grpc.NewServer()

	// Register it with Proto service
	pb.RegisterUavControlServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
