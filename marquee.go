package main

import (
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/firmata"
	"github.com/hybridgroup/gobot/platforms/gpio"
	"strconv"
)

func main() {
	gbot := gobot.NewGobot()

	firmataAdaptor := firmata.NewFirmataAdaptor("firmata", "/dev/tty.usbmodem1411")
	const num = 7
	var leds [num]*gpio.LedDriver
	for i := 1; i <= num; i++ {
		leds[i-1] = gpio.NewLedDriver(firmataAdaptor, "led", strconv.Itoa(i+1))
	}

	work := func() {
		gobot.Every(400*time.Millisecond, func() {
			leds[0].Toggle()
		})

		gobot.After(100*time.Millisecond, func() {
			gobot.Every(400*time.Millisecond, func() {
				leds[1].Toggle()
			})
		})
		gobot.After(200*time.Millisecond, func() {
			gobot.Every(400*time.Millisecond, func() {
				leds[2].Toggle()
			})
		})
		gobot.After(300*time.Millisecond, func() {
			gobot.Every(400*time.Millisecond, func() {
				leds[3].Toggle()
			})
		})
		gobot.After(400*time.Millisecond, func() {
			gobot.Every(400*time.Millisecond, func() {
				leds[4].Toggle()
			})
		})
		gobot.After(500*time.Millisecond, func() {
			gobot.Every(400*time.Millisecond, func() {
				leds[5].Toggle()
			})
		})
		gobot.After(600*time.Millisecond, func() {
			gobot.Every(400*time.Millisecond, func() {
				leds[6].Toggle()
			})
		})

	}

	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{leds[0], leds[1], leds[2], leds[3], leds[4], leds[5], leds[6]},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
