package car

import (
	"context"
	"log"
	"net"

	api "github.com/JoshSharpe/rc_car/api/v1"
	pi "github.com/stianeikeland/go-rpio"
	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1"
	port    = "8350"

	direction1APinNumber = 50
	direction1BPinNumber = 51
	direction2APinNumber = 52
	direction2BPinNumber = 53
	speed1PinNumber      = 3
	speed2PinNumber      = 4

	direction1Pin = pi.Pin(direction1APinNumber)
	direction2Pin = pi.Pin(direction2APinNumber)
	speed1Pin     = pi.Pin(speed1PinNumber)
	speed2Pin     = pi.Pin(speed2PinNumber)
)

// server is used to implement car.MoveForward.
type car struct {
	server *grpc.Server
}

func init() {
	pi.Open()
	direction1Pin.Output()
	direction1Pin.Output()
	speed1Pin.Pwm()
	speed2Pin.Pwm()
}

func ShutDown() {
	pi.Close()
}

func startCar() *car {
	c := &car{}

	fullAddress := address + ":" + port

	lis, err := net.Listen("tcp", fullAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	api.RegisterCarServer(s, c)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	c.server = s

	return c
}

func (c *car) Move(ctx context.Context, fv *api.MoveVector) (*api.CurrentData, error) {
	direction1Pin.High()
	direction2Pin.High()
	speed1Pin.High()
	speed2Pin.High()

	return &api.CurrentData{
		Location: &api.Vector{},
	}, nil
}

func (c *car) GetSensorData(ctx context.Context, fv *api.SensorParameters) (*api.CurrentData, error) {
	return &api.CurrentData{
		Location: &api.Vector{},
	}, nil
}

func (c *car) ShutDown() {
	c.server.GracefulStop()
}
