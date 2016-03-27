package geom

import (
    . "github.com/intdxdt/simplex/geom/point"
    . "github.com/intdxdt/simplex/util/math"
    "math"
)
//Distance betwen two segments
func (self *Segment) Distance(other *Segment) float64{
    var x1, y1 = (*self.A)[x], (*self.A)[y]
    var x2, y2 = (*self.B)[x], (*self.B)[y]

    var x3, y3 = (*other.A)[x], (*other.A)[y]
    var x4, y4 = (*other.B)[x], (*other.B)[y]

    var dist, mua, mub float64
    var denom, numera, numerb float64
    var aminx, amaxx, aminy, amaxy, bminx, bmaxx, bminy, bmaxy float64
    var dx, dy *float64
    var pta, ptb *Point

    var eps = 1e-10

    denom = (y4 - y3) * (x2 - x1) - (x4 - x3) * (y2 - y1)
    numera = (x4 - x3) * (y1 - y3) - (y4 - y3) * (x1 - x3)
    numerb = (x2 - x1) * (y1 - y3) - (y2 - y1) * (x1 - x3)

    if math.Abs(denom) < eps {
        /* are the lines parallel */
        aminx, amaxx = math.Min(x1, x2), math.Max(x1, x2)
        aminy, amaxy = math.Min(y1, y2), math.Max(y1, y2)

        bminx, bmaxx = math.Min(x3, x4), math.Max(x3, x4)
        bminy, bmaxy = math.Min(y3, y4), math.Max(y3, y4)

        if amaxx < bminx {
            vx := bminx - amaxx
            dx = &vx
        } else if aminx > bmaxx {
            vx := aminx - bmaxx
            dx = &vx
        }

        if amaxy < bminy {
            vy := bminy - amaxy
            dy = &vy
        }  else if aminy > bmaxy {
            vy := aminy - bmaxy
            dy = &vy
        }

        if dx == nil && dy == nil {
            /*calculate the perpendicular distance*/
            var m = (y2 - y1) / (x2 - x1)
            var dc = (y1 - m * x1) - (y3 - m * x3)

            if dc == 0.0 {
                dist = 0.0
            } else {
                dist = math.Abs(dc) / math.Sqrt(m * m + 1)
            }
        } else if (aminx == amaxx && aminy == amaxy) || (bminx == bmaxx && bminy == bmaxy) {
            var isa = (aminx == amaxx && aminy == amaxy)
            var isb = (bminx == bmaxx && bminy == bmaxy)
            var ln *Segment

            if isa {
                pta = self.A
                ln = other
            } else if isb {
                pta = other.A
                ln = self
            }
            dist = ln.DistanceToPoint(pta)
        } else {
            var _dx, _dy = 0.0, 0.0
            if (dx != nil) {
                _dx = *dx
            }
            if (dy != nil) {
                _dy = *dy
            }
            dist = math.Hypot(_dx, _dy)
        }
    } else {
        /* the is intersection along the the segments */
        mua = numera / denom
        mub = numerb / denom

        if mua < 0 || mua > 1 || mub < 0 || mub > 1 {
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
                    other.DistanceToPoint(pta),
                    self.DistanceToPoint(ptb),
                )
            } else if pta != nil {
                dist = other.DistanceToPoint(pta)
            } else {
                dist = self.DistanceToPoint(ptb)
            }
        }  else {
            dist = 0.0
        }
    }

    return dist
}


//Minimum distance from segement to point
func (self *Segment) DistanceToPoint(pt *Point) float64 {

    var cPt *Point
    var dPt = self.B.Sub(self.A)
    var dx, dy = (*dPt)[x], (*dPt)[y]

    if FloatEqual(dx, 0.0) &&  FloatEqual(dy, 0.0) {
        //line with zero length
        return pt.Distance(self.A)
    }
    var u = pt.Sub(self.A).DotProduct(dPt) / dPt.SquareMagnitude()

    if u < 0 {
        cPt = self.A
    } else if u > 1 {
        cPt = self.B
    } else {
        cPt = NewPointXY((*self.A)[x] + u * dx, (*self.A)[y] + u * dy)
    }
    return pt.Distance(cPt)
}

