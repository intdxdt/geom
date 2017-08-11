package geom

import (
	"testing"
	"simplex/util/math"
	"github.com/franela/goblin"
	"simplex/geom/mbr"
)

func TestLineString(t *testing.T) {
	g := goblin.Goblin(t)
	pts := []*Point{
		{5.6, 7.9}, {5.6, 8.9}, {6.6, 8.9}, {6.6, 7.9}, {5.6, 7.9},
	}
	pt_array := [][]float64{
		{5.6, 7.9, 0}, {5.6, 8.9, 0}, {6.6, 8.9, 0}, {6.6, 7.9, 0}, {5.6, 7.9, 0},
	}

	pts_closed := []*Point{{5.538, 8.467}, {5.498, 8.559}, {5.858, 8.987}, {6.654, 8.638}, {6.549, 8.024}, {5.765, 8.082}, {5.538, 8.467}}
	pts_open := []*Point{{5.538, 8.467}, {5.498, 8.559}, {5.858, 8.987}, {6.654, 8.638}, {6.549, 8.024}, {5.765, 8.082}}

	ln := NewLineString(pts)
	ln2 := NewLineString(pts_closed)
	ln3 := NewLineString(pts_open)
	ply := NewPolygon(pts_closed)

	cln := ln.Clone()
	pt_lnstr := NewLineStringFromPoint(pts[0])

	g.Describe("Linestring", func() {
		g.It("should test length", func() {
			g.Assert(ln.Type().IsLineString()).IsTrue()
			g.Assert(ln.Length() == 4.0).IsTrue()

			g.Assert(pt_lnstr.Length() == 0.0).IsTrue()
			g.Assert(ln.IsRing()).IsTrue()
			g.Assert(math.Round(ln.Area(), 5)).Equal(1.0)
			g.Assert(ln.len(len(ln.coordinates)-1, 0) == ln.Length()).IsTrue()
			g.Assert(ln.chain_length(ln.chains[0])).Equal(ln.chain_length(ln.chains[1]))
			g.Assert(ln.chain_length(ln.chains[2])).Equal(ln.chain_length(ln.chains[3]))
			g.Assert(cln.Length() == 4.0).IsTrue()

			g.Assert(ln3.Area()).Equal(0.0)
			g.Assert(ln2.Area()).Equal(ply.Area())

		})

		g.It("should throw if empty coordinates", func(done goblin.Done) {
			defer func() {
				r := recover()
				if r != nil {
					g.Assert(r != nil).Equal(true)
				} else {
					g.Fail("did not throw")
				}
				done()
			}()
			pts := make([]*Point, 0)
			NewLineString(pts)
		})

		g.It("should be slice of array", func() {
			ln.build_index()
			g.Assert(ln.ToArray()).Eql(pt_array)
			g.Assert(cln.ToArray()).Eql(pt_array)
			ln.build_index()
			g.Assert(ln.ToArray()).Eql(pt_array)
		})

	})

	g.Describe("Linestring - Coords", func() {

		g.It("should be slice of points", func() {
			ln.build_index()
			g.Assert(ln.Coordinates()).Eql(pts)
			g.Assert(cln.Coordinates()).Eql(pts)
			ln.build_index()
			g.Assert(ln.Coordinates()).Eql(pts)
		})

		g.It("should test coords indexing", func() {
			g.Assert(ln.VertexAt(0)).Eql(pts[0])
			g.Assert(ln.VertexAt(ln.LenVertices() - 1)).Eql(pts[len(pts)-1])
			g.Assert(ln.LenVertices()).Eql(len(pts))
		})

		g.It("should test envelope", func() {
			box := mbr.NewMBR(pts[0][x], pts[0][y], pts[0][x], pts[0][y])
			for _, v := range pts[1:] {
				box.ExpandIncludeXY(v.X(), v.Y())
			}
			g.Assert(ln.Envelope()).Eql(box)
		})

	})

	g.Describe("Linestring - WKT", func() {
		g.It("should test wkt string", func() {
			lnstr := "LINESTRING (5.6 7.9, 5.6 8.9, 6.6 8.9, 6.6 7.9, 5.6 7.9)"
			g.Assert(ln.WKT()).Eql(lnstr)
		})
	})
}

