package gobotSphero

func (sd *SpheroDriver) SetRGBApi(params map[string]interface{}) {
	r := uint8(params["r"].(float64))
	g := uint8(params["g"].(float64))
	b := uint8(params["b"].(float64))
	sd.SetRGB(r, g, b)
}

func (sd *SpheroDriver) RollApi(params map[string]interface{}) {
	speed := uint8(params["speed"].(float64))
	heading := uint16(params["heading"].(float64))
	sd.Roll(speed, heading)
}

func (sd *SpheroDriver) StopApi() {
	sd.Stop()
}
