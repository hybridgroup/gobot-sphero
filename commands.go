package gobotSphero

func (sd *SpheroDriver) SetRGBApi(params map[string]interface{}) {
	r := uint8(params["r"].(float64))
	g := uint8(params["g"].(float64))
	b := uint8(params["b"].(float64))
	sd.SetRGB(r, g, b)
}
