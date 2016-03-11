package point

const (
	x = iota
	y
)

type Point [2]float64


//New constructor of Point
func NewPoint(array []float64) *Point {
	pt := new(Point)
	if len(array) == 1 {
		pt[x] = array[x]
	}else if len(array) >= 2 {
		pt[x], pt[y] = array[x], array[y]
	}
	return pt
}

//Clone point
func (self *Point) Clone() *Point {
	return &Point{self[x], self[y]}
}

//X gets the x coordinate of a point same as point[0]
func (self *Point) X() float64 {
	return self[x]
}

//Y gets the y coordinate of a point , same as self[1]
func (self *Point) Y() float64 {
	return self[y]
}

//AsArray converts Point to [2]float64
func (self *Point) AsArray() [2]float64 {
	return [2]float64{self[x], self[y]}
}
