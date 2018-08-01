package geom

//Clone linestring
func (self *LineString) Clone() *LineString {
	return NewLineString(self.Coordinates)
}

//Clone linestring
func (self *LinearRing) Clone() *LinearRing {
	return &LinearRing{self.LineString.Clone()}
}


//Clone polygon
func (self *Polygon) Clone() *Polygon {
	rings := self.AsLinearRings()
	for i := range rings {
		rings[i] = rings[i].Clone()
	}
	return newPolygonFromRings(rings...)
}
