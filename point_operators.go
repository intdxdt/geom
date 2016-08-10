package geom

import (
    "simplex/cart2d"
)

//Add creates a new point by adding to other point
func (a *Point) Add(b cart2d.Cart2D) *Point {
    cx, cy := cart2d.Add(a, b)
    return NewPointXY(cx, cy)
}

//Sub creates a new point by adding to other point
func (a *Point) Sub(b cart2d.Cart2D) *Point {
    cx, cy := cart2d.Sub(a, b)
    return NewPointXY(cx, cy)
}

//KProduct create new point by multiplying point by a scaler  k
func (pt *Point) KProduct(k float64) *Point {
    cx, cy := cart2d.KProduct(pt, k)
    return NewPointXY(cx, cy)
}

//Dot Product of two points as vectors
func (pt *Point) DotProduct(other *Point) float64 {
    return cart2d.DotProduct(pt, other)
}

//Neg create new point by finding the negation of point
func (self *Point) Neg() *Point {
    return self.KProduct(-1)
}