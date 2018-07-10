package geom

import (
	"github.com/intdxdt/mbr"
)

//pop chain from chainl list
func pop_mono_mbr(a []MonoMBR) (MonoMBR, []MonoMBR) {
	var v MonoMBR
	var n int
	if len(a) == 0 {
		return MonoMBR{}, a
	}
	n = len(a) - 1
	v, a[n] = a[n], MonoMBR{}
	return v, a[:n]
}

//pop chain from chainl list
func pop_coords(a []Point) (Point, []Point) {
	var v Point
	var n int
	if len(a) == 0 {
		return NullPt, a
	}
	n = len(a) - 1
	v, a[n] = a[n], NullPt
	return v, a[:n]
}

//compare bbox, ptr
func is_bbox(a *mbr.MBR, b *mbr.MBR) bool {
	return a == b
}
