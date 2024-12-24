package device

import (
	"fmt"

	"github.com/google/gousb"
)

const (
	// TODO: document source of these values
	ControlRequestType = uint8(gousb.ControlClass | gousb.ControlInterface)

	BacklightColourVal = uint16(0x307)

	SetupPacketRequest = uint8(9)

	SetupPacketIndex = uint16(0)
)

func (d *G13Device) SetBacklightColour(r, g, b uint8) error {
	// TODO: set context with timeout
	data := []byte{5, r, g, b, 0}
	n, err := d.dev.Control(ControlRequestType, SetupPacketRequest, BacklightColourVal, SetupPacketIndex, data)
	if err != nil {
		return fmt.Errorf("failed setting backlight colour %+v: %w", data, err)
	}
	if n != len(data) {
		return fmt.Errorf("sent %d bytes but wrote %d while setting backlight colour", len(data), n)
	}
	return nil
}

func (d *G13Device) ResetBacklightColour() error {
	return d.SetBacklightColour(uint8(0), uint8(0), uint8(0))
}
