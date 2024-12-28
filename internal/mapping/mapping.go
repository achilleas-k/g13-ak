// Package mapping provides functionality for mapping G13 buttons to keyboard
// keys.
package mapping

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/achilleas-k/g13-ak/internal/device"
	"github.com/achilleas-k/g13-ak/internal/keyboard"
)

type keyMap map[device.KeyBit]int

// Mapping maps G13 keys to uinput key codes.
type Mapping struct {
	keyMap keyMap
}

// NewEmpty returns an empty [Mapping].
func NewEmpty() *Mapping {
	return &Mapping{
		keyMap: make(keyMap, len(device.AllKeys())),
	}
}

// NewFromFile returns a [Mapping] initialised from the file at the given path.
func NewFromFile(path string) (*Mapping, error) {
	km, err := loadConfig(path)
	if err != nil {
		return nil, err
	}

	return &Mapping{
		keyMap: km,
	}, nil
}

// SetKey maps a G13 key to the given keyboard key.
func (m *Mapping) SetKey(gkey device.KeyBit, kbKey int) {
	m.keyMap[gkey] = kbKey
}

// SetKeys maps one or more G13 keys to the given keyboard key. It does not
// override any mappings not present in keyMap.
func (m *Mapping) SetKeys(km keyMap) {
	for gkey, kbkey := range km {
		m.keyMap[gkey] = kbkey
	}
}

// UnsetKey unmaps a gkey.
func (m *Mapping) UnsetKey(gkey device.KeyBit) {
	delete(m.keyMap, gkey)
}

// Reset unmaps all G13 keys.
func (m *Mapping) Reset() {
	m.keyMap = make(keyMap, len(device.AllKeys()))
}

// GetKeyStates returns the state of each mapped keyboard key for the given
// input (from [device.ReadInput]). The result maps a keyboard keycode to a
// state, true for down (pressed) and false for up (released).
func (m *Mapping) GetKeyStates(input uint64) map[int]bool {
	kbkeys := make(map[int]bool, len(m.keyMap))
	for gkey, kbkey := range m.keyMap {
		kbkeys[kbkey] = (gkey.Uint64() & input) != 0
	}
	return kbkeys
}

func loadConfig(path string) (keyMap, error) {
	fileConfig := map[string]string{}
	configFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed opening config file %q: %w", path, err)
	}

	decoder := json.NewDecoder(configFile)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&fileConfig); err != nil {
		return nil, fmt.Errorf("failed decoding config file %q: %w", path, err)
	}

	km := make(keyMap, len(fileConfig))
	for gKeyStr, kbKeyStr := range fileConfig {
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
	return km, nil
}
