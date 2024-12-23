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

type Device interface {
	Close()
	ReadBytes() []byte
	ReadInput() (uint64, error)
}

type G13Device struct {
	ctx  *gousb.Context
	dev  *gousb.Device
	cfg  *gousb.Config
	intf *gousb.Interface
	IEP  *gousb.InEndpoint
}

// New returns an initialised [G13Device] for a connected G13 gameboard. It
// contains an initialised [gousb.InEndpoint] which is used by
// [G13Device.ReadBytes] and [G13Device.ReadInput] for reading button presses.
func New() (Device, error) {
	ctx := gousb.NewContext()

	d := G13Device{}
	dev, err := ctx.OpenDeviceWithVIDPID(0x046d, 0xc21c)
	if err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to open device: %s", err)
	}
	d.dev = dev

	if dev == nil {
		d.Close()
		return nil, fmt.Errorf("device not found")
	}

	cfg, err := dev.Config(1)
	if err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to initialise config: %s", err)
	}
	d.cfg = cfg

	if err := dev.SetAutoDetach(true); err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to enable automatic kernel driver detachment: %s", err)
	}

	intf, err := cfg.Interface(0, 0)
	if err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to select interface 0: %s", err)
	}
	d.intf = intf

	ep, err := intf.InEndpoint(1)
	if err != nil {
		d.Close()
		return nil, fmt.Errorf("failed to initialise input endpoint: %s", err)
	}

	// Probably unnecessary, but good to be sure
	ep.Desc.TransferType = gousb.TransferTypeInterrupt
	d.IEP = ep

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
	buf := make([]byte, 1*d.IEP.Desc.MaxPacketSize)
	if _, err := d.IEP.Read(buf); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(buf), nil
}

func (d *G13Device) ReadBytes() []byte {
	buf := make([]byte, 1*d.IEP.Desc.MaxPacketSize)
	if _, err := d.IEP.Read(buf); err != nil {
		panic(err)
	}
	return buf
}
