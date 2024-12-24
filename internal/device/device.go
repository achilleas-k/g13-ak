// Package device provides the [Device] interface for interacting with the G13
// gameboard.
//
// Decoding logic adapted from https://github.com/khampf/g13 which was
// originally forked and adapted from https://github.com/ecraven/g13/.
package device

import (
	"encoding/binary"
	"fmt"

	"github.com/google/gousb"
)

const (
	g13VendorID  = 0x046d
	g13ProductID = 0xc21c
)

type Device interface {
	Close()
	ReadBytes() ([]byte, error)
	ReadInput() (uint64, error)
	SetBacklightColour(r, g, b uint8) error
}

type G13Device struct {
	ctx  *gousb.Context
	dev  *gousb.Device
	cfg  *gousb.Config
	intf *gousb.Interface
	iep  *gousb.InEndpoint
}

// New returns an initialised [G13Device] for a connected G13 gameboard. It
// contains an initialised [gousb.InEndpoint] which is used by
// [G13Device.ReadBytes] and [G13Device.ReadInput] for reading button presses.
func New() (Device, error) {
	ctx := gousb.NewContext()

	d := G13Device{}
	dev, err := ctx.OpenDeviceWithVIDPID(g13VendorID, g13ProductID)
	if err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to open device: %w", err)
	}
	d.dev = dev

	if dev == nil {
		d.Close()
		return nil, fmt.Errorf("device not found")
	}

	cfg, err := dev.Config(1)
	if err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to initialise config: %w", err)
	}
	d.cfg = cfg

	if err := dev.SetAutoDetach(true); err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to enable automatic kernel driver detachment: %w", err)
	}

	intf, err := cfg.Interface(0, 0)
	if err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to select interface 0: %w", err)
	}
	d.intf = intf

	ep, err := intf.InEndpoint(1)
	if err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to initialise input endpoint: %w", err)
	}

	// Probably unnecessary, but good to be sure
	ep.Desc.TransferType = gousb.TransferTypeInterrupt
	d.iep = ep

	return &d, nil
}

func (d *G13Device) Close() {
	if d.ctx != nil {
		defer d.ctx.Close()
	}
	if d.dev != nil {
		defer d.dev.Close()
	}
	if d.cfg != nil {
		defer d.cfg.Close()
	}
	if d.intf != nil {
		defer d.intf.Close()
	}
}

func (d *G13Device) ReadInput() (uint64, error) {
	buf, err := d.ReadBytes()
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(buf), nil
}

func (d *G13Device) ReadBytes() ([]byte, error) {
	buf := make([]byte, 1*d.iep.Desc.MaxPacketSize)
	if _, err := d.iep.Read(buf); err != nil {
		return nil, fmt.Errorf("failed reading from device: %w", err)
	}
	return buf, nil
}
