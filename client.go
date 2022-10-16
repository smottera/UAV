// This is the client stub (UAV/Robot backend)
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "./protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:50051"

	controlport = ":50052"
	noOfSteps   = 50
)

type server struct {
	pb.UnimplementedUavControlServer
}

func (s *server) SendDroneControl(stream pb.UavControl_SendDroneControlClient, in *pb.Acknowledged) {
	fmt.Println("send drone control func called ... ")
	log.Printf("Got request for mor....")
	log.Printf("a: $%s", in.A)

	// Send streams here
	for i := 0; i < noOfSteps; i++ {

		time.Sleep(time.Microsecond * 1)

		if err := stream.Send(&pb.ControlPacket{
			Throttle:   1500,
			Rudder:     1501,
			Aileron:    1502,
			Elevator:   1503,
			MotorPower: 1522,
			Aux1:       1533,
			Aux2:       1544,
			Aux3:       1555}); err != nil {

			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}

	}

	log.Printf("Successfully transfered amount $%v ", in.A)
}

// ReceiveStream listens to the stream contents and use them
func ReceiveStream(client pb.UavControlClient, request *pb.Acknowledged) {

	log.Println("Started listening to the server stream!")

	stream, err := client.GetTelemetry(context.Background(), request)

	if err != nil {
		log.Fatalf("%v.GetTelemetry(_) = _, %v", client, err)
	}

	// Listen to the stream of messages
	for {

		response, err := stream.Recv()
		if err == io.EOF {
			// If there are no more messages, get out of loop
			break
		}

		if err != nil {
			log.Fatalf("%v.GetTelemetry(_) = _, %v", client, err)
		}

		fmt.Println("Battery Voltage: %f", "Current Draw: %f", "Longitude: %f", "Latitude: %f",
			"Altitude: %f", "Temperature: %f", "MotorRPM: %f", "Gyro: %f", "Accel: %f",
			response.BatteryVoltage, response.CurrentDraw, response.Longitude, response.Latitude,
			response.Altitude, response.Temperature, response.MotorRPM, response.Gyro, response.Accel)
	}
}

func testControlLink() error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewUavControlClient(conn)

	ReceiveStream(client, &pb.Acknowledged{A: "guccii"})

	//start server
	lis, err := net.Listen("tcp", controlport)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUavControlServer(s, &pb.UnimplementedUavControlServer{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

	return nil
}

func main() {

	testControlLink()

}
