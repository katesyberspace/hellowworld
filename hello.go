package main

import (
        "time"

        "gobot.io/x/gobot"
        "gobot.io/x/gobot/drivers/gpio"
        "gobot.io/x/gobot/platforms/firmata"
)

func main() {
        firmataAdaptor := firmata.NewAdaptor("/dev/ttyACM0")
        lcd := i2c.NewJHD1313M1Driver(adaptor)
		
		work := func(h *JHD1313M1Driver) {
                Write("hello world")
        }

        robot := gobot.NewRobot("bot",
                []gobot.Connection{firmataAdaptor},
                []gobot.Device{lcd},
                work,
        )

        robot.Start()
}