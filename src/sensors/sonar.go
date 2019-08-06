package sensors

import "time"

const (
	convertToCentimeters = 58.0
	convertToInches      = 148.0
)

type sonar struct {
	signalPin pi.Pin
	echoPin   pi.Pin
}

func NewSonar(signalPin, echoPin pi.Pin) *sonar {
	return &sonar{
		signalPin: signalPin,
		echoPin:   echoPin,
	}
}

func (s *sonar) GetDistance() float64 {
	s.signalPin.High()
	time.Sleep(time.Microsecond * 10)
	s.signalPin.Low()

	for s.echoPin.Read() == pi.Low {
	}
	initTime := time.Now()
	for s.echoPin.Read() == pi.High {
	}
	diff := time.Now().Sub(initTime)

	return float64(diff.Nanoseconds()/1000) / convertToCentimeters
}
