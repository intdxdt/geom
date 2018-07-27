package geom

import "github.com/intdxdt/geom/mono"

//pop chain from chainl list
func pop_mono_mbr(a []mono.MBR) (mono.MBR, []mono.MBR) {
	var v mono.MBR
	var n int
	if len(a) == 0 {
		return mono.MBR{}, a
	}
	n = len(a) - 1
	v, a[n] = a[n], mono.MBR{}
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
