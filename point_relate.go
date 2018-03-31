package geom

import (
	"simplex/side"
	"github.com/intdxdt/cart"
	"github.com/intdxdt/math"
)

//Equals evaluates whether two points are the same
func (pt *Point) Equals2D(o *Point) bool {
	return math.FloatEqual(pt[X], o[X]) &&
		math.FloatEqual(pt[Y], o[Y])
}

func (pt *Point) Equals3D(o *Point) bool {
	return math.FloatEqual(pt[X], o[X]) &&
		math.FloatEqual(pt[Y], o[Y]) &&
		math.FloatEqual(pt[Z], o[Z])
}

//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(o *Point) bool {
	return !pt.Equals2D(o)
}

//compare points as items - x | y ordering
func (self *Point) Compare(o *Point) int {
	d := self[X] - o[X]
	if math.FloatEqual(d, 0.0) {
		//x's are close enougth to each other
		d = self[Y] - o[Y]
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
