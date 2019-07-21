package car

import (
	"time"

	pi "github.com/stianeikeland/go-rpio"
)

const (
	convertToCentimeters = 58.0
	convertToInches      = 148.0
)

type sonar struct {
	signalPin pi.Pin
	echoPin   pi.Pin
}

func NewSonar(signalPin, echoPin int) *sonar {
	outputPin := pi.Pin(signalPin)
	outputPin.Output()

	inputPin := pi.Pin(echoPin)
	inputPin.Input()

	return &sonar{
		signalPin: outputPin,
		echoPin:   inputPin,
	}
}

func (s *sonar) GetDistance() float64 {
	s.signalPin.High()
	s.signalPin.Low()
	initTime := time.Now()

	for s.echoPin.Read() == pi.Low {
	}
	diff := time.Now().Sub(initTime)

	return float64(diff.Nanoseconds()) / convertToCentimeters
}
