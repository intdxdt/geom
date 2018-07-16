package geom

//get geometry type
func (self *LineString) Type() *geoType {
	return new_geoType(GeoTypeLineString)
}

//get geometry interface
func (self *LineString) Geometry() Geometry {
	return self
}
