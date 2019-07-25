package main

import (
	"time"

	"github.com/JoshSharpe/rc_car/src/car"
	pi "github.com/stianeikeland/go-rpio"
)

func main() {
	ledPin := pi.Pin(18)
	ledPin.Output()

	led := car.NewLED(18)

	for i := 0; i < 10; i++ {
		led.Toggle()
		time.Sleep(time.Second)
	}


	pi.Close()
}
