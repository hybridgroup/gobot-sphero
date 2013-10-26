package main

import (
  "github.com/hybridgroup/gobot"
  "gobot-sphero"
  "time"
)

func main() {

  spheroAdaptor := new(gobotSphero.SpheroAdaptor)
  spheroAdaptor.Name = "Sphero"
  spheroAdaptor.Port = "127.0.0.1:4567"

  sphero := gobotSphero.NewSphero(spheroAdaptor)
  sphero.Driver = gobot.Driver{
    Name: "Sphero",
  }

  connections := []interface{} {
    spheroAdaptor,
  }
  devices := []interface{} {
    sphero,
  }

  work := func(){
    sphero.Stop()
    gobot.Every(1000 * time.Millisecond, func(){ sphero.Roll(100,100) })
  }
  
  robot := gobot.Robot{
      Connections: connections, 
      Devices: devices,
      Work: work,
  }

  robot.Start()
}
