package geom

import (
	"testing"
	"github.com/franela/goblin"
)

func TestPolygonInters(t *testing.T) {
	g := goblin.Goblin(t)
	lnwkt := "LINESTRING ( 350 710, 400 770, 450 770, 480 810, 570 820, 670 730, 720 760, 930 800 )"
	lnwkt2 := "LINESTRING ( 620 620, 720 690, 790 680, 820 630, 870 630, 910 600, 900 530, 800 450, 730 390, 680 420, 640 460, 600 480, 650 540, 690 570, 780 570, 730 630, 680 600, 610 570, 550 610 )"

	plywkt := "POLYGON (( 720 760, 860 770, 950 700, 930 640, 800 600, 740 580, 730 500, 760 440, 720 360, 620 390, 510 480, 460 570, 440 630, 450 740, 480 810, 570 820, 570 770, 580 740, 670 730, 720 760 ), ( 630 670, 580 650, 590 600, 650 580, 710 600, 710 670, 630 670 ), ( 780 650, 800 640, 850 710, 830 720, 780 650 ))"
	plywkt2 := "POLYGON (( 860 920, 950 880, 860 800, 930 720, 880 690, 830 700, 810 730, 790 790, 820 840, 810 870, 860 920 ), ( 840 750, 860 750, 850 800, 830 800, 840 750 ))"

	ptAwkt := "POINT ( 620 620 )"
	ptBwkt := "POINT ( 710 600 )"
	ptCwkt := "POINT ( 722.1298042987639 582.0334837046336 )"
	ptDwkt := "POINT ( 720 360 )"

	g.Describe("Polygon", func() {
		g.It("polygon intersection", func() {
			ln := NewLineStringFromWKT(lnwkt)
			ln2 := NewLineStringFromWKT(lnwkt2)
			ply := NewPolygonFromWKT(plywkt)
			ply2 := NewPolygonFromWKT(plywkt2)

			ptA := NewPointFromWKT(ptAwkt)
			ptB := NewPointFromWKT(ptBwkt)
			ptC := NewPointFromWKT(ptCwkt)
			ptD := NewPointFromWKT(ptDwkt)

			g.Assert(ln.IsSimple()).IsTrue()
			g.Assert(ply.IsSimple()).IsTrue()
			g.Assert(len(ply.Intersection(nil))).Equal(0)
			g.Assert(len(ply.Intersection(ln))).Equal(4)
			g.Assert(len(ply.Intersection(ln2))).Equal(22)
			g.Assert(len(ply.Intersection(ply2))).Equal(10)
			g.Assert(len(ply.Intersection(ptA))).Equal(0)
			g.Assert(len(ptA.Intersection(ply))).Equal(0)

			g.Assert(len(ply.Intersection(ptB))).Equal(1)
			g.Assert(len(ptB.Intersection(ply))).Equal(1)

			g.Assert(len(ply.Intersection(ptC))).Equal(1)
			g.Assert(len(ptC.Intersection(ply))).Equal(1)

			g.Assert(len(ply.Intersection(ptD))).Equal(1)
			g.Assert(len(ptD.Intersection(ply))).Equal(1)

			g.Assert(len(ptA.Intersection(nil))).Equal(0)
			g.Assert(len(ptA.Intersection(ptB))).Equal(0)
			g.Assert(len(ptA.Intersection(ptA))).Equal(1)
			g.Assert(len(ptA.Intersection(ln2))).Equal(1)
			g.Assert(len(ln2.Intersection(ptA))).Equal(1)
			g.Assert(len(ln2.Intersection(ptB))).Equal(0)
		})

	})
}
