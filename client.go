package main

import (
	"io"
	"log"

	pb "./protofiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

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

		log.Printf("Stus: %f, Open: %f", response.BatteryVoltage, response.CurrentDraw)
	}
}

func main() {

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewUavControlClient(conn)

	// Prepare data. Get this from clients like Front-end or Android App

	// Contact the server and print out its response.
	ReceiveStream(client, &pb.Acknowledged{A: "guccii"})
}
