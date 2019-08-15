package main

import (
	"fmt"
	"log"
	"time"

	"github.com/JoshSharpe/rc_car/src/car"
	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/host"
)

func main() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal("Unable to create imu. Err: ", err)
	}
	defer b.Close()

	imu := car.NewIMU(b)

	for i := 0; i < 2; i++ {
		err := imu.ReadData()
		if err != nil {
			log.Fatal("Unable to read imu. Err: ", err)
		}

		fmt.Printf("Acceleration: %f\n", imu.Acceleration)
		fmt.Printf("Rotation: %f\n", imu.Rotation)
		fmt.Printf("Temp: %f\n", imu.Temperature)

		time.Sleep(time.Second * 1)
	}
}
