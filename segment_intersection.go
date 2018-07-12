package geom

import (
	"sort"
	"github.com/intdxdt/mbr"
)

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func SegSegIntersection(sa, sb, oa, ob *Point) []*InterPoint {
	var coords []*InterPoint
	var a, b, d = segseg_abd(sa[:], sb[:], oa[:], ob[:])

	//snap to zero if near -0 or 0
	a = snap_to_zero(a)
	b = snap_to_zero(b)
	d = snap_to_zero(d)

	// Are the line coincident?
	if d == 0 {
		return coincident_segs(sa, sb, oa, ob, coords, a, b)
	}

	// is the intersection along the the segments
	var ua = snap_to_zero_or_one(a / d)
	var ub = snap_to_zero_or_one(b / d)

	var ua_0_1 = 0.0 <= ua && ua <= 1.0
	var ub_0_1 = 0.0 <= ub && ub <= 1.0

	if ua_0_1 && ub_0_1 {
		var pt = &InterPoint{
			Point: PointXY(
				sa[X]+ua*(sb[X]-sa[X]),
				sa[Y]+ua*(sb[Y]-sa[Y]),
			),
			Inter: interRelation(ua, ub),
		}
		coords = append(coords, pt)
	}
	return coords
}

func segseg_abd(sa, sb, oa, ob []float64) (float64, float64, float64) {
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

func coincident_segs(sa, sb, oa, ob *Point, coords []*InterPoint, a, b float64) []*InterPoint {
	if a == 0 && b == 0 {
		var selfBox = BBox(sa, sb)
		var otherBox = BBox(oa, ob)
		if selfBox.Intersects(&otherBox) {
			update_coords_inbounds(&otherBox, sa, &coords, SelfA)
			update_coords_inbounds(&otherBox, sb, &coords, SelfB)
			update_coords_inbounds(&selfBox, oa, &coords, OtherA)
			update_coords_inbounds(&selfBox, ob, &coords, OtherB)
		}
	}
	//lexical sort
	sort.Sort(IntPts(coords))

	var points []*InterPoint
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

//updates coords that are in bounds
func update_coords_inbounds(bounds *mbr.MBR, point *Point, intpts *[]*InterPoint, vbits VBits) {
	if bounds.ContainsXY(point[X], point[Y]) {
		*intpts = append(*intpts, &InterPoint{*point, vbits})
	}
}
