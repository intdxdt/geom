package geom

import (
	"sort"
	"github.com/intdxdt/math"
)

type Coordinates []*Point

//len of coordinates - sort interface
func (s Coordinates) Len() int {
	return len(s)
}
//swap - sort interface
func (s Coordinates) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//less - 2d compare - sort interface
func (s Coordinates) Less(i, j int) bool {
	return (s[i][0] < s[j][0]) || (
		math.FloatEqual(s[i][0], s[j][0]) && s[i][1] < s[j][1])
}

//2D sort
func (s Coordinates) Sort() Coordinates{
	sort.Sort(s)
	return s
}

//pop point from
func (s Coordinates) Pop() (*Point, Coordinates) {
	var v *Point
	var n int
	if len(s) == 0 {
		return nil, s
	}
	n = len(s) - 1
	v, s[n] = s[n], nil
	return v, s[:n]
}


//get copy of coordinates of linestring
func (self *Point) Coordinates() *Point {
	return self.Clone()
}

//get copy of coordinates of linestring
func (self *LineString) Coordinates() []*Point {
	return CloneCoordinates(self.coordinates)
}

//get copy of coordinates of polygon
func (self *Polygon) Coordinates() [][]*Point {
	lns := self.AsLinear()
	coords := make([][]*Point, len(lns))
	for i, ln := range lns {
		coords[i] = ln.Coordinates()
	}
	return coords
}

//checks if a point is a ring , by def every point is a ring
// which concides on itself
func (self *Point) IsRing() bool {
	return true
}

//Checks if line string is a ring
func (self *LineString) IsRing() bool {
	return IsRing(self.coordinates)
}

//Checks if polygon is a ring - default to true since all polygons are closed ring(s)
func (self *Polygon) IsRing() bool {
	return true
}

//------------------------------------------------------------------------------
//Is coordinates a ring : P0 == Pn
func IsRing(coordinates []*Point) bool {
	if len(coordinates) < 2 {
		return false
	}
	return coordinates[0].Equals2D(
		coordinates[len(coordinates)-1],
	)
}

//Coordinates returns a copy of linestring coordinates
func CloneCoordinates(coordinates []*Point) []*Point {
	n := len(coordinates)
	clone := make([]*Point, n)
	for i := 0; i < n; i++ {
		clone[i] = coordinates[i].Clone()
	}
	return clone
}

//
func CoordinatesAsFloat2D(coordinates []*Point) [][]float64 {
	var n = len(coordinates)
	var coords = make([][]float64, n)
	for i := 0; i < n; i++ {
		coords[i] = []float64{coordinates[i][X], coordinates[i][Y]}
	}
	return coords
}
