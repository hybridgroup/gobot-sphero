package main

import (
  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot-sphero"
  "fmt"
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
    sphero.Name = "Sphero"
    sphero.Interval = "0.1s"

    work := func(){
      sphero.Stop()

      go func() {
        for{
          gobot.On(sphero.Events["Collision"])
          fmt.Println("Collision Detected!")
        }
      }()    

      gobot.Every("1s", func() { 
        sphero.Roll(100, uint16(gobot.Random(0, 360))) 
      })

      gobot.Every("0.5s", func() { 
        sphero.SetRGB(uint8(gobot.Random(0, 255)), uint8(gobot.Random(0, 255)), uint8(gobot.Random(0, 255))) 
      })
    }

    robots = append(robots, gobot.Robot{Connections: []interface{} {spheroAdaptor}, Devices: []interface{} {sphero}, Work: work,})
  }

  gobot.Work(robots)
}
