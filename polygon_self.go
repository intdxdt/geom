package geom

//get geometry type
func (self *Polygon) Type() *geoType {
	return new_geoType(GeoTypePolygon)
}

//get geometry interface
func (self *Polygon) Geometry() Geometry {
	return self
}

