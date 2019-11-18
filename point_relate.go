package geom

import (
	"github.com/intdxdt/robust"
	"github.com/intdxdt/side"
)

//Equals evaluates whether two points are the same
func (pt *Point) Equals2D(o *Point) bool {
	return feq(pt[X], o[X]) && feq(pt[Y], o[Y])
}

func (pt *Point) Equals3D(o *Point) bool {
	return feq(pt[X], o[X]) && feq(pt[Y], o[Y]) && feq(pt[Z], o[Z])
}

//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(o *Point) bool {
	return !pt.Equals2D(o)
}

//compare points as items - x | y ordering
func (self *Point) Compare(o *Point) int {
	var d = self[X] - o[X]
	if feq(d, 0.0) {
		d = self[Y] - o[Y]
	}

	if feq(d, 0.0) {
		return 0
	} else if d < 0 {
		return -1
	}
	return 1
}

//position of Pnts relative to line AB
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

//2D cross product of AB and AC vectors given A, B, and Pnts as points,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if ABC makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func (c *Point) Orientation2D(a, b *Point) float64 {
	return robust.Orientation2D(a[:2], b[:2], c[:2])
}
