package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"strconv"
	"time"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "/dev/tty.usbmodem1411")
	const num = 3
	var servos [num]*gpio.ServoDriver
	for i := 0; i < num; i++ {
		servos[i] = gpio.NewServoDriver(firmataAdaptor, "servo", strconv.Itoa(i+7))
	}

	work := func() {
		gobot.Every(300*time.Millisecond, func() { //right up
			left := uint8(40)
			right := uint8(80)

			err := servos[0].Move(left)
			if err != nil {
				fmt.Println("01:", err)
			}
			time.Sleep(100 * time.Millisecond)
			err = servos[0].Move(right)
			if err != nil {
				fmt.Println("02:", err)
			}
		})

		gobot.After(100*time.Millisecond, func() {
			gobot.Every(300*time.Millisecond, func() { //left up
				left := uint8(80)
				right := uint8(110)

				err := servos[1].Move(right)
				if err != nil {
					fmt.Println("11:", err)
				}
				time.Sleep(100 * time.Millisecond)
				err = servos[1].Move(left)
				if err != nil {
					fmt.Println("12:", err)
				}
			})
		})

		// gobot.Every(3000*time.Millisecond, func() { //left down
		// 	down := uint8(30)
		// 	up := uint8(70)
		// 	err := servos[2].Move(down)
		// 	if err != nil {
		// 		fmt.Println("21:", err)
		// 	}
		// 	time.Sleep(1000 * time.Millisecond)
		// 	err = servos[2].Move(up)
		// 	if err != nil {
		// 		fmt.Println("22:", err)
		// 	}
		// })
	}

	robot := gobot.NewRobot("servoBot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{servos[0], servos[1], servos[2]},
		work,
	)

	gbot.AddRobot(robot)
	gbot.Start()
}
