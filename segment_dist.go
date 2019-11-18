package geom

import (
	"github.com/intdxdt/math"
)

type hypotFunc func(float64, float64) float64

//Length of segment
func (self *Segment) Length() float64 {
	var a, b = self.A(), self.B()
	return math.Hypot(a[X]-b[X], a[Y]-b[Y])
}

//Distance betwen two segments
func (self *Segment) SegSegDistance(other *Segment) float64 {
	return SegSegDistance(self.A(), self.B(), other.A(), other.B())
}

//Distance betwen two segments
func (self *Segment) SquareSegSegDistance(other *Segment) float64 {
	return SquareSegSegDistance(self.A(), self.B(), other.A(), other.B())
}

//Minimum distance from segement to point
func (self *Segment) DistanceToPoint(pt *Point) float64 {
	return DistanceToPoint(self.A(), self.B(), pt)
}

//Minimum square distance from segement to point
func (self *Segment) SquareDistanceToPoint(pt *Point) float64 {
	return SquareDistanceToPoint(self.A(), self.B(), pt)
}

//Distance betwen two segments
func SegSegDistance(sa, sb, oa, ob *Point) float64 {
	return segsegDistance(sa, sb, oa, ob, hypot)
}

//Distance betwen two segments
func SquareSegSegDistance(sa, sb, oa, ob *Point) float64 {
	return segsegDistance(sa, sb, oa, ob, hypotSqr)
}

//Distance from segment endpoints to point
func DistanceToPoint(sa, sb, pt *Point) float64 {
	return distanceToPoint(sa, sb, pt, hypot)
}

//Square Distance from segment endpoints to point
func SquareDistanceToPoint(sa, sb, pt *Point) float64 {
	return distanceToPoint(sa, sb, pt, hypotSqr)
}

//Distance betwen two segments with custom hypot function
func segsegDistance(sa, sb, oa, ob *Point, hypotFn hypotFunc) float64 {
	var dist = nan
	var x1, y1 = sa[X], sa[Y]
	var x2, y2 = sb[X], sb[Y]

	var x3, y3 = oa[X], oa[Y]
	var x4, y4 = ob[X], ob[Y]

	var pta, ptb *Point
	var mua, mub float64
	var is_aspt_a, is_aspt_b bool
	var denom, numera, numerb float64

	denom = (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	numera = (x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)
	numerb = (x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)

	if math.Abs(denom) < math.EPSILON {
		is_aspt_a = sa.Equals2D(sb)
		is_aspt_b = oa.Equals2D(ob)

		if is_aspt_a && is_aspt_b {
			dist = hypotFn(x1-x4, y1-y4)
		} else if is_aspt_a || is_aspt_b {
			var lna, lnb *Point

			if is_aspt_a {
				pta = sa
				lna, lnb = oa, ob
			} else if is_aspt_b {
				pta = oa
				lna, lnb = sa, sb
			}
			dist = distanceToPoint(lna, lnb, pta, hypotFn)
		} else {
			dist = minDistSegmentEndpoints(sa, sb, oa, ob, hypotFn)
		}

	} else {
		mua = numera / denom
		mua = snap_to_zero_or_one(mua)

		mub = numerb / denom
		mub = snap_to_zero_or_one(mub)

		if mua < 0 || mua > 1 || mub < 0 || mub > 1 {
			//the is intersection along the the segments
			if mua < 0 {
				pta = sa
			} else if mua > 1 {
				pta = sb
			}

			if mub < 0 {
				ptb = oa
			} else if mub > 1 {
				ptb = ob
			}

			if pta != nil && ptb != nil {
				dist = minf64(
					distanceToPoint(oa, ob, pta, hypotFn),
					distanceToPoint(sa, sb, ptb, hypotFn),
				)
			} else if pta != nil {
				dist = distanceToPoint(oa, ob, pta, hypotFn)
			} else {
				dist = distanceToPoint(sa, sb, ptb, hypotFn)
			}
		} else {
			dist = 0 //lines intersect
		}
	}

	return dist
}

func minDistSegmentEndpoints(sa, sb, oa, ob *Point, fn hypotFunc) float64 {
	var o_sa = distanceToPoint(oa, ob, sa, fn)
	var o_sb = distanceToPoint(oa, ob, sb, fn)
	var s_oa = distanceToPoint(sa, sb, oa, fn)
	var s_ob = distanceToPoint(sa, sb, ob, fn)
	return minf64(minf64(o_sa, o_sb), minf64(s_oa, s_ob))
}

//Distance from segment endpoints to point
func distanceToPoint(sa, sb, pt *Point, hypotFunc func(float64, float64) float64) float64 {
	var dist = nan
	var cPtx, cPty, u float64
	var ax, ay = sa[X], sa[Y]
	var bx, by = sb[X], sb[Y]
	var px, py = pt[X], pt[Y]
	var dx, dy = bx - ax, by - ay
	var isz_x = feq(dx, 0)
	var isz_y = feq(dy, 0)

	if isz_x && isz_y {
		//line with zero length
		dist = hypotFunc(px-ax, py-ay)
	} else {
		u = (((px - ax) * dx) + ((py - ay) * dy)) / (dx*dx + dy*dy)

		if u < 0 {
			cPtx, cPty = ax, ay
		} else if u > 1 {
			cPtx, cPty = bx, by
		} else {
			cPtx, cPty = ax+u*dx, ay+u*dy
		}
		dist = hypotFunc(px-cPtx, py-cPty)
	}

	return dist
}
