package geom

import (
	"github.com/intdxdt/math"
)

//Distance computes distance between two points
func (self Point) Distance(other Geometry) float64 {
	var dist = math.NaN()
	if IsNullGeometry(other) {
		return dist
	}

	if self.Intersects(other) {
		dist = 0.0
	} else if other.Type().IsPoint() {
		var pt = CastAsPoint(other)
		dist = self.Magnitude(&pt)
	} else {
		dist = distAsLines(self, other)
	}
	return dist
}

//Distance computes distance from segment to other geometry
func (self *Segment) Distance(other Geometry) float64 {
	return self.AsLineString().Distance(other)
}

// Computes the distance between wktreg and another linestring
// the distance between intersecting linestrings is 0.  Otherwise, the
// distance is the Euclidean distance between the closest points.
func (self *Polygon) Distance(other Geometry) float64 {
	var dist = math.NaN()
	if IsNullGeometry(other) {
		return dist
	}

	if self.Intersects(other) {
		dist = 0.0
	} else {
		dist = distAsLines(self, other)
	}
	return dist
}

//Computes the distance between wktreg and another linestring
func (self *LineString) Distance(other Geometry) float64 {
	var dist = math.NaN()
	if IsNullGeometry(other) {
		return dist
	}
	if self.Intersects(other) {
		dist = 0.0
	} else {
		dist = distAsLines(self, other)
	}
	return dist
}

//Computes the distance between a linestring and another linestring
//the distance between intersecting linestrings is 0.  Otherwise, the
//distance is the Euclidean distance between the closest segments.
func (self *LineString) line_line_dist(other *LineString) float64 {
	//TODO(titus):this could be improved KNN in Rtree
	// go bruteforce dist(seg , seg)
	return self.mindistBruteforce(other)
}

// brute force distance
func (self *LineString) mindistBruteforce(other *LineString) float64 {
	var dist, d float64 = math.MaxFloat64, 0
	var bln = false
	var ln = self.Coordinates
	var ln2 = other.Coordinates
	for i := 0; !bln && i < ln.Len()-1; i++ {
		for j := 0; !bln && j < ln2.Len()-1; j++ {
			d = SegSegDistance(ln.Pt(i), ln.Pt(i+1), ln2.Pt(j), ln2.Pt(j+1))
			dist = minf64(d, dist)
			bln = dist == 0
		}
	}
	return dist
}

//Computes the distance between wktreg and another linestring
func distAsLines(self, other Geometry) float64 {
	var dist = nan
	var lns1 = self.AsLinear()
	var lns2 = other.AsLinear()

	for i := 0; i < len(lns1); i++ {
		for j := 0; j < len(lns2); j++ {
			d := lns1[i].line_line_dist(lns2[j])
			if math.IsNaN(dist) {
				dist = d
			} else {
				dist = minf64(d, dist)
			}
		}
	}
	return dist
}
