package geom

import (
	"sort"
)

func Coordinates(c []Point) Coords {
	var n = len(c)
	var coords = Coords{Pnts: c[:n:n], Idxs: make([]int, n)}
	for i := range coords.Pnts {
		coords.Idxs[i] = i
	}
	return coords
}

func CoordinatesFromArray(array [][]float64) Coords {
	var pts = PointsFromArray(array)
	var coords = Coords{Pnts: pts, Idxs: make([]int, len(pts))}
	for i := range coords.Pnts {
		coords.Idxs[i] = i
	}
	return coords
}

func CoordinatesFrom2DArray(array [][2]float64) Coords {
	var pts = PointsFromArray2D(array)
	var coords = Coords{Pnts: pts, Idxs: make([]int, len(pts))}
	for i := range coords.Pnts {
		coords.Idxs[i] = i
	}
	return coords
}

type Coords struct {
	Pnts []Point
	Idxs []int
}

// Point at index
func (s *Coords) Pt(i int) *Point {
	return &s.Pnts[s.Idxs[i]]
}

// Point at index
func (s Coords) DataView() []Point {
	return s.Pnts[s.Idxs[0] : s.Idxs[s.Len()-1]+1]
}

// Point at index
func (s *Coords) Slice(i, j int) Coords {
	var coords = *s
	coords.Idxs = s.Idxs[i:j]
	return coords
}

// Point at index
func (s Coords) Points() []Point {
	var pts = make([]Point, 0, len(s.Idxs))
	for _, i := range s.Idxs {
		pts = append(pts, s.Pnts[i])
	}
	return pts
}

// First index at 0
func (s *Coords) FirstIndex() int {
	return s.Idxs[0]
}

// First index at len(indices)-1
func (s *Coords) LastIndex() int {
	return s.Idxs[s.Len()-1]
}

// Point at index 0
func (s *Coords) First() *Point {
	return &s.Pnts[s.FirstIndex()]
}

// Point at index 0
func (s *Coords) Last() *Point {
	return &s.Pnts[s.LastIndex()]
}

// len of Coords - sort interface
func (s Coords) Len() int {
	return len(s.Idxs)
}

// swap - sort interface
func (s Coords) Swap(i, j int) {
	s.Idxs[i], s.Idxs[j] = s.Idxs[j], s.Idxs[i]
}

// less - 2d compare - sort interface
func (s Coords) Less(i, j int) bool {
	i, j = s.Idxs[i], s.Idxs[j]
	return (s.Pnts[i][0] < s.Pnts[j][0]) ||
		(feq(s.Pnts[i][0], s.Pnts[j][0]) && s.Pnts[i][1] < s.Pnts[j][1])
}

// 2D sort
func (s *Coords) Sort() *Coords {
	sort.Sort(s)
	return s
}

// Clone coordinates
func (s Coords) Clone() Coords {
	var clone = Coords{
		Pnts: make([]Point, len(s.Pnts)),
		Idxs: make([]int, len(s.Idxs)),
	}
	copy(clone.Pnts, s.Pnts)
	copy(clone.Idxs, s.Idxs)
	return clone
}

// Shallow clone of coordinates with optional slice indices
func (s Coords) ShallowClone(slice ...int) Coords {
	var i, j = 0, s.Len()
	if len(slice) == 1 {
		j = slice[0]
	} else if len(slice) > 1 {
		i, j = slice[0], slice[1]
	}
	var o = Coords{Pnts: s.Pnts, Idxs: make([]int, 0, j-i)}
	for _, v := range s.Idxs[i:j] {
		o.Idxs = append(o.Idxs, v)
	}
	return o
}

func (s *Coords) Append(pt Point) *Coords {
	s.Pnts = append(s.Pnts, pt)
	s.Idxs = append(s.Idxs, len(s.Pnts)-1)
	return s
}

// pop point from
func (s *Coords) Pop() (bool, Point) {
	var v Point
	var n int
	if len(s.Idxs) == 0 {
		return false, NullPt
	}
	n = len(s.Idxs) - 1
	v, s.Idxs[n] = s.Pnts[s.Idxs[n]], -1
	s.Idxs = s.Idxs[:n]
	return true, v
}

// get copy of Coords of polygon
func (self *Polygon) Coordinates() []Coords {
	var lns = self.AsLinear()
	var coords = make([]Coords, len(lns))
	for i, ln := range lns {
		coords[i] = ln.Coordinates
	}
	return coords
}

// checks if a point is a ring , by def every point is a ring
// which concides on itself
func (self *Point) IsRing() bool {
	return true
}

// Checks if line string is a ring
func (self *LineString) IsRing() bool {
	return IsRing(self.Coordinates)
}

// Checks if polygon is a ring - default to true since all polygons are closed ring(s)
func (self *Polygon) IsRing() bool {
	return true
}

// ------------------------------------------------------------------------------
// Is Coords a ring : P0 == Pn
func IsRing(coordinates Coords) bool {
	if coordinates.Len() < 2 {
		return false
	}
	return coordinates.First().Equals2D(coordinates.Last())
}
