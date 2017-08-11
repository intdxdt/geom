package geom


//GeomType
func (self *Segment) Type() *geoType {
	return new_geoType(GeoType_Segment)
}


//get geometry interface
func (self *Segment) Geometry() Geometry {
	return self
}


//checks if polygon is simple
func (self *Segment) IsSimple() bool {
	return true
}
