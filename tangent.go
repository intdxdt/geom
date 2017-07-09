package geom

import (
	"robust"
)

//true if Vi is above Vj
func above(P, Vi, Vj *Point) bool {
	return isLeft(P, Vi, Vj) > 0
}

//true if Vi is below Vj
func below(P, Vi, Vj *Point) bool {
	return isLeft(P, Vi, Vj) < 0
}

//isLeft(): test if a point is Left|On|Right of an infinite line.
//Input:  three points P0, P1, && P2
//Return: >0 for c left of the line through a && b
//        =0 for c on the line
//        <0 for c right of the line
func isLeft(a, b, c *Point) float64 {
	o := robust.Orientation2D(a[:], b[:], c[:])
	if o < 0 {
		o = 1
	} else if o > 0 {
		o = -1
	}
	return o
}

//tangent_PointPoly(): find any polygon's exterior tangents
//Input:  P = a 2D point (exterior to the polygon)
//        V = array of vertices for any 2D polygon with V[n]=V[0]
//Output: rtan = index of rightmost tangent point V[rtan]
//        ltan = index of leftmost tangent point  V[ltan]
func TangentPointToPoly(pt *Point, coords [] *Point) (int, int) {
	if !IsRing(coords) {
		coords = CloneCoordinates(coords)
		coords = append(coords, coords[0].Clone())
	}
	v := coords
	n := len(v) - 1

	// eprev, enext  - V[i] previous && next edge turn direction
	eprev := isLeft(v[0], v[1], pt)
	var rtan, ltan int // initially assume V[0] = both tangents

	for i := 1; i < n; i++ {
		enext := isLeft(v[i], v[i+1], pt)

		if (eprev <= 0) && (enext > 0) {
			if !below(pt, v[i], v[rtan]) {
				rtan = i
			}
		} else if (eprev > 0) && (enext <= 0) {
			if !above(pt, v[i], v[ltan]) {
				ltan = i
			}
		}

		eprev = enext
	}

	return rtan, ltan
}
