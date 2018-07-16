package geom

//get geometry type
func (self *Point) Type() *geoType {
	return newGeoType(GeoTypePoint)
}

//get geometry interface
func (self *Point) Geometry() Geometry {
	return self
}
