package vkeyboard

import (
	"fmt"

	"github.com/bendahl/uinput"
)

type VKeyboard struct {
	kb uinput.Keyboard
}

// New returns a [VKeyboard] instance with a [uinput.Keyboard] initialised with
// the provided name.
func New(name string) (*VKeyboard, error) {
	kb, err := uinput.CreateKeyboard("/dev/uinput", []byte(name))
	if err != nil {
		return nil, err
	}
	return &VKeyboard{
		kb: kb,
	}, nil
}

func (vkb *VKeyboard) Close() error {
	if !vkb.hasKeyboard() {
		// just do nothing
		return nil
	}
	return vkb.kb.Close()
}

func (vkb *VKeyboard) KeyPress(k int) error {
	if !vkb.hasKeyboard() {
		return fmt.Errorf("key press before initialising keyboard")
	}
	return vkb.kb.KeyPress(k)
}

func (vkb *VKeyboard) KeyDown(k int) error {
	if !vkb.hasKeyboard() {
		return fmt.Errorf("key down before initialising keyboard")
	}
	return vkb.kb.KeyDown(k)
}

func (vkb *VKeyboard) KeyUp(k int) error {
	if !vkb.hasKeyboard() {
		return fmt.Errorf("key up before initialising keyboard")
	}
	return vkb.kb.KeyUp(k)
}

func (vkb *VKeyboard) hasKeyboard() bool {
	return vkb.kb != nil
}
