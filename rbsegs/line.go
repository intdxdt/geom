package main

import (
	"github.com/intdxdt/geom"
	"fmt"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
)

const (
	CreateRED = iota
	CreateBLUE
	RemoveRED
	RemoveBLUE
)

func main() {
	//var awkt = "LINESTRING ( 800 1600, 1000 1800, 1200 1600, 1400 2000, 1000 2200 )"
	//var bwkt = "LINESTRING ( 1000 2000, 1600 1600, 1200 2200, 1600 2000, 1402.6469565217394 1490.912173913043, 875.6904347826086 1716.3034782608693 )"

	var c = geom.Point{1.5, -2}
	var d = geom.Point{-1.5, 2}
	var h = geom.Point{0.484154648492778, -0.645539531323704}
	var i = geom.Point{0.925118053504632, -1.233490738006176}
	var ln_cd = NewLineString(geom.Coordinates([]geom.Point{c, d}))
	var ln_hi = NewLineString(geom.Coordinates([]geom.Point{h, i}))

	//var aln = NewLineString(geom.NewLineStringFromWKT(awkt).Coordinates)
	//var bln = NewLineString(geom.NewLineStringFromWKT(bwkt).Coordinates)
	var inters = ln_cd.RBIntersection(ln_hi)

	fmt.Println(inters)
}

type LineString struct {
	Coordinates geom.Coords
	bbox        mono.MBR
	rbEvent     []event
	bfList      bfList
}

//New LineString from a given Coords {Array} [[x,y], ....[x,y]]
func NewLineString(coordinates geom.Coords) *LineString {
	var n = coordinates.Len()
	if n < 2 {
		panic("a linestring must have at least 2 coordinates")
	}
	var ln = &LineString{
		Coordinates: coordinates,
		rbEvent:     make([]event, 0, 2*(n-1)),
		bfList:      createBruteForceList(n - 1),
	}
	return ln.prepEvents()
}

func (self *LineString) prepEvents() *LineString {
	var n = self.Coordinates.Len() - 1
	var a, b *geom.Point
	var x, y float64
	a = self.Coordinates.Pt(0)
	self.bbox.MBR = mbr.MBR{a[0], a[1], a[0], a[1]}
	self.bbox.I = self.Coordinates.Idxs[0]
	self.bbox.J = self.Coordinates.Idxs[n]

	for i := 0; i < n; i++ {
		a, b = self.Coordinates.Pt(i), self.Coordinates.Pt(i+1)
		x, y = a[0], b[0]
		self.rbEvent = append(self.rbEvent, event{val: minf64(x, y)})
		self.rbEvent = append(self.rbEvent, event{val: maxf64(x, y)})
		self.bbox.MBR.ExpandIncludeXY(b[0], b[1])
	}
	return self
}

func (self *LineString) RBIntersection(other *LineString) [][]int {
	var crossings [][]int
	//self.bfList.reset()
	//other.bfList.reset()
	var visit = func(i, j int) bool {
		crossings = append(crossings, []int{i, j})
		return false
	}

	RedBlueLineSegmentIntersection(self, other, visit)
	return crossings
}

func RedBlueLineSegmentIntersection(red, blue *LineString, visit func(int, int) bool) bool {
	var nr = red.Coordinates.Len() - 1
	var nb = blue.Coordinates.Len() - 1
	var n = nr + nb
	var ne = 2 * n
	var ret bool

	var events = prepareEvents(red, blue)

	var redList = &red.bfList
	var blueList = &blue.bfList

	for i := 0; i < ne; i++ {
		var ev, index = events[i].ev, events[i].idx

		if ev == CreateRED {
			ret = addSegment(index, red, redList, blue, blueList, visit, false)
		} else if ev == CreateBLUE {
			ret = addSegment(index, blue, blueList, red, redList, visit, true)
		} else if ev == RemoveRED {
			redList.remove(index)
		} else if ev == RemoveBLUE {
			blueList.remove(index)
		}

		if ret {
			break
		}
	}

	return ret
}

func maxf64(x, y float64) float64 {
	if y > x {
		return y
	}
	return x
}

func minf64(x, y float64) float64 {
	if y < x {
		return y
	}
	return x
}
