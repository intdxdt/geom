package geom

import (
	"simplex/cart2d"
	. "simplex/side"
	. "simplex/struct/item"
	"simplex/util/math"
)

//Equals evaluates whether two points are the same
func (pt *Point) Equals2D(point *Point) bool {
	return math.FloatEqual(pt[x], point[x]) && math.FloatEqual(pt[y], point[y])
}

func (pt *Point) Equals3D(point *Point) bool {
	return math.FloatEqual(pt[x], point[x]) && math.FloatEqual(pt[y], point[y]) &&
		math.FloatEqual(pt[z], point[z])
}

//Disjoint evaluates whether points are not coincident
func (pt *Point) Disjoint(point *Point) bool {
	return !(pt.Intersects(point))
}

//compare points as items - x | y ordering
func (self *Point) Compare(o Item) int {
	pt := o.(*Point)
	d := self[x] - pt[x]
	if math.FloatEqual(d, 0.0) {
		//x's are close enougth to each other
		d = self[y] - pt[y]
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
func (c *Point) SideOf(a, b *Point) *Side {
	s := NewSide()
	ccw := cart2d.CCW(a, b, c)
	if math.FloatEqual(ccw, 0) {
		s.AsOn()
	} else if ccw > 0 {
		s.AsLeft()
	} else if ccw < 0 {
		s.AsRight()
	}
	return s
}
