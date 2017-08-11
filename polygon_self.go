package geom

//get geometry type
func (self *Polygon) Type() *geoType {
	return new_geoType(GeoType_Polygon)
}

//get geometry interface
func (self *Polygon) Geometry() Geometry {
	return self
}

//checks if polygon is simple
func (self *Polygon) IsSimple() bool {
	return self.Shell.IsSimple()
}
