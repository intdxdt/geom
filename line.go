package geom

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/rtree"
)

const bucketSize = 8

type LineString struct {
	chains      []MonoMBR
	coordinates []Point
	length      float64
	monosize    int
	bbox        MonoMBR
	index       *rtree.RTree
}

//New LineString from a given coordinates {Array} [[x,y], ....[x,y]]
//optional clone coords : make a copy of input coordinates
func NewLineString(coordinates []Point) *LineString {
	var n = len(coordinates)
	if n < 2 {
		panic("a linestring must have at least 2 coordinate")
	}
	var mSize = int(math.Log2(float64(n) + 1.0))
	var ln = LineString{
			chains:      make([]MonoMBR, 0, mSize),
			coordinates: coordinates[:n:n],
			monosize:    mSize,
			index:       rtree.NewRTree(bucketSize),
		}
	ln.processChains(0, n-1)
	ln.buildIndex()
	return &ln
}

//New line string from array
func NewLineStringFromArray(array [][]float64) *LineString {
	return NewLineString(AsPointArray(array))
}

//create a new linestring from wkt string
//empty wkt are not allowed
func NewLineStringFromWKT(wkt string) *LineString {
	return NewLineStringFromArray(
		readWKT(wkt, GeoTypeLineString).ToArray()[0],
	)
}

//Point to LineString
func NewLineStringFromPoint(pt Point) *LineString {
	return NewLineString([]Point{pt, pt})
}

//builds rtree index of chains
func (self *LineString) buildIndex() *LineString {
	if !self.index.IsEmpty() {
		self.index.Clear()
	}
	var data = make([]*rtree.Obj, 0, len(self.chains))
	for i := range self.chains {
		data = append(data, rtree.Object(i, self.chains[i].MBR, self.chains[i]))
	}
	self.index.Load(data) //bulkload
	return self
}

//get copy of chains of linestring
func (self *LineString) MonoChains() []MonoMBR {
	var chains = make([]MonoMBR, 0, len(self.chains))
	for i := range self.chains {
		chains = append(chains, self.chains[i])
	}
	return chains
}

//ConvexHull computes slice of vertices as points forming convex hull
func (self *LineString) ConvexHull() *Polygon {
	return NewPolygon(ConvexHull(self.coordinates))
}

//number of vertices
func (self *LineString) LenVertices() int {
	return len(self.coordinates)
}

//vertex at given index
func (self *LineString) VertexAt(i int) *Point {
	return &self.coordinates[i]
}



