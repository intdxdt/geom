package point

import "math"

const (
	x = iota
	y
	z
)

type Point [2]float64

var NullPt = Point{math.NaN(),math.NaN()}

//New constructor of Point
func New(array []float64) Point {
	var pt Point
	if len(array) == 1 {
		pt[x] = array[x]
	}else if len(array) >= 2 {
		pt[x], pt[y] = array[x], array[y]
	}
	return pt
}

//Clone point
func (self *Point) Clone() Point {
	var pt Point
	pt[x], pt[y] = self[x], self[y]
	return pt
}

//X gets the x coordinate of a point same as point[0]
func (self *Point) X() float64 {
	return self[x]
}

//Y gets the y coordinate of a point , same as self[1]
func (self *Point) Y() float64 {
	return self[y]
}

//As_array converts Point to [2]float64
func (self *Point) As_array() [2]float64 {
	var pt  [2]float64
	pt[x], pt[y] = self[x], self[y]
	return pt
}
