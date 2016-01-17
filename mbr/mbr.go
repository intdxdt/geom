package mbr

import (
	point "github.com/intdxdt/simplex/geom/point"
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

func (self *MBR ) Clone() MBR {
	return New(self.ll, self.ur)
}
