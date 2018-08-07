package geom

import (
	"github.com/intdxdt/math"
)

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
func SegSegDistance(sa, sb, oa, ob *Point) float64 {
	var dist = math.NaN()
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
			dist = hypot(x1-x4, y1-y4)
		} else if is_aspt_a || is_aspt_b {
			var lna, lnb *Point

			if is_aspt_a {
				pta = sa
				lna, lnb = oa, ob
			} else if is_aspt_b {
				pta = oa
				lna, lnb = sa, sb
			}
			dist = DistanceToPoint(lna, lnb, pta)
		} else {
			dist = minf64(
				minf64(
					DistanceToPoint(oa, ob, sa),
					DistanceToPoint(oa, ob, sb)),
				minf64(
					DistanceToPoint(sa, sb, oa),
					DistanceToPoint(sa, sb, ob),
				))
		}

	} else {
		//if close to zero or one , snap
		mua = numera / denom
		if (mua == 0) || math.Abs(mua) < math.EPSILON { //a == b || Abs(a - b) < EPSILON
			mua = 0
		} else if (mua == 1) || math.Abs(mua-1) < math.EPSILON {
			mua = 1
		}

		mub = numerb / denom
		if (mub == 0) || math.Abs(mub) < math.EPSILON {
			mub = 0
		} else if (mub == 1) || math.Abs(mub-1) < math.EPSILON {
			mub = 1
		}

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
					DistanceToPoint(oa, ob, pta),
					DistanceToPoint(sa, sb, ptb),
				)
			} else if pta != nil {
				dist = DistanceToPoint(oa, ob, pta)
			} else {
				dist = DistanceToPoint(sa, sb, ptb)
			}
		} else {
			//lines intersect
			dist = 0
		}
	}

	return dist
}

//Minimum distance from segement to point
func (self *Segment) DistanceToPoint(pt *Point) float64 {
	return DistanceToPoint(self.A(), self.B(), pt)
}

func DistanceToPoint(sa, sb, pt *Point) float64 {
	var dist = math.NaN()
	//var cPt *Point
	var ax, ay = sa[X], sa[Y]
	var bx, by = sb[X], sb[Y]
	var px, py = pt[X], pt[Y]
	//var dab = sb.Sub(sa)
	var dx, dy = bx-ax, by-ay
	//a == b || Abs(a - b) < EPSILON
	var isz_x = (dx == 0) || math.Abs(dx) < math.EPSILON
	var isz_y = (dy == 0) || math.Abs(dy) < math.EPSILON

	if isz_x && isz_y {
		//line with zero length
		dist = hypot(px-ax, py-ay)
	} else {
		var cPtx, cPty float64
		var u = (((px - ax) * dx) + ((py - ay) * dy)) / (dx*dx + dy*dy)

		if u < 0 {
			cPtx, cPty = ax, ay
		} else if u > 1 {
			cPtx, cPty = bx, by
		} else {
			cPtx, cPty = ax+u*dx, ay+u*dy
		}
		dist = hypot(px-cPtx, py-cPty)
	}

	return dist
}
