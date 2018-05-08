package geom

import (
	"github.com/intdxdt/side"
	"github.com/intdxdt/math"
	"github.com/intdxdt/robust"
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
	var s = side.NewSide()
	var ccw = c.Orientation2D(a, b)
	if ccw == 0 {
		s.AsOn()
	} else if ccw < 0 {
		s.AsLeft()
	} else if ccw > 0 {
		s.AsRight()
	}
	return s
}

//2D cross product of AB and AC vectors given A, B, and C as points,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if ABC makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func (c *Point) Orientation2D(a, b *Point) float64 {
	return robust.Orientation2D(a[:2], b[:2], c[:2])
}
