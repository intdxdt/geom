package geom

//Checks if pt intersects other geometry
func (pt *Point) Intersects(other Geometry) bool {
	//checks for non-geometry types
	if IsNullGeometry(other) {
		return false
	}
	if p, ok := IsPoint(other); ok {
		return pt.Equals2D(p)
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

	bln := false
	_, ispoly := IsPolygon(other)
	_, isline := IsLineString(other)
	_, isseg := IsSegment(other)
	_, ispoint := IsPoint(other)

	other_lns := other.AsLinear()
	shell := other_lns[0]

	if self.bbox.Disjoint(shell.bbox.MBR) {
		bln = false
	} else if ispoly {
		bln = self.intersects_polygon(other_lns)
	} else if isline || isseg || ispoint {
		bln = self.intersectsLinestring(shell)
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

	_, ispoint := IsPoint(other)
	_, isline := IsLineString(other)
	_, isseg := IsSegment(other)

	//reverse intersect line inter poly
	if isseg || isline || ispoint {

		ln = other.AsLinear()[0]
		within_bounds = self.Shell.bbox.Intersects(ln.bbox.MBR)
		rings = self.AsLinear()
		bln = within_bounds && ln.intersects_polygon(rings)

	} else if other_poly, ok := IsPolygon(other); ok {

		if self.Shell.bbox.Disjoint(other_poly.Shell.bbox.MBR) {
			bln = false
		}
		var small, big *Polygon

		if self.Shell.bbox.Area() < other_poly.Shell.bbox.Area() {
			small, big = self, other_poly
		} else {
			small, big = other_poly, self
		}

		ln = small.Shell.LineString
		rings = big.AsLinear()
		within_bounds = ln.bbox.Intersects(rings[0].bbox.MBR)
		bln = within_bounds && ln.intersects_polygon(rings)

	}
	return bln
}
