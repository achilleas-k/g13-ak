package main

import (
	"fmt"
	"os"

	"github.com/achilleas-k/g13-ak/internal/config"
	"github.com/achilleas-k/g13-ak/internal/device"
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

	dev, err := device.New()
	if err != nil {
		return fmt.Errorf("device initialisation failed: %w", err)
	}
	defer dev.Close()

	vkb, err := keyboard.New("g13-vkb")
	if err != nil {
		return fmt.Errorf("virtual keyboard initialisation failed: %w", err)
	}

	if err := dev.SetBacklightColour(255, 0, 0); err != nil {
		return err
	}

	fmt.Println("Ready")
	for {
		input, err := dev.ReadInput()
		if err != nil {
			fmt.Fprintf(os.Stderr, "e: %s\n", err)
		}
		for kbkey, isDown := range g13cfg.GetKeyStates(input) {
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

func main() {
	cmd := mkcmd()
	if err := cmd.Execute(); err != nil {
		// Don't print anything: Cobra will print error message with usage if
		// necessary
		os.Exit(1)
	}
}
