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


//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(point *Point) bool {
    return !(pt.Intersects(point))
}



