package geom

//DistanceSquare computes distance squared between two points
//Has possible overflow with squared x, y components
func (self *Point) SquareDistance(pt *Point) float64 {
	return self.MagnitudeSquare(pt)
}

//Computes vector magnitude of pt as vector: x , y as components
func (self *Point) Magnitude(o *Point) float64 {
	return MagnitudeXY(o[X]-self[X], o[Y]-self[Y])
}

//Computes the square vector magnitude of pt as vector: x , y as components
//This has a potential overflow problem based on coordinates of pt x^2 + y^2
func (self *Point) MagnitudeSquare(o *Point) float64 {
	return MagnitudeSquareXY(o[X]-self[X], o[Y]-self[Y])

}
