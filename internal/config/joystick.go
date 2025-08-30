package config

type StickPosition struct {
	posX uint8
	posY uint8
}

func (sp *StickPosition) Position() (uint8, uint8) {
	return sp.posX, sp.posY
}

func (sp *StickPosition) X() uint8 {
	return sp.posX
}

func (sp *StickPosition) Y() uint8 {
	return sp.posY
}

func (sp *StickPosition) UinputPosition() (float32, float32) {
	return sp.UinputX(), sp.UinputY()
}

func (sp *StickPosition) UinputX() float32 {
	return (float32(sp.posX) - float32(127)) / float32(127)
}

func (sp *StickPosition) UinputY() float32 {
	return (float32(sp.posY) - float32(127)) / float32(127)
}
