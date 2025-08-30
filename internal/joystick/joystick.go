package joystick

import (
	"fmt"

	"github.com/bendahl/uinput"
)

type Joystick interface {
	Close() error
	ButtonPress(b int) error
	ButtonDown(b int) error
	ButtonUp(b int) error
	StickPosition(x, y float32) error
}

type UinputJoystick struct {
	js uinput.Gamepad
}

func New(name string) (Joystick, error) {
	js, err := uinput.CreateGamepad("/dev/uinput", []byte(name), 12, 12)
	if err != nil {
		return nil, err
	}
	return &UinputJoystick{
		js: js,
	}, nil
}

func (vjs *UinputJoystick) Close() error {
	if !vjs.hasJoystick() {
		// just do nothing
		return nil
	}
	return vjs.js.Close()
}

func (vjs *UinputJoystick) ButtonPress(k int) error {
	if !vjs.hasJoystick() {
		return fmt.Errorf("button press before initialising joystick")
	}
	return vjs.js.ButtonPress(k)
}

func (vjs *UinputJoystick) ButtonDown(k int) error {
	if !vjs.hasJoystick() {
		return fmt.Errorf("button down before initialising joystick")
	}
	return vjs.js.ButtonDown(k)
}

func (vjs *UinputJoystick) ButtonUp(k int) error {
	if !vjs.hasJoystick() {
		return fmt.Errorf("button up before initialising joystick")
	}
	return vjs.js.ButtonUp(k)
}

func (vjs *UinputJoystick) StickPosition(x, y float32) error {
	if !vjs.hasJoystick() {
		return fmt.Errorf("stick position set before initialising joystick")
	}
	return vjs.js.LeftStickMove(x, y)
}

func (vjs *UinputJoystick) hasJoystick() bool {
	return vjs.js != nil
}
