// Package mapping provides functionality for mapping G13 buttons to keyboard
// keys.
package mapping

import "github.com/achilleas-k/g13-ak/internal/device"

// Mapping maps G13 keys to uinput key codes.
type Mapping struct {
	keyMap map[device.KeyBit]int
}

// New empty [Mapping].
func New() *Mapping {
	return &Mapping{
		keyMap: make(map[device.KeyBit]int, len(device.AllKeys())),
	}
}

// SetKey maps a G13 key to the given keyboard key.
func (m *Mapping) SetKey(gkey device.KeyBit, kbKey int) {
	m.keyMap[gkey] = kbKey
}

// SetKeys maps one or more G13 keys to the given keyboard key. It does not
// override any mappings not present in keyMap.
func (m *Mapping) SetKeys(keyMap map[device.KeyBit]int) {
	for gkey, kbkey := range keyMap {
		m.keyMap[gkey] = kbkey
	}
}

// UnsetKey unmaps a gkey.
func (m *Mapping) UnsetKey(gkey device.KeyBit) {
	delete(m.keyMap, gkey)
}

// Reset unmaps all G13 keys.
func (m *Mapping) Reset() {
	m.keyMap = make(map[device.KeyBit]int, len(device.AllKeys()))
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