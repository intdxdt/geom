package geom

import (
    . "github.com/intdxdt/simplex/util/math"
    . "github.com/intdxdt/simplex/struct/item"
)

//Equals evaluates whether two points are the same
func (pt *Point) Equals(point *Point) bool {
    return (
        FloatEqual(pt[x], point[x]) &&
            FloatEqual(pt[y], point[y]))
}


//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(point *Point) bool {
    return !(pt.Intersects(point))
}

//SideOf point (Left|On|Right : -1, 0, 1 ) to an infinite line through a and b
//Input:  two points a, b forming begin and end of line segment
//Return: Side Obj with Side.s :
//        -1 pt is left of the line through a and b
//         0 pt on the line
//         1 pt right of the line
func (pt *Point) SideOf(a, b *Point) *Side {
    v := pt.CrossProduct(a, b)

    if FloatEqual(v, 0.0) {
        v = 0.0
    }
    var o = NewSide()
    if v > 0 {
        o.AsLeft()
    } else if v < 0 {
        o.AsRight()
    }
    return o
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


//compare points as items - x & y ordering
//func (self *Point) Compare(o Item) int {
//    pt := o.(*Point)
//    dx := self[x] - pt[x]
//    dy := self[y] - pt[y]
//
//    //check if close enougth ot zero
//    if FloatEqual(dx, 0.0) && FloatEqual(dy, 0.0) {
//        return 0
//    } else if dx < 0 && dy < 0 {
//        return -1
//    } else if dy > 0 && dy > 0 {
//        return 1
//    }
//    //check lexical
//    d := self[x] - pt[x]
//    if FloatEqual(self[x], pt[x]) {
//        //x's are close enougth to each other
//        d = self[y] - pt[y]
//    } else {
//        d = self[x] - pt[x];
//    }
//    if FloatEqual(d, 0.0) {
//        //check if close enougth ot zero
//        return 0
//    } else if d < 0 {
//        return -1
//    }
//    return 1
//}


