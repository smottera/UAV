// /This is the client stub (UAV/Robot backend)
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "./protofiles"
)

const (
	telemetryPort = "localhost:50068"
	noOfSteps     = 50
)

var (
	//controls
	throttle   int32 = 1500 //Channel 0
	rudder     int32 = 1501
	aileron    int32 = 1502
	elevator   int32 = 1503
	motorPower int32 = 1504
	aux1       int32 = 1505
	aux2       int32 = 1506
	aux3       int32 = 1507 //Channel 7
)

func runBidirectional() error {

	conn, err := grpc.Dial(telemetryPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewUavControlClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()
	stream, err := client.BidComm(ctx)
	if err != nil {
		return err
	}

	// Send initial data to server
	err = stream.Send(&pb.ToDrone{
		Throttle:   throttle,
		Rudder:     rudder,
		Aileron:    aileron,
		Elevator:   elevator,
		MotorPower: motorPower,
		Aux1:       aux1,
		Aux2:       aux2,
		Aux3:       aux3,
	})
	if err != nil {
		return err
	}

	// Receive processed data from server
	for i := 0; i < noOfSteps; i++ {
		processedData, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		fmt.Println("Received processed data: \n", processedData)
	}

	return nil
}

func main() {
	fmt.Println("TOP")
	flag.Parse()
	fmt.Println(runBidirectional())
	fmt.Println("bottom")

}
