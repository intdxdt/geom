package geom

import "github.com/intdxdt/math"

//do two lines intersect line segments a && b with
//vertices sa, sb, oa, ob
func SegSegIntersects(sa, sb, oa, ob *Point) bool {
	var bln = false
	var a, b, d = segsegABD(sa, sb, oa, ob)

	//snap to zero if near -0 or 0
	if a == 0 || math.Abs(a) < math.EPSILON {
		a = 0
	}
	if b == 0 || math.Abs(b) < math.EPSILON {
		b = 0
	}
	if d == 0 || math.Abs(d) < math.EPSILON {
		d = 0
	}

	if d == 0 {
		if a == 0.0 && b == 0.0 {
			bln = bounds_intersects(sa, sb, oa, ob)
		}
		return bln
	}
	//intersection along the the seg or extended seg
	var ua = a / d
	if ua == 0 || math.Abs(ua) < math.EPSILON {
		ua = 0
	} else if ua == 1 || math.Abs(ua-1) < math.EPSILON {
		ua = 1
	}

	var ub = b / d
	if ub == 0 || math.Abs(ub) < math.EPSILON {
		ub = 0
	} else if ub == 1 || math.Abs(ub-1) < math.EPSILON {
		ub = 1
	}

	return (0 <= ua && ua <= 1) && (0 <= ub && ub <= 1)
}
