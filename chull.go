package geom

// description computes the convex hull of a point set.
// param points An array of [X, Y] coordinates
func ConvexHull(points Coords) Coords {
	var pnts = ShallowClone(points)
	//trivial case less than three coordinates
	if points.Len() < 3 {
		return pnts
	}
	var N = pnts.Len()

	pnts.Sort()

	var lower = makeCoords(pnts, 0, N/2)
	var upper = makeCoords(pnts, 0, N/2)

	lower = build_hull(lower, pnts, 0, 1, N)
	upper = build_hull(upper, pnts, N-1, -1, -1)

	upper.Pop()
	lower.Pop()

	for _, v := range upper.Idxs {
		lower.Idxs = append(lower.Idxs, v)
	}

	return lower
}

//build boundary
func build_hull(hb, points Coords, start, step, stop int) Coords {
	var pnt *Point
	var i = start
	var idx int
	for i != stop {
		idx, pnt = points.Idxs[i], points.Pt(i)
		//pnt.CrossProduct(boundary[n - 2], boundary[n - 1])
		for n := hb.Len(); n >= 2 && pnt.SideOf(hb.Pt(n-2), hb.Pt(n-1)).IsOnOrRight(); n = hb.Len() {
			hb.Pop()
		}
		hb.Idxs = append(hb.Idxs, idx)
		i += step
	}
	return hb
}

//Coords returns a copy of linestring coordinates
func makeCoords(coordinates Coords, i, j int) Coords {
	var o = Coords{_c: coordinates._c, Idxs: make([]int, 0, j-i+1)}
	return o
}

//Coords returns a copy of linestring coordinates
func ShallowClone(coordinates Coords, slice ...int) Coords {
	var i, j = 0, coordinates.Len()
	if len(slice) == 1 {
		j = slice[0]
	} else if len(slice) > 1 {
		i, j = slice[0], slice[1]
	}

	var o = Coords{_c: coordinates._c, Idxs: make([]int, 0,  j-i)}
	for _, v := range coordinates.Idxs[i:j]{
		o.Idxs = append(o.Idxs, v)
	}
	return o
}
