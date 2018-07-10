package geom

import (
	"testing"
	"github.com/franela/goblin"
)

func TestGeom(t *testing.T) {
	g := goblin.Goblin(t)
	p := PointXY(4.0, 5.0)
	ln := NewLineString([]Point{{0, 0}, {1, 1}})
	var ply *Polygon
	var pnt *Point
	var line *LineString
	var rng = NewLinearRing([]Point{{0, 0}, {1, 1}})
	var pt_wkt = "POINT (30 10)"
	var ln_wkt = "LINESTRING (30 10, 10 30, 40 40)"
	var ply_wkt = "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))"

	g.Describe("Geometry", func() {
		g.It("it should test new geometry ", func() {
			pt := NewGeometry(pt_wkt)
			ln := NewGeometry(ln_wkt)
			ply := NewGeometry(ply_wkt)
			g.Assert(pt.Type().IsPoint()).IsTrue()
			g.Assert(ln.Type().IsLineString()).IsTrue()
			g.Assert(ply.Type().IsPolygon()).IsTrue()
		})

		g.It("it should test NullGeometry", func() {
			g.Assert(IsNullGeometry(&p)).IsFalse()
			g.Assert(IsNullGeometry(ln)).IsFalse()
			g.Assert(IsNullGeometry(pnt)).IsTrue()
			g.Assert(IsNullGeometry(ply)).IsTrue()
			g.Assert(IsNullGeometry(line)).IsTrue()

			ring, ok := IsLinearRing(rng)
			g.Assert(ok).IsTrue()
			g.Assert(ring).Eql(rng)

		})
		g.It("it should panic if NOT one of fundermental types ", func() {
			defer func() {
				g.Assert(recover() != nil).IsTrue()
			}()
			g.Assert(IsNullGeometry(rng)).IsTrue()
		})
	})
}
