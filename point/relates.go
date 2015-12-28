package point

import (
	util "github.com/intdxdt/simplex/util/math"
)

//Equals evaluates whether two points are the same
func (pt *Point) Equals(point Point) bool{
	return (util.Float_equal(pt[x], point[x]) &&
		util.Float_equal(pt[y], point[y]))
}

//Intersects evaluates whether two points are the same
func (pt *Point) Intersects(point Point) bool{
	return pt.Equals(point)
}

//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(point Point) bool{
	return !(pt.Intersects(point))
}

