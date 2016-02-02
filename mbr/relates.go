package mbr

import (
	point "github.com/intdxdt/simplex/geom/point"
	"math"
)

func (self *MBR) Equals(other MBR) bool {
	return self.ll.Equals(other.ll) && self.ur.Equals(other.ur)
}

func (self *MBR) IsNull() bool {
	return self.ll.IsNull() || self.ur.IsNull()
}

func (self *MBR) Intersection(other MBR) (MBR, bool) {
	nan := math.NaN()

	minx, miny := nan, nan
	maxx, maxy := nan, nan

	inters := self.Intersects(other)

	if inters {
		if self.ll[x] > other.ll[x] {
			minx = self.ll[x]
		} else {
			minx = other.ll[x]
		}

		if self.ll[y] > other.ll[y] {
			miny = self.ll[y]
		}else {
			miny = other.ll[y]
		}

		if self.ur[x] < other.ur[x] {
			maxx = self.ur[x]
		}else {
			maxx = other.ur[x]
		}

		if self.ur[y] < other.ur[y] {
			maxy = self.ur[y]
		}else {
			maxy = other.ur[y]
		}

	}

	return New(point.Point{minx, miny}, point.Point{maxx, maxy}), inters
}

func (self *MBR) Intersects(other MBR) bool {
	//not disjoint
	return ! (
	other.ll[x] > self.ur[x] ||
	other.ur[x] < self.ll[x] ||
	other.ll[y] > self.ur[y] ||
	other.ur[y] < self.ll[y])
}

func (self *MBR) IntersectsPoint(p point.Point) bool {
	return self.ContainsXY(p[x], p[y])
}

func (self *MBR) IntersectsBounds(q1, q2 point.Point) bool {

	minq := math.Min(q1[x], q2[x])
	maxq := math.Max(q1[x], q2[x])

	if (self.ll[x] > maxq || self.ur[x] < minq) {
		return false
	}

	minq = math.Min(q1[y], q2[y])
	maxq = math.Max(q1[y], q2[y])

	// not disjoint
	return !(self.ll[y] > maxq || self.ur[y] < minq)
}

func (self *MBR)  Contains(other MBR) bool {
	return (
	(other.ll[x] >= self.ll[x]) &&
	(other.ur[x] <= self.ur[x]) &&
	(other.ll[y] >= self.ll[y]) &&
	(other.ur[y] <= self.ur[y]))
}

func (self *MBR) ContainsXY(x, y float64) bool {
	return (
	(x >= self.ll[0]) &&
	(x <= self.ur[0]) &&
	(y >= self.ll[1]) &&
	(y <= self.ur[1]))
}

//CompletelyContainsXY is true if mbr completely contains location with {x, y}
func (self *MBR) CompletelyContainsXY(x, y float64) bool {
	return (
	(x > self.ll[0]) &&
	(x < self.ur[0]) &&
	(y > self.ll[1]) &&
	(y < self.ur[1]))
}

//CompletelyContainsMBR is true if mbr completely contains other
func (self *MBR) CompletelyContainsMBR(other MBR) bool {
	return (
	(other.ll[x] > self.ll[x]) &&
	(other.ur[x] < self.ur[x]) &&
	(other.ll[y] > self.ll[y]) &&
	(other.ur[y] < self.ur[y]))
}

//Disjoint of mbr do not intersect
func (self *MBR) Disjoint(m MBR) bool {
	return !(self.Intersects(m))
}

