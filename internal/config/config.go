// Package config provides functionality for loading and defining the device
// configuration, which includes mapping G13 buttons to keyboard keys.
package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/achilleas-k/g13-ak/internal/device"
	"github.com/achilleas-k/g13-ak/internal/keyboard"
)

// G13Config maps G13 keys to uinput key codes.
type G13Config struct {
	keyMap    keyMap
	backlight [3]uint8
}

type keyMap map[device.KeyBit]int

// NewEmpty returns an empty [G13Config].
func NewEmpty() *G13Config {
	return &G13Config{
		keyMap: make(keyMap, len(device.AllKeys())),
	}
}

// NewFromFile returns a [G13Config] initialised from the file at the given path.
func NewFromFile(path string) (*G13Config, error) {
	km, err := loadConfig(path)
	if err != nil {
		return nil, err
	}

	return convertConfig(km)
}

// SetKey maps a G13 key to the given keyboard key.
func (m *G13Config) SetKey(gkey device.KeyBit, kbKey int) {
	m.keyMap[gkey] = kbKey
}

// SetKeys maps one or more G13 keys to the given keyboard key. It does not
// override any mappings not present in keyMap.
func (m *G13Config) SetKeys(km keyMap) {
	for gkey, kbkey := range km {
		m.keyMap[gkey] = kbkey
	}
}

// UnsetKey unmaps a gkey.
func (m *G13Config) UnsetKey(gkey device.KeyBit) {
	delete(m.keyMap, gkey)
}

// Reset unmaps all G13 keys.
func (m *G13Config) Reset() {
	m.keyMap = make(keyMap, len(device.AllKeys()))
}

// GetKeyStates returns the state of each mapped keyboard key for the given
// input (from [device.ReadInput]). The result maps a keyboard keycode to a
// state, true for down (pressed) and false for up (released).
func (m *G13Config) GetKeyStates(input uint64) map[int]bool {
	kbkeys := make(map[int]bool, len(m.keyMap))
	for gkey, kbkey := range m.keyMap {
		kbkeys[kbkey] = (gkey.Uint64() & input) != 0
	}
	return kbkeys
}

func (cfg *G13Config) GetBacklight() [3]uint8 {
	return cfg.backlight
}

// fileConfig describes the on-disk file format for the config file.
type fileConfig struct {
	Mapping   map[string]string   `json:"mapping"`
	Backlight backlightFileConfig `json:"backlight"`
}

type backlightFileConfig struct {
	Red   uint8 `json:"red"`
	Green uint8 `json:"green"`
	Blue  uint8 `json:"blue"`
}

func loadConfig(path string) (*fileConfig, error) {
	configFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed opening config file %q: %w", path, err)
	}

	cfg := fileConfig{}
	decoder := json.NewDecoder(configFile)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed decoding config file %q: %w", path, err)
	}

	return &cfg, nil
}

// Convert the on-disk fileConfig to a G13Config.
func convertConfig(cfg *fileConfig) (*G13Config, error) {
	km := make(keyMap, len(cfg.Mapping))
	for gKeyStr, kbKeyStr := range cfg.Mapping {
		gKey := device.KeyCode(gKeyStr)
		if gKey == 0 {
			return nil, fmt.Errorf("unknown G13 key name: %s", gKeyStr)
		}
		kbKey := keyboard.KeyCode(kbKeyStr)
		if kbKey == 0 {
			return nil, fmt.Errorf("unknown keyboard key name: %s", kbKeyStr)
		}
		km[gKey] = kbKey
	}

	backlight := [3]uint8{cfg.Backlight.Red, cfg.Backlight.Green, cfg.Backlight.Blue}

	return &G13Config{
		keyMap:    km,
		backlight: backlight,
	}, nil
}
