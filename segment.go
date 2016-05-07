package geom

import (
    . "simplex/util/math"
    . "simplex/geom/mbr"
    "simplex/struct/sset"
    "simplex/struct/item"
)

type Segment struct {
    A *Point
    B *Point
}
//New Segment constructor
func NewSegment(a, b *Point) *Segment {
    return &Segment{a, b}
}

//Side of pt to segement
func (self *Segment)SideOf(pt *Point) *Side {
    return pt.SideOf(self.A, self.B)
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
    var set = sset.NewSSet()
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
                update_coords_inbounds(abox, x3, y3, x4, y4, set)
                update_coords_inbounds(bbox, x1, y1, x2, y2, set)
            }
        }
        set.Each(func(o item.Item) {
            coords = append(coords, o.(*Point))
        })
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
        if !InCoordinates(coords, pt) {
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
func update_coords_inbounds(bounds *MBR, x1, y1, x2, y2 float64, set *sset.SSet) {
    if bounds.ContainsXY(x1, y1) {
        set.Add(&Point{x1, y1})
    }
    if bounds.ContainsXY(x2, y2) {
        set.Add(&Point{x2, y2})
    }
}








