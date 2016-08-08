package geom

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

//2D cross product of OA and OB vectors,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if OAB makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func (o *Point) CCW(a, b *Point) float64 {
    return (b[x] - a[x]) * (o[y] - a[y]) - (b[y] - a[y]) * (o[x] - a[x])
}


//Neg create new point by finding the negation of point
func (self *Point) Neg() *Point {
    return self.KProduct(-1)
}