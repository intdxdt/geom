package geom

import (
	"github.com/intdxdt/math"
)

// Project point on linestring
func (self *LineString) Project(point *Point, normalized ...bool) float64 {
	var distSqr = math.MaxFloat64
	var coords = self.Coordinates
	var di, dj int
	var val, dist float64
	var a, b, dPt, pt *Point
	var normValue = 1.0
	if len(normalized) > 0 && normalized[0] {
		normValue = self.Length()
	}

	var intersects = false
	var n = coords.Len() - 1
	for i, j := 0, 0; !intersects && i < n; i++ {
		j = i + 1
		val, pt = distanceToPoint(coords.Pt(i), coords.Pt(j), point, hypotSqr)
		if val < distSqr {
			di, dj, dPt, distSqr = i, j, pt, val
		}
		intersects = distSqr == 0
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

	return dist / normValue
}

// Interpolate point on linestring given distance along linear string
func (self *LineString) Interpolate(distance float64, normalized ...bool) *Point {
	var coords = self.Coordinates
	if len(normalized) > 0 && normalized[0] {
		distance = self.Length() * distance
	}

	var dist float64
	var idxDist float64
	var n = coords.Len() - 1
	var idx = 0
	for i := 0; i < n; i++ {
		dist += coords.Pt(i).Magnitude(coords.Pt(i + 1))
		if dist < distance {
			idx = i + 1
			idxDist = dist
		} else {
			break
		}
	}

	var pt *Point
	if math.FloatEqual(distance, 0) {
		pt = coords.Pt(0).Clone()
	} else if math.FloatEqual(dist, distance) {
		pt = coords.Pt(idx + 1).Clone()
	} else {
		var delta = distance - idxDist
		var a = coords.Pt(idx)
		var b = coords.Pt(idx + 1)
		var vx, vy = b.Sub(a[X], a[Y])
		var cx, cy = Extend(vx, vy, delta, 0, false)
		pt = &Point{a[X] + cx, a[Y] + cy}
	}

	return pt
}
