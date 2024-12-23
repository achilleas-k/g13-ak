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
	allKeys = []KeyBit{
		G1, G2, G3, G4, G5, G6, G7, G8,
		G9, G10, G11, G12, G13, G14, G15, G16,
		G17, G18, G19, G20, G21, G22,
		BD, L1, L2, L3, L4, M1, M2, M3,
		MR, LEFT, DOWN, TOP,
	}

	keyNames = map[KeyBit]string{
		G1: "G1", G2: "G2", G3: "G3", G4: "G4", G5: "G5", G6: "G6", G7: "G7", G8: "G8",
		G9: "G9", G10: "G10", G11: "G11", G12: "G12", G13: "G13", G14: "G14", G15: "G15", G16: "G16",
		G17: "G17", G18: "G18", G19: "G19", G20: "G20", G21: "G21", G22: "G22", UNDEF1: "UNDEF1", LIGHT_STATE: "LIGHT_STATE",
		BD: "BD", L1: "L1", L2: "L2", L3: "L3", L4: "L4", M1: "M1", M2: "M2", M3: "M3",
		MR: "MR", LEFT: "LEFT", DOWN: "DOWN", TOP: "TOP", UNDEF3: "UNDEF3", LIGHT: "LIGHT", LIGHT2: "LIGHT2", MISC_TOGGLE: "MISC_TOGGLE",
	}

	// Mask (LE order) for button states (1 down, 0 up).
	buttonStateMask = uint64(0b00001111_11111111_00111111_11111111_11111111_00000000_00000000_00000000)
)

func (kb KeyBit) String() string {
	return keyNames[kb]
}

func (kb KeyBit) Uint64() uint64 {
	return uint64(kb)
}

func AllKeys() []KeyBit {
	return allKeys
}

// MaskDataForInput returns the given data masked to only contain the bits
// relevant for reading button states. The input is assumed to be 8 bytes and
// LE ordered, as it is read from [Device.ReadInput].
func MaskDataForInput(data uint64) uint64 {
	return (data & buttonStateMask) >> 24
}

func btoiLE(b []byte) (i uint64) {
	for idx := len(b) - 1; idx >= 0; idx-- {
		i <<= 8
		i += uint64(b[idx])
	}
	return
}
