package main

import (
	"fmt"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "/dev/tty.usbmodem1411")
	servo := gpio.NewServoDriver(firmataAdaptor, "servo", "7")

	work := func() {
		gobot.Every(500*time.Millisecond, func() {
			i := uint8(gobot.Rand(100) + 50)
			fmt.Println("Turning", i)
			servo.Move(i)
		})
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{servo},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
