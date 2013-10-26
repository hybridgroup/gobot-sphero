package main
import (
  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot-sphero"
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

    gobot.Every("1s", func(){ 
      sphero.Roll(100, uint16(gobot.Random(0, 360))) 
    })

    gobot.Every("0.5s", func(){ 
      sphero.SetRGB(uint8(gobot.Random(0, 255)), uint8(gobot.Random(0, 255)), uint8(gobot.Random(0, 255))) 
    })

  }
  
  robot := gobot.Robot{
      Connections: connections, 
      Devices: devices,
      Work: work,
  }

  robot.Start()
}
