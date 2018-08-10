package geom

import (
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/mbr"
)

type LineString struct {
	Coordinates Coords
	bbox        mono.MBR
	rbEvents    []float64
}

//New LineString from a given Coords {Array} [[x,y], ....[x,y]]
func NewLineString(coordinates Coords) *LineString {
	var n = coordinates.Len()
	if n < 2 {
		panic("a linestring must have at least 2 coordinates")
	}
	var ln = &LineString{
		Coordinates: coordinates,
		rbEvents:    make([]float64, 0, 2*(n-1)),
	}
	return ln.prepEvents()
}

func (self *LineString) prepEvents() *LineString {
	var n = self.Coordinates.Len() - 1
	var a, b *Point
	var x, y float64
	a = self.Coordinates.Pt(0)
	self.bbox.MBR = mbr.MBR{a[0], a[1], a[0], a[1]}
	self.bbox.I = self.Coordinates.Idxs[0]
	self.bbox.J = self.Coordinates.Idxs[n]

	for i := 0; i < n; i++ {
		a, b = self.Coordinates.Pt(i), self.Coordinates.Pt(i+1)
		x, y = a[0], b[0]
		self.rbEvents = append(self.rbEvents, minf64(x, y))
		self.rbEvents = append(self.rbEvents, maxf64(x, y))
		self.bbox.MBR.ExpandIncludeXY(b[0], b[1])
	}
	return self
}

func redblueLineSegmentIntersection(red, blue *LineString, visit func(int, int) bool) bool {

	var nr = red.Coordinates.Len() - 1
	var nb = blue.Coordinates.Len() - 1
	var n = nr + nb
	var ne = 2 * n
	var ret bool

	var redList = createBruteForceList(nr)
	var blueList = createBruteForceList(nb)

	var events = prepareEvents(red, blue)

	var ev, index int

	for i := 0; !ret && i < ne; i++ {
		ev, index = events[i].ev, events[i].idx

		if ev == CreateRED {
			ret = addSegment(index, red, &redList, blue, &blueList, visit, false)
		} else if ev == CreateBLUE {
			ret = addSegment(index, blue, &blueList, red, &redList, visit, true)
		} else if ev == RemoveRED {
			redList.remove(index)
		} else if ev == RemoveBLUE {
			blueList.remove(index)
		}
	}

	return ret
}

//New line string from array
func NewLineStringFromArray(array Coords) *LineString {
	return NewLineString(array)
}

//create a new linestring from wkt string
//empty wkt are not allowed
func NewLineStringFromWKT(wkt string) *LineString {
	return NewLineStringFromArray(
		readWKT(wkt, GeoTypeLineString).ToCoordinates()[0],
	)
}

//Point to LineString
func NewLineStringFromPoint(pt Point) *LineString {
	return NewLineString(Coordinates([]Point{pt, pt}))
}

//Point at index i
func (self *LineString) Pt(i int) *Point {
	return self.Coordinates.Pt(i)
}

//ConvexHull computes slice of vertices as points forming convex hull
func (self *LineString) ConvexHull() *Polygon {
	return NewPolygon(ConvexHull(self.Coordinates))
}

//number of vertices
func (self *LineString) LenVertices() int {
	return self.Coordinates.Len()
}
