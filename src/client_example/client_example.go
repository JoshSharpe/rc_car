package main

import (
	"context"
	"log"
	"time"

	pb "github.com/JoshSharpe/rc_car/src/protobuff"
	"google.golang.org/grpc"
)

const (
	address = "localhost"
	port    = "8350"
)

func main() {
	fullAddress := address + ":" + port
	// Set up a connection to the server.
	conn, err := grpc.Dial(fullAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewMovementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.MoveForward(ctx, &pb.ForwardVector{})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("New location: <%f, %f, %f>", r.GetLocation().GetX(), r.GetLocation().GetY(), r.GetLocation().GetZ())
}
