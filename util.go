package geom

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/geom/mono"
)

//Checks if geometry type is one of the fundermental types
//panics if geometry is not Point, Segment, LineString or Polygon
//NOTE: type embedding of any of these types does satisfy the Geometry
//Interface but is a null geometry.
func IsNullGeometry(g Geometry) bool {
	var bln bool
	//get underlying geometry type with g.Geometry()
	if g.Type().IsPoint() {
		bln = false //Point{} is same as Point{0, 0}
	} else if g.Type().IsSegment() {
		bln = CastAsSegment(g) == nil
	} else if g.Type().IsLineString() {
		bln = CastAsLineString(g) == nil
	} else if g.Type().IsPolygon() {
		bln = CastAsPolygon(g) == nil
	}
	return bln
}

//Is linearing
func IsLinearRing(g Geometry) (*LinearRing, bool) {
	ln, ok := g.(*LinearRing)
	return ln, ok
}

//Insersection of two intersecting mono bounding boxes
func mono_intersection(mbr, other *mono.MBR) (float64, float64, float64, float64) {
	var minx, miny, maxx, maxy = other.MBR[0], other.MBR[1], other.MBR[2], other.MBR[3]

	if mbr.MBR[0] > other.MBR[0] {
		minx = mbr.MBR[0]
	}

	if mbr.MBR[1] > other.MBR[1] {
		miny = mbr.MBR[1]
	}

	if mbr.MBR[2] < other.MBR[2] {
		maxx = mbr.MBR[2]
	}

	if mbr.MBR[3] < other.MBR[3] {
		maxy = mbr.MBR[3]
	}

	return minx, miny, maxx, maxy
}


//Intersects bounding box defined by two points pt1 & pt2
func intersectsBounds(minx, miny, maxx, maxy float64, pt1, pt2 *Point) bool {
	if minx > math.MaxF64(pt1[0], pt2[0]) || maxx < math.MinF64(pt1[0], pt2[0]) {
		return false
	}
	return !(miny > math.MaxF64(pt1[1], pt2[1]) || maxy < math.MinF64(pt1[1], pt2[1]))
}
