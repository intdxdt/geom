package geom

import (
	"sort"
	"github.com/intdxdt/math"
)

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func SegSegIntersection(sa, sb, oa, ob *Point) []InterPoint {
	var coords []InterPoint
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

	// Are the line coincident?
	if d == 0 {
		return coincidentSegs(sa, sb, oa, ob, coords, a, b)
	}

	// is the intersection along the the segments
	var ua = a / d
	if ua == 0 || math.Abs(ua) < math.EPSILON { //feq(ua, 0.0) {
		ua = 0
	} else if ua == 1 || math.Abs(ua-1) < math.EPSILON { //feq(ua, 1.0) {
		ua = 1
	}

	var ub = b / d
	if ub == 0 || math.Abs(ub) < math.EPSILON { //feq(ub, 0.0) {
		ub = 0
	} else if ub == 1 || math.Abs(ub-1) < math.EPSILON { //feq(ub, 1.0) {
		ub = 1
	}

	var ua_0_1 = 0.0 <= ua && ua <= 1.0
	var ub_0_1 = 0.0 <= ub && ub <= 1.0

	if ua_0_1 && ub_0_1 {
		coords = append(coords, InterPoint{
			Pt(sa[X]+ua*(sb[X]-sa[X]), sa[Y]+ua*(sb[Y]-sa[Y])),
			interRelation(ua, ub),
		})
	}
	return coords
}

func segsegABD(sa, sb, oa, ob *Point) (float64, float64, float64) {
	var x1, y1, x2, y2, x3, y3, x4, y4, d, a, b float64

	x1, y1 = sa[X], sa[Y]
	x2, y2 = sb[X], sb[Y]

	x3, y3 = oa[X], oa[Y]
	x4, y4 = ob[X], ob[Y]

	d = ((y4 - y3) * (x2 - x1)) - ((x4 - x3) * (y2 - y1))
	a = ((x4 - x3) * (y1 - y3)) - ((y4 - y3) * (x1 - x3))
	b = ((x2 - x1) * (y1 - y3)) - ((y2 - y1) * (x1 - x3))

	return a, b, d
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
		var i, j = idx, idx+1
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
