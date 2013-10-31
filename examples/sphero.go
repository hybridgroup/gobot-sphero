package main
import (
  "github.com/hybridgroup/gobot"
  "github.com/hybridgroup/gobot-sphero"
  "fmt"
)

func main() {

  spheroAdaptor := new(gobotSphero.SpheroAdaptor)
  spheroAdaptor.Name = "Sphero"
  spheroAdaptor.Port = "127.0.0.1:4560"

  sphero := gobotSphero.NewSphero(spheroAdaptor)
  sphero.Name = "Sphero"

  connections := []interface{} {
    spheroAdaptor,
  }
  devices := []interface{} {
    sphero,
  }

  work := func(){

    sphero.Stop()

    go func() {
      for{
        gobot.On(sphero.Events["Collision"])
        fmt.Println("Collision Detected!")
      }
    }()    
    
    gobot.Every("2s", func(){ 
      dir := uint16(gobot.Random(0, 360))
      sphero.Roll(100, dir) 
    })

    gobot.Every("3s", func(){ 
      r := uint8(gobot.Random(0, 255))
      g := uint8(gobot.Random(0, 255))
      b := uint8(gobot.Random(0, 255))
      sphero.SetRGB(r, g, b)
    })
  }
  
  robot := gobot.Robot{
      Connections: connections, 
      Devices: devices,
      Work: work,
  }

  robot.Start()
}
