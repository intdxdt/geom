package main

import (
	"sort"
	"fmt"
	"github.com/intdxdt/math"
)

var nan = math.NaN()
var feq = math.FloatEqual

type Point [3]float64

var NullPt = Point{nan, nan, nan}

func main() {
	var pts = []Point{{5, 6}, {7, 8}, {3, 4},  {9, 10}, {1, 2}, }
	var a = NewCoordinates(pts)
	var b = a
	var r = makeLnrRing(b)
	setZero(r)
	b.Sort()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(r)

}
func setZero(a Coordinates){
	a._c[0][0] = -9
	a._c[0][1] = -9
	a.Idxs[0] = -1
}

func makeLnrRing(coords Coordinates) Coordinates{
	var n = len(coords._c) - 1
	var a, b = coords._c[0], coords._c[n]
	if !(a[0] == b[0] && a[1] == b[1]) {
		coords.Idxs = coords.Idxs[0:len(coords.Idxs):len(coords.Idxs)]
		coords.Idxs = append(coords.Idxs, 0)
	}
	return coords
}

func NewCoordinates(c []Point) Coordinates {
	var coords = Coordinates{_c: c, Idxs: make([]int, 0, len(c))}
	for i := range coords._c {
		coords.Idxs = append(coords.Idxs, i)
	}
	return coords
}

type Coordinates struct {
	_c   []Point
	Idxs []int
}

//coordinate at location
func (s Coordinates) Pt(i int) *Point {
	return &s._c[s.Idxs[i]]
}

//len of coordinates - sort interface
func (s Coordinates) Len() int {
	return len(s.Idxs)
}

//swap - sort interface
func (s Coordinates) Swap(i, j int) {
	s.Idxs[i], s.Idxs[j] = s.Idxs[j], s.Idxs[i]
}

//less - 2d compare - sort interface
func (s Coordinates) Less(i, j int) bool {
	i, j = s.Idxs[i], s.Idxs[j]
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
	if len(s.Idxs) == 0 {
		return NullPt, s
	}
	n = len(s.Idxs) - 1
	v, s.Idxs[n] = s._c[s.Idxs[n]], -1
	s.Idxs = s.Idxs[:n]
	return v, s
}
