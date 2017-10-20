package geom

import (
	"simplex/side"
	"github.com/intdxdt/cart"
	"github.com/intdxdt/math"
)

//Equals evaluates whether two points are the same
func (pt *Point) Equals2D(point *Point) bool {
	return math.FloatEqual(pt[X], point[X]) && math.FloatEqual(pt[Y], point[Y])
}

func (pt *Point) Equals3D(point *Point) bool {
	return math.FloatEqual(pt[X], point[X]) && math.FloatEqual(pt[Y], point[Y]) &&
		math.FloatEqual(pt[Z], point[Z])
}

//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(point *Point) bool {
	return !(pt.Intersects(point))
}

//compare points as items - x | y ordering
func (self *Point) Compare(pt *Point) int {
	d := self[X] - pt[X]
	if math.FloatEqual(d, 0.0) {
		//x's are close enougth to each other
		d = self[Y] - pt[Y]
	}

	if math.FloatEqual(d, 0.0) {
		//check if close enougth ot zero
		return 0
	} else if d < 0 {
		return -1
	}
	return 1
}

//position of C relative to line AB
func (c *Point) SideOf(a, b *Point) *side.Side {
	s := side.NewSide()
	ccw := cart.Orientation2D(a, b, c)
	if ccw ==  0 {
		s.AsOn()
	} else if ccw < 0 {
		s.AsLeft()
	} else if ccw > 0 {
		s.AsRight()
	}
	return s
}
