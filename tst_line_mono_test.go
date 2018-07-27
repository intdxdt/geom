package geom

import (
	"testing"
	"github.com/intdxdt/mbr"
	"github.com/franela/goblin"
	"github.com/intdxdt/geom/mono"
)

func TestLineStringMono(t *testing.T) {
	var g   = goblin.Goblin(t)
	var pts = []Point{{5.78, 8.07}, {6.44, 9.09}, {7.87, 9.61}}
	var ln  = NewLineString(pts)
	var n   = ln.LenVertices()

	g.Describe("Linestring", func() {
		g.It("should test mono mbr", func() {
			bounds := mbr.CreateMBR(
				pts[0][X], pts[0][Y],
				pts[n-1][X], pts[n-1][Y],
			)

			mbox := mono.MBR{bounds, 0, n - 1}
			g.Assert(mbox.I).Eql(ln.bbox.I)
			g.Assert(mbox.J).Eql(ln.bbox.J)
			g.Assert(ln.BBox()).Eql(mbox.BBox())
			g.Assert(ln.BBox()).Eql(mbox.BBox())
			g.Assert(ln.Bounds()).Eql(mbox.MBR)

			mbox.UpdateIndex(-1, n)
			g.Assert(mbox.I).Eql(ln.bbox.I - 1)
			g.Assert(mbox.J).Eql(ln.bbox.J + 1)

			mbox.UpdateIndex(-1+1, n-1)
			g.Assert(mbox.I).Eql(ln.bbox.I)
			g.Assert(mbox.J).Eql(ln.bbox.J)

			mono_boxes := []mono.MBR{mbox}
			g.Assert(len(mono_boxes)).Equal(1)

			box, mono_boxes := pop_mono_mbr(mono_boxes)
			g.Assert(box).Eql(mbox)
			g.Assert(len(mono_boxes)).Equal(0)

			box, mono_boxes = pop_mono_mbr(mono_boxes)
			g.Assert(box == mono.MBR{}).IsTrue()
			g.Assert(len(mono_boxes)).Equal(0)
		})
	})
}
