package geom

import (
	"fmt"
	"github.com/franela/goblin"
	"testing"
)

func TestHullGen(t *testing.T) {
	g := goblin.Goblin(t)
	var A = NewPointXY(172.0, 224.0)
	var B = NewPointXY(180.0, 158.0)
	var C = NewPointXY(266.0, 46.0)
	var D = NewPointXY(374.0, 38.0)
	var E = NewPointXY(480.0, 100.0)
	var F = NewPointXY(500.0, 200.0)
	var G = NewPointXY(440.0, 300.0)
	var H = NewPointXY(340.0, 280.0)
	var I = NewPointXY(200.0, 240.0)

	coords := []*Point{A, B, C, D, E, F, G, H, I}
	hull := NewHull(coords)
	g.Describe("Geometry", func() {
		g.It("it should test Hull Antipodal", func() {
			g.Assert(hull.Antipodal(2, 3)).Equal(6)
			g.Assert(hull.Antipodal(0, 1)).Equal(5)
			g.Assert(hull.Antipodal(4, 5)).Equal(0)
		})
	})

}

func TestHullHexagon(t *testing.T) {
	g := goblin.Goblin(t)
	var A = NewPointXY(20, 20)
	var B = NewPointXY(40, 20)
	var C = NewPointXY(45, 35)
	var D = NewPointXY(40, 50)
	var E = NewPointXY(20, 50)
	var F = NewPointXY(15, 35)

	coords := []*Point{A, B, C, D, E, F}
	hull := NewHull(ConvexHull(coords))
	g.Describe("Geometry", func() {
		g.It("it should test Hull Antipodal Hexagon", func() {
			g.Assert(hull.Antipodal(0, 1)).Equal(4)
			g.Assert(hull.Antipodal(1, 2)).Equal(5)
			g.Assert(hull.Antipodal(2, 3)).Equal(5)
			g.Assert(hull.Antipodal(3, 4)).Equal(1)
			g.Assert(hull.Antipodal(4, 5)).Equal(2)
		})
	})

}

func TestHullHexagonPlus1(t *testing.T) {
	g := goblin.Goblin(t)
	var A = NewPointXY(20, 20)
	var B = NewPointXY(40, 20)
	var C = NewPointXY(45, 35)
	var D = NewPointXY(40, 50)
	var E = NewPointXY(20, 50)
	var F = NewPointXY(15, 35)
	var I = NewPointXY(30, 55)

	coords := []*Point{A, B, C, D, I, E, F}
	hull := NewHull(ConvexHull(coords))
	g.Describe("Geometry", func() {
		g.It("it should test Hull Antipodal Hexagon + I", func() {
			g.Assert(hull.Antipodal(0, 1)).Equal(4)
			g.Assert(hull.Antipodal(1, 2)).Equal(5)
			g.Assert(hull.Antipodal(2, 3)).Equal(6)
			g.Assert(hull.Antipodal(3, 4)).Equal(1)
			g.Assert(hull.Antipodal(4, 5)).Equal(1)
			g.Assert(hull.Antipodal(5, 6)).Equal(2)
			g.Assert(hull.Antipodal(6, 0)).Equal(2)
		})
	})

}

func TestHullOnLine(t *testing.T) {
	g := goblin.Goblin(t)
	var A = NewPointXY(70, 55)
	var B = NewPointXY(75, 55)
	var C = NewPointXY(80, 55)
	var D = NewPointXY(85, 55)

	coords := []*Point{A, B, C, D}
	hull := NewHull(ConvexHull(coords))
	fmt.Println(hull)
	g.Describe("Geometry", func() {
		g.It("it should test Hull Antipodal Hexagon + I", func() {
			g.Assert(hull.Antipodal(0, 1)).Equal(1)
			g.Assert(hull.Antipodal(1, 0)).Equal(0)
		})
		g.It("it shoud panic - hull indexer util", func() {
			defer func() {
				r := recover()
				g.Assert(r != nil).IsTrue()
			}()
			h := Hull{}
			idxer := h.indexer(0, 9)
			idxer(-1)
		})
		g.It("it shoud panic - hull chain indexer util", func() {
			defer func() {
				r := recover()
				g.Assert(r != nil).IsTrue()
			}()
			h := Hull{}
			chindxer := h.chainIndexer(0, 9)
			chindxer(10)
		})
	})

}
