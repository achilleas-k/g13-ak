package config_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/achilleas-k/g13-ak/internal/config"
	"github.com/achilleas-k/g13-ak/internal/device"
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
			configData: `{"mapping":{"keys":{"G1":"Key1","G22":"KeyT"}}}`,
			expectedKeymap: map[device.KeyBit]int{
				device.G1:  uinput.Key1,
				device.G22: uinput.KeyT,
			},
		},
		"full": {
			configData: `{
	"mapping": {
		"keys": {
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
		}
	}
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
"mapping":{
	"keys": {
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
	}
}}`,
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

			m, err := config.NewFromFile(cfgPath)
			assert.NoError(err)

			expectedMapping := config.NewEmpty()
			expectedMapping.SetKeys(tc.expectedKeymap)
			assert.Equal(expectedMapping, m)
		})
	}
}

func TestNewFromFileError(t *testing.T) {
	t.Run("file-not-found", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")
		_, err := config.NewFromFile(cfgPath)
		assert.ErrorContains(err, "no such file or directory")
	})

	t.Run("bad-g13-key", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{"mapping":{"keys":{"G23":"KeyA"}}}`), 0o660)
		assert.NoError(err)

		_, err = config.NewFromFile(cfgPath)
		assert.EqualError(err, "failed reading config file: unknown G13 key name: G23")
	})

	t.Run("bad-kb-key", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{"mapping":{"keys":{"G2":"NotAKey"}}}`), 0o660)
		assert.NoError(err)

		_, err = config.NewFromFile(cfgPath)
		assert.EqualError(err, "failed reading config file: unknown keyboard key name: NotAKey")
	})

	t.Run("permission-denied", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{}`), 0o000)
		assert.NoError(err)

		_, err = config.NewFromFile(cfgPath)
		assert.ErrorContains(err, "failed opening config file")
		assert.ErrorContains(err, "permission denied")
	})

	t.Run("wrong-format", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{"bad-top-level-config-key":{}}`), 0o660)
		assert.NoError(err)

		_, err = config.NewFromFile(cfgPath)
		assert.ErrorContains(err, "failed decoding config file")
		assert.ErrorContains(err, "unknown field \"bad-top-level-config-key\"")
	})

	t.Run("image-does-not-exist", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{"image_file":"1c606e45-c739-43a8-bcc4-39b8c9845794.bmp"}`), 0o660)
		assert.NoError(err)

		_, err = config.NewFromFile(cfgPath)
		assert.ErrorContains(err, "failed reading config file: image file")
		assert.ErrorContains(err, "set in config file does not exist")
	})

	t.Run("bad-stick-mode", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")
		err := os.WriteFile(cfgPath, []byte(`{"mapping":{"stick":{"mode":"bad"}}}`), 0o660)
		assert.NoError(err)

		_, err = config.NewFromFile(cfgPath)
		assert.ErrorContains(err, "unknown stick mode: bad")
	})

	t.Run("bad-stick-key", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")
		err := os.WriteFile(cfgPath, []byte(`{"mapping":{"stick":{"mode":"keys","keys":{"Up":"up"}}}}`), 0o660)
		assert.NoError(err)

		_, err = config.NewFromFile(cfgPath)
		assert.ErrorContains(err, "unknown keyboard key name: up")
	})
}

func TestDefaultConfig(t *testing.T) {
	cfgPath := "../../configs/default.json"
	_, err := config.NewFromFile(cfgPath)
	assert.NoError(t, err)
}

func TestGetImageErrors(t *testing.T) {
	t.Run("no-image-in-config", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte("{}"), 0o660)
		assert.NoError(err)

		cfg, err := config.NewFromFile(cfgPath)
		assert.NoError(err)

		assert.Equal("", cfg.GetImagePath())
		_, err = cfg.GetImage()
		assert.EqualError(err, "no image file defined in config")
	})

	t.Run("cant-open-image-file", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{"image_file":"0.bmp"}`), 0o660)
		assert.NoError(err)

		fp, err := os.OpenFile(filepath.Join(tmpdir, "0.bmp"), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0000)
		assert.NoError(err)
		_ = fp.Close()

		cfg, err := config.NewFromFile(cfgPath)
		assert.NoError(err)

		_, err = cfg.GetImage()
		assert.ErrorContains(err, "failed to open image file")
		assert.ErrorContains(err, "permission denied")
	})

	t.Run("unreadable-image-file", func(t *testing.T) {
		assert := assert.New(t)

		tmpdir := t.TempDir()
		cfgPath := filepath.Join(tmpdir, "mapping.json")

		err := os.WriteFile(cfgPath, []byte(`{"image_file":"/dev/zero"}`), 0o660)
		assert.NoError(err)

		cfg, err := config.NewFromFile(cfgPath)
		assert.NoError(err)

		_, err = cfg.GetImage()
		assert.ErrorContains(err, "failed to read image file")
		assert.ErrorContains(err, "invalid format")
	})
}
