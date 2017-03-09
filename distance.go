package geom

import (
	"math"
)

//Distance computes distance between two points
func (self *Point) Distance(other Geometry) float64 {
	var dist = math.NaN()
	if IsNullGeometry(other) {
		return dist
	}
	pt, ispoint := IsPoint(other)

	if self.Intersects(other) {
		dist = 0.0
	} else if ispoint {
		dist = self.Magnitude(pt)
	} else {
		dist = dist_as_lines(self, other)
	}
	return dist
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

//Computes the distance between wktreg and another linestring
//the distance between intersecting linestrings is 0.  Otherwise, the
//distance is the Euclidean distance between the closest points.
func (self *LineString) line_line_dist(other *LineString) float64 {
	//TODO(titus):this could be improved KNN in Rtree
	// go bruteforce dist(seg , seg)
	return self.mindist_bruteforce(other)
}

// brute force distance
func (self *LineString) mindist_bruteforce(other *LineString) float64 {
	var dist = math.NaN()
	var bln = false
	var ln = self.coordinates
	var ln2 = other.coordinates
	for i := 0; !bln && i < len(ln)-1; i++ {
		for j := 0; !bln && j < len(ln2)-1; j++ {

			segA := &Segment{A: ln[i], B: ln[i+1]}
			segB := &Segment{A: ln2[j], B: ln2[j+1]}

			d := segA.Distance(segB)

			if math.IsNaN(dist) {
				dist = d
			} else {
				dist = math.Min(d, dist)
			}
			bln = (dist == 0.0)
		}
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
				dist = math.Min(d, dist)
			}
		}
	}
	return dist
}
