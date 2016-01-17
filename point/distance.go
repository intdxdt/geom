package point

import "math"


//Distance computes distance between two points
func (self *Point ) Distance(pt Point) float64 {
	return math.Hypot(self[x] - pt[x], self[y] - pt[y])
}

//DistanceSquare computes distance squared between two points
func (self *Point ) DistanceSquare(pt Point) float64 {
	dx := self[x] - pt[x]
	dy := self[y] - pt[y]
	//posible overflow
	return (dx * dx) + (dy * dy)
}
