package geom

import (
	"github.com/intdxdt/geom/index"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/math"
	"github.com/intdxdt/mbr"
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

//Computes the distance between geometries
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

//Computes the distance between a linestring and another linestring
//the distance between intersecting linestrings is 0.  Otherwise, the
//distance is the Euclidean distance between the closest segments.
func (self *LineString) line_line_dist(other *LineString) float64 {
	if self.Coordinates.Len() < 16 && other.Coordinates.Len() < 16 {
		return self.mindistBruteforce(other)
	}
	return knnMinLinearDistance(self.Coordinates, other.Coordinates)

}

// brute force distance
func (self *LineString) mindistBruteforce(other *LineString) float64 {
	var dist = math.MaxFloat64
	var bln = false
	var ln = self.Coordinates
	var ln2 = other.Coordinates
	var n1, n2 = ln.Len() - 1, ln2.Len() - 1
	var d float64
	for i := 0; !bln && i < n1; i++ {
		for j := 0; !bln && j < n2; j++ {
			d = SegSegDistance(ln.Pt(i), ln.Pt(i+1), ln2.Pt(j), ln2.Pt(j+1))
			if d < dist {
				dist = d
			}
			//dist = minf64(dist, d)
			bln = dist == 0
		}
	}
	return dist
}

func knnMinLinearDistance(a, b Coords) float64 {
	if a.Len() > b.Len() {
		a, b = b, a
	}
	var db = segmentDB(b)
	var queries = queryBounds(a)

	var dist = math.MaxFloat64
	var d float64
	for q := range queries {
		db.KnnMinDist(&queries[q],
			func(query *mono.MBR, item *mono.MBR) (float64, float64) {
				d = SegSegDistance(
					&a.Pnts[query.I], &a.Pnts[query.J], &b.Pnts[item.I], &b.Pnts[item.J],
				)

				if d < dist {
					dist = d
				}
				return d, dist
			},
			func(o *index.KObj) bool {
				return o.Distance > dist || dist == 0 //add to neibs, stop
			})
	}

	return dist
}

func queryBounds(coords Coords) []mono.MBR {
	var n = coords.Len() - 1
	var I, J int
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		I, J = coords.Idxs[i], coords.Idxs[i+1]
		items = append(items, mono.MBR{
			MBR: mbr.CreateMBR(
				coords.Pnts[I][X], coords.Pnts[I][Y],
				coords.Pnts[J][X], coords.Pnts[J][Y]), I: I, J: J,
		})
	}
	return items
}

func segmentDB(coords Coords) *index.Index {
	var tree = index.NewIndex()
	var n = coords.Len() - 1
	var I, J int
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		I, J = coords.Idxs[i], coords.Idxs[i+1]
		items = append(items, mono.MBR{
			MBR: mbr.CreateMBR(
				coords.Pnts[I][X], coords.Pnts[I][Y],
				coords.Pnts[J][X], coords.Pnts[J][Y]), I: I, J: J,
		})
	}
	tree.Load(items)
	return tree
}
