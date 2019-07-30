package main

import (
	"fmt"
	"time"
	"os"

	"github.com/JoshSharpe/rc_car/src/car"
	pi "github.com/stianeikeland/go-rpio"
)

func main() {
	inputPin := pi.Pin(21)
	outputPin := pi.Pin(20)
	if err := pi.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer pi.Close()

	inputPin.Input()
	outputPin.Output()

	sonar := car.NewSonar(outputPin, inputPin)

	for i := 0; i < 10; i++ {
		dist := sonar.GetDistance()
		fmt.Printf("Distance: %f\n", dist)

		time.Sleep(time.Second)
	}
}
