package gobotSphero

import (
  "github.com/hybridgroup/gobot"
  "net"
)

type SpheroAdaptor struct {
  gobot.Adaptor
  TcpPort net.Conn
}

func (sa *SpheroAdaptor) Connect() {
  sa.TcpPort, _ = net.Dial("tcp", sa.Adaptor.Port)
}

func (sa *SpheroAdaptor) Disconnect() {
}