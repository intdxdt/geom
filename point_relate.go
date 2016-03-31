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

//SideOf point (Left|On|Right : -1, 0, 1 ) to an infinite line through a and b
//Input:  two points a, b forming begin and end of line segment
//Return: Orient with side value of
//        -1 pt is left of the line through a and b
//         0 pt on the line
//         1 pt right of the line
func (pt *Point) SideOf(a, b *Point) *Orient{
    v := pt.CrossProduct(a, b);
    var o  = &Orient{'o'};
    if v > 0 {
        o.side = 'l'
    } else if v < 0 {
        o.side = 'r'
    }
    return o
}