func TestLineStringRelate(t *testing.T) {
	g := goblin.Goblin(t)

	var coords = []*Point{{0.5, 0.5}, {0.06, -0.1}, {0.26, -0.61}, {0, -1}, {-1.5, -1}, {-0.5, -0.5}}
	var coords2 = []*Point{{0.64, 1.72}, {1.18, 1.87}, {1.68, 1.43}, {0.54, 1.38}}
	plywkt := "POLYGON (( 0.64 1.72, 1.18 1.87, 1.68 1.43, 0.54 1.38, 0.64 1.72 ), (0.9471 1.5300, 0.9471 1.7102, 1.0653 1.7102, 1.0653 1.5300, 0.9471 1.5300 ))"
	plywktc := "POLYGON (( 0.9694190834241365 1.6351888097521738, 0.9963995357624527 1.6647388289798535, 1.013101720543315 1.6467518607543095, 1.032373472213541 1.6608844786458083, 1.0465060901050398 1.6454670773096276, 1.0278767301571547 1.6313344594181287, 1.0400821728816312 1.6152746663596074, 0.9880484433720215 1.6094931408585396, 0.9694190834241365 1.6351888097521738 ))"

	plywktd := "POLYGON (( 1.06137745847723 1.0766292071767967, 0.9394836291630517 0.8815990802741116, 1.3140301228738902 0.7752190110544651, 1.593277804575462 1.0034929095882898, 1.2453263281695353 1.185225527838519, 1.06137745847723 1.0766292071767967 ),( 1.2364613224012313 1.0832779615030246, 1.1212162474132812 1.0012766581462138, 1.2364613224012313 0.9303566119997828, 1.3472738945050298 0.9613841321888464, 1.4093289348831568 1.0300879268932013, 1.3384088887367258 1.0877104643871764, 1.2364613224012313 1.0832779615030246 ),( 1.1721900305810282 0.850571560085048, 1.1721900305810282 0.8838153317161875, 1.3517063973891816 0.8838153317161875, 1.3517063973891816 0.850571560085048, 1.1721900305810282 0.850571560085048 ))"
	plywkte := "POLYGON (( -0.2405548235983036 -0.1291889913629033, 0.3266242131459507 0.0813804713804726, 0.5032308593178143 -0.0442819499341227, 0.5747608247298315 -0.5006764018668276, 0.3368130581174044 -0.6318386766212843, 0.3979461279461264 -0.2718328209632546, 0.224735763431414 -0.1461703996486594, -0.2099882886839426 -0.292210510906162, -0.2405548235983036 -0.1291889913629033 ))"
	plywktf := "POLYGON (( -0.277913921826967 -0.5367427902210501, -0.4850871029131916 -0.6997643097643087, -0.3424432733128402 -0.8152378861074503, -0.2337622602840011 -0.7303308446786697, -0.1080998389694059 -0.8661821109647186, 0.0141663006880382 -0.7371234079929722, -0.1182886839408595 -0.5842907334211672, -0.277913921826967 -0.5367427902210501 ))"
	plywktg := "POLYGON (( 0.1161332552173457 -0.4654208754208744, 0.1161332552173457 -0.2431398008042315, 0.2824725516029848 -0.2431398008042315, 0.2824725516029848 -0.4654208754208744, 0.1161332552173457 -0.4654208754208744 ))"

	lna := NewLineString(coords)
	lnb := NewLineString(coords2)
	plya := NewPolygon(coords2)

	plyb := NewPolygonFromWKT(plywkt)
	plyc := NewPolygonFromWKT(plywktc)
	plyd := NewPolygonFromWKT(plywktd)
	plye := NewPolygonFromWKT(plywkte)
	plyf := NewPolygonFromWKT(plywktf)
	plyg := NewPolygonFromWKT(plywktg)

	var pnt_null *Point
	var ln_null *LineString
	var ply_null *Polygon

	g.Describe("Linestring - Relate", func() {
		g.It("should linestring relate", func() {
			g.Assert(lna.Envelope().Equals(lna.bbox.MBR)).IsTrue()
			g.Assert(lna.Intersects(pnt_null)).IsFalse()
			g.Assert(lna.Geometry().Intersects(pnt_null)).IsFalse()
			g.Assert(lna.Intersects(ln_null)).IsFalse()
			g.Assert(lna.Intersects(ply_null)).IsFalse()
			g.Assert(lna.Intersects(lnb)).IsFalse()
			g.Assert(lna.Intersects(plya)).IsFalse()
			g.Assert(lna.Intersects(plya)).IsFalse()

			g.Assert(plya.Intersects(ply_null)).IsFalse()
			g.Assert(plyb.Intersects(plyc)).IsFalse()

			g.Assert(plya.Intersects(lna)).IsFalse()
			g.Assert(plyb.Intersects(lna)).IsFalse()

			g.Assert(plyd.Intersects(plyb)).IsFalse()
			g.Assert(plyb.Intersects(plyd)).IsFalse()

			g.Assert(plyd.Intersects(plyc)).IsFalse()
			g.Assert(plyc.Intersects(plyd)).IsFalse()

			g.Assert(plye.Intersects(plyg)).IsFalse()
			g.Assert(plyg.Intersects(plye)).IsFalse()

			g.Assert(lna.Intersects(plyb)).IsFalse()
			g.Assert(lna.Intersects(plye)).IsTrue()
			g.Assert(lna.Intersects(plye.Geometry())).IsTrue()
			g.Assert(lna.Geometry().Intersects(plye)).IsTrue()
			g.Assert(lna.Intersects(plyf)).IsFalse()
		})
	})

}
