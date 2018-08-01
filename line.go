package geom

import (
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/geom/index"
)

type LineString struct {
	chains      []mono.MBR
	Coordinates Coords
	index       *index.Index
	bbox        mono.MBR
}

//New LineString from a given Coords {Array} [[x,y], ....[x,y]]
func NewLineString(coordinates Coords) *LineString {
	//var n =
	if coordinates.Len() < 2 {
		panic("a linestring must have at least 2 coordinates")
	}
	var ln = &LineString{
		Coordinates: coordinates,
		index:       index.NewIndex(),
	}
	return ln.processChains().buildIndex()
}

//New LineString from a given Coords
func NewLineStringFromCoords(coordinates Coords) *LineString {
	if coordinates.Len() < 2 {
		panic("a linestring must have at least 2 coordinate")
	}
	var ln = &LineString{
		Coordinates: coordinates,
		index:       index.NewIndex(),
	}
	return ln.processChains().buildIndex()
}

//New line string from array
func NewLineStringFromArray(array Coords) *LineString {
	return NewLineStringFromCoords(array)
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
	return NewPolygonFromCoords(ConvexHull(self.Coordinates))
}

//number of vertices
func (self *LineString) LenVertices() int {
	return self.Coordinates.Len()
}
