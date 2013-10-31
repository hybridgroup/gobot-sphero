package gobotSphero

import (
  "github.com/hybridgroup/gobot"
  "fmt"
)

type packet struct {
  header []uint8
  body []uint8
  checksum uint8
}

type SpheroDriver struct {
  gobot.Driver
  SpheroAdaptor *SpheroAdaptor
  seq uint8
  messages [][]uint8
}

func NewSphero(sa *SpheroAdaptor) *SpheroDriver {
  s := new(SpheroDriver)
  s.Events = make(map[string]chan interface{})
  s.SpheroAdaptor = sa
  return s
}

func (sd *SpheroDriver) StartDriver() {
  sd.ConfigureCollisionDetection()

  gobot.Every(sd.Interval, func() {
    sd.handleMessageEvents()
  })
}

func (sd *SpheroDriver) Roll(speed uint8, heading uint16) {
  packet := new(packet)
  packet.body = []uint8{speed, uint8(heading >> 8), uint8(heading & 0xFF), 0x01}
  dlen := len(packet.body) + 1
  packet.header = []uint8{0xFF, 0xFF, 0x02, 0x30, sd.seq, uint8(dlen)}
  packet.checksum = sd.calculateChecksum(packet)
  sd.write(packet)
}

func (sd *SpheroDriver) Stop() {
  sd.Roll(0, 0)
}

func (sd *SpheroDriver) SetRGB(r uint8, g uint8, b uint8) {
  packet := new(packet)
  packet.body = []uint8{r, g, b, 0x01}
  dlen := len(packet.body) + 1
  packet.header = []uint8{0xFF, 0xFF, 0x02, 0x20, sd.seq, uint8(dlen)}
  packet.checksum = sd.calculateChecksum(packet)
  sd.write(packet)
}

func (sd *SpheroDriver) ConfigureCollisionDetection() {
  packet := new(packet)
  packet.body = []uint8{0x01, 0x40, 0x40, 0x50, 0x50, 0x50}
  dlen := len(packet.body) + 1
  packet.header = []uint8{0xFF, 0xFF, 0x02, 0x12, sd.seq, uint8(dlen)}
  packet.checksum = sd.calculateChecksum(packet)
  sd.Events["Collision"] = make(chan interface{})
  sd.write(packet) 
}

func (sd *SpheroDriver) write(packet *packet) []uint8 {
  var header []uint8
  var body []uint8
  buf := append(packet.header, packet.body...)
  buf = append(buf, packet.checksum)
  length, err := sd.SpheroAdaptor.TcpPort.Write(buf)
  if err != nil {
    fmt.Println(sd.Name, err)
    sd.SpheroAdaptor.Disconnect()
    sd.SpheroAdaptor.Connect()
    return nil
  } else if length != len(buf) {
    fmt.Println("Not enough bytes written", sd.Name)
  }
  sd.seq += 1

  header = sd.readHeader()
  if header != nil {
    body = sd.readBody(header[len(header)-1])
  }

  for header != nil && header[1] == 0xFE {
    async := append(header, body...)
    sd.messages = append(sd.messages, async)

    header := sd.readHeader()
    if header != nil {
      body = sd.readBody(header[len(header)-1])
    } else {
      body = nil
    }
  }

  if len(header) != 0 && header[2] == 0 {
    return append(header, body...)
  } else {
    fmt.Println("Unable to write to Sphero!", sd.Name)
    return nil
  }
}

func (sd *SpheroDriver) calculateChecksum(packet *packet) uint8 {
  buf := append(packet.header, packet.body...)
  buf = buf[2:]
  var calculatedChecksum uint16
  for i := range buf {
    calculatedChecksum += uint16(buf[i])
  }
  return uint8(^(calculatedChecksum % 256))
}

func (sd *SpheroDriver) handleMessageEvents() {
  var evt []uint8
  for len(sd.messages) != 0{
    evt, sd.messages = sd.messages[len(sd.messages)-1], sd.messages[:len(sd.messages)-1]
    if evt[2] == 0x07 {
      sd.handleCollisionDetected(evt)
    }
  } 
}

func (sd *SpheroDriver) handleCollisionDetected(data []uint8) {
  sd.Events["Collision"] <- data
}

func (sd *SpheroDriver) readHeader() []uint8 {
  headerLen := uint8(5)
  data := sd.readNextChunk(headerLen)
  if data == nil || uint8(len(data)) != headerLen {
    return nil
  } else {
    return data
  }
}

func (sd *SpheroDriver) readBody(length uint8) []uint8 {
  data := sd.readNextChunk(length)
  if data == nil || uint8(len(data)) != length {
    return nil
  } else {
    return data
  }
}

func (sd *SpheroDriver) readNextChunk(length uint8) []uint8 {
  var read = make([]uint8, int(length))
  _, err := sd.SpheroAdaptor.TcpPort.Read(read[:])
  if err != nil {
    return nil
  }
  return read
}
