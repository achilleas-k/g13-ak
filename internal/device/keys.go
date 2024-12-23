package device

// KeyBit defines, for each button on the G13, the corresponding single bit
// mask that can be applied to the value returned by [MaskDataForInput].
// For example:
//
//	G1  = 0b1
//	G2  = 0b10
//	G2  = 0b100
//	G14 = 0b10000000000000
type KeyBit uint64

const (
	G1 KeyBit = 1 << iota
	G2
	G3
	G4
	G5
	G6
	G7
	G8
	G9
	G10
	G11
	G12
	G13
	G14
	G15
	G16
	G17
	G18
	G19
	G20
	G21
	G22
	UNDEF1
	LIGHT_STATE
	BD
	L1
	L2
	L3
	L4
	M1
	M2
	M3
	MR
	LEFT
	DOWN
	TOP
	UNDEF3
	LIGHT
	LIGHT2
	MISC_TOGGLE
)

var (
	KeyNames = []string{
		"G1", "G2", "G3", "G4", "G5", "G6", "G7", "G8",
		"G9", "G10", "G11", "G12", "G13", "G14", "G15", "G16",
		"G17", "G18", "G19", "G20", "G21", "G22", "UNDEF1", "LIGHT_STATE",
		"BD", "L1", "L2", "L3", "L4", "M1", "M2", "M3",
		"MR", "LEFT", "DOWN", "TOP", "UNDEF3", "LIGHT", "LIGHT2", "MISC_TOGGLE",
	}

	// Mask (LE order) for button states (1 down, 0 up).
	buttonStateMask = uint64(0b00001111_11111111_00111111_11111111_11111111_00000000_00000000_00000000)
)

// MaskDataForInput returns the given data masked to only contain the bits
// relevant for reading button states. The input is assumed to be 8 bytes and
// LE ordered, as it is read from [G13Device.ReadInput].
func MaskDataForInput(data uint64) uint64 {
	return (data & buttonStateMask) >> 24
}

func btoiBE(b []byte) (i uint64) {
	for _, bi := range b {
		i <<= 8
		i += uint64(bi)
	}
	return
}

func btoiLE(b []byte) (i uint64) {
	for idx := len(b) - 1; idx >= 0; idx-- {
		i <<= 8
		i += uint64(b[idx])
	}
	return
}

func itobBE(i uint64) (b []byte) {
	b = make([]byte, 8)
	for idx := 7; idx >= 0; idx-- {
		b[idx] = byte(i & 0xFF)
		i >>= 8
	}
	return
}

func itobLE(i uint64) (b []byte) {
	b = make([]byte, 8)
	for idx := 0; idx < 8; idx++ {
		b[idx] = byte(i & 0xFF)
		i >>= 8
	}
	return
}
