package geom

//Clone linestring
func (self *LineString) Clone() *LineString {
	return NewLineString(self.coordinates)
}

//Clone linestring
func (self *LinearRing) Clone() *LinearRing {
	return &LinearRing{self.LineString.Clone()}
}

//Clone point
func (self *Point) Clone() *Point {
	return NewPoint(self[:])
}

//Clone polygon
func (self *Polygon) Clone() *Polygon {
	rings := self.AsLinearRings()
	for i := range rings {
		rings[i] = rings[i].Clone()
	}
	return NewPolygonFromRings(rings...)
}
