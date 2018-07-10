package geom

import (
	"testing"
	"github.com/intdxdt/mbr"
	"github.com/franela/goblin"
)

func TestLineStringMono(t *testing.T) {
	var g   = goblin.Goblin(t)
	var pts = []Point{{5.78, 8.07}, {6.44, 9.09}, {7.87, 9.61}}
	var ln  = NewLineString(pts)
	var n   = ln.LenVertices()

	g.Describe("Linestring", func() {
		g.It("should test mono mbr", func() {
			bounds := mbr.NewMBR(
				pts[0][X], pts[0][Y],
				pts[n-1][X], pts[n-1][Y],
			)

			mbox := MonoMBR{bounds, 0, n - 1}
			g.Assert(mbox.i).Eql(ln.bbox.i)
			g.Assert(mbox.j).Eql(ln.bbox.j)
			g.Assert(ln.Envelope()).Eql(mbox.BBox())
			g.Assert(ln.Envelope()).Eql(mbox.Clone().BBox())

			mbox.update_index(-1, n)
			g.Assert(mbox.i).Eql(ln.bbox.i - 1)
			g.Assert(mbox.j).Eql(ln.bbox.j + 1)

			mbox.update_index(-1+1, n-1)
			g.Assert(mbox.i).Eql(ln.bbox.i)
			g.Assert(mbox.j).Eql(ln.bbox.j)

			mono_boxes := []MonoMBR{mbox}
			g.Assert(len(mono_boxes)).Equal(1)

			box, mono_boxes := pop_mono_mbr(mono_boxes)
			g.Assert(box).Eql(mbox)
			g.Assert(len(mono_boxes)).Equal(0)

			box, mono_boxes = pop_mono_mbr(mono_boxes)
			g.Assert(box == MonoMBR{}).IsTrue()
			g.Assert(len(mono_boxes)).Equal(0)
		})
	})
}
