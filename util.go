package geom

import (
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/math"
)

const (
	X = iota
	Y
	Z
	null = -9
)

var nan = math.NaN()
var feq = math.FloatEqual

func hypotSqr(p, q float64) float64 {
	return (p * p) + (q * q)
}

func hypot(p, q float64) float64 {
	if p < 0 {
		p = -p
	}
	if q < 0 {
		q = -q
	}
	if p < q {
		p, q = q, p
	}
	if p == 0 {
		return 0
	}
	q = q / p
	return p * math.Sqrt(1+q*q)
}

func snap_to_zero(x float64) float64 {
	if feq(x, 0) {
		x = 0
	}
	return x
}

func snap_to_zero_or_one(x float64) float64 {
	if feq(x, 0) {
		x = 0
	} else if feq(x, 1) {
		x = 1
	}
	return x
}

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

//Insersection of two intersecting mono bounding boxes
func monoIntersection(mbr, other *mono.MBR) (float64, float64, float64, float64) {
	var minx, miny = other.MinX, other.MinY
	var maxx, maxy = other.MaxX, other.MaxY

	if mbr.MinX > other.MinX {
		minx = mbr.MinX
	}

	if mbr.MinY > other.MinY {
		miny = mbr.MinY
	}

	if mbr.MaxX < other.MaxX {
		maxx = mbr.MaxX
	}

	if mbr.MaxY < other.MaxY {
		maxy = mbr.MaxY
	}

	return minx, miny, maxx, maxy
}

//Checks if two bounding boxes intesect
func intersects(
	m_minx, m_miny, m_maxx, m_maxy float64,
	o_minx, o_miny, o_maxx, o_maxy float64,
) bool {
	//not disjoint
	return !(o_minx > m_maxx || o_maxx < m_minx || o_miny > m_maxy || o_maxy < m_miny)
}

//Checks if two bounding boxes intesect
func boundsIntersects(sa, sb, oa, ob *Point) bool {
	var s_minx, s_miny, s_maxx, s_maxy = bounds(sa, sb)
	var o_minx, o_miny, o_maxx, o_maxy = bounds(oa, ob)
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
