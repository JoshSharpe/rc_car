package main

import (
	"time"

	"github.com/JoshSharpe/rc_car/src/car"
)

func main() {
	led := car.NewLED(18)

	for i := 0; i < 10; i++ {
		led.Toggle()
		time.Sleep(time.Second)
	}

	car.ShutDown()

}
