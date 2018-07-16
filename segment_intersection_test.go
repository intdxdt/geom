package geom

import (
	"time"
	"testing"
	"github.com/franela/goblin"
)

func segment(ln string) *Segment {
	var coords = NewLineStringFromWKT(ln).Coordinates()
	return NewSegment(&coords[0], &coords[1])
}

func TestToSegmentIntersection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("planar intersection", func() {
		g.It("should seg seg intersection", func() {
			g.Timeout(1 * time.Hour)
			var l0 = segment("LINESTRING ( 350 350, 450 350 )")
			var l1 = segment("LINESTRING ( 400.5652173913044 350, 550 350 )")
			var l2 = segment("LINESTRING ( 350 450, 350 300 )")
			var l3 = segment("LINESTRING ( 450 350, 450 450 )")
			var l4 = segment("LINESTRING ( 400 300, 500 400 )")
			var l5 = segment("LINESTRING ( 400 450, 400 250 )")
			var l6 = segment("LINESTRING ( 300 450, 350 350 )")
			var l7 = segment("LINESTRING ( 450 350, 600 350 )")
			var l8 = segment("LINESTRING ( 450 350, 450 350 )")
			var l9 = segment("LINESTRING (350 350, 350 350 )")
			var l10 = segment("LINESTRING ( 350 350, 450 350 )")
			var l11 = segment("LINESTRING (350 350, 350 350 )")
			var l12 = segment("LINESTRING ( 400 350, 450 350 )")

			var intpts = SegSegIntersection(l0.A, l0.B, l1.A, l1.B)
			g.Assert(len(intpts)).Equal(2)
			g.Assert(intpts[0].IsVertex()).IsTrue()
			g.Assert(intpts[0].Inter).Equal(OtherA)
			g.Assert(intpts[0].IsVertexSelf()).IsFalse()
			g.Assert(intpts[0].IsVertexOther()).IsTrue()
			g.Assert(intpts[0].IsVerteXOR()).IsTrue()

			g.Assert(intpts[1].Inter).Equal(SelfB)
			g.Assert(intpts[1].IsVertex()).IsTrue()
			g.Assert(intpts[1].IsVertexSelf()).IsTrue()
			g.Assert(intpts[1].IsVertexOther()).IsFalse()
			g.Assert(intpts[1].IsVerteXOR()).IsTrue()

			intpts = SegSegIntersection(l0.A, l0.B, l2.A, l2.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].IsVertex()).IsTrue()
			g.Assert(intpts[0].Inter).Equal(SelfA)

			intpts = SegSegIntersection(l0.A, l0.B, l3.A, l3.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].IsVertex()).IsTrue()
			g.Assert(intpts[0].Inter).Equal(SelfB | OtherA)
			g.Assert(intpts[0].IsVerteXOR()).IsFalse()

			intpts = SegSegIntersection(l0.A, l0.B, l4.A, l4.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].IsVertex()).IsTrue()
			g.Assert(intpts[0].Inter).Equal(SelfB)

			intpts = SegSegIntersection(l0.A, l0.B, l5.A, l5.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].IsIntersection()).IsTrue()
			g.Assert(intpts[0].Inter).Equal(InterX)
			g.Assert(intpts[0].IsVerteXOR()).IsFalse()

			intpts = SegSegIntersection(l0.A, l0.B, l6.A, l6.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].IsVertex()).IsTrue()
			g.Assert(intpts[0].Inter).Equal(SelfA | OtherB)

			intpts = SegSegIntersection(l0.A, l0.B, l7.A, l7.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].Inter).Equal(SelfB | OtherA)
			g.Assert(intpts[0].IsVertex()).IsTrue()

			intpts = SegSegIntersection(l0.A, l0.B, l8.A, l8.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].Inter).Equal(SelfB | OtherA | OtherB)

			intpts = SegSegIntersection(l0.A, l0.B, l9.A, l9.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].Inter).Equal(SelfA | OtherA | OtherB)

			intpts = SegSegIntersection(l0.A, l0.B, l10.A, l10.B)
			g.Assert(len(intpts)).Equal(2)
			g.Assert(intpts[0].Inter).Equal(SelfA | OtherA)
			g.Assert(intpts[1].Inter).Equal(SelfB | OtherB)

			intpts = SegSegIntersection(l9.A, l9.B, l11.A, l11.B)
			g.Assert(len(intpts)).Equal(1)
			g.Assert(intpts[0].Inter).Equal(SelfA | SelfB | OtherA | OtherB)

			intpts = SegSegIntersection(l0.A, l0.B, l12.A, l12.B)
			g.Assert(len(intpts)).Equal(2)
			g.Assert(intpts[0].Inter).Equal(OtherA)
			g.Assert(intpts[1].Inter).Equal(SelfB | OtherB)

			g.Assert(intpts[0].String()).Equal("[400, 350, 0, 0010]")
			g.Assert(intpts[1].String()).Equal("[450, 350, 0, 0101]")
		})
	})
}
