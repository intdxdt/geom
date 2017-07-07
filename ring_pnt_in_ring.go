package geom

import (
	"simplex/util/math"
)

//point completely in ring
func (self *LinearRing) PointCompletelyInRing(pnt *Point) bool {
	return self.Envelope().IntersectsPoint(pnt[:]) &&
			self.completely_in_ring(pnt)
}

/*
 Test whether a point lies inside a ring.
 The ring may be oriented in either direction.
 If the point lies on the ring boundary the result of this method is unspecified.
 This algorithm does not attempt to first check the point against the envelope
 of the ring.
 param p{Point} point to check for ring inclusion
 param ring{LinearRing} assumed to have first point identical to last point
 return {boolean}
*/
func (self *LinearRing) completely_in_ring(pnt *Point) bool {
	var i, i1 int
	var p1, p2 *Point
	var x1, y1, x2, y2, xInt float64
	var p = *pnt
	// For each segment l = (i-1, i), see if it crosses ray from test point in positive x direction.
	var crossings = 0 // number of segment/ray crossings
	for i = 1; i < self.LenVertices(); i++ {
		i1 = i - 1
		p1 = self.VertexAt(i)
		p2 = self.VertexAt(i1)

		if ((p1[y] > p[y]) && (p2[y] <= p[y])) || ((p2[y] > p[y]) && (p1[y] <= p[y])) {
			x1, y1 = p1[x]-p[x], p1[y]-p[y]
			x2, y2 = p2[x]-p[x], p2[y]-p[y]
			//segment straddles x axis, so compute intersection with x-axis.
			xInt = float64(math.SignOfDet2(x1, y1, x2, y2)) / (y2 - y1)
			//xsave = xInt
			//  crosses ray if strictly positive intersection.
			if xInt > 0.0 {
				crossings++
			}
		}
	}
	//  p is inside if number of crossings is odd.
	return (crossings % 2) == 1
}
