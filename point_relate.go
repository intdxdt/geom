package geom

import (
    . "simplex/util/math"
    . "simplex/struct/item"
    . "simplex/side"
    "simplex/cart2d"
)

//Equals evaluates whether two points are the same
func (pt *Point) Equals(point *Point) bool {
    return FloatEqual(pt[x], point[x]) &&
        FloatEqual(pt[y], point[y]) &&
        FloatEqual(pt[z], point[z])
}


//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(point *Point) bool {
    return !(pt.Intersects(point))
}

//compare points as items - x | y ordering
func (self *Point) Compare(o Item) int {
    pt := o.(*Point)
    d := self[x] - pt[x]
    if FloatEqual(d, 0.0) {
        //x's are close enougth to each other
        d = self[y] - pt[y]
    }

    if FloatEqual(d, 0.0) {
        //check if close enougth ot zero
        return 0
    } else if d < 0 {
        return -1
    }
    return 1
}

//position of C relative to line AB
func (c *Point) SideOf(a, b *Point) *Side {
    s := NewSide()
    ccw := cart2d.CCW(a, b, c)
    if FloatEqual(ccw, 0) {
        s.AsOn()
    } else if (ccw > 0) {
        s.AsLeft()
    } else if (ccw < 0) {
        s.AsRight()
    }
    return s
}