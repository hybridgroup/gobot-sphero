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
	sa.TcpPort = gobot.ConnectTo(sa.Adaptor.Port)
}

func (sa *SpheroAdaptor) Disconnect() {
	sa.TcpPort.Close()
}
