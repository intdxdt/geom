package geom

//get geometry type
func (self *Point) Type() *geoType {
	return new_geoType(GeoType_Point)
}

//get geometry interface
func (self *Point) Geometry() Geometry {
	return self
}
