package geom

// description computes the convex hull of a point set.
// param points An array of [X, Y] coordinates
func ConvexHull(points Coords) Coords {
	var pnts = points.ShallowClone()
	//trivial case less than three coordinates
	if points.Len() < 3 {
		return pnts
	}
	var N = pnts.Len()

	pnts.Sort()

	var lower = makeCoords(pnts, 0, N/2)
	var upper = makeCoords(pnts, 0, N/2)

	lower = buildHull(lower, pnts, 0, 1, N)
	upper = buildHull(upper, pnts, N-1, -1, -1)

	upper.Pop()
	lower.Pop()

	for _, v := range upper.Idxs {
		lower.Idxs = append(lower.Idxs, v)
	}

	return lower
}

//build boundary
func buildHull(hb, points Coords, start, step, stop int) Coords {
	var pnt *Point
	var i = start
	var idx int
	for i != stop {
		idx, pnt = points.Idxs[i], points.Pt(i)
		//pnt.CrossProduct(boundary[n - 2], boundary[n - 1])
		for n := hb.Len(); hb.Len() >= 2 && pnt.SideOf(hb.Pt(n-2), hb.Pt(n-1)).IsOnOrRight(); n = hb.Len() {
			hb.Pop()
		}
		hb.Idxs = append(hb.Idxs, idx)
		i += step
	}
	return hb
}

//Coords returns a copy of linestring coordinates
func makeCoords(coordinates Coords, i, j int) Coords {
	var o = Coords{Pnts: coordinates.Pnts, Idxs: make([]int, 0, j-i+1)}
	return o
}

