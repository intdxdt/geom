package geom

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
			bln = BBox(sa, sb).Intersects(BBox(oa, ob))
		}
		return bln
	}
	//intersection along the the seg or extended seg
	ua := snap_to_zero_or_one(a / d)
	ub := snap_to_zero_or_one(b / d)
	return (0 <= ua && ua <= 1) && (0 <= ub && ub <= 1)
}
