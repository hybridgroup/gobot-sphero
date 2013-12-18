package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"
)

func main() {

	spheroAdaptor := new(gobotSphero.SpheroAdaptor)
	spheroAdaptor.Name = "Sphero"
	spheroAdaptor.Port = "127.0.0.1:4560"

	sphero := gobotSphero.NewSphero(spheroAdaptor)
	sphero.Name = "sphero"

	work := func() {

		sphero.Stop()

		gobot.On(sphero.Events["Collision"], func(data interface{}) {
			fmt.Println("Collision Detected!")
		})

		gobot.Every("2s", func() {
			dir := uint16(gobot.Random(0, 360))
			sphero.Roll(100, dir)
		})

		gobot.Every("3s", func() {
			r := uint8(gobot.Random(0, 255))
			g := uint8(gobot.Random(0, 255))
			b := uint8(gobot.Random(0, 255))
			sphero.SetRGB(r, g, b)
		})
	}

	robot := gobot.Robot{
		Connections: []interface{}{spheroAdaptor},
		Devices:     []interface{}{sphero},
		Work:        work,
	}

	robot.Start()
}
