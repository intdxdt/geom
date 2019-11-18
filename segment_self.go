package geom

//GeomType
func (self *Segment) Type() GeoType {
	return GeoType(GeoTypeSegment)
}

//get geometry interface
func (self *Segment) Geometry() Geometry {
	return self
}

//checks if polygon is simple
func (self *Segment) IsSimple() bool {
	return true
}
