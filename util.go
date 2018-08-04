package geom

import (
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

//Checks if two bounding boxes intesect
func intersects(
	m_minx, m_miny, m_maxx, m_maxy float64,
	o_minx, o_miny, o_maxx, o_maxy float64) bool {
	//not disjoint
	return !(o_minx > m_maxx || o_maxx < m_minx || o_miny > m_maxy || o_maxy < m_miny)
}

//Checks if two bounding boxes intesect
func bounds_intersects(sa, sb, oa, ob *Point) bool {
	var  s_minx, s_miny, s_maxx, s_maxy = bounds(sa, sb)
	var  o_minx, o_miny, o_maxx, o_maxy = bounds(oa, ob)
	//not disjoint
	return !(o_minx > s_maxx || o_maxx < s_minx || o_miny > s_maxy || o_maxy < s_miny)
}


//Intersects bounding box defined by two points pt1 & pt2
func intersectsBounds(minx, miny, maxx, maxy float64, pt1, pt2 *Point) bool {
	if minx > maxf64(pt1[0], pt2[0]) || maxx < minf64(pt1[0], pt2[0]) {
		return false
	}
	return !(miny > maxf64(pt1[1], pt2[1]) || maxy < minf64(pt1[1], pt2[1]))
}

//bounds contains x, y
func containsXY(minx, miny, maxx, maxy, x, y float64) bool {
	return (x >= minx) && (x <= maxx) && (y >= miny) && (y <= maxy)
}

//envelope of segment
func bounds(a, b *Point) (float64, float64, float64, float64) {
	var minx, miny, maxx, maxy = a[X], a[Y], b[X], b[Y]
	return minf64(minx, maxx), minf64(miny, maxy), maxf64(minx, maxx), maxf64(miny, maxy)
}

//max
func maxf64(x, y float64) float64 {
	if y > x {
		return y
	}
	return x
}

//min
func minf64(x, y float64) float64 {
	if y < x {
		return y
	}
	return x
}
