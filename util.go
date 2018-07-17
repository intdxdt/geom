package geom

//Checks if geometry type is one of the fundermental types
//panics if geometry is not Point, Segment, LineString or Polygon
//NOTE: type embedding of any of these types does satisfy the Geometry
//Interface but is a null geometry.
func IsNullGeometry(g Geometry) bool {
	var bln bool
	if g.Type().IsPoint() {
		bln = false //Point{} is same as Point{0, 0}
	} else if g.Type().IsSegment() {
		bln = g.(*Segment) == nil
	} else if g.Type().IsLineString() {
		bln = g.(*LineString) == nil
	} else if g.Type().IsPolygon() {
		bln = g.(*Polygon) == nil
	}
	return bln
}

//Is linearing
func IsLinearRing(g Geometry) (*LinearRing, bool) {
	ln, ok := g.(*LinearRing)
	return ln, ok
}
