package geom

import (
	"simplex/util/math"
)

//Distance betwen two segments
func (self *Segment) Distance(other *Segment) float64 {
	var dist = math.NaN()
	var x1, y1 = self.A[x], self.A[y]
	var x2, y2 = self.B[x], self.B[y]

	var x3, y3 = other.A[x], other.A[y]
	var x4, y4 = other.B[x], other.B[y]
	var mua, mub float64
	var denom, numera, numerb float64
	var is_aspt_a, is_aspt_b bool
	var pta, ptb *Point

	denom = (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	numera = (x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)
	numerb = (x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)

	if math.Abs(denom) < math.Eps {
		is_aspt_a = self.A.Equals2D(self.B)
		is_aspt_b = other.A.Equals2D(other.B)

		if is_aspt_a && is_aspt_b {
			dist = self.A.Magnitude(other.B)
		} else if is_aspt_a || is_aspt_b {
			var ln *Segment

			if is_aspt_a {
				pta = self.A
				ln = other
			} else if is_aspt_b {
				pta = other.A
				ln = self
			}
			dist = ln.DistanceToPoint(pta)
		} else {
			dist = math.MinF64(
				math.MinF64(
					other.DistanceToPoint(self.A),
					other.DistanceToPoint(self.B)),
				math.MinF64(
					self.DistanceToPoint(other.A),
					self.DistanceToPoint(other.B),
				))
		}

	} else {
		//if close to zero or one , snap
		mua = snap_to_zero_or_one(numera / denom)
		mub = snap_to_zero_or_one(numerb / denom)

		if mua < 0 || mua > 1 || mub < 0 || mub > 1 {
			//the is intersection along the the segments
			if mua < 0 {
				pta = self.A
			} else if mua > 1 {
				pta = self.B
			}

			if mub < 0 {
				ptb = other.A
			} else if mub > 1 {
				ptb = other.B
			}

			if pta != nil && ptb != nil {
				dist = math.MinF64(
					other.DistanceToPoint(pta),
					self.DistanceToPoint(ptb),
				)
			} else if pta != nil {
				dist = other.DistanceToPoint(pta)
			} else {
				dist = self.DistanceToPoint(ptb)
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
	var dist = math.NaN()
	var cPt *Point
	var dab = self.B.Sub(self.A)

	if dab.IsZero() {
		//line with zero length
		dist = pt.Magnitude(self.A)
	} else {
		var dx, dy = dab[x], dab[y]
		var u = pt.Sub(self.A).DotProduct(dab) / (dx*dx + dy*dy)

		if u < 0 {
			cPt = self.A
		} else if u > 1 {
			cPt = self.B
		} else {
			cPt = NewPointXY(self.A[x]+u*dx, self.A[y]+u*dy)
		}
		dist = pt.Magnitude(cPt)
	}

	return dist
}
