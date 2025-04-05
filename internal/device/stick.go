package device

const (
	XMask = 255 << 8
	YMask = 255 << 16
)

// StickPosition returns the x, y position of the stick decoded from the input.
func StickPosition(input uint64) (uint8, uint8) {
	x := (input & XMask) >> 8  // (input & (255 << 8) >> 8)
	y := (input & YMask) >> 16 // (input & (255 << 16) >> 16)
	return uint8(x), uint8(y)
}
