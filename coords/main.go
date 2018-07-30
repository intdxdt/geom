package main

import (
	"sort"
	"github.com/intdxdt/math"
	"fmt"
)

var nan = math.NaN()
var feq = math.FloatEqual

type Point [3]float64

var NullPt = Point{nan, nan, nan}

func main() {
	var pts = []Point{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 10}}
	var a = NewCoordinates(pts)
	var b = a
	var r = makeLnrRing(b)
	setZero(r)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(r)
}
func setZero(a Coordinates){
	a._c[0][0] = -9
	a._c[0][1] = -9
	a.idxs[0] = -1
}
func makeLnrRing(coords Coordinates) Coordinates{
	var n = len(coords._c) - 1
	var a, b = coords._c[0], coords._c[n]
	if !(a[0] == b[0] && a[1] == b[1]) {
		coords.idxs = coords.idxs[0:len(coords.idxs):len(coords.idxs)]
		coords.idxs = append(coords.idxs, 0)
	}
	return coords
}

func NewCoordinates(c []Point) Coordinates {
	var coords = Coordinates{_c: c, idxs: make([]int, 0, len(c))}
	for i := range coords._c {
		coords.idxs = append(coords.idxs, i)
	}
	return coords
}

type Coordinates struct {
	_c   []Point
	idxs []int
}

//coordinate at location
func (s Coordinates) Pt(i int) *Point {
	return &s._c[s.idxs[i]]
}

//len of coordinates - sort interface
func (s Coordinates) Len() int {
	return len(s.idxs)
}

//swap - sort interface
func (s Coordinates) Swap(i, j int) {
	s.idxs[i], s.idxs[j] = s.idxs[j], s.idxs[i]
}

//less - 2d compare - sort interface
func (s Coordinates) Less(i, j int) bool {
	i, j = s.idxs[i], s.idxs[j]
	return (s._c[i][0] < s._c[j][0]) ||
		(feq(s._c[i][0], s._c[j][0]) && s._c[i][1] < s._c[j][1])
}

//2D sort
func (s Coordinates) Sort() Coordinates {
	sort.Sort(s)
	return s
}

//pop point from
func (s Coordinates) Pop() (Point, Coordinates) {
	var v Point
	var n int
	if len(s.idxs) == 0 {
		return NullPt, s
	}
	n = len(s.idxs) - 1
	v, s.idxs[n] = s._c[s.idxs[n]], -1
	s.idxs = s.idxs[:n]
	return v, s
}
