package geom

import (
	"github.com/intdxdt/math"
)

// Project point on linestring
func (self *LineString) Project(point *Point, normalized ...bool) float64 {
	var bln = false
	var distSqr = math.MaxFloat64
	var coords = self.Coordinates
	var n = coords.Len() - 1
	var di, dj int
	var val, dist float64
	var a, b, dPt, pt *Point
	var norm bool
	if len(normalized) > 0 {
		norm = normalized[0]
	}

	for i, j := 0, 0; !bln && i < n; i++ {
		j = i + 1
		val, pt = distanceToPoint(coords.Pt(i), coords.Pt(j), point, hypotSqr)
		if val < distSqr {
			di, dj, dPt, distSqr = i, j, pt, val
		}
		bln = distSqr == 0
	}

	for i := 0; i < di; i++ {
		a, b = coords.Pt(i), coords.Pt(i+1)
		dist += a.Magnitude(b)
	}

	a, b = coords.Pt(di), coords.Pt(dj)
	if a == dPt {
		dist += 0
	} else {
		dist += a.Magnitude(dPt)
	}

	if norm {
		return dist / self.Length()
	}
	return dist
}
