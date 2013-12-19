package gobotSphero

import (
	"github.com/hybridgroup/gobot"
)

type SpheroAdaptor struct {
	gobot.Adaptor
	sp gobot.Port
}

func (sa *SpheroAdaptor) Connect() bool {
	if gobot.IsUrl(sa.Adaptor.Port) {
		sa.sp = gobot.ConnectToTcp(sa.Adaptor.Port)
	} else {
		sa.sp = gobot.ConnectToSerial(sa.Adaptor.Port, 115200)
	}
	sa.Connected = true
	return true
}

func (sa *SpheroAdaptor) Reconnect() bool {
	if sa.Connected == true {
		sa.Disconnect()
	}
	return sa.Connect()
}

func (sa *SpheroAdaptor) Disconnect() bool {
	sa.sp.Close()
	sa.Connected = false
	return true
}
func (sa *SpheroAdaptor) Finalize() bool { return true }
