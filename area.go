package geom

import "github.com/intdxdt/math"

//Area of point
func (self *Point) Area() float64 {
	return 0.0
}

//Area of line string
func (self *LineString) Area() float64 {
	if self.IsRing() {
		ring := &LinearRing{self}
		return ring.Area()
	}
	return 0.0
}

//Area of linear ring
func (self *LinearRing) Area() float64 {
	var coords = self.LineString.coordinates
	var n = len(coords)
	var a, b *Point
	var area = 0.0
	b = coords[n-1]
	for i := 0; i < n; i++ {
		a = b
		b = coords[i]
		area += a[y]*b[x] - a[x]*b[y]
	}
	return math.Abs(area * 0.5)
}

//Area of polygon
func (self *Polygon) Area() float64 {
	var a = self.Shell.Area()
	for i := 0; i < len(self.Holes); i++ {
		a -= self.Holes[i].Area()
	}
	return a
}
