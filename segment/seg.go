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

func NewSegment(a, b *Point) *Segment{
    return &Segment{a, b}
}
//do two lines intersect line segments a && b with
//vertices lna0, lna1, lnb0, lnb1
func (self *Segment) Intersects(other *Segment, extln bool) bool {
    var bln = false
    var a, b, d,
    x1, y1, x2, y2,
    x3, y3, x4, y4 = seg_intersect_abdxy(self, other)

    //snap to zero if near -0 or 0
    snap_to_zero(&a)
    snap_to_zero(&b)
    snap_to_zero(&d)

    if d == 0 {
        if a == 0.0 && b == 0.0 {
            abox := NewMBR(x1, y1, x2, y2)
            bbox := NewMBR(x3, y3, x4, y4)
            bln = abox.Intersects(bbox)
        }
        return bln
    }
    //intersection along the the seg or extended seg
    ua := a / d
    ub := b / d
    ua_0_1 := (0.0 <= ua  && ua <= 1.0)
    ub_0_1 := (0.0 <= ub  && ub <= 1.0)
    bln = ua_0_1 && ub_0_1 || extln
    return bln
}

//do two lines intersect line segments a && b with
//vertices lna0, lna1 and lnb0, lnb1
func (self *Segment) Intersection(other *Segment, extln bool) ([]*Point, bool) {
    var coords = make([]*Point, 0)
    var bln = false
    var a, b, d,
    x1, y1, x2, y2,
    x3, y3, x4, y4 = seg_intersect_abdxy(self, other)

    //snap to zero if near -0 or 0
    snap_to_zero(&a)
    snap_to_zero(&b)
    snap_to_zero(&d)

    // Are the line coincident?
    if d == 0 {
        if a == 0 && b == 0 {
            abox := NewMBR(x1, y1, x2, y2)
            bbox := NewMBR(x3, y3, x4, y4)
            if abox.Intersects(bbox) {
                update_coords_inbounds(abox, x3, y3, x4, y4, &coords)
                update_coords_inbounds(bbox, x1, y1, x2, y2, &coords)
            }
        }
        bln = (len(coords) > 0)
        return coords, bln
    }
    // is the intersection along the the segments
    ua := a / d
    ub := b / d
    ua_0_1 := 0.0 <= ua  && ua <= 1.0
    ub_0_1 := 0.0 <= ub  && ub <= 1.0

    if ua_0_1 && ub_0_1 || extln {
        // instersection point is within range of lna && lnb ||  by extension
        bln = true
        pt := &Point{x1 + ua * (x2 - x1), y1 + ua * (y2 - y1)}
        if !contains_point(coords, pt) {
            coords = append(coords, pt)
        }
    }

    return coords, bln
}

func seg_intersect_abdxy(self, other *Segment) (float64, float64, float64,
float64, float64, float64, float64,
float64, float64, float64, float64) {

    var x1, y1, x2, y2, x3, y3, x4, y4, d, a, b float64

    x1, y1 = self.A[x], self.A[y]
    x2, y2 = self.B[x], self.B[y]

    x3, y3 = other.A[x], other.A[y]
    x4, y4 = other.B[x], other.B[y]

    d = ((y4 - y3) * (x2 - x1)) - ((x4 - x3) * (y2 - y1))
    a = ((x4 - x3) * (y1 - y3)) - ((y4 - y3) * (x1 - x3))
    b = ((x2 - x1) * (y1 - y3)) - ((y2 - y1) * (x1 - x3))
    return a, b, d, x1, y1, x2, y2, x3, y3, x4, y4
}

//clamp to zero if float is near zero
func snap_to_zero(v *float64) {
    if FloatEqual(*v, 0.0) {
        *v = 0.0
    }
}

//updates coords that are in bounds
func update_coords_inbounds(bounds *MBR,
x1, y1, x2, y2 float64, coords *[]*Point) {

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




