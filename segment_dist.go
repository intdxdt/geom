package geom

import (
    . "github.com/intdxdt/simplex/util/math"
    "math"
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

    denom = (y4 - y3) * (x2 - x1) - (x4 - x3) * (y2 - y1)
    numera = (x4 - x3) * (y1 - y3) - (y4 - y3) * (x1 - x3)
    numerb = (x2 - x1) * (y1 - y3) - (y2 - y1) * (x1 - x3)

    if math.Abs(denom) < Eps {
        is_aspt_a = self.A.Equals(self.B)
        is_aspt_b = other.A.Equals(other.B)

        if is_aspt_a && is_aspt_b {
            dist = self.A.Sub(other.B).Magnitude()
        } else if is_aspt_a || is_aspt_b {
            var ln *Segment

            if is_aspt_a {
                pta = self.A
                ln = other
            } else if is_aspt_b {
                pta = other.A
                ln = self
            }
            dist = ln.segpt_mindist(pta)
        } else {
            dist = math.Min(
                math.Min(
                    other.segpt_mindist(self.A),
                    other.segpt_mindist(self.B)),
                math.Min(
                    self.segpt_mindist(other.A),
                    self.segpt_mindist(other.B),
                ))
        }

    } else {
        mua = numera / denom
        mub = numerb / denom

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

            if pta != nil  && ptb != nil {
                dist = math.Min(
                    other.segpt_mindist(pta),
                    self.segpt_mindist(ptb),
                )
            } else if pta != nil {
                dist = other.segpt_mindist(pta)
            } else {
                dist = self.segpt_mindist(ptb)
            }
        } else {
            //lines intersect
            dist = 0.0
        }
    }

    return dist
}


//Minimum distance from segement to point
func (self *Segment) segpt_mindist(pt *Point) float64 {
    var dist = math.NaN()
    var cPt *Point
    var dln = self.B.Sub(self.A)
    var dx, dy = dln[x], dln[y]

    if dln.IsZero() {
        //line with zero length
        dist = pt.Sub(self.A).Magnitude()
    } else {
        var u = pt.Sub(self.A).DotProduct(dln) / dln.SquareMagnitude()

        if u < 0 {
            cPt = self.A
        } else if u > 1 {
            cPt = self.B
        } else {
            cPt = NewPointXY(self.A[x] + u * dx, self.A[y] + u * dy)
        }
        dist = pt.Sub(cPt).Magnitude()
    }

    return dist
}