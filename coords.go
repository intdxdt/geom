package geom

import (
	"sort"
)

func Coordinates(c []Point) Coords {
	var n = len(c)
	var coords = Coords{_c: c[:n:n], Idxs: make([]int, len(c))}
	for i := range coords._c {
		coords.Idxs[i] = i
	}
	return coords
}

type Coords struct {
	_c   []Point
	Idxs []int
}

//Point at index
func (s *Coords) Pt(i int) *Point {
	return &s._c[s.Idxs[i]]
}

//Point at index
func (s Coords) Points() []Point {
	var pts = make([]Point, 0, len(s.Idxs))
	for _, i := range s.Idxs {
		pts = append(pts, s._c[i])
	}
	return pts
}

//Point at index 0
func (s *Coords) First() *Point {
	return &s._c[s.Idxs[0]]
}

//Point at index 0
func (s *Coords) Last() *Point {
	return &s._c[s.Idxs[s.Len()-1]]
}

//len of Coords - sort interface
func (s Coords) Len() int {
	return len(s.Idxs)
}

//swap - sort interface
func (s Coords) Swap(i, j int) {
	s.Idxs[i], s.Idxs[j] = s.Idxs[j], s.Idxs[i]
}

//less - 2d compare - sort interface
func (s Coords) Less(i, j int) bool {
	i, j = s.Idxs[i], s.Idxs[j]
	return (s._c[i][0] < s._c[j][0]) ||
		(feq(s._c[i][0], s._c[j][0]) && s._c[i][1] < s._c[j][1])
}

//2D sort
func (s *Coords) Sort() *Coords {
	sort.Sort(s)
	return s
}

//pop point from
func (s *Coords) Pop() (bool, Point) {
	var v Point
	var n int
	if len(s.Idxs) == 0 {
		return false, NullPt
	}
	n = len(s.Idxs) - 1
	v, s.Idxs[n] = s._c[s.Idxs[n]], -1
	s.Idxs = s.Idxs[:n]
	return true, v
}

//get copy of Coords of polygon
func (self *Polygon) Coordinates() []Coords {
	var lns = self.AsLinear()
	var coords = make([]Coords, len(lns))
	for i, ln := range lns {
		coords[i] = ln.Coordinates
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
	return IsRing(self.Coordinates)
}

//Checks if polygon is a ring - default to true since all polygons are closed ring(s)
func (self *Polygon) IsRing() bool {
	return true
}

//------------------------------------------------------------------------------
//Is Coords a ring : P0 == Pn
func IsRing(coordinates Coords) bool {
	if coordinates.Len() < 2 {
		return false
	}
	return coordinates.First().Equals2D(coordinates.Last())
}
