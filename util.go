package geom

//Checks if geometry type is one of the fundermental types
//panics if geometry is not Point, Segment, LineString or Polygon
//NOTE: type embedding of any of these types does satisfy the Geometry
//Interface but is a null geometry.
func IsNullGeometry(geom Geometry) bool {
	if g, ok := IsPoint(geom); ok {
		return g == nil
	} else if g, ok := IsSegment(geom); ok {
		return g == nil
	} else if g, ok := IsLineString(geom); ok {
		return g == nil
	} else if g, ok := IsPolygon(geom); ok {
		return g == nil
	}
	panic("unknown geometry type")
}

//Is point
func IsPoint(g Geometry) (*Point, bool) {
	pt, ok := g.(*Point)
	return pt, ok
}

//Is polygon
func IsPolygon(g Geometry) (*Polygon, bool) {
	ply, ok := g.(*Polygon)
	return ply, ok
}

//Is segment
func IsSegment(g Geometry) (*Segment, bool) {
	seg, ok := g.(*Segment)
	return seg, ok
}

//Is linestring
func IsLineString(g Geometry) (*LineString, bool) {
	ln, ok := g.(*LineString)
	return ln, ok
}

//Is linearing
func IsLinearRing(g Geometry) (*LinearRing, bool) {
	ln, ok := g.(*LinearRing)
	return ln, ok
}

