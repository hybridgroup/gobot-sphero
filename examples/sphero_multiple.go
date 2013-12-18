package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-sphero"
)

func main() {

	var robots []gobot.Robot
	spheros := []string{
		"127.0.0.1:4560",
		"127.0.0.1:4561",
		"127.0.0.1:4562",
		"127.0.0.1:4563",
	}

	for s := range spheros {
		spheroAdaptor := new(gobotSphero.SpheroAdaptor)
		spheroAdaptor.Name = "Sphero"
		spheroAdaptor.Port = spheros[s]

		sphero := gobotSphero.NewSphero(spheroAdaptor)
		sphero.Name = "Sphero" + spheros[s]
		sphero.Interval = "0.5s"

		work := func() {
			sphero.Stop()

			gobot.On(sphero.Events["Collision"], func(data interface{}) {
				fmt.Println("Collision Detected!")
			})

			gobot.Every("1s", func() {
				sphero.Roll(100, uint16(gobot.Random(0, 360)))
			})
			gobot.Every("3s", func() {
				sphero.SetRGB(uint8(gobot.Random(0, 255)), uint8(gobot.Random(0, 255)), uint8(gobot.Random(0, 255)))
			})
		}

		robots = append(robots, gobot.Robot{Connections: []interface{}{spheroAdaptor}, Devices: []interface{}{sphero}, Work: work})
	}

	gobot.Work(robots)
}
