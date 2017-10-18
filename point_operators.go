package geom

import (
	"github.com/intdxdt/cart"
)

//Add creates a new point by adding to other point
func (a *Point) Add(b cart.Coord2D) *Point {
	cx, cy := cart.Add(a, b)
	return NewPointXY(cx, cy)
}

//Sub creates a new point by adding to other point
func (a *Point) Sub(b cart.Coord2D) *Point {
	cx, cy := cart.Sub(a, b)
	return NewPointXY(cx, cy)
}

//KProduct create new point by multiplying point by a scaler  k
func (pt *Point) KProduct(k float64) *Point {
	cx, cy := cart.KProduct(pt, k)
	return NewPointXY(cx, cy)
}

//Dot Product of two points as vectors
func (pt *Point) DotProduct(other *Point) float64 {
	return cart.DotProduct(pt, other)
}

//Neg create new point by finding the negation of point
func (self *Point) Neg() *Point {
	return self.KProduct(-1)
}
