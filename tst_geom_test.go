package geom

import (
	"github.com/franela/goblin"
	"testing"
)

func TestGeom(t *testing.T) {
	g := goblin.Goblin(t)
	p := NewPointXY(4.0, 5.0)
	ln := NewLineString([]*Point{{0, 0}, {1, 1}})
	var ply *Polygon
	var pnt *Point
	var line *LineString
	var rng = NewLinearRing([]*Point{{0, 0}, {1, 1}})

	g.Describe("Geometry", func() {
		g.It("it should test NullGeometry", func() {
			g.Assert(IsNullGeometry(p)).IsFalse()
			g.Assert(IsNullGeometry(ln)).IsFalse()
			g.Assert(IsNullGeometry(pnt)).IsTrue()
			g.Assert(IsNullGeometry(ply)).IsTrue()
			g.Assert(IsNullGeometry(line)).IsTrue()

			ring, ok := IsLinearRing(rng)
			g.Assert(ok).IsTrue()
			g.Assert(ring).Eql(rng)

			g.Assert(IsNullGeometry(rng)).IsTrue()
		})
	})

}
