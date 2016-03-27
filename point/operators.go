package point

//Add creates a new point by adding to other point
func (self *Point) Add(pt *Point) *Point {
	return &Point{self[x] + pt[x], self[y] + pt[y]}
}

//Sub creates a new point by adding to other point
func (self *Point) Sub(pt *Point) *Point {
	return &Point{self[x] - pt[x], self[y] - pt[y]}
}

//KProduct create new point by multiplying point by a scaler  k
func (self *Point) KProduct(k float64) *Point {
	return &Point{k * self[x], k * self[y]}
}

//Dot Product of two points as vectors
func (self *Point) DotProduct(other *Point) float64 {
    return (self[x] * other[x]) + (self[y] * other[y])
}

//Unit vector of point
func (self *Point) UnitVector () *Point {
    m := self.Magnitude()
    return NewPointXY(self[x] / m, self[y] / m)
}


//Neg create new point by finding the negation of point
func (self *Point) Neg() *Point {
	return self.KProduct(-1)
}