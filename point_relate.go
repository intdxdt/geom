package geom

import (
    util "github.com/intdxdt/simplex/util/math"
)

//Equals evaluates whether two points are the same
func (pt *Point) Equals(point *Point) bool {
    return (
    util.FloatEqual(pt[x], point[x]) &&
    util.FloatEqual(pt[y], point[y]))
}

//Intersects evaluates whether two points are the same
func (pt *Point) Intersects(point *Point) bool {
    return pt.Equals(point)
}

//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(point *Point) bool {
    return !(pt.Intersects(point))
}


//linear search if point is a member of list of points
func is_point_inlist(coords []*Point, pt *Point) bool {
    bln := false
    n := len(coords)
    for i := 0; !bln && i < n; i++ {
        bln = pt.Equals(coords[i])
    }
    return bln
}

