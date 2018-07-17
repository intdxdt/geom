package geom

//get geometry type
func (self Point) Type() GeoType {
	return GeoType(GeoTypePoint)
}

//get geometry interface
func (self *Point) Geometry() Geometry {
	return *self
}
