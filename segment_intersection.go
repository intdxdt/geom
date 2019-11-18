package geom

import (
	"sort"
)

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func SegSegIntersection(sa, sb, oa, ob *Point) []InterPoint {
	var coords []InterPoint
	var a = ((ob[0] - oa[0]) * (sa[1] - oa[1])) - ((ob[1] - oa[1]) * (sa[0] - oa[0]))
	var b = ((sb[0] - sa[0]) * (sa[1] - oa[1])) - ((sb[1] - sa[1]) * (sa[0] - oa[0]))
	var d = ((ob[1] - oa[1]) * (sb[0] - sa[0])) - ((ob[0] - oa[0]) * (sb[1] - sa[1]))

	//snap to zero if near -0 or 0
	a, b, d = snap_to_zero(a), snap_to_zero(b), snap_to_zero(d)

	// are the line coincident?
	if d == 0 {
		return coincidentSegs(sa, sb, oa, ob, coords, a, b)
	}

	// is the intersection along the the segments
	var ua, ub = a / d, b / d
	ua = snap_to_zero_or_one(ua)
	ub = snap_to_zero_or_one(ub)

	var ua_0_1 = 0 <= ua && ua <= 1
	var ub_0_1 = 0 <= ub && ub <= 1

	if ua_0_1 && ub_0_1 {
		coords = append(coords, InterPoint{
			Pt(sa[X]+ua*(sb[X]-sa[X]), sa[Y]+ua*(sb[Y]-sa[Y])),
			interRelation(ua, ub),
		})
	}
	return coords
}

func interRelation(ua, ub float64) VBits {
	var sa, sb, oa, ob VBits

	if ua == 0 {
		sa = SelfA
	} else if ua == 1 {
		sb = SelfB
	}

	if ub == 0 {
		oa = OtherA
	} else if ub == 1 {
		ob = OtherB
	}

	return sa | sb | oa | ob
}

func coincidentSegs(sa, sb, oa, ob *Point, coords []InterPoint, a, b float64) []InterPoint {
	if a == 0 && b == 0 {
		var s_minx, s_miny, s_maxx, s_maxy = bounds(sa, sb)
		var o_minx, o_miny, o_maxx, o_maxy = bounds(oa, ob)
		if intersects(s_minx, s_miny, s_maxx, s_maxy, o_minx, o_miny, o_maxx, o_maxy) {
			updateCoordsInbounds(o_minx, o_miny, o_maxx, o_maxy, sa, &coords, SelfA)
			updateCoordsInbounds(o_minx, o_miny, o_maxx, o_maxy, sb, &coords, SelfB)
			updateCoordsInbounds(s_minx, s_miny, s_maxx, s_maxy, oa, &coords, OtherA)
			updateCoordsInbounds(s_minx, s_miny, s_maxx, s_maxy, ob, &coords, OtherB)
		}
	}

	//lexical sort
	sort.Sort(IntPts(coords))

	var points []InterPoint
	var last = false
	var n = len(coords) - 1

	for idx := 0; idx < n; idx++ { //O(n)
		var i, j = idx, idx + 1
		var pt = coords[i]
		for i < n && coords[i].Equals2D(&coords[j].Point) {
			coords[j].Inter = coords[i].Inter | coords[j].Inter
			last = j == n
			pt = coords[j]
			i = j
			j = i + 1
		}
		idx = i
		points = append(points, pt)
	}

	if n >= 0 && !last {
		points = append(points, coords[n])
	}
	return points
}

//updates Coords that are in bounds
func updateCoordsInbounds(minx, miny, maxx, maxy float64, point *Point, intpts *[]InterPoint, vbits VBits) {
	if containsXY(minx, miny, maxx, maxy, point[X], point[Y]) {
		*intpts = append(*intpts, InterPoint{*point, vbits})
	}
}
