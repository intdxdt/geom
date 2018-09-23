package geom

import (
	"github.com/intdxdt/geom/index"
	"github.com/intdxdt/geom/mono"
)

type LineString struct {
	Coordinates Coords
	bbox        mono.MBR
	chains      []mono.MBR
	index       *index.Index
}

//New LineString from a given Coords {Array} [[x,y], ....[x,y]]
func NewLineString(coordinates Coords) *LineString {
	if coordinates.Len() < 2 {
		panic("a linestring must have at least 2 coordinates")
	}
	var ln = &LineString{
		Coordinates: coordinates,
		index:       index.NewIndex(),
	}
	return ln.processChains().buildIndex()
}

//New line string from array
func NewLineStringFromArray(array [][]float64) *LineString {
	return NewLineString(CoordinatesFromArray(array))
}

//create a new linestring from wkt string
//empty wkt are not allowed
func NewLineStringFromWKT(wkt string) *LineString {
	return NewLineString(
		ReadWKT(wkt, GeoTypeLineString).ToCoordinates()[0],
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

//builds rtree index of chains
func (self *LineString) buildIndex() *LineString {
	if !self.index.IsEmpty() {
		self.index.Clear()
	}
	self.index.Load(self.chains) //bulkload
	return self
}

//get copy of chains of linestring
func (self *LineString) MonoChains() []mono.MBR {
	var chains = make([]mono.MBR, 0, len(self.chains))
	for i := range self.chains {
		chains = append(chains, self.chains[i])
	}
	return chains
}

//ConvexHull computes slice of vertices as points forming convex hull
func (self *LineString) ConvexHull() *Polygon {
	return NewPolygon(ConvexHull(self.Coordinates))
}

//number of vertices
func (self *LineString) LenVertices() int {
	return self.Coordinates.Len()
}
