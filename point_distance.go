package geom

import "github.com/intdxdt/math"

//DistanceSquare computes distance squared between two points
//Has possible overflow with squared x, y components
func (self *Point) SquareDistance(pt *Point) float64 {
	return self.SquareMagnitude(pt)
}

//Computes vector magnitude of pt as vector: x , y as components
func (self *Point) Magnitude(o *Point) float64 {
	return math.Hypot(o[X]-self[X], o[Y]-self[Y])
}

//Computes the square vector magnitude of pt as vector: x , y as components
//This has a potential overflow problem based on coordinates of pt x^2 + y^2
func (self *Point) SquareMagnitude(o *Point) float64 {
	dx := o[X] - self[X]
	dy := o[Y] - self[Y]
	return (dx * dx) + (dy * dy)
}
