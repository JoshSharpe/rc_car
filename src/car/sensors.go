package car

import (
	"bytes"
	"encoding/binary"
	"log"
	"time"

	"leolaunches.com/assembly/physics"

	"periph.io/x/periph/conn/i2c"

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

	return float64(diff.Nanoseconds() / 1000) / convertToCentimeters
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

/*
	Constants for the imu MPU-6050
*/
const (
	MPU6050_RA_ACCEL_XOUT_H        = 0x3B
	MPU6050_RA_PWR_MGMT_1          = 0x6B
	MPU6050_PWR1_CLKSEL_BIT        = 2
	MPU6050_PWR1_CLKSEL_LENGTH     = 3
	MPU6050_CLOCK_PLL_XGYRO        = 0x01
	MPU6050_GYRO_FS_250            = 0x00
	MPU6050_RA_GYRO_CONFIG         = 0x1B
	MPU6050_GCONFIG_FS_SEL_LENGTH  = 2
	MPU6050_GCONFIG_FS_SEL_BIT     = 4
	MPU6050_RA_ACCEL_CONFIG        = 0x1C
	MPU6050_ACONFIG_AFS_SEL_BIT    = 4
	MPU6050_ACONFIG_AFS_SEL_LENGTH = 2
	MPU6050_ACCEL_FS_2             = 0x00
	MPU6050_PWR1_SLEEP_BIT         = 6
	MPU6050_PWR1_ENABLE_BIT        = 0
)

// IMU Represents the MPU-6050 device
type IMU struct {
	device       *i2c.Dev
	Acceleration *physics.Vector
	Rotation     *physics.Vector
	Temperature  float64
}

// NewIMU create a new imu object that is initialized and ready to read data.
func NewIMU(bus i2c.BusCloser) *IMU {
	d := &i2c.Dev{
		Addr: 0x68,
		Bus:  bus,
	}

	// Set Clock
	_, err := d.Write([]byte{
		MPU6050_RA_PWR_MGMT_1,
		MPU6050_PWR1_CLKSEL_BIT,
		MPU6050_PWR1_CLKSEL_LENGTH,
		MPU6050_CLOCK_PLL_XGYRO,
	})
	if err != nil {
		log.Fatal("A) Unable to create imu. Err: ", err)
	}

	// Set Full Scale Gryo Range
	_, err = d.Write([]byte{
		MPU6050_RA_GYRO_CONFIG,
		MPU6050_GCONFIG_FS_SEL_BIT,
		MPU6050_GCONFIG_FS_SEL_LENGTH,
		MPU6050_GYRO_FS_250,
	})
	if err != nil {
		log.Fatal("B) Unable to create imu. Err: ", err)
	}

	// Set Full Scale Acc Range
	_, err = d.Write([]byte{
		MPU6050_RA_ACCEL_CONFIG,
		MPU6050_ACONFIG_AFS_SEL_BIT,
		MPU6050_ACONFIG_AFS_SEL_LENGTH,
		MPU6050_ACCEL_FS_2,
	})
	if err != nil {
		log.Fatal("C) Unable to create imu. Err: ", err)
	}

	// Set sleep disabled
	_, err = d.Write([]byte{
		MPU6050_RA_PWR_MGMT_1,
		MPU6050_PWR1_ENABLE_BIT,
		0,
	})
	if err != nil {
		log.Fatal("D) Unable to create imu. Err: ", err)
	}

	return &IMU{
		device:       d,
		Acceleration: physics.NewVector(0, 0, 0),
		Rotation:     physics.NewVector(0, 0, 0),
	}
}

// ReadData triggers a read that will populate accel, gyro, and temp data.
func (i *IMU) ReadData() error {
	var ax, ay, az int16
	var rx, ry, rz int16
	var temp int16

	write := []byte{
		MPU6050_RA_ACCEL_XOUT_H,
	}
	read := make([]byte, 14)

	err := i.device.Tx(write, read)

	if err != nil {
		log.Println("Unable to write read command.")
		return err
	}

	buf := bytes.NewBuffer(read)
	err = binary.Read(buf, binary.BigEndian, &ax)
	if err != nil {
		log.Println("Unable to unpackage ax.")
		return err
	}

	err = binary.Read(buf, binary.BigEndian, &ay)
	if err != nil {
		log.Println("Unable to unpackage ay.")
		return err
	}
	err = binary.Read(buf, binary.BigEndian, &az)
	if err != nil {
		log.Println("Unable to unpackage az.")
		return err
	}
	i.Acceleration = physics.NewVector(float64(ax), float64(ay), float64(az))
	// log.Printf("Acceleration: <%d, %d, %d>", ax, ay, az)

	err = binary.Read(buf, binary.BigEndian, &temp)
	if err != nil {
		log.Println("Unable to unpackage Temperature.")
		return err
	}
	i.Temperature = (float64(temp) + 12412) / 340
	// log.Printf("Temperature: %d", i.Temperature)

	err = binary.Read(buf, binary.BigEndian, &rx)
	if err != nil {
		log.Println("Unable to unpackage Rotation.")
	}
	err = binary.Read(buf, binary.BigEndian, &ry)
	if err != nil {
		log.Println("Unable to unpackage Rotation.")
	}
	err = binary.Read(buf, binary.BigEndian, &rz)
	if err != nil {
		log.Println("Unable to unpackage Rotation.")
	}
	i.Rotation = physics.NewVector(float64(rx), float64(ry), float64(rz))
	return err
}
