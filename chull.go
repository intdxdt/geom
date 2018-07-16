package geom

// description computes the convex hull of a point set.
// param points An array of [X, Y] coordinates
func ConvexHull(points []Point, clone_coords ...bool) []Point {
	var clone = true
	if len(clone_coords) > 0 {
		clone = clone_coords[0]
	}
	var pnts = points
	//copy points into mutable container
	if clone {
		pnts = CloneCoordinates(points)
	}

	//trivial case less than three coordinates
	if len(points) < 3 {
		return pnts
	}
	var N = len(pnts)

	Coordinates(pnts).Sort()

	var lower = make(Coordinates, 0, N/2)
	var upper = make(Coordinates, 0, N/2)
	lower = build_hull(lower, pnts, 0, 1, N)
	upper = build_hull(upper, pnts, N-1, -1, -1)
	_, upper = upper.Pop()
	_, lower = lower.Pop()

	for _, v := range upper {
		lower = append(lower, v)
	}

	return lower
}

//build boundary
func build_hull(hb, points Coordinates, start, step, stop int) Coordinates {
	var pnt *Point
	var i = start
	for i != stop {
		pnt = &points[i]
		//pnt.CrossProduct(boundary[n - 2], boundary[n - 1])
		for n := len(hb); n >= 2 &&
			pnt.SideOf(&hb[n-2], &hb[n-1]).IsOnOrRight(); n = len(hb) {
			_, hb = hb.Pop()
		}
		hb = append(hb, *pnt)
		i += step
	}
	return hb
}

//SimpleHull(): Melkman's 2D simple polyline O(n) convex hull algorithm
//Input:  coords[] = array of 2D vertex points for a simple polyline
//             n   = the number of points in V[]
//Output: H[] = output convex hull array of vertices (max is n)
//Return: h   = the number of points in H[]
//http://geomalgorithms.com/a12-_hull-3.html
func SimpleHull(coords []Point, clone_coords ...bool) []Point {
	var clone = true
	if len(clone_coords) > 0 {
		clone = clone_coords[0]
	}
	//copy points into mutable container
	if clone {
		coords = CloneCoordinates(coords)
	}

	//trivial case less than three coordinates
	if len(coords) < 3 {
		return coords
	}

	var n = len(coords)
	// initialize a deque dQ[] from bottom to top so that the
	// 1st three vertices of coords[] are a ccw triangle
	var dQ = make([]Point, 2*n+1)
	var bot = n - 2
	var top = bot + 3                       // initial bottom and top deque indices
	dQ[bot], dQ[top] = coords[2], coords[2] // 3rd vertex is at both bot and top

	if coords[2].SideOf(&coords[0], &coords[1]).IsLeft() {
		dQ[bot+1] = coords[0]
		dQ[bot+2] = coords[1] // ccw vertices are: 2,0,1,2
	} else {
		dQ[bot+1] = coords[1]
		dQ[bot+2] = coords[0] // ccw vertices are: 2,1,0,2
	}

	// compute the hull on the deque dQ[]
	for i := 3; i < n; i++ {
		// process the rest of vertices
		// test if next vertex is inside the deque hull
		//is left
		if coords[i].SideOf(&dQ[bot], &dQ[bot+1]).IsLeft() &&
			coords[i].SideOf(&dQ[top-1], &dQ[top]).IsLeft() {
			continue // skip an interior vertex
		}

		// incrementally add an exterior vertex to the deque hull
		// get the rightmost tangent at the deque bot
		//coords[i] right of dq[bot] --- dq[bot+1]
		for coords[i].SideOf(&dQ[bot], &dQ[bot+1]).IsOnOrRight() {
			bot++ // remove bot of deque
		}
		bot--
		dQ[bot] = coords[i] // insert coords[i] at bot of deque

		// get the leftmost tangent at the deque top
		for coords[i].SideOf(&dQ[top-1], &dQ[top]).IsOnOrRight() {
			top-- // pop top of deque
		}
		top++
		dQ[top] = coords[i] // push coords[i] onto top of deque
	}

	// transcribe deque dQ[] to the output hull array H[]
	var hull []Point
	for h := 0; h <= (top - bot); h++ {
		hull = append(hull, dQ[bot+h])
	}
	return hull
}
