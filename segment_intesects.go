package geom

//do two lines intersect line segments a && b with
//vertices sa, sb, oa, ob
func SegSegIntersects(sa, sb, oa, ob *Point) bool {
	var bln = false
	var a, b, d = segsegABD(sa[:], sb[:], oa[:], ob[:])

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
	var ua = a / d
	if feq(ua, 0) {
		ua = 0.0
	} else if feq(ua, 1) {
		ua = 1.0
	}

	var ub = b / d
	if feq(ub, 0) {
		ub = 0.0
	} else if feq(ub, 1) {
		ub = 1.0
	}

	return (0 <= ua && ua <= 1) && (0 <= ub && ub <= 1)
}
