package geom

import "github.com/intdxdt/math"

type LinearRing struct {
	*LineString
}

//new linear ring
func NewLinearRing(coords Coords) *LinearRing {
	var n = coords.Len()
	if n > 1 {
		if !IsRing(coords) {
			coords.Idxs = coords.Idxs[:n:n] //trigger a copy on append
			coords.Idxs = append(coords.Idxs, coords.Idxs[0])
		}
	}
	return &LinearRing{NewLineString(coords)}
}

//Contains point
func (self *LinearRing) containsPoint(pnt *Point) bool {
	return self.bbox.IntersectsPoint(pnt[:]) &&
		self.PointCompletelyInRing(pnt)
}

//Contains line
func (self *LinearRing) containsLine(ln *LineString) bool {
	if self.bbox.Disjoint(&ln.bbox.MBR) { //disjoint
		return false
	}
	var bln = true
	for i := 0; bln && i < ln.Coordinates.Len(); i++ {
		bln = self.containsPoint(ln.Pt(i))
	}
	return bln
}

//Contains polygon
func (self *LinearRing) containsPolygon(polygon *Polygon) bool {
	return self.containsLine(polygon.Shell.LineString)
}

//point completely in ring
func (self *LinearRing) PointCompletelyInRing(pnt *Point) bool {
	return self.LineString.bbox.IntersectsPoint(pnt[:]) && self.completelyInRing(pnt)
}

//Test whether a point lies inside a ring.
//The ring may be oriented in either direction.
//If the point lies on the ring boundary the result of this method is unspecified.
//This algorithm does not attempt to first check the point against the envelope of the ring.
func (self *LinearRing) completelyInRing(p *Point) bool {
	var i, i1 int
	var p1, p2 *Point
	var x1, y1, x2, y2, xInt float64
	// for each segment l = (i-1, i), see if it crosses ray from test point in positive x direction.
	var crossings = 0 // number of segment/ray crossings
	for i = 1; i < self.LenVertices(); i++ {
		i1 = i - 1
		p1 = self.Pt(i)
		p2 = self.Pt(i1)

		if ((p1[Y] > p[Y]) && (p2[Y] <= p[Y])) || ((p2[Y] > p[Y]) && (p1[Y] <= p[Y])) {
			x1, y1 = p1[X]-p[X], p1[Y]-p[Y]
			x2, y2 = p2[X]-p[X], p2[Y]-p[Y]
			//segment straddles x axis, so compute intersection with x-axis.
			xInt = float64(math.SignOfDet2(x1, y1, x2, y2)) / (y2 - y1)
			//xsave = xInt
			//crosses ray if strictly positive intersection.
			if xInt > 0.0 {
				crossings++
			}
		}
	}
	//  p is inside if number of crossings is odd.
	return (crossings % 2) == 1
}
