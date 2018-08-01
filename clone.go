package geom

//Clone linestring
func (self *LineString) Clone() *LineString {
	var coords = make([]Point, 0, self.Coordinates.Len())
	for _, i := range self.Coordinates.Idxs {
		coords = append(coords, self.Coordinates._c[i])
	}
	return NewLineString(Coordinates(coords))
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
