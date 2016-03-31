package geom

// description computes the convex hull of a point set.
// param points An array of [X, Y] coordinates

func ConvexHull(points []*Point) []*Point {
    //copy points into mutable container
    pnts := CloneCoordinates(points)
    var N = len(pnts)
    //trivial case 0 or 1 point
    if len(pnts) < 2 {
        return pnts
    }
    XYCoordinates{pnts}.Sort()

    var lower = make(Coordinates, 0)
    var upper = make(Coordinates, 0)
    lower = build_hull(lower, pnts, 0, 1, N)
    upper = build_hull(upper, pnts, N - 1, -1, -1)
    _, upper = upper.Pop()
    _, lower = lower.Pop()

    //two reapeated pnts
    if len(upper) == 1 && len(lower) == 1 {
        if (upper[0][x] == lower[0][x]) &&  (upper[0][y] == lower[0][y]) {
            _, upper = upper.Pop()
        }
    }

    var hull = append(lower, upper...)
    //close ring
    if len(hull) > 1 {
        hull = append(hull, hull[0])
    }

    return hull
}

//build boundary
func build_hull(hb, points Coordinates, start, step, stop int) Coordinates {
    var pnt *Point
    var i = start
    for i != stop {
        pnt = points[i]
        //pnt.CrossProduct(boundary[n - 2], boundary[n - 1])
        for n := len(hb);
        n >= 2 && pnt.SideOf(hb[n - 2], hb[n - 1]).IsOnOrRight(); n =len(hb) {
            _, hb = hb.Pop()
        }
        hb = append(hb, pnt)
        i += step
    }
    return hb
}

/*
 orient boundary
 param array
 param pnt
 returns {*}
 */
func orient(array Coordinates, pnt *Point) float64 {
    var n = len(array)
    return pnt.CrossProduct(array[n - 2], array[n - 1])
}




// simpleHull_2D(): Melkman's 2D simple polyline O(n) convex hull algorithm
//    Input:  coords[] = array of 2D vertex points for a simple polyline
//            n   = the number of points in V[]
//    Output: H[] = output convex hull array of vertices (max is n)
//    Return: h   = the number of points in H[]
//http://geomalgorithms.com/a12-_hull-3.html

func SimpleHull(coords []*Point) []*Point {
    var n = len(coords)
    // initialize a deque dQ[] from bottom to top so that the
    // 1st three vertices of coords[] are a ccw triangle
    var dQ = make([]*Point, 2 * n + 1);
    var bot = n - 2
    var top = bot + 3    // initial bottom and top deque indices
    dQ[bot], dQ[top] = coords[2], coords[2]; // 3rd vertex is at both bot and top
    if coords[2].SideOf(coords[0], coords[1]).IsLeft() {
        dQ[bot + 1] = coords[0];
        dQ[bot + 2] = coords[1]; // ccw vertices are: 2,0,1,2
    } else {
        dQ[bot + 1] = coords[1];
        dQ[bot + 2] = coords[0]; // ccw vertices are: 2,1,0,2
    }

    // compute the hull on the deque dQ[]
    for i := 3; i < n; i++ {
        // process the rest of vertices
        // test if next vertex is inside the deque hull
        if (//is left
        coords[i].SideOf(dQ[bot], dQ[bot + 1]).IsLeft() &&
        coords[i].SideOf(dQ[top - 1], dQ[top]).IsLeft()) {
            continue; // skip an interior vertex
        }

        // incrementally add an exterior vertex to the deque hull
        // get the rightmost tangent at the deque bot
        //coords[i] right of dq[bot] --- dq[bot+1]
        for coords[i].SideOf(dQ[bot], dQ[bot + 1]).IsOnOrRight() {
            bot++; // remove bot of deque
        }
        bot--
        dQ[bot] = coords[i]; // insert coords[i] at bot of deque

        // get the leftmost tangent at the deque top
        for (coords[i].SideOf(dQ[top - 1], dQ[top]).IsOnOrRight()) {
            top--; // pop top of deque
        }
        top++
        dQ[top] = coords[i]; // push coords[i] onto top of deque
    }

    // transcribe deque dQ[] to the output hull array H[]
    n = (top - bot) + 1
    var hull = make([]*Point, n)
    var h int; // hull vertex counter
    for h = 0; h < n; h++ {
        hull[h] = dQ[bot + h];
    }

    return hull
}






