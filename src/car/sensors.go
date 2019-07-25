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
	time.Sleep(time.Microsecond * 10)
	s.signalPin.Low()

	initTime := time.Now()

	for s.echoPin.Read() == pi.Low {
	}
	diff := time.Now().Sub(initTime)

	return float64(diff.Nanoseconds()) / convertToCentimeters
}

type led struct {
	pinNumber pi.Pin
	isOn      bool
}

func NewLED(p pi.Pin) *led {
	return &led{
		pinNumber: p,
		isOn:      false,
	}
}

func (l *led) Toggle() {
	l.pinNumber.Toggle()
	l.isOn = !l.isOn
}
