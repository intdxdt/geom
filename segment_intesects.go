package geom

//do two lines intersect line segments a && b with
//vertices sa, sb, oa, ob
func SegSegIntersects(sa, sb, oa, ob *Point) bool {
	var bln = false
	var a = ((ob[0] - oa[0]) * (sa[1] - oa[1])) - ((ob[1] - oa[1]) * (sa[0] - oa[0]))
	var b = ((sb[0] - sa[0]) * (sa[1] - oa[1])) - ((sb[1] - sa[1]) * (sa[0] - oa[0]))
	var d = ((ob[1] - oa[1]) * (sb[0] - sa[0])) - ((ob[0] - oa[0]) * (sb[1] - sa[1]))

	//snap to zero if near -0 or 0
	a, b, d = snap_to_zero(a), snap_to_zero(b), snap_to_zero(d)

	if d == 0 {
		if a == 0 && b == 0 {
			bln = boundsIntersects(sa, sb, oa, ob)
		}
		return bln
	}

	var ua, ub = a / d, b / d
	ua, ub = snap_to_zero_or_one(ua), snap_to_zero_or_one(ub)

	return (0 <= ua && ua <= 1) && (0 <= ub && ub <= 1)
}
