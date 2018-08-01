package geom

import (
	"testing"
	"github.com/franela/goblin"
)

func TestGeom(t *testing.T) {
	var g = goblin.Goblin(t)
	var p = PointXY(4.0, 5.0)
	var ln = NewLineString(Coordinates([]Point{{0, 0}, {1, 1}}))
	var ply *Polygon
	var line *LineString
	var pnt Point
	var rng = NewLinearRing(Coordinates([]Point{{0, 0}, {1, 1}}))
	var pt_wkt = "POINT (30 10)"
	var ln_wkt = "LINESTRING (30 10, 10 30, 40 40)"
	var ply_wkt = "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))"

	g.Describe("Geometry", func() {
		g.It("it should test new geometry ", func() {
			var pt = ReadGeometry(pt_wkt)
			var ln = ReadGeometry(ln_wkt)
			var ply = ReadGeometry(ply_wkt)
			var seg = NewSegmentAB(&p, &p)
			g.Assert(CastAsPoint(pt) == pt).IsTrue()
			g.Assert(CastAsLineString(ln) == ln).IsTrue()
			g.Assert(CastAsPolygon(ply) == ply).IsTrue()
			g.Assert(CastAsSegment(seg) == seg).IsTrue()

			g.Assert(pt.Type().IsPoint()).IsTrue()
			g.Assert(ln.Type().IsLineString()).IsTrue()
			g.Assert(ply.Type().IsPolygon()).IsTrue()
		})

		g.It("it should test NullGeometry", func() {
			g.Assert(IsNullGeometry(&p)).IsFalse()
			g.Assert(IsNullGeometry(ln)).IsFalse()
			g.Assert(IsNullGeometry(pnt)).IsFalse()
			g.Assert(IsNullGeometry(ply)).IsTrue()
			g.Assert(IsNullGeometry(line)).IsTrue()

			ring, ok := IsLinearRing(rng)
			g.Assert(ok).IsTrue()
			g.Assert(ring).Eql(rng)

		})
		g.It("test geom type from composed type", func() {
			g.Assert(IsNullGeometry(rng)).IsFalse()
		})
	})
}
