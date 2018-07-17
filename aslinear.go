package geom

//Point as line strings
func (self Point) AsLinear() []*LineString {
	return self.AsLineStrings()
}

//Segment as lineString
func (self *Segment) AsLinear() []*LineString {
	return []*LineString{self.AsLineString()}
}

//Linestring as line strings
func (self *LineString) AsLinear() []*LineString {
	return []*LineString{self}
}

//polygon as  array of line strings
func (self *Polygon) AsLinear() []*LineString {
	var rings = self.AsLinearRings()
	var lns = make([]*LineString, len(rings))
	for i, ln := range rings {
		lns[i] = ln.LineString
	}
	return lns
}
