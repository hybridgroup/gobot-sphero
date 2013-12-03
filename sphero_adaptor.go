package gobotSphero

import (
	"github.com/hybridgroup/gobot"
)

type SpheroAdaptor struct {
	gobot.Adaptor
	sp gobot.Port
}

func (sa *SpheroAdaptor) Connect() {
	if gobot.IsUrl(sa.Adaptor.Port) {
		sa.sp = gobot.ConnectToTcp(sa.Adaptor.Port)
	} else {
		sa.sp = gobot.ConnectToSerial(sa.Adaptor.Port, 115200)
	}
}

func (sa *SpheroAdaptor) Disconnect() {
	sa.sp.Close()
}
