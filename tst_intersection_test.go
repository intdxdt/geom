package geom

import (
	"testing"
	"github.com/franela/goblin"
	"time"
)

func TestIntersection(t *testing.T) {
	g := goblin.Goblin(t)

	type Seg struct {
		*Segment
	}
	lnwkt := "LINESTRING ( 350 710, 400 770, 450 770, 480 810, 570 820, 670 730, 720 760, 930 800 )"
	lnwkt2 := "LINESTRING ( 620 620, 720 690, 790 680, 820 630, 870 630, 910 600, 900 530, 800 450, 730 390, 680 420, 640 460, 600 480, 650 540, 690 570, 780 570, 730 630, 680 600, 610 570, 550 610 )"

	plywkt := "POLYGON (( 720 760, 860 770, 950 700, 930 640, 800 600, 740 580, 730 500, 760 440, 720 360, 620 390, 510 480, 460 570, 440 630, 450 740, 480 810, 570 820, 570 770, 580 740, 670 730, 720 760 ), ( 630 670, 580 650, 590 600, 650 580, 710 600, 710 670, 630 670 ), ( 780 650, 800 640, 850 710, 830 720, 780 650 ))"
	plywkt2 := "POLYGON (( 860 920, 950 880, 860 800, 930 720, 880 690, 830 700, 810 730, 790 790, 820 840, 810 870, 860 920 ), ( 840 750, 860 750, 850 800, 830 800, 840 750 ))"

	ptAwkt := "POINT ( 620 620 )"
	ptBwkt := "POINT ( 710 600 )"
	ptCwkt := "POINT ( 722.1298042987639 582.0334837046336 )"
	ptDwkt := "POINT ( 720 360 )"

	polyAwkt := "POLYGON ((730 410, 920 500, 930 540, 930 580, 900 640, 810 650, 750 520, 730 410))"
	polyBwkt := "POLYGON ((630 620, 730 410, 890 410, 1040 510, 1080 620, 1020 720, 690 720, 630 620))"

	ln := NewLineStringFromWKT(lnwkt)
	ln2 := NewLineStringFromWKT(lnwkt2)
	ply := NewPolygonFromWKT(plywkt)
	ply2 := NewPolygonFromWKT(plywkt2)

	plyA := NewPolygonFromWKT(polyAwkt)
	plyB := NewPolygonFromWKT(polyBwkt)

	ptA := NewPointFromWKT(ptAwkt)
	ptB := NewPointFromWKT(ptBwkt)
	ptC := NewPointFromWKT(ptCwkt)
	ptD := NewPointFromWKT(ptDwkt)

	segAA := NewSegment(ptA, ptA)
	segAB := NewSegment(ptA, ptB)
	segNoneGeom_AB := Seg{segAB}
	var nilG *Polygon

	g.Describe("Intersection with pt, seg, ln, poly", func() {

		g.It("intersection", func() {
			g.Timeout(1 * time.Hour)
			inters := plyA.Intersection(plyB)
			g.Assert(len(inters)).Equal(7)

			g.Assert(ln.IsSimple()).IsTrue()
			g.Assert(ply.IsSimple()).IsTrue()
			g.Assert(len(ply.Intersection(nilG))).Equal(0)
			g.Assert(len(ply.Intersection(ln))).Equal(4)
			g.Assert(len(ply.Intersection(ln2))).Equal(22)
			g.Assert(len(ply.Intersection(ply2))).Equal(13)

			g.Assert(len(ptA.Intersection(nilG))).Equal(0)
			g.Assert(len(ptA.Intersection(ply))).Equal(0)
			g.Assert(len(ply.Intersection(ptA))).Equal(0)

			g.Assert(len(ply.Intersection(ptB))).Equal(1)
			g.Assert(len(ptB.Intersection(ply))).Equal(1)

			g.Assert(len(segAB.Intersection(nilG))).Equal(0)
			g.Assert(len(segAA.Intersection(ply))).Equal(0)
			g.Assert(len(ply.Intersection(segAA))).Equal(0)
			g.Assert(len(segAB.Intersection(ply))).Equal(1)

			g.Assert(len(ptA.Intersection(ln))).Equal(0)
			g.Assert(len(ln.Intersection(ptA))).Equal(0)
			g.Assert(len(segAB.Intersection(ptA))).Equal(1)
			g.Assert(len(ptA.Intersection(segAB))).Equal(1)
			g.Assert(len(ply.Intersection(segAB))).Equal(1)

			g.Assert(len(ply.Intersection(ptC))).Equal(1)
			g.Assert(len(ptC.Intersection(ply))).Equal(1)

			g.Assert(len(ply.Intersection(ptD))).Equal(1)
			g.Assert(len(ptD.Intersection(ply))).Equal(1)

			g.Assert(len(ptA.Intersection(nilG))).Equal(0)
			g.Assert(len(ptA.Intersection(ptB))).Equal(0)
			g.Assert(len(ptA.Intersection(ptA))).Equal(1)
			g.Assert(len(ptA.Intersection(ln2))).Equal(1)

			g.Assert(len(ln2.Intersection(nilG))).Equal(0)
			g.Assert(len(ln2.Intersection(ptA))).Equal(1)
			g.Assert(len(ln2.Intersection(ptB))).Equal(0)
		})

		g.It("polygon intersection other not segment ", func() {
			defer func() {
				g.Assert(recover() != nil)
			}()
			ply.Intersection(segNoneGeom_AB)
		})
		g.It("segment intersection other not segment ", func() {
			defer func() {
				g.Assert(recover() != nil)
			}()
			segAB.Intersection(segNoneGeom_AB)
		})
		g.It("line intersection other not segment ", func() {
			defer func() {
				g.Assert(recover() != nil)
			}()
			ln.Intersection(segNoneGeom_AB)
		})
		g.It("pt intersection other not segment ", func() {
			defer func() {
				g.Assert(recover() != nil)
			}()
			ptA.Intersection(segNoneGeom_AB)
		})
	})
}
