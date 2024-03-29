/*__________________________________________________________________________
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
	  Fully Automatic Piloting System ------->  |	      |
	                                        |         | 3-------> sbus output
                                                | TxRx.go |           (directly talk to FC) (Must be integrated with Golang uart packages
	                                        |         |
	                                        |         | 4-------> Traffic Management memebership
		                                | TxRx.go |          (depends on a PostgreSQL DB)
                                                |         |
                                                |_________| 5-------> image/video transmission
								      									(frames are captured, compressed and minced before dispatch to cloud)
                                                                      	(images/frames need to be memcached in Redis
__________________________________________________________________________

-----TxRx.go mini feature list

Dashboard/Frontend GUI
WebGL, OpenCV, Google map APIs for orthomosaic manipulation feature.
Functional and optimized bidirectional gRPC comms.
Live video streaming.
Mission and airtraffic management.
GPS-RTK usable centimeter-level accuracy.
S.Bus
Computer Controlled UAVs and USVs system (autopilot, waypoint, pathPlanning, etc).
Fully functioning and necessary backend drivers for databases(redis and postgres).


-----Misc:
test under different network circumstances
watchdog function
Design and build a scheduler?


//setup gateway, test across the internet
// test with multiple server / client instances
// test what happens to stream during connection loss
// txrx.go should be able to heal connection after loss ASAP
//get list of of all peripherals connected
//check wireless connections
//check communication with HQ/cloud
//check if flight controls are working correctly
//check if data is correctly being logged
//check health of all cells
//log data
//calculate range, lifespan, next service date
//inputs: mission start time, deadline (time), description, distance,
// path[[gcode]], landing bool?, payload info, attempts, failure?complete?,
//return address
//backup return address
//check mission queue
//stop motion, reset attitude
//plan path for landing
//begin path to decend
//check sensors for possibilty of landing
//continue phase 2 of decent
//touch down
//poweroff main motor
//Disarm
//inputs:starting,ending points,
//regions to avoid, regions to pass through
//collion prevention with other UAVs
//path dependent on time?
//output: path coordinates, keep track of timing
//check ping/latency with all comms
//measure signal strength, in proximity UAVs and HQs
//log everything
//check if connection with HQ is okay
//check if sbus output to microcontroller is okay
//in case of signal loss: continue with current path? path plan a new Return to home then land


*/

//protoc --go_out=. *.proto

//this works
//protoc -I ./ protofiles/person.proto --go_out=plugins=grpc:.
//protoc -I ./ protofiles/person.proto --go-grpc_out=plugins=grpc:.
//health checks, load balancing and goroutines, monitor the network, implement client side retries
// circuit breakers, authorization and authentication.

package main

import (
	"fmt"
	"io"
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
	telemetryPort = ":50051"
	controlPort   = ":50052"
	noOfSteps     = 50
)

var (
	uniqueID string

	//priority 0
	batteryLow   bool = true
	safe2Land    bool = false
	returnToHome bool = false
	disarm       bool = true
	mode         int
	homeLat      float32
	homeLon      float32
	homeAlt      float32

	//controls
	throttle   int = 1500 //Channel 0
	rudder     int = 1500
	aileron    int = 1500
	elevator   int = 1500
	motorPower int = 1500
	aux1       int = 1500
	aux2       int = 1500
	aux3       int = 1500 //Channel 7

	//telemetry
	batteryVoltage float32 = 1.0
	currentDraw    float32 = 2.0
	longitude      float32 = 3.0
	latitude       float32 = 4.0
	altitude       float32 = 5.0
	temperature    float32 = 40.3
	motorRPM       float32 = 9000.23
	gyro           float32 = 123.21
	accel          float32 = 432.12

	//mission packet
	missionID  uint
	mLatitude  float32
	mLongitude float32
	mAltitude  float32
	repeat     bool
)

type server struct {
	pb.UnimplementedUavControlServer
}

func (s *server) GetTelemetry(in *pb.Acknowledged, stream pb.UavControl_GetTelemetryServer) error {
	log.Printf("Got request for mor....")
	log.Printf("a: $%s", in.A)
	// Send streams here
	for i := 0; i < noOfSteps; i++ {

		time.Sleep(time.Microsecond * 1)
		if err := stream.Send(&pb.TelemetryPacket{
			BatteryVoltage: batteryVoltage,
			CurrentDraw:    currentDraw,
			Altitude:       altitude,
			Longitude:      longitude,
			Temperature:    temperature,
			MotorRPM:       motorRPM,
			Gyro:           gyro,
			Accel:          accel}); err != nil {

			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}

		droneDummyDataGenerator()
	}

	log.Printf("Successfully transfered amount $%v ", in.A)
	return nil
}

//receive client stream
func ReceiveStream(stream pb.UavControl_SendDroneControlServer) error {

	log.Println("Started listening to the stream!")

	values, err := stream.Recv()
	//stream, err := client.GetTelemetry(context.Background(), request)

	if err == io.EOF {
		// Close the connection and return the response to the client
		return err
	}

	// Listen to the stream of messages
	for {
		fmt.Println(values.Throttle, values.Aileron, values.Rudder, values.Elevator)
	}
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

func startTelemetry() {
	lis, err := net.Listen("tcp", telemetryPort)
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

func controlClient() {

	conn, err := grpc.Dial("localhost"+controlPort, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewUavControlClient(conn)

	ReceiveStream(client, &pb.Acknowledged{A: "Kiti is a randi. Client sent this."})
}

func droneDummyDataGenerator() {

	//init these variables

	rand1 := float32(rand.Float64())
	rand2 := int(rand.Int())
	//timestamp := time.Now()

	//priority 0
	batteryLow = true
	safe2Land = false
	returnToHome = false
	disarm = true
	mode = 10 * rand2
	homeLat = 1501 * rand1
	homeLon = 1505 * rand1
	homeAlt = 1801 * rand1

	//Control params
	throttle = 1500 //Channel 0
	rudder = 1501 * rand2
	aileron = 1501 * rand2
	elevator = 1501 * rand2
	motorPower = 1501 * rand2
	aux1 = 1501 * rand2
	aux2 = 1501 * rand2
	aux3 = 1501 * rand2 //Channel 7

	//telemetry randz
	batteryVoltage = 1500 * rand1
	currentDraw = 1501 * rand1
	longitude = 1502 * rand1
	latitude = 1503 * rand1
	altitude = 1504 * rand1
	temperature = 1505 * rand1
	motorRPM = 1506 * rand1
	gyro = 1507 * rand1
	accel = 1508 * rand1

	//fmt.Println("Battery Voltage", batteryVoltage, currentDraw, homeLat, homeLon, timestamp)

}

func initSys() {
	fmt.Println("System Initialization has started ... ")
	//example mission Travel from point a to b

}

func main() {

	//initSys()
	//testProtoMarshalling()

	startTelemetry()
	//testControl()

}
