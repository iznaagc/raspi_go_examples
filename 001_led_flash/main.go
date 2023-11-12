package main

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func main()  {
	fmt.Println("start!")
	raspberry := raspi.NewAdaptor()
	led := gpio.NewLedDriver(raspberry, "11")

	work := func() {
		gobot.Every(2*time.Second, func() {
			led.Toggle()
		})
	}


        robot := gobot.NewRobot("blinkLED",
		[]gobot.Connection{raspberry},
		[]gobot.Device{led},
		work,
	)

	robot.Start()
}
