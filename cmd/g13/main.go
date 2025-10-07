package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/achilleas-k/g13-ak/internal/config"
	"github.com/achilleas-k/g13-ak/internal/device"
	"github.com/achilleas-k/g13-ak/internal/joystick"
	"github.com/achilleas-k/g13-ak/internal/keyboard"
	"github.com/spf13/cobra"
)

func mkcmd() *cobra.Command {
	rootCmd := cobra.Command{
		Use:                   "g13 <config>",
		Args:                  cobra.ExactArgs(1),
		Long:                  "Userspace Linux driver for the Logitech G13 gameboard",
		Version:               "devel",
		RunE:                  g13,
		DisableFlagsInUseLine: true, // don't put [flags] at the end of the Use line
	}

	return &rootCmd
}

func setCleanupHandler(cleanup func()) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for sig := range signalChan {
			if sig == os.Interrupt {
				fmt.Println("Stopping...")
				cleanup()
				break
			}
		}
		os.Exit(0)
	}()
}

func initialise(g13cfg *config.G13Config) (device.Device, keyboard.Keyboard, joystick.Joystick, error) {
	dev, err := device.New()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("device initialisation failed: %w", err)
	}
	setCleanupHandler(dev.Close)

	vkb, err := keyboard.New("g13-vkb")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("virtual keyboard initialisation failed: %w", err)
	}

	vjs, err := joystick.New("g13-vjs")
	if err != nil {
		return nil, nil, nil, fmt.Errorf("virtual joystick initialisation failed: %w", err)
	}

	backlight := g13cfg.GetBacklight()
	if err := dev.SetBacklightColour(backlight[0], backlight[1], backlight[2]); err != nil {
		return nil, nil, nil, err
	}

	if g13cfg.GetImagePath() != "" {
		lcdImg, err := g13cfg.GetImage()
		if err != nil {
			return nil, nil, nil, err
		}
		if err := dev.SetLCD(lcdImg); err != nil {
			return nil, nil, nil, err
		}
	}
	return dev, vkb, vjs, nil
}

func g13(cmd *cobra.Command, args []string) error {
	// SilenceUsage if the command executed correctly.
	// Argument parsing has already succeeded, so any error returned here
	// shouldn't show usage instructions but just print the error message.
	cmd.SilenceUsage = true

	configPath := args[0]
	g13cfg, err := config.NewFromFile(configPath)
	if err != nil {
		return err
	}

	dev, vkb, vjs, err := initialise(g13cfg)
	if err != nil {
		return err
	}

	defer func() {
		dev.Close()
		if err := vkb.Close(); err != nil {
			fmt.Fprintf(os.Stderr, "error closing keyboard during shutdown: %s", err)
		}
	}()

	fmt.Println("Ready")
	var consecutiveReadErrors uint8 = 0
	for {
		input, err := dev.ReadInput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "e: %s (%d)\n", err, consecutiveReadErrors)
			consecutiveReadErrors++
			// wait a bit before continuing to try to read
			time.Sleep(500 * time.Millisecond)

			// TODO: shut down if consecutive errors reaches a limit
			continue
		}

		consecutiveReadErrors = 0
		for kbkey, isDown := range g13cfg.GetKeyStates(input) {
			if isDown {
				if err := vkb.KeyDown(kbkey); err != nil {
					fmt.Fprintf(os.Stderr, "keyboard error pressing %d: %s\n", kbkey, err)
				}
			} else if err := vkb.KeyUp(kbkey); err != nil {
				fmt.Fprintf(os.Stderr, "keyboard error releasing %d: %s\n", kbkey, err)
			}
		}

		stickPos := g13cfg.GetStickPosition(input)
		if stickPos != nil {
			xOutput, yOutput := stickPos.UinputPosition()
			if err := vjs.StickPosition(xOutput, yOutput); err != nil {
				fmt.Fprintf(os.Stderr, "joystick error setting position %f %f", xOutput, yOutput)
			}
		}

		if consecutiveReadErrors > 0 {
			// device is back after read errors
			fmt.Println("Reinitialising device")
			dev.Close()
			dev = nil
			if err := vkb.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "error closing vkb: %s\n", err)
			}
			// After 3 consecutive read errors, try to reinitialise the device.
			// This is primarily meant to handle device disconnections.
			dev, vkb, vjs, err = initialise(g13cfg)
			if err != nil {
				return err
			}
			consecutiveReadErrors = 0
			fmt.Println("Device restored")
		}
	}
}

func main() {
	cmd := mkcmd()
	if err := cmd.Execute(); err != nil {
		// Don't print anything: Cobra will print error message with usage if
		// necessary
		os.Exit(1)
	}
}
