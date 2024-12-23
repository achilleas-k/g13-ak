package main

import (
	"fmt"
	"os"

	"github.com/achilleas-k/g13-ak/internal/device"
	"github.com/achilleas-k/g13-ak/internal/keyboard"
	"github.com/achilleas-k/g13-ak/internal/mapping"
	"github.com/bendahl/uinput"
)

var myDefaultMapping = map[device.KeyBit]int{
	device.G1: uinput.Key1,
	device.G2: uinput.Key2,
	device.G3: uinput.KeyQ,
	device.G4: uinput.KeyW,
	device.G5: uinput.KeyE,
	device.G6: uinput.KeyR,
	device.G7: uinput.KeyT,

	device.G8:  uinput.Key3,
	device.G9:  uinput.Key4,
	device.G10: uinput.KeyA,
	device.G11: uinput.KeyS,
	device.G12: uinput.KeyD,
	device.G13: uinput.KeyF,
	device.G14: uinput.KeyG,

	device.G15: uinput.KeyLeftshift,
	device.G16: uinput.KeyZ,
	device.G17: uinput.KeyX,
	device.G18: uinput.KeyC,
	device.G19: uinput.KeyV,
	device.G20: uinput.KeyLeftctrl,
	device.G21: uinput.KeyTab,
	device.G22: uinput.KeyLeftalt,

	device.LEFT: uinput.KeySpace,
	device.DOWN: uinput.KeyEsc,
}

func main() {
	dev, err := device.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initialising device: %s\n", err)
		os.Exit(1)
	}
	defer dev.Close()

	vkb, err := keyboard.New("g13-vkb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initialising virtual keyboard: %s\n", err)
	}

	keyMap := mapping.New()
	keyMap.SetKeys(myDefaultMapping)
	fmt.Println("Ready")
	for {
		input, err := dev.ReadInput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "e: %s\n", err)
		}
		for kbkey, isDown := range keyMap.GetKeyStates(input) {
			if isDown {
				if err := vkb.KeyDown(kbkey); err != nil {
					fmt.Fprintf(os.Stderr, "keyboard error pressing %d: %s\n", kbkey, err)
				}
			} else if err := vkb.KeyUp(kbkey); err != nil {
				fmt.Fprintf(os.Stderr, "keyboard error releasing %d: %s\n", kbkey, err)
			}
		}
	}
}
