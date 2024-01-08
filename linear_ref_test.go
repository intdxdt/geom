package geom

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/math"
	"testing"
	"time"
)

func TestLinearReference(t *testing.T) {
	var g = goblin.Goblin(t)
	g.Describe("geom.linear.ref", func() {
		g.It("tests linear ref", func() {
			g.Timeout(1 * time.Hour)

			var lnA = NewLineStringFromWKT("LINESTRING ( 300 500, 300 600, 500 600, 500 400, 300 500 )")
			var ptA = PointFromWKT("POINT ( 400 500 )")
			g.Assert(math.FloatEqual(lnA.Project(&ptA), 634.1640786499873)).IsTrue()
			g.Assert(lnA.Interpolate(634.1640786499873).WKT()).Equal("POINT (380 460)")
			g.Assert(lnA.Interpolate(0).WKT()).Equal("POINT (300 500)")
			g.Assert(lnA.Interpolate(723.606797749979).WKT()).Equal("POINT (300 500)")
			g.Assert(lnA.Interpolate(750.5).WKT()).Equal("POINT (300 500)")
			g.Assert(lnA.Interpolate(1, true).WKT()).Equal("POINT (300 500)")

			ptA = PointFromWKT("POINT ( 250 450 )")
			g.Assert(math.FloatEqual(lnA.Project(&ptA), 0.0)).IsTrue()

			ptA = PointFromWKT("POINT ( 250 650 )")
			g.Assert(math.FloatEqual(lnA.Project(&ptA), 100.0)).IsTrue()

			lnA = NewLineStringFromWKT("LINESTRING (0 0, 0 1, 1 1)")
			ptA = PointFromWKT("POINT (0.5 1)")
			g.Assert(math.FloatEqual(lnA.Project(&ptA), 1.5)).IsTrue()
			g.Assert(math.FloatEqual(lnA.Project(&ptA, true), 0.75)).IsTrue()

			lnA = NewLineStringFromWKT("LINESTRING (0 0, 0 1, 1 1)")
			ptA = PointFromWKT("POINT (0.5 1)")
			g.Assert(lnA.Interpolate(-2.5).WKT()).Equal("POINT (0 0)")
			g.Assert(lnA.Interpolate(2.5).WKT()).Equal("POINT (1 1)")
			g.Assert(lnA.Interpolate(1.5).WKT()).Equal("POINT (0.5 1)")
			g.Assert(lnA.Interpolate(0).WKT()).Equal("POINT (0 0)")
			g.Assert(lnA.Interpolate(2.0).WKT()).Equal("POINT (1 1)")
			g.Assert(lnA.Interpolate(1.0).WKT()).Equal("POINT (0 1)")
			g.Assert(lnA.Interpolate(0.75, true).WKT()).Equal("POINT (0.5 1)")

			lnA = NewLineStringFromWKT("LINESTRING (0 0, 0 0.1, 0.2 0.1, 0.3 0.1, 0.5 0.1, 0.7 0.1)")
			g.Assert(lnA.Interpolate(0.3).WKT()).Equal("POINT (0.2 0.1)")
			g.Assert(lnA.Interpolate(0.4).WKT()).Equal("POINT (0.3 0.1)")
			g.Assert(lnA.Interpolate(0.6).WKT()).Equal("POINT (0.5 0.1)")
			g.Assert(lnA.Interpolate(0.8).WKT()).Equal("POINT (0.7 0.1)")

			lnA = NewLineStringFromWKT("LINESTRING (0 0, 0 0.1, 0.2 0.1, 0.3 0.1, 0.5 0.1, 0.7 0.1)")
			var tokens = lnA.SplitLineString(0.35)
			g.Assert(len(tokens)).Equal(2)
			g.Assert(tokens[0].Length()).Equal(0.35)
			g.Assert(math.FloatEqual(tokens[1].Length(), lnA.Length()-0.35)).IsTrue()

			tokens = lnA.SplitLineString(0.3)
			g.Assert(len(tokens)).Equal(2)
			g.Assert(math.FloatEqual(tokens[0].Length(), 0.3)).IsTrue()
			g.Assert(math.FloatEqual(tokens[1].Length(), lnA.Length()-0.3)).IsTrue()

			tokens = lnA.SplitLineString(-0.001)
			g.Assert(len(tokens)).Equal(1)
			g.Assert(tokens[0].Length()).Equal(lnA.Length())

			tokens = lnA.SplitLineString(0.0)
			g.Assert(len(tokens)).Equal(1)
			g.Assert(tokens[0].Length()).Equal(lnA.Length())

			tokens = lnA.SplitLineString(lnA.Length())
			g.Assert(len(tokens)).Equal(1)
			g.Assert(tokens[0].Length()).Equal(lnA.Length())

			tokens = lnA.SplitLineString(lnA.Length() + .01)
			g.Assert(len(tokens)).Equal(1)
			g.Assert(tokens[0].Length()).Equal(lnA.Length())

			lnA = NewLineStringFromWKT("LINESTRING ( 500 500, 500 400, 600 400, 600 500, 500 500 )")
			tokens = lnA.SplitLineString(250)
			g.Assert(len(tokens)).Equal(2)
			g.Assert(tokens[0].Length()).Equal(250.0)
			g.Assert(tokens[1].Length()).Equal(150.0)

		})
	})

}
