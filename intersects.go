package geom

//Checks if pt intersects other geometry
func (pt Point) Intersects(other Geometry) bool {
	//checks for non-geometry types
	if IsNullGeometry(other) {
		return false
	}
	if other.Type().IsPoint() {
		var p = CastAsPoint(other)
		return pt.Equals2D(&p)
	}
	return other.Intersects(pt)
}

//Segment Intersects other geometry
func (self *Segment) Intersects(other Geometry) bool {
	return self.AsLineString().Intersects(other)
}

//Checks if linestring intersecs other geometry
func (self *LineString) Intersects(other Geometry) bool {
	//checks for non-geometry types
	if IsNullGeometry(other) {
		return false
	}
	var bln = false
	var other_lns = other.AsLinear()
	var shell = other_lns[0]

	if self.bbox.Disjoint(&shell.bbox.MBR) {
		bln = false
	} else if other.Type().IsPolygon() {
		bln = self.intersects_polygon(other_lns)
	} else if other.Type().IsLineString() ||
		other.Type().IsSegment() || other.Type().IsPoint() {
		bln = self.intersects_linestring(shell)
	}

	return bln
}

//Checks if polygon intersects other geometry
func (self *Polygon) Intersects(other Geometry) bool {
	//checks for non-geometry types
	if IsNullGeometry(other) {
		return false
	}

	var bln = false
	var within_bounds bool
	var rings []*LineString
	var ln *LineString

	//reverse intersect line inter poly
	if other.Type().IsSegment() ||
		other.Type().IsLineString() || other.Type().IsPoint() {

		ln = other.AsLinear()[0]
		within_bounds = self.Shell.bbox.Intersects(&ln.bbox.MBR)
		rings = self.AsLinear()
		bln = within_bounds && ln.intersects_polygon(rings)

	} else if other.Type().IsPolygon() {
		var ply = CastAsPolygon(other)
		if self.Shell.bbox.Intersects(&ply.Shell.bbox.MBR) {
			var small, big *Polygon

			if self.Shell.bbox.Area() < ply.Shell.bbox.Area() {
				small, big = self, ply
			} else {
				small, big = ply, self
			}
			bln = small.Shell.LineString.intersects_polygon( big.AsLinear())
		}
	}

	return bln
}
