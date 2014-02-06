package gobotSphero

import (
	"fmt"
	"github.com/hybridgroup/gobot"
)

type packet struct {
	header   []uint8
	body     []uint8
	checksum uint8
}

type SpheroDriver struct {
	gobot.Driver
	SpheroAdaptor *SpheroAdaptor
	seq           uint8
	messages      [][]uint8
}

func NewSphero(sa *SpheroAdaptor) *SpheroDriver {
	s := new(SpheroDriver)
	s.Events = make(map[string]chan interface{})
	s.SpheroAdaptor = sa
	s.Commands = []string{
		"SetRGBC",
		"RollC",
		"StopC",
		"GetRGBC",
		"SetBackLEDC",
		"SetHeadingC",
		"SetStabilizationC",
	}
	return s
}

func (sd *SpheroDriver) Start() bool {
	sd.ConfigureCollisionDetection()

	gobot.Every(sd.Interval, func() {
		sd.handleMessageEvents()
	})
	return true
}

func (sd *SpheroDriver) handleMessageEvents() {
	var evt []uint8
	for len(sd.messages) != 0 {
		evt, sd.messages = sd.messages[len(sd.messages)-1], sd.messages[:len(sd.messages)-1]
		if evt[2] == 0x07 {
			sd.handleCollisionDetected(evt)
		}
	}
}

func (sd *SpheroDriver) handleCollisionDetected(data []uint8) {
	gobot.Publish(sd.Events["Collision"], data)
}

func (sd *SpheroDriver) GetRGB() []uint8 {
	return sd.write(sd.craftPacket([]uint8{}, 0x22))
}

func (sd *SpheroDriver) SetBackLED(level uint8) {
	sd.write(sd.craftPacket([]uint8{level}, 0x21))
}

func (sd *SpheroDriver) SetHeading(heading uint16) {
	sd.write(sd.craftPacket([]uint8{uint8(heading >> 8), uint8(heading & 0xFF)}, 0x01))
}

func (sd *SpheroDriver) SetStabilization(on bool) {
	b := uint8(0x01)
	if on == false {
		b = 0x00
	}
	sd.write(sd.craftPacket([]uint8{b}, 0x02))
}

func (sd *SpheroDriver) Roll(speed uint8, heading uint16) {
	sd.write(sd.craftPacket([]uint8{speed, uint8(heading >> 8), uint8(heading & 0xFF), 0x01}, 0x30))
}

func (sd *SpheroDriver) Stop() {
	sd.Roll(0, 0)
}

func (sd *SpheroDriver) SetRGB(r uint8, g uint8, b uint8) {
	sd.write(sd.craftPacket([]uint8{r, g, b, 0x01}, 0x20))
}

func (sd *SpheroDriver) ConfigureCollisionDetection() {
	sd.Events["Collision"] = make(chan interface{})
	sd.write(sd.craftPacket([]uint8{0x01, 0x40, 0x40, 0x50, 0x50, 0x60}, 0x12))
}

func (sd *SpheroDriver) craftPacket(body []uint8, cid byte) *packet {
	packet := new(packet)
	packet.body = body
	dlen := len(packet.body) + 1
	packet.header = []uint8{0xFF, 0xFF, 0x02, cid, sd.seq, uint8(dlen)}
	packet.checksum = sd.calculateChecksum(packet)
	return packet
}

func (sd *SpheroDriver) write(packet *packet) []uint8 {
	var header []uint8
	var body []uint8
	buf := append(packet.header, packet.body...)
	buf = append(buf, packet.checksum)
	length, err := sd.SpheroAdaptor.sp.Write(buf)
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
		body = sd.readBody(header[4])
	}

	for header != nil && header[1] == 0xFE {
		async := append(header, body...)
		sd.messages = append(sd.messages, async)

		header := sd.readHeader()
		if header != nil {
			body = sd.readBody(header[4])
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

func (sd *SpheroDriver) readHeader() []uint8 {
	data := sd.readNextChunk(5)
	if data == nil {
		return nil
	} else {
		return data
	}
}

func (sd *SpheroDriver) readBody(length uint8) []uint8 {
	data := sd.readNextChunk(length)
	if data == nil {
		return nil
	} else {
		return data
	}
}

func (sd *SpheroDriver) readNextChunk(length uint8) []uint8 {
	var read = make([]uint8, int(length))
	l, err := sd.SpheroAdaptor.sp.Read(read[:])
	if err != nil || length != uint8(l) {
		return nil
	} else {
		return read
	}
}
