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
	var mua, mub float64
	var denom, numera, numerb float64
	var is_aspt_a, is_aspt_b bool
	var pta, ptb *Point

	denom = (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	numera = (x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)
	numerb = (x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)

	if math.Abs(denom) < math.EPSILON {
		is_aspt_a = sa.Equals2D(sb)
		is_aspt_b = oa.Equals2D(ob)

		if is_aspt_a && is_aspt_b {
			//dist = sa.Magnitude(ob)
			dist = math.Hypot(x1-x4, y1-y4)
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
			dist = math.MinF64(
				math.MinF64(
					DistanceToPoint(oa, ob, sa),
					DistanceToPoint(oa, ob, sb)),
				math.MinF64(
					DistanceToPoint(sa, sb, oa),
					DistanceToPoint(sa, sb, ob),
				))
		}

	} else {
		//if close to zero or one , snap
		mua = snap_to_zero_or_one(numera / denom)
		mub = snap_to_zero_or_one(numerb / denom)

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
				dist = math.MinF64(
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
			dist = 0.0
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
	var isz_x = feq(dx, 0)
	var isz_y = feq(dy, 0)

	if isz_x && isz_y {
		//line with zero length
		//dist = pt.Magnitude(sa)
		dist = math.Hypot(px-ax, py-ay)
	} else {
		var cPtx, cPty float64
		//var dx, dy = dab[X], dab[Y]
		var pax, pay = px-ax, py-ay
		//(pax * dx) + (pay * dy)
		//var u = pt.Sub(sa).DotProduct(dab) / (dx*dx + dy*dy)
		var u = ((pax * dx) + (pay * dy)) / (dx*dx + dy*dy)

		if u < 0 {
			//cPt = sa
			cPtx, cPty = ax, ay
		} else if u > 1 {
			//cPt = sb
			cPtx, cPty = bx, by
		} else {
			//cPt = PointXY(sa[X]+u*dx, sa[Y]+u*dy)
			cPtx, cPty = ax+u*dx, ay+u*dy
		}
		//dist = pt.Magnitude(cPt)
		dist = math.Hypot(px-cPtx, py-cPty)
	}

	return dist
}
