// Package config provides functionality for loading and defining the device
// configuration, which includes mapping G13 buttons to keyboard keys.
package config

import (
	"encoding/json"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/achilleas-k/g13-ak/internal/device"
	"github.com/achilleas-k/g13-ak/internal/keyboard"
	"golang.org/x/image/bmp"
)

// G13Config maps G13 keys to uinput key codes.
type G13Config struct {
	// mapping from G keys to keyboard keycodes
	keyMap keyMap

	// backlight rgb
	backlight [3]uint8

	// path to image configured for the display
	lcdImage string
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
	cfg, err := loadConfig(path)
	if err != nil {
		return nil, err
	}

	return cfg, nil
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

func (cfg *G13Config) GetImagePath() string {
	return cfg.lcdImage
}

func (cfg *G13Config) GetImage() (image.Image, error) {
	path := cfg.lcdImage
	if path == "" {
		return nil, fmt.Errorf("no image file defined in config")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open image file %q: %w", path, err)
	}
	img, err := bmp.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file %q: %w", path, err)
	}

	return img, nil

}

// fileConfig describes the on-disk file format for the config file.
type fileConfig struct {
	Mapping   map[string]string   `json:"mapping"`
	Backlight backlightFileConfig `json:"backlight"`
	ImageFile string              `json:"image_file"`
}

type backlightFileConfig struct {
	Red   uint8 `json:"red"`
	Green uint8 `json:"green"`
	Blue  uint8 `json:"blue"`
}

func loadConfig(path string) (*G13Config, error) {
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

	errPrefix := "failed reading config file"
	km := make(keyMap, len(cfg.Mapping))
	for gKeyStr, kbKeyStr := range cfg.Mapping {
		gKey := device.KeyCode(gKeyStr)
		if gKey == 0 {
			return nil, fmt.Errorf("%s: unknown G13 key name: %s", errPrefix, gKeyStr)
		}
		kbKey := keyboard.KeyCode(kbKeyStr)
		if kbKey == 0 {
			return nil, fmt.Errorf("%s: unknown keyboard key name: %s", errPrefix, kbKeyStr)
		}
		km[gKey] = kbKey
	}

	backlight := [3]uint8{cfg.Backlight.Red, cfg.Backlight.Green, cfg.Backlight.Blue}

	imageFile := cfg.ImageFile

	if imageFile != "" {
		// The image file, if defined, should be relative to the config file
		// (unless it's already absolute)
		if !filepath.IsAbs(imageFile) {
			cfgDir, err := filepath.Abs(filepath.Dir(path))
			if err != nil {
				return nil, fmt.Errorf("failed to get absolute path of config file %q: %w", path, err)
			}
			imageFile = filepath.Clean(filepath.Join(cfgDir, imageFile))
		}

		// Check if the image file exists and is stat-able if it's set; no
		// need for any extra validation right now
		_, err := os.Stat(imageFile)
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%s: image file %q (%s) set in config file does not exist", errPrefix, cfg.ImageFile, imageFile)
		}
		if err != nil {
			return nil, fmt.Errorf("%s: %w", errPrefix, err)
		}
	}

	return &G13Config{
		keyMap:    km,
		backlight: backlight,
		lcdImage:  imageFile,
	}, nil
}
