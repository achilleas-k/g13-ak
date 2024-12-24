package main

import (
	"fmt"
	"os"

	"github.com/achilleas-k/g13-ak/internal/device"
	"github.com/achilleas-k/g13-ak/internal/keyboard"
	"github.com/achilleas-k/g13-ak/internal/mapping"
)

func main() {
	configPath := os.Args[1]

	dev, err := device.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initialising device: %s\n", err)
		os.Exit(1)
	}
	defer dev.Close()

	vkb, err := keyboard.New("g13-vkb")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error initialising virtual keyboard: %s\n", err)
		os.Exit(1)
	}

	keyMap, err := mapping.NewFromFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
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
