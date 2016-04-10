package geom

import (
	"math"
)

//Distance computes distance between two points
func (self *Point) Distance(other Geometry) float64 {
	var dist = -1.0
	if IsNullGeometry(other) {
		return dist
	}
	pt, ispoint := IsPoint(other)

	if self.Intersects(other) {
		dist = 0.0
	} else if ispoint {
		dist = self.Sub(pt).Magnitude()
	} else {
		var bln = false
		var lns1 = self.AsLinear()
		var lns2 = other.AsLinear()

		for i := 0; !bln && i < len(lns1); i++ {
			for j := 0; !bln && j < len(lns2); j++ {
				d := lns1[i].Distance(lns2[j])
				if dist < 0 {
					dist = d
				} else {
					dist = math.Min(d, dist)
				}
			}
		}
	}
	return dist
}

// Computes the distance between self and another linestring
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
		var bln = false
		var lns1 = self.AsLinear()
		var lns2 = other.AsLinear()

		for i := 0; !bln && i < len(lns1); i++ {
			for j := 0; !bln && j < len(lns2); j++ {
				d := lns1[i].Distance(lns2[j])
				if math.IsNaN(dist) {
					dist = d
				} else {
					dist = math.Min(d, dist)
				}
			}
		}
	}
	return dist
}

//description  Computes the distance between self and another linestring
//the distance between intersecting linestrings is 0.  Otherwise, the
//distance is the Euclidean distance between the closest points.
func (self *LineString) Distance(other *LineString) float64 {
	var dist = -1.0
	if self.Intersects(other) {
		dist = 0.0
	} else {
		//TODO(titus):this could be improved KNN in Rtree
		// go bruteforce dist(seg , seg)
		dist = self.mindist_bruteforce(other)
	}
	return dist
}

// brute force distance
func (self *LineString) mindist_bruteforce(other *LineString) float64 {
	var bln = false
	var ln = self.coordinates
	var ln2 = other.coordinates
	var dist = -1.0
	for i := 0; !bln && i < len(ln)-1; i++ {
		for j := 0; !bln && j < len(ln2)-1; j++ {

			segA := &Segment{ln[i], ln[i+1]}
			segB := &Segment{ln2[j], ln2[j+1]}

			d := segA.Distance(segB)

			if dist < 0 {
				dist = d
			} else {
				dist = math.Min(d, dist)
			}
			bln = (dist == 0.0)
		}
	}
	return dist
}

//minimum distance
func segseg_mindist(segsa, segsb []*Segment) float64 {
	var bln = false
	var dist = -1.0
	var _dist float64
	for a := 0; !bln && a < len(segsa); a++ {
		for b := 0; !bln && b < len(segsb); b++ {
			bln = segsa[a].Intersects(segsb[b], false)
			if bln {
				dist = 0.0
			} else {
				_dist = segsa[a].Distance(segsb[b])
				if dist < 0 {
					dist = _dist
				} else {
					dist = math.Min(_dist, dist)
				}
			}
		}
	}
	return dist
}
