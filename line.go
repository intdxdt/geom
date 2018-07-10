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
	return (
		&LineString{
			chains:      make([]MonoMBR, 0, mSize),
			coordinates: coordinates[:n:n],
			monosize:    mSize,
			index:       rtree.NewRTree(bucketSize),
		}).process_chains(0, n-1).build_index()
}

//New line string from array
func NewLineStringFromArray(array [][]float64) *LineString {
	return NewLineString(AsPointArray(array))
}

//create a new linestring from wkt string
//empty wkt are not allowed
func NewLineStringFromWKT(wkt_geom string) *LineString {
	return NewLineStringFromArray(
		ReadWKT(wkt_geom).ToArray()[0],
	)
}

//Point to LineString
func NewLineStringFromPoint(pt Point) *LineString {
	return NewLineString([]Point{pt, pt})
}

//get copy of chains of linestring
func (self *LineString) MonoChains() []*MonoMBR {
	chains := make([]*MonoMBR, len(self.chains))
	for i := range self.chains {
		chains[i] = self.chains[i].Clone()
	}
	return chains
}
