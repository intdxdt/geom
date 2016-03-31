package geom

import "math"


//DistanceSquare computes distance squared between two points
//Has possible overflow with squared x, y components
func (self *Point ) SquareDistance(pt *Point) float64 {
    return self.Sub(pt).SquareMagnitude()
}

//Computes vector magnitude of pt as vector: x , y as components
func (self *Point) Magnitude() float64{
    return math.Hypot(self[x], self[y])
}

//Computes the square vector magnitude of pt as vector: x , y as components
//This has a potential overflow problem based on coordinates of pt x^2 + y^2
func (self *Point)  SquareMagnitude() float64{
    return (self[x] * self[x]) + (self[y] * self[y])
}