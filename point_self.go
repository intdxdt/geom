package geom

//get geometry type
func (self *Point) Type() *geoType {
	return new_geoType(GeoTypePoint)
}

//get geometry interface
func (self *Point) Geometry() Geometry {
	return self
}
