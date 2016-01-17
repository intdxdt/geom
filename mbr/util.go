package mbr

import (
	point "github.com/intdxdt/simplex/geom/point"
)



func (self *MBR) AsArray() []float64 {
	return []float64{self.ll[x], self.ll[y], self.ur[x], self.ur[y]}
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
