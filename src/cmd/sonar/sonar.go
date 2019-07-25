package main

import (
	"fmt"
	"time"

	"github.com/JoshSharpe/rc_car/src/car"
)

func main() {
	sonar := car.NewSonar(20, 21)

	for i := 0; i < 1000; i++ {
		dist := sonar.GetDistance()
		fmt.Printf("Distance: %f\n", dist)

		time.Sleep(time.Millisecond * 10)
	}

	car.ShutDown()

}
