package geom

//get geometry type
func (self *LineString) Type() GeoType {
	return GeoType(GeoTypeLineString)
}

//get geometry interface
func (self *LineString) Geometry() Geometry {
	return self
}
