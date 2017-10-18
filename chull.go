package geom

import (
	"github.com/intdxdt/cart"
	"github.com/intdxdt/math"
)

const RightAngle = 90.0
const LeftAngle = 270.0

type Hull struct {
	H []*Point
}

func NewHull(coords []*Point) *Hull {
	return &Hull{H: coords}
}

func (self *Hull) Antipodal(i, j int) int {
	var n = len(self.H) - 1
	var fn = self.chainIndexer(i, n)
	var idxer = self.indexer(i, n)

	var ptI, ptJ = self.H[i], self.H[j]
	var cmpIJ = ptJ.Sub(ptI)

	var start, end = fn(i), fn(i-1)

	var mid = (start + end) / 2
	var pt = self.H[idxer(mid)]
	var uvect = func(m int) *Point {
		return self.H[m].Sub(ptJ)
	}

	var side = pt.SideOf(ptI, ptJ)

	var angl = math.Deg2rad(RightAngle)
	if side.IsOn() {
		return idxer(end)
	} else if side.IsLeft() {
		angl = math.Deg2rad(LeftAngle)
	}

	orth := self.orthvector(cmpIJ, angl)

	for {
		if start == end {
			mid = start
			break
		}
		mid = (start + end) / 2

		cur := self.offset(uvect(idxer(mid)), orth)
		next := self.offset(uvect(idxer(mid+1)), orth)

		if math.FloatEqual(cur, next) {
			m := idxer(mid)
			m1 := idxer(mid + 1)
			d := cart.DistanceToPoint(ptI, ptJ, self.H[m])
			d1 := cart.DistanceToPoint(ptI, ptJ, self.H[m1])
			if math.FloatEqual(d, d1) || d1 > d {
				mid += 1
			}
			break
		} else {
			if cur < next {
				start = mid + 1
			} else if cur > next {
				end = mid
			}
		}
	}
	return idxer(mid)
}

func (self *Hull) offset(u, v *Point) float64 {
	return cart.Project(u, v)
}

func (self *Hull) orthvector(v *Point, angle float64) *Point {
	cx, cy := cart.Extend(v, 1.0, angle, true)
	return NewPointXY(cx, cy)
}

func (self *Hull) indexer(origin, max int) func(k int) int {
	return func(k int) int {
		if k >= origin && k <= max {
			return k
		} else if k > max {
			return k - max - 1
		}
		panic("index out of bounds")
	}
}

func (self *Hull) chainIndexer(origin, max int) func(k int) int {
	return func(k int) int {
		if k >= origin && k <= max {
			return k
		} else if k < origin {
			return max + k + 1
		}
		panic("index out of bounds")
	}
}

// description computes the convex hull of a point set.
// param points An array of [X, Y] coordinates
func ConvexHull(points []*Point, clone_coords ...bool) []*Point {
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

	XYCoordinates{pnts}.Sort()

	var lower = make(Coordinates, 0)
	var upper = make(Coordinates, 0)
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
		pnt = points[i]
		//pnt.CrossProduct(boundary[n - 2], boundary[n - 1])
		for n := len(hb); n >= 2 &&
			pnt.SideOf(hb[n-2], hb[n-1]).IsOnOrRight(); n = len(hb) {
			_, hb = hb.Pop()
		}
		hb = append(hb, pnt)
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
func SimpleHull(coords []*Point, clone_coords ...bool) []*Point {
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
	var dQ = make([]*Point, 2*n+1)
	var bot = n - 2
	var top = bot + 3                       // initial bottom and top deque indices
	dQ[bot], dQ[top] = coords[2], coords[2] // 3rd vertex is at both bot and top

	if coords[2].SideOf(coords[0], coords[1]).IsLeft() {
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
		if //is left
		coords[i].SideOf(dQ[bot], dQ[bot+1]).IsLeft() &&
			coords[i].SideOf(dQ[top-1], dQ[top]).IsLeft() {
			continue // skip an interior vertex
		}

		// incrementally add an exterior vertex to the deque hull
		// get the rightmost tangent at the deque bot
		//coords[i] right of dq[bot] --- dq[bot+1]
		for coords[i].SideOf(dQ[bot], dQ[bot+1]).IsOnOrRight() {
			bot++ // remove bot of deque
		}
		bot--
		dQ[bot] = coords[i] // insert coords[i] at bot of deque

		// get the leftmost tangent at the deque top
		for coords[i].SideOf(dQ[top-1], dQ[top]).IsOnOrRight() {
			top-- // pop top of deque
		}
		top++
		dQ[top] = coords[i] // push coords[i] onto top of deque
	}

	// transcribe deque dQ[] to the output hull array H[]
	var hull = make([]*Point, 0)
	for h := 0; h <= (top - bot); h++ {
		hull = append(hull, dQ[bot+h])
	}
	return hull
}
