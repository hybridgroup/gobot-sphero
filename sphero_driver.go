package gobotSphero

import (
  "github.com/hybridgroup/gobot"
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
}

func NewSphero(sa *SpheroAdaptor) *SpheroDriver {
  s := new(SpheroDriver)
  s.SpheroAdaptor = sa
  return s
}

func (sd *SpheroDriver) StartDriver() {
  f := func(){ 
  }

  gobot.Every(sd.Interval, f)
}

func (sd *SpheroDriver) Roll(speed uint8, heading uint16) {
  packet := new(packet)
  packet.body = []byte{speed, uint8(heading >> 8), uint8(heading & 0xFF), 0x01}
  dlen := len(packet.body) + 1
  packet.header = []byte{0xFF, 0xFF, 0x02, 0x30, sd.seq, uint8(dlen)}
  packet.checksum = sd.calculateChecksum(packet)
  sd.write(packet)
}

func (sd *SpheroDriver) Stop() {
  sd.Roll(0, 0)
}

func (sd *SpheroDriver) SetRGB(r uint8, g uint8, b uint8) {
  packet := new(packet)
  packet.body = []byte{r, g, b, 0x01}
  dlen := len(packet.body) + 1
  packet.header = []byte{0xFF, 0xFF, 0x02, 0x20, sd.seq, uint8(dlen)}
  packet.checksum = sd.calculateChecksum(packet)
  sd.write(packet)
}

func (sd *SpheroDriver) write(packet *packet) {
  buf := append(packet.header, packet.body...)
  buf = append(buf, packet.checksum)
  sd.SpheroAdaptor.TcpPort.Write(buf)
  sd.seq += 1
}

func (ds *SpheroDriver) calculateChecksum(packet *packet) uint8 {
  buf := append(packet.header, packet.body...)
  var calculatedChecksum uint16
  buf = buf[2:]
  for i := range buf {
    calculatedChecksum += uint16(buf[i])
  }
  return uint8(^(calculatedChecksum % 256))
}
