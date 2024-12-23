package device_test

import (
	"testing"

	"github.com/achilleas-k/g13-ak/internal/device"
	"github.com/stretchr/testify/assert"
)

type testData struct {
	data     uint64
	keyNames []string
}

var (
	// short sequence of recorded data from device
	smallDataSet = []testData{
		{
			data:     0x8000800001707801,
			keyNames: []string{"G1"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800002707801,
			keyNames: []string{"G2"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x800400707801,
			keyNames: []string{"G11"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800400707801,
			keyNames: []string{"G11"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x800400707801,
			keyNames: []string{"G11"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x800800707801,
			keyNames: []string{"G12"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x801000707801,
			keyNames: []string{"G13"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800040707801,
			keyNames: []string{"G7"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000a00000707801,
			keyNames: []string{"G22"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x900000707801,
			keyNames: []string{"G21"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8200800000707801,
			keyNames: []string{"LEFT"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x400800000707801,
			keyNames: []string{"DOWN"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800002707801,
			keyNames: []string{"G2"},
		},
		{
			data:     0x8000800006707801,
			keyNames: []string{"G2", "G3"},
		},
		{
			data:     0x8000800004707801,
			keyNames: []string{"G3"},
		},
		{
			data:     0x800080000c707801,
			keyNames: []string{"G3", "G4"},
		},
		{
			data:     0x8000800008707801,
			keyNames: []string{"G4"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
	}

	// (almost) all buttons pressed and released sequentially
	allButtons = []testData{
		{
			data:     0x8000800001707801,
			keyNames: []string{"G1"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x800002707801,
			keyNames: []string{"G2"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800004707801,
			keyNames: []string{"G3"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800008707801,
			keyNames: []string{"G4"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x800010707801,
			keyNames: []string{"G5"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x800020707801,
			keyNames: []string{"G6"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800040707801,
			keyNames: []string{"G7"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800080707801,
			keyNames: []string{"G8"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800100707801,
			keyNames: []string{"G9"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x800200707801,
			keyNames: []string{"G10"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x800400707801,
			keyNames: []string{"G11"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000800800707801,
			keyNames: []string{"G12"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000801000707801,
			keyNames: []string{"G13"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000802000707801,
			keyNames: []string{"G14"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x804000707801,
			keyNames: []string{"G15"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x808000707801,
			keyNames: []string{"G16"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000810000707801,
			keyNames: []string{"G17"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000820000707801,
			keyNames: []string{"G18"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x840000707801,
			keyNames: []string{"G19"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x880000707801,
			keyNames: []string{"G20"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x900000707801,
			keyNames: []string{"G21"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8000a00000707801,
			keyNames: []string{"G22"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8200800000707801,
			keyNames: []string{"LEFT"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8400800000707801,
			keyNames: []string{"DOWN"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8020800000707801,
			keyNames: []string{"M1"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x40800000707801,
			keyNames: []string{"M2"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8080800000707801,
			keyNames: []string{"M3"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x100800000707801,
			keyNames: []string{"MR"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8001800000707801,
			keyNames: []string{"BD"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x2800000707801,
			keyNames: []string{"L1"},
		},
		{
			data:     0x8000800000707801,
			keyNames: []string{},
		},
		{
			data:     0x4800000707801,
			keyNames: []string{"L2"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x8008800000707801,
			keyNames: []string{"L3"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
		{
			data:     0x10800000707801,
			keyNames: []string{"L4"},
		},
		{
			data:     0x800000707801,
			keyNames: []string{},
		},
	}
)

func TestButtonIdentification(t *testing.T) {
	// This test ensures that any internal button representation change doesn't
	// affect the identification of pressed buttons.
	assert := assert.New(t)
	for idx, testItem := range smallDataSet {
		data := testItem.data
		expectedKeyNames := testItem.keyNames

		var decodedKeyNames []string
		for _, key := range device.AllKeys() {
			if key.Uint64()&data != 0 {
				// key is pressed
				decodedKeyNames = append(decodedKeyNames, key.String())
			}
		}
		assert.ElementsMatch(expectedKeyNames, decodedKeyNames, "smallDataSet[%d]: %#v", idx, data)
	}

	for idx, testItem := range allButtons {
		data := testItem.data
		expectedKeyNames := testItem.keyNames

		var decodedKeyNames []string
		for _, key := range device.AllKeys() {
			if key.Uint64()&data != 0 {
				// key is pressed
				decodedKeyNames = append(decodedKeyNames, key.String())
			}
		}
		assert.ElementsMatch(expectedKeyNames, decodedKeyNames, "allButtons[%d]: %#v", idx, data)
	}
}
