package geom

import (
	"github.com/intdxdt/math"
)

//do two lines intersect line segments a && b with
//vertices sa, sb, oa, ob
func SegSegIntersects(sa, sb, oa, ob *Point) bool {
	var bln = false
	var a, b, d = segseg_abd(sa[:], sb[:], oa[:], ob[:])

	//snap to zero if near -0 or 0
	a = snap_to_zero(a)
	b = snap_to_zero(b)
	d = snap_to_zero(d)

	if d == 0 {
		if a == 0.0 && b == 0.0 {
			bln = bounds_intersects(sa, sb, oa, ob)
		}
		return bln
	}
	//intersection along the the seg or extended seg
	ua := snap_to_zero_or_one(a / d)
	ub := snap_to_zero_or_one(b / d)
	return (0 <= ua && ua <= 1) && (0 <= ub && ub <= 1)
}

//Checks if two bounding boxes intesect
func bounds_intersects(sa, sb, oa, ob *Point) bool {
	var  s_minx, s_miny, s_maxx, s_maxy = min_bounds_rect(sa, sb)
	var  o_minx, o_miny, o_maxx, o_maxy = min_bounds_rect(oa, ob)
	//not disjoint
	return !(
		o_minx > s_maxx ||
		o_maxx < s_minx ||
		o_miny > s_maxy ||
		o_maxy < s_miny)
}

func min_bounds_rect(a, b *Point) (float64, float64, float64, float64) {
	var minx, maxx = math.MinF64(a[X], b[X]), math.MaxF64(a[X], b[X])
	var miny, maxy = math.MinF64(a[Y], b[Y]), math.MaxF64(a[Y], b[Y])
	return minx, miny, maxx, maxy
}
