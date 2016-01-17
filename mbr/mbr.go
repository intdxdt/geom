package mbr

import (
	point "github.com/intdxdt/simplex/geom/point"
	"math"
)

const (
	x = iota
	y
)

type MBR struct {
	ll, ur point.Point
}

func New(ll, ur  point.Point) MBR {
	x1, y1 := ll[x], ll[y]
	x2, y2 := ur[x], ur[y]

	var minx, maxx, miny, maxy float64

	var self MBR

	if x1 < x2 {
		minx = x1
		maxx = x2
	}else {
		minx = x2
		maxx = x1
	}

	if y1 < y2 {
		miny = y1
		maxy = y2
	}else {
		miny = y2
		maxy = y1
	}
	self.ll = point.Point{minx, miny}
	self.ur = point.Point{maxx, maxy}

	return self
}

func (self *MBR) String() string {
	ll, ur := self.ll, self.ur
	ul, lr := point.Point{ll[x], ur[y]}, point.Point{ur[x], ll[y]}

	return "POLYGON ((" +
	ll.String() + ", " +
	ul.String() + ", " +
	ur.String() + ", " +
	lr.String() + ", " +
	ll.String() +
	"))"
}

func (self *MBR ) Clone() MBR {
	return New(self.ll, self.ur)
}

func (self *MBR) As_array() []float64 {
	return []float64{self.ll[x], self.ll[y], self.ur[x], self.ur[y]}
}

func (self *MBR) Equals(other MBR) bool {
	return self.ll.Equals(other.ll) && self.ur.Equals(other.ur)
}

func (self *MBR) Is_null() bool {
	return self.ll.IsNull() || self.ur.IsNull()
}

func (self *MBR) Width() float64 {
	return self.ur[x] - self.ll[x]
}

func (self *MBR) Height() float64 {
	return self.ur[y] - self.ll[y]
}

func (self *MBR) Area() float64 {
	return self.Height() * self.Width()
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

//Expand to include other mbr
func (self *MBR) Expand(other MBR) *MBR {

	if other.ll[x] < self.ll[x] {
		self.ll[x] = other.ll[x]
	}
	if other.ur[x] > self.ur[x] {
		self.ur[x] = other.ur[x]
	}
	if other.ll[y] < self.ll[y] {
		self.ll[y] = other.ll[y]
	}
	if other.ur[y] > self.ur[y] {
		self.ur[y] = other.ur[y]
	}
	return self
}

//ExpandBy expands mbr by change in x and y
func (self *MBR) ExpandBy(dx, dy float64) *MBR {

	minx, miny := self.ll[x] - dx, self.ll[y] - dy
	maxx, maxy := self.ur[x] + dx, self.ur[y] + dy

	minx, maxx = math.Min(minx, maxx), math.Max(minx, maxx)
	miny, maxy = math.Min(miny, maxy), math.Max(miny, maxy)

	self.ll[x], self.ll[y] = minx, miny
	self.ur[x], self.ur[y] = maxx, maxy

	return self
}

//ExpandXY expands mbr to include x and y
func (self *MBR) ExpandXY(x_coord, y_coord float64) *MBR {

	if x_coord < self.ll[x] {
		self.ll[x] = x_coord
	}else if x_coord > self.ur[x] {
		self.ur[x] = x_coord
	}

	if y_coord < self.ll[y] {
		self.ll[y] = y_coord
	}else if y_coord > self.ur[y] {
		self.ur[y] = y_coord
	}

	return self
}

//Translate mbr  by change in x and y
func (self *MBR)Translate(dx, dy float64) MBR {
	return New(
		point.Point{self.ll[x] + dx, self.ll[y] + dy},
		point.Point{self.ur[x] + dx, self.ur[y] + dy})
}

func (self *MBR) Center() point.Point {
	return point.Point{
		(self.ll[x] + self.ur[x]) / 2.0,
		(self.ll[y] + self.ur[y]) / 2.0}
}

func (self *MBR) Distance(other MBR) float64 {

	if self.Intersects(other) {
		return 0.0
	}

	dx := 0.0
	dy := 0.0
	//find closest edge by x
	if (self.ur[x] < other.ll[x]) {
		dx = other.ll[x] - self.ur[x]
	}else if (self.ll[x] > other.ur[x]) {
		dx = self.ll[x] - other.ur[x]
	}
	//find closest edge by y
	if (self.ur[y] < other.ll[y]) {
		dy = other.ll[y] - self.ur[y]
	}else if (self.ll[y] > other.ur[y]) {
		dy = self.ll[y] - other.ur[y]
	}

	return math.Hypot(dx, dy)
}