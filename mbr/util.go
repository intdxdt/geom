package mbr

import (
	"github.com/intdxdt/simplex/geom/point"
)

func (self *MBR) AsArray() []float64 {
	return []float64{self[x1], self[y1], self[x2], self[y2]}
}

func (self *MBR) Width() float64 {
	return self[x2] - self[x1]
}

func (self *MBR) Height() float64 {
	return self[y2] - self[y1]
}

func (self *MBR) Area() float64 {
	return self.Height() * self.Width()
}



//Translate mbr  by change in x and y
func (self *MBR)Translate(dx, dy float64) MBR {
	return New(
		self[x1] + dx, self[y1] + dy,
		self[x2] + dx, self[y2] + dy)
}

func (self *MBR) Center() point.Point {
	return point.Point{
		(self[x1] + self[x2]) / 2.0,
		(self[y1] + self[y2]) / 2.0}
}
