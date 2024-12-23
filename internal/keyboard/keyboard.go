package keyboard

import (
	"fmt"

	"github.com/bendahl/uinput"
)

type Keyboard interface {
	Close() error
	KeyPress(k int) error
	KeyDown(k int) error
	KeyUp(k int) error
}

type UinputKeyboard struct {
	kb uinput.Keyboard
}

// New returns a [Keyboard] instance with a [uinput.Keyboard] initialised with
// the provided name.
func New(name string) (Keyboard, error) {
	kb, err := uinput.CreateKeyboard("/dev/uinput", []byte(name))
	if err != nil {
		return nil, err
	}
	return &UinputKeyboard{
		kb: kb,
	}, nil
}

func (vkb *UinputKeyboard) Close() error {
	if !vkb.hasKeyboard() {
		// just do nothing
		return nil
	}
	return vkb.kb.Close()
}

func (vkb *UinputKeyboard) KeyPress(k int) error {
	if !vkb.hasKeyboard() {
		return fmt.Errorf("key press before initialising keyboard")
	}
	return vkb.kb.KeyPress(k)
}

func (vkb *UinputKeyboard) KeyDown(k int) error {
	if !vkb.hasKeyboard() {
		return fmt.Errorf("key down before initialising keyboard")
	}
	return vkb.kb.KeyDown(k)
}

func (vkb *UinputKeyboard) KeyUp(k int) error {
	if !vkb.hasKeyboard() {
		return fmt.Errorf("key up before initialising keyboard")
	}
	return vkb.kb.KeyUp(k)
}

func (vkb *UinputKeyboard) hasKeyboard() bool {
	return vkb.kb != nil
}
