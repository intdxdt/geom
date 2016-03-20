package segment

import (
    . "github.com/intdxdt/simplex/util/math"
    . "github.com/intdxdt/simplex/geom/point"
    . "github.com/intdxdt/simplex/geom/mbr"
)

const (
    x = iota
    y
)

type Segment struct {
    A *Point
    B *Point
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1, lnb0, lnb1
func (self *Segment) Intersects(other *Segment, extln bool) bool {
    _, ok := self.Intersection(other, extln)
    return ok
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func (self *Segment) Intersection(other *Segment, extln bool) ([]*Point, bool) {

    var x1, y1, x2, y2, x3, y3, x4, y4, d, a, b, ua, ub float64
    var coords = make([]*Point, 0)
    var zero_denum, bln, ua_0_1, ub_0_1 bool

    x1, y1 = self.A[x], self.A[y]
    x2, y2 = self.B[x], self.B[y]

    x3, y3 = other.A[x], other.A[y]
    x4, y4 = other.B[x], other.B[y]

    d = ((y4 - y3) * (x2 - x1)) - ((x4 - x3) * (y2 - y1))
    a = ((x4 - x3) * (y1 - y3)) - ((y4 - y3) * (x1 - x3))
    b = ((x2 - x1) * (y1 - y3)) - ((y2 - y1) * (x1 - x3))

    zero_denum = FloatEqual(d, 0.0)
    bln = (zero_denum && FloatEqual(a, 0.0) && FloatEqual(b, 0.0))
    if bln {
        abox := NewMBR(x1, y1, x2, y2)
        bbox := NewMBR(x3, y3, x4, y4)
        update_coords_inbounds(abox, x3, y3, x4, y4, &coords)
        update_coords_inbounds(bbox, x1, y1, x2, y2, &coords)
        return coords, bln
    }

    // Are the line coincident?
    if !zero_denum {
        // is the intersection along the the segments
        ua = a / d
        ub = b / d
        ua_0_1 = 0.0 <= ua  && ua <= 1.0
        ub_0_1 = 0.0 <= ub  && ub <= 1.0

        if ua_0_1 && ub_0_1 || extln {
            // instersection point is within range of lna && lnb ||  by extension
            bln = true
            pt := &Point{x1 + ua * (x2 - x1), y1 + ua * (y2 - y1)}
            if !contains_point(coords, pt) {
                coords = append(coords, pt)
            }
        }
    }

    return coords, bln
}

//updates coords that are in bounds
func update_coords_inbounds(bounds *MBR, x1, y1, x2, y2 float64, coords *[]*Point) {
    var a, b *Point

    if bounds.ContainsXY(x1, y1) {
        a = &Point{x1, y1}
    }
    if bounds.ContainsXY(x2, y2) {
        b = &Point{x2, y2}
    }

    if a != nil && !contains_point(*coords, a) {
        *coords = append(*coords, a)//a
    }
    if b != nil && !contains_point(*coords, b) {
        *coords = append(*coords, b)//b
    }
}

//linear search if point is a member of list of points
func contains_point(coords []*Point, pt *Point) bool {
    bln := false
    n := len(coords)
    for i := 0; !bln && i < n; i++ {
        bln = pt.Equals(coords[i])
    }
    return bln
}




