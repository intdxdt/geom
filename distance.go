package geom

import (
	"github.com/intdxdt/math"
)

//Distance computes distance between two points
func (self *Point) Distance(other Geometry) float64 {
	var dist = math.NaN()
	if IsNullGeometry(other) {
		return dist
	}

	if self.Intersects(other) {
		dist = 0.0
	} else if pt, ok := IsPoint(other); ok {
		dist = self.Magnitude(pt)
	} else {
		dist = dist_as_lines(self, other)
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
		dist = dist_as_lines(self, other)
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
		dist = dist_as_lines(self, other)
	}
	return dist
}

//Computes the distance between a linestring and another linestring
//the distance between intersecting linestrings is 0.  Otherwise, the
//distance is the Euclidean distance between the closest segments.
func (self *LineString) line_line_dist(other *LineString) float64 {
	//TODO(titus):this could be improved KNN in Rtree
	// go bruteforce dist(seg , seg)
	return self.mindist_bruteforce(other)
}

// brute force distance
func (self *LineString) mindist_bruteforce(other *LineString) float64 {
	var dist, d float64 = math.MaxFloat64, 0
	var bln, exe = false, false
	var ln = self.coordinates
	var ln2 = other.coordinates
	for i := 0; !bln && i < len(ln)-1; i++ {
		for j := 0; !bln && j < len(ln2)-1; j++ {
			exe = true //execution of segment coords happened
			d = SegSegDistance(ln[i], ln[i+1], ln2[j], ln2[j+1])
			dist = math.MinF64(d, dist)
			bln = dist == 0.0
		}
	}
	if !exe {
		dist = math.NaN()
	}
	return dist
}

//Computes the distance between wktreg and another linestring
func dist_as_lines(self, other Geometry) float64 {
	var dist = math.NaN()
	var lns1 = self.AsLinear()
	var lns2 = other.AsLinear()

	for i := 0; i < len(lns1); i++ {
		for j := 0; j < len(lns2); j++ {
			d := lns1[i].line_line_dist(lns2[j])
			if math.IsNaN(dist) {
				dist = d
			} else {
				dist = math.MinF64(d, dist)
			}
		}
	}
	return dist
}
