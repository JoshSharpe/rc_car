package main

import (
	"time"
	"fmt"
	"os"

	"github.com/JoshSharpe/rc_car/src/car"
	pi "github.com/stianeikeland/go-rpio"
)

func main() {
	ledPin := pi.Pin(17)
	if err := pi.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer pi.Close()

	ledPin.Output()

	led := car.NewLED(ledPin)

	for i := 0; i < 10; i++ {
		led.Toggle()
		time.Sleep(time.Second)
	}


	// pi.Close()
}
