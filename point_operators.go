package geom

//Add creates a new point by adding to other point
func (a *Point) Add(x, y float64) Point {
	return Pt(a[X]+x, a[Y]+y)
}

//Sub creates a new point by adding to other point
func (a *Point) Sub(x, y float64) Point {
	return Pt(a[X]-x, a[Y]-y)
}

//KProduct create new point by multiplying point by a scaler  k
func (pt *Point) KProduct(k float64) Point {
	return Pt(k*pt[X], k*pt[Y])
}

//Dot Product of two points as vectors
func (pt *Point) DotProduct(other *Point) float64 {
	return DotProductXY(pt[X],pt[Y], other[X], other[Y])
}

//Neg create new point by finding the negation of point
func (self *Point) Neg() Point {
	return self.KProduct(-1)
}
