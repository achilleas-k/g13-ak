package mapping_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/achilleas-k/g13-ak/internal/device"
	"github.com/achilleas-k/g13-ak/internal/mapping"
	"github.com/bendahl/uinput"
	"github.com/stretchr/testify/assert"
)

func TestNewFromFile(t *testing.T) {

	type testCase struct {
		configData     string
		expectedKeymap map[device.KeyBit]int
	}

	testCases := map[string]testCase{
		"empty": {
			configData:     "{}",
			expectedKeymap: map[device.KeyBit]int{},
		},
		"simple": {
			configData: `{"G1":"Key1","G22":"KeyT"}`,
			expectedKeymap: map[device.KeyBit]int{
				device.G1:  uinput.Key1,
				device.G22: uinput.KeyT,
			},
		},
		"full": {
			configData: `{
	"G1": "KeyLeftctrl",
	"G2": "KeyRightbrace",
	"G3": "Key4",
	"G4": "KeyO",
	"G5": "Key2",
	"G6": "KeyP",
	"G7": "KeyBackspace",
	"G8": "KeyF2",
	"G9": "KeyEsc",
	"G10": "KeyLeftbrace",
	"G11": "KeyF",
	"G12": "KeyRightshift",
	"G13": "KeyU",
	"G14": "KeyCapslock",
	"G15": "KeySlash",
	"G16": "KeyG",
	"G17": "Key7",
	"G18": "KeyR",
	"G19": "KeyK",
	"G20": "KeyV",
	"G21": "KeyE",
	"G22": "KeyC",
	"L1": "KeyY",
	"L2": "KeyH",
	"L3": "Key8",
	"L4": "KeyX",
	"LEFT": "KeyA",
	"DOWN": "Key5",
	"BD": "KeyGrave",
	"M1": "KeyTab",
	"M2": "KeyF6",
	"M3": "KeyL",
	"MR": "KeyI",
	"TOP": "KeyApostrophe"
}`,
			expectedKeymap: map[device.KeyBit]int{
				device.G1:   uinput.KeyLeftctrl,
				device.G2:   uinput.KeyRightbrace,
				device.G3:   uinput.Key4,
				device.G4:   uinput.KeyO,
				device.G5:   uinput.Key2,
				device.G6:   uinput.KeyP,
				device.G7:   uinput.KeyBackspace,
				device.G8:   uinput.KeyF2,
				device.G9:   uinput.KeyEsc,
				device.G10:  uinput.KeyLeftbrace,
				device.G11:  uinput.KeyF,
				device.G12:  uinput.KeyRightshift,
				device.G13:  uinput.KeyU,
				device.G14:  uinput.KeyCapslock,
				device.G15:  uinput.KeySlash,
				device.G16:  uinput.KeyG,
				device.G17:  uinput.Key7,
				device.G18:  uinput.KeyR,
				device.G19:  uinput.KeyK,
				device.G20:  uinput.KeyV,
				device.G21:  uinput.KeyE,
				device.G22:  uinput.KeyC,
				device.L1:   uinput.KeyY,
				device.L2:   uinput.KeyH,
				device.L3:   uinput.Key8,
				device.L4:   uinput.KeyX,
				device.LEFT: uinput.KeyA,
				device.DOWN: uinput.Key5,
				device.BD:   uinput.KeyGrave,
				device.M1:   uinput.KeyTab,
				device.M2:   uinput.KeyF6,
				device.M3:   uinput.KeyL,
				device.MR:   uinput.KeyI,
				device.TOP:  uinput.KeyApostrophe,
			},
		},
		"full-with-dupes": {
			configData: `{
	"G1": "KeyLeftctrl",
	"G2": "KeyRightbrace",
	"G3": "Key3",
	"G4": "Key3",
	"G5": "Key3",
	"G6": "KeyP",
	"G7": "KeyBackspace",
	"G8": "KeyF2",
	"G9": "KeyEsc",
	"G10": "KeyLeftbrace",
	"G11": "KeyF",
	"G12": "KeyRightshift",
	"G13": "KeyU",
	"G14": "KeyCapslock",
	"G15": "KeySlash",
	"G16": "KeyG",
	"G17": "Key7",
	"G18": "KeyR",
	"G19": "KeyK",
	"G20": "KeyV",
	"G21": "KeyE",
	"G22": "KeyC",
	"L1": "KeyY",
	"L2": "KeyH",
	"L3": "Key8",
	"L4": "KeyX",
	"LEFT": "KeyA",
	"DOWN": "Key5",
	"BD": "KeyGrave",
	"M1": "KeyTab",
	"M2": "KeyF6",
	"M3": "KeyL",
	"MR": "KeyI",
	"TOP": "KeyApostrophe"
}`,
			expectedKeymap: map[device.KeyBit]int{
				device.G1:   uinput.KeyLeftctrl,
				device.G2:   uinput.KeyRightbrace,
				device.G3:   uinput.Key3,
				device.G4:   uinput.Key3,
				device.G5:   uinput.Key3,
				device.G6:   uinput.KeyP,
				device.G7:   uinput.KeyBackspace,
				device.G8:   uinput.KeyF2,
				device.G9:   uinput.KeyEsc,
				device.G10:  uinput.KeyLeftbrace,
				device.G11:  uinput.KeyF,
				device.G12:  uinput.KeyRightshift,
				device.G13:  uinput.KeyU,
				device.G14:  uinput.KeyCapslock,
				device.G15:  uinput.KeySlash,
				device.G16:  uinput.KeyG,
				device.G17:  uinput.Key7,
				device.G18:  uinput.KeyR,
				device.G19:  uinput.KeyK,
				device.G20:  uinput.KeyV,
				device.G21:  uinput.KeyE,
				device.G22:  uinput.KeyC,
				device.L1:   uinput.KeyY,
				device.L2:   uinput.KeyH,
				device.L3:   uinput.Key8,
				device.L4:   uinput.KeyX,
				device.LEFT: uinput.KeyA,
				device.DOWN: uinput.Key5,
				device.BD:   uinput.KeyGrave,
				device.M1:   uinput.KeyTab,
				device.M2:   uinput.KeyF6,
				device.M3:   uinput.KeyL,
				device.MR:   uinput.KeyI,
				device.TOP:  uinput.KeyApostrophe,
			},
		},
	}

	for name := range testCases {
		tc := testCases[name]
		t.Run(name, func(t *testing.T) {
			assert := assert.New(t)

			tmpdir := t.TempDir()
			cfgPath := filepath.Join(tmpdir, "mapping.json")

			err := os.WriteFile(cfgPath, []byte(tc.configData), 0o660)
			assert.NoError(err)

			m, err := mapping.NewFromFile(cfgPath)
			assert.NoError(err)

			expectedMapping := mapping.NewEmpty()
			expectedMapping.SetKeys(tc.expectedKeymap)
			assert.Equal(m, expectedMapping)
		})
	}
}

func TestNewFromFileError(t *testing.T) {
	t.Run("file-not-found", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")
		_, err := mapping.NewFromFile(cfgPath)
		assert.ErrorContains(err, "no such file or directory")
	})

	t.Run("bad-g13-key", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{"G23":"KeyA"}`), 0o660)
		assert.NoError(err)

		_, err = mapping.NewFromFile(cfgPath)
		assert.EqualError(err, "unknown G13 key name: G23")
	})

	t.Run("bad-kb-key", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{"G2":"NotAKey"}`), 0o660)
		assert.NoError(err)

		_, err = mapping.NewFromFile(cfgPath)
		assert.EqualError(err, "unknown keyboard key name: NotAKey")
	})
}

func TestDefaultConfig(t *testing.T) {
	cfgPath := "../../configs/default.json"
	_, err := mapping.NewFromFile(cfgPath)
	assert.NoError(t, err)
}
