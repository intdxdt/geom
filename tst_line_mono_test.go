package geom

import (
	"testing"
	"simplex/geom/mbr"
	"github.com/franela/goblin"
)

func TestLineStringMono(t *testing.T) {
	g := goblin.Goblin(t)
	pts := []*Point{
		{5.78, 8.07},
		{6.44, 9.09},
		{7.87, 9.61},
	}

	ln := NewLineString(pts)
	n := ln.LenVertices()

	g.Describe("Linestring", func() {
		g.It("should test mono mbr", func() {
			bounds := mbr.NewMBR(
				pts[0][x], pts[0][y],
				pts[n-1][x], pts[n-1][y],
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

			mono_boxes := []*MonoMBR{&mbox}
			g.Assert(len(mono_boxes)).Equal(1)

			box, mono_boxes := pop_mono_mbr(mono_boxes)
			g.Assert(box).Eql(&mbox)
			g.Assert(len(mono_boxes)).Equal(0)

			box, mono_boxes = pop_mono_mbr(mono_boxes)
			g.Assert(box == nil).IsTrue()
			g.Assert(len(mono_boxes)).Equal(0)
		})
	})
}
