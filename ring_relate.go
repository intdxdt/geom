package geom

//Contains point
func (self *LinearRing) contains_point(pnt *Point) bool {
	return self.bbox.IntersectsPoint(pnt[:]) &&
		self.PointCompletelyInRing(pnt)
}

//Contains line
func (self *LinearRing) contains_line(ln *LineString) bool {
	if self.bbox.Disjoint(&ln.bbox.MBR) { //disjoint
		return false
	}
	var bln = true
	for i := 0; bln && i < len(ln.coordinates); i++ {
		bln = self.contains_point(&ln.coordinates[i])
	}
	return bln
}

//Contains polygon
func (self *LinearRing) contains_polygon(polygon *Polygon) bool {
	return self.contains_line(polygon.Shell.LineString)
}
