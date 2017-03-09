package mbr

import (
	"math"
	umath "simplex/util/math"
)

func (self *MBR) Equals(other *MBR) bool {
	return umath.FloatEqual(self[x1], other[x1]) &&
		umath.FloatEqual(self[y1], other[y1]) &&
		umath.FloatEqual(self[x2], other[x2]) &&
		umath.FloatEqual(self[y2], other[y2])
}

func (self *MBR) Intersection(other *MBR) (*MBR, bool) {
	nan := math.NaN()

	minx, miny := nan, nan
	maxx, maxy := nan, nan

	inters := self.Intersects(other)

	if inters {
		if self[x1] > other[x1] {
			minx = self[x1]
		} else {
			minx = other[x1]
		}

		if self[y1] > other[y1] {
			miny = self[y1]
		} else {
			miny = other[y1]
		}

		if self[x2] < other[x2] {
			maxx = self[x2]
		} else {
			maxx = other[x2]
		}

		if self[y2] < other[y2] {
			maxy = self[y2]
		} else {
			maxy = other[y2]
		}

	}

	return NewMBR(minx, miny, maxx, maxy), inters
}

func (self *MBR) Intersects(other *MBR) bool {
	//not disjoint
	return !(other[x1] > self[x2] ||
		other[x2] < self[x1] ||
		other[y1] > self[y2] ||
		other[y2] < self[y1])
}

func (self *MBR) IntersectsPoint(p []float64) bool {
	if len(p) < 2 {
		return false
	}
	return self.ContainsXY(p[x1], p[y1])
}

func (self *MBR) IntersectsBounds(q1, q2 []float64) bool {
	if len(q1) < 2 || len(q2) < 2 {
		return false
	}
	minq := math.Min(q1[x1], q2[x1])
	maxq := math.Max(q1[x1], q2[x1])

	if self[x1] > maxq || self[x2] < minq {
		return false
	}

	minq = math.Min(q1[y1], q2[y1])
	maxq = math.Max(q1[y1], q2[y1])

	// not disjoint
	return !(self[y1] > maxq || self[y2] < minq)
}

func (self *MBR) Contains(other *MBR) bool {
	return (other[x1] >= self[x1]) &&
		(other[x2] <= self[x2]) &&
		(other[y1] >= self[y1]) &&
		(other[y2] <= self[y2])
}

func (self *MBR) ContainsXY(x, y float64) bool {
	return (x >= self[x1]) &&
		(x <= self[x2]) &&
		(y >= self[y1]) &&
		(y <= self[y2])
}

//CompletelyContainsXY is true if mbr completely contains location with {x, y}
func (self *MBR) CompletelyContainsXY(x, y float64) bool {
	return (x > self[x1]) &&
		(x < self[x2]) &&
		(y > self[y1]) &&
		(y < self[y2])
}

//CompletelyContainsMBR is true if mbr completely contains other
func (self *MBR) CompletelyContainsMBR(other *MBR) bool {
	return (other[x1] > self[x1]) &&
		(other[x2] < self[x2]) &&
		(other[y1] > self[y1]) &&
		(other[y2] < self[y2])
}

//Disjoint of mbr do not intersect
func (self *MBR) Disjoint(m *MBR) bool {
	return !(self.Intersects(m))
}
