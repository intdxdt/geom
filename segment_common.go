package geom

import (
	"github.com/intdxdt/mbr"
)

type VBits uint8

const (
	OtherB VBits = 1 << iota // 1 << 0 == 0001
	OtherA                   // 1 << 1 == 0010
	SelfB                    // 1 << 2 == 0100
	SelfA                    // 1 << 3 == 1000
)
const InterX VBits = 0

const (
	SelfMask  = SelfA | SelfB
	OtherMask = OtherA | OtherB
)

//clamp to zero if float is near zero
func snap_to_zero(v float64) float64 {
	if feq(v, 0.0) {
		v = 0.0
	}
	return v
}

//clamp to zero or one
func snap_to_zero_or_one(v float64) float64 {
	if feq(v, 0.0) {
		v = 0.0
	} else if feq(v, 1.0) {
		v = 1.0
	}
	return v
}

//envelope of segment
func BBox(a, b *Point) mbr.MBR {
	return mbr.CreateMBR(a[X], a[Y], b[X], b[Y])
}
