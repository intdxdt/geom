package geom

import (
	"fmt"
	"github.com/franela/goblin"
	"github.com/intdxdt/math"
	"github.com/intdxdt/mbr"
	"testing"
	"time"
)

func sqr(x float64) float64 {
	return x * x
}

func TestSegment(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Segment", func() {
		g.It("should test segment intersection", func() {
			wkt := "POLYGON (( -0.3604422185430426 -10, -0.3604422185430426 0.5291138245033155, 10 0.5291138245033155, 10 -10, -0.3604422185430426 -10 ))"
			ply := NewPolygonFromWKT(wkt)
			var a = PointXY(0, 0)
			var b = PointXY(-3, 4)
			var c = PointXY(1.5, -2)
			var d = PointXY(-1.5, 2)
			var e = PointXY(0.5, 3)

			//f := &Point{-2, -2}
			var gk = Point{-1.5, -2.5}
			var h = Point{0.484154648492778, -0.645539531323704}
			var i = Point{0.925118053504632, -1.233490738006176}
			var k = Point{2, 2}
			var n = Point{1, 5}
			n.BBox()

			seg_ab := NewSegmentAB(a, b)
			ln_ab := NewLineString(Coordinates([]Point{a, b}))
			seg_de := NewSegmentAB(d, e)
			seg_cd := NewSegmentAB(c, d)
			seg_gkh := NewSegmentAB(gk, h)
			seg_hi := NewSegmentAB(h, i)
			seg_ak := NewSegmentAB(a, k)
			seg_kn := NewSegmentAB(k, n)

			g.Assert(seg_ab.Type().IsSegment()).IsTrue()
			g.Assert(seg_ab.IsSimple()).IsTrue()
			g.Assert(seg_ab.Type().IsLineString()).IsFalse()
			var box = mbr.CreateMBR(0, 0, -3, 4)
			var seg_ab_box = seg_ab.BBox()
			g.Assert(seg_ab_box.Equals(&box)).IsTrue()
			g.Assert(seg_ab.AsLinear()).Eql([]*LineString{ln_ab})
			g.Assert(seg_ab.WKT()).Eql(ln_ab.WKT())
			g.Assert(seg_ab.Intersects(k)).IsFalse()
			g.Assert(seg_ab.Intersects(seg_kn)).IsFalse()
			g.Assert(seg_ab.Geometry().Intersects(seg_kn)).IsFalse()
			g.Assert(seg_ab.Intersects(seg_kn.Geometry())).IsFalse()
			g.Assert(seg_ab.Intersects(seg_ak)).IsTrue()
			g.Assert(seg_ab.Geometry().Intersects(seg_ak)).IsTrue()
			g.Assert(seg_ab.Intersects(seg_ak.Geometry())).IsTrue()
			g.Assert(seg_ab.Intersects(ply)).IsTrue()
			g.Assert(ply.Intersects(seg_ab)).IsTrue()

			g.Assert(seg_kn.Intersects(ply)).IsFalse()
			g.Assert(ply.Intersects(seg_kn)).IsFalse()

			g.Assert(seg_ab.Intersection(seg_ak)).Eql([]Point{a})
			g.Assert(seg_ab.Distance(seg_ak)).Equal(0.0)
			fmt.Println(seg_ab.Distance(seg_kn))
			g.Assert(feq(seg_ab.Distance(seg_kn), 2.8)).IsTrue()

			pts := seg_ab.SegSegIntersection(seg_de)
			g.Assert(pts[0].Point).Equal(Point{-1.5, 2})

			pts = seg_ab.SegSegIntersection(seg_cd)
			ok := len(pts) > 0
			g.Assert(ok).IsTrue()
			g.Assert(pts[0].Point).Equal(Point{-1.5, 2})
			g.Assert(pts[1].Point).Equal(Point{0.0, 0.0})

			pts = seg_gkh.SegSegIntersection(seg_cd)
			g.Assert(len(pts)).Equal(1) //at h

			pts = seg_hi.SegSegIntersection(seg_cd)
			g.Assert(len(pts)).Equal(2) //at h, i

			pts = seg_hi.SegSegIntersection(seg_ab)
			ok = len(pts) > 0
			g.Assert(seg_hi.SegSegIntersects(seg_ab)).Equal(ok)
			g.Assert(ok).IsFalse()
			g.Assert(len(pts)).Equal(0) //empty

			pts = seg_ak.SegSegIntersection(seg_kn)
			ok = len(pts) > 0
			g.Assert(seg_ak.SegSegIntersects(seg_kn)).Equal(ok)
			g.Assert(ok).IsTrue()
			g.Assert(len(pts)).Equal(1)                  //at k
			g.Assert(pts[0].Point.Equals2D(&k)).IsTrue() //k
		})
	})

}

func TestSegDist(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("SegSeg and SegToPoint", func() {
		g.It("should test segment distance to point and seg", func() {
			g.Timeout(1 * time.Hour)

			var a = Pt(-0.8, -2.6)
			var b = Pt(-1, 1)
			var c = Pt(-2, 3)
			var d = Pt(7, -3)
			var e = Pt(1.6, 0.6)
			var f = Pt(-8, 4)
			var gi = Pt(10, -8)
			var j = Pt(-3.5, 4)
			var k = Pt(-5, 5)
			var l = Pt(8.5, -4)
			var m = Pt(10, -5)
			var t = Pt(1, 6)
			var u = Pt(6, 4)

			var n = Pt(1, 3)
			var o = Pt(6, 5)
			var expects = math.Round(1.1094003924504583, 12)

			var seg_ab = NewSegmentAB(a, b)
			var seg_ba = NewSegmentAB(b, a)
			var seg_cd = NewSegmentAB(c, d)
			var seg_dc = NewSegmentAB(d, c)
			var seg_dd = NewSegmentAB(d, d)
			var seg_ff = NewSegmentAB(f, f)
			var seg_ef = NewSegmentAB(e, f)
			var seg_fg = NewSegmentAB(f, gi)
			var seg_jk = NewSegmentAB(j, k)
			var seg_jj = NewSegmentAB(j, j)
			var seg_lm = NewSegmentAB(l, m)
			var seg_ll = NewSegmentAB(l, l)
			var seg_no = NewSegmentAB(n, o)
			var seg_tu = NewSegmentAB(t, u)

			g.Assert(math.Round(seg_ab.SegSegDistance(seg_ab), 12)).Equal(0.0)
			g.Assert(math.Round(seg_ab.SegSegDistance(seg_cd), 12)).Equal(expects)
			g.Assert(math.Round(seg_ab.SegSegDistance(seg_dc), 12)).Equal(expects)
			g.Assert(math.Round(seg_ba.SegSegDistance(seg_cd), 12)).Equal(expects)
			g.Assert(math.Round(seg_cd.SegSegDistance(seg_ab), 12)).Equal(expects)

			g.Assert(math.Round(seg_dc.SegSegDistance(seg_ef), 12)).Equal(0.0)
			g.Assert(seg_dd.SegSegDistance(seg_ff)).Equal(d.Distance(f))
			g.Assert(seg_ff.Length()).Equal(0.0)
			g.Assert(seg_ff.Distance(seg_jj)).Equal(seg_ff.A().Distance(seg_jj.A()))
			g.Assert(seg_ab.Length()).Equal(seg_ab.A().Distance(seg_ab.B()))

			g.Assert(math.Round(seg_dc.SegSegDistance(seg_fg), 12)).Equal(
				math.Round(2.496150883013531, 12),
			)

			g.Assert(math.Round(seg_dc.SegSegDistance(seg_lm), 12)).Equal(
				math.Round(d.Distance(&l), 12),
			)

			g.Assert(math.Round(seg_cd.SegSegDistance(seg_lm), 12)).Equal(
				math.Round(d.Distance(&l), 12),
			)
			g.Assert(math.Round(seg_dc.SegSegDistance(seg_ll), 12)).Equal(
				math.Round(d.Distance(&l), 12),
			)
			g.Assert(math.Round(seg_cd.SegSegDistance(seg_ll), 12)).Equal(
				math.Round(d.Distance(&l), 12),
			)

			g.Assert(math.Round(seg_dc.SegSegDistance(seg_jk), 12)).Equal(
				math.Round(c.Distance(&j), 12),
			)
			g.Assert(math.Round(seg_dc.SegSegDistance(seg_jj), 12)).Equal(
				math.Round(c.Distance(&j), 12),
			)
			g.Assert(math.Round(seg_jj.SegSegDistance(seg_dc), 12)).Equal(
				math.Round(c.Distance(&j), 12),
			)

			g.Assert(math.Round(seg_ab.SegSegDistance(seg_no), 12)).Equal(
				math.Round(b.Distance(&n), 12),
			)
			g.Assert(math.Round(seg_no.SegSegDistance(seg_ab), 12)).Equal(
				math.Round(n.Distance(&b), 12),
			)
			//no intersects tu
			g.Assert(math.Round(seg_no.SegSegDistance(seg_tu), 12)).Equal(0.0)

			a = Point{16.82295, 10.44635}
			b = Point{28.99656, 15.76452}
			on_ab := Point{25.32, 14.16}

			tpoints := []Point{
				{30., 0.},
				{15.78786, 25.26468},
				{-2.61504, -3.09018},
				{28.85125, 27.81773},
				a, b, on_ab,
			}

			var t_dists = []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
			var tvect = NewSegmentAB(a, b)
			var dists = make([]float64, len(tpoints))

			for i, tp := range tpoints {
				dists[i] = tvect.DistanceToPoint(&tp)
			}

			var pt1_out = PointFromWKT("POINT ( 49.8322373906287 49.1670033843562 )")
			var pt2_out = PointFromWKT("POINT (  26.70508112717612 29.46609249326697 )")
			var pnt3_in = PointFromWKT("POINT ( 27.439276564111122 38.76590136111034 )")
			var poly = NewPolygonFromWKT("POLYGON (( 35 10, 45 45, 15 40, 10 20, 35 10 ), ( 20 30, 35 35, 30 20, 20 30 ))")
			var lnr = NewLineStringFromWKT("LINESTRING ( 35 10, 45 45, 15 40, 10 20, 35 10 )")
			var pt_online = PointXY(45.00000, 45.000000000000000000000000001)
			ln := poly.Shell.AsLinear()[0]

			g.Assert(math.Round(ln.Distance(poly), 12)).Equal(math.Round(0, 12))
			g.Assert(math.Round(ln.Distance(&pt_online), 12)).Equal(math.Round(0, 12))
			g.Assert(math.Round(pt1_out.Distance(ln), 12)).Equal(math.Round(6.380786425247758, 12))
			g.Assert(math.Round(ln.Distance(&pt1_out), 12)).Equal(math.Round(6.380786425247758, 12))
			g.Assert(math.Round(pt1_out.Distance(poly), 12)).Equal(math.Round(6.380786425247758, 12))
			g.Assert(math.Round(poly.Distance(&pt1_out), 12)).Equal(math.Round(6.380786425247758, 12))
			g.Assert(math.Round(pt2_out.Distance(poly), 12)).Equal(math.Round(2.626841960149983, 12))
			g.Assert(math.Round(poly.Distance(&pt2_out), 12)).Equal(math.Round(2.626841960149983, 12))
			g.Assert(math.Round(pnt3_in.Distance(poly), 12)).Equal(math.Round(0.0, 12))
			g.Assert(math.Round(poly.Distance(&pnt3_in), 12)).Equal(math.Round(0.0, 12))

			var null_poly *Polygon
			var null_ln *LineString

			g.Assert(math.IsNaN(pt2_out.Distance(null_poly))).IsTrue()
			g.Assert(math.IsNaN(poly.Distance(null_poly))).IsTrue()
			g.Assert(math.IsNaN(poly.Distance(null_ln))).IsTrue()
			g.Assert(math.IsNaN(lnr.Distance(null_ln))).IsTrue()

			var seg_aa = NewSegmentAB(a, a)
			g.Assert(seg_aa.DistanceToPoint(&a)).Equal(0.0)
			g.Assert(a.SideOf(&a, &b).IsOn()).IsTrue()
			g.Assert(b.SideOf(&a, &b).IsOn()).IsTrue()

			seg_ab = NewSegmentAB(a, b)
			g.Assert(seg_ab.SideOf(&a).IsOn()).IsTrue()
			g.Assert(seg_ab.SideOf(&b).IsOn()).IsTrue()

			for i := range tpoints {
				g.Assert(math.Round(dists[i], 2)).Equal(math.Round(t_dists[i], 2))
			}
		})

		g.It("should test segment square distance to point and seg", func() {
			g.Timeout(1 * time.Hour)
			var a = Pt(-0.8, -2.6)
			var b = Pt(-1, 1)
			var c = Pt(-2, 3)
			var d = Pt(7, -3)
			var e = Pt(1.6, 0.6)
			var f = Pt(-8, 4)
			var gi = Pt(10, -8)
			var j = Pt(-3.5, 4)
			var k = Pt(-5, 5)
			var l = Pt(8.5, -4)
			var m = Pt(10, -5)
			var t = Pt(1, 6)
			var u = Pt(6, 4)

			var n = Pt(1, 3)
			var o = Pt(6, 5)
			var expects = math.Round(1.1094003924504583*1.1094003924504583, 12)

			seg_ab := NewSegmentAB(a, b)
			seg_ba := NewSegmentAB(b, a)
			seg_cd := NewSegmentAB(c, d)
			seg_dc := NewSegmentAB(d, c)
			seg_dd := NewSegmentAB(d, d)
			seg_ff := NewSegmentAB(f, f)
			seg_ef := NewSegmentAB(e, f)
			seg_fg := NewSegmentAB(f, gi)
			seg_jk := NewSegmentAB(j, k)
			seg_jj := NewSegmentAB(j, j)
			seg_lm := NewSegmentAB(l, m)
			seg_ll := NewSegmentAB(l, l)
			seg_no := NewSegmentAB(n, o)
			seg_tu := NewSegmentAB(t, u)

			g.Assert(math.Round(seg_ab.SquareSegSegDistance(seg_ab), 12)).Equal(0.0)
			g.Assert(math.Round(seg_ab.SquareSegSegDistance(seg_cd), 12)).Equal(expects)
			g.Assert(math.Round(seg_ab.SquareSegSegDistance(seg_dc), 12)).Equal(expects)
			g.Assert(math.Round(seg_ba.SquareSegSegDistance(seg_cd), 12)).Equal(expects)
			g.Assert(math.Round(seg_cd.SquareSegSegDistance(seg_ab), 12)).Equal(expects)

			g.Assert(math.Round(seg_dc.SquareSegSegDistance(seg_ef), 12)).Equal(0.0)
			g.Assert(seg_dd.SquareSegSegDistance(seg_ff)).Equal(d.Distance(f) * d.Distance(f))
			g.Assert(seg_ff.Length()).Equal(0.0)
			g.Assert(seg_ff.Distance(seg_jj)).Equal(seg_ff.A().Distance(seg_jj.A()))
			g.Assert(seg_ab.Length()).Equal(seg_ab.A().Distance(seg_ab.B()))

			g.Assert(math.Round(seg_dc.SquareSegSegDistance(seg_fg), 12)).Equal(
				math.Round(sqr(2.496150883013531), 12),
			)

			g.Assert(math.Round(seg_dc.SquareSegSegDistance(seg_lm), 12)).Equal(
				math.Round(sqr(d.Distance(&l)), 12),
			)

			g.Assert(math.Round(seg_cd.SquareSegSegDistance(seg_lm), 12)).Equal(
				math.Round(sqr(d.Distance(&l)), 12),
			)
			g.Assert(math.Round(seg_dc.SquareSegSegDistance(seg_ll), 12)).Equal(
				math.Round(sqr(d.Distance(&l)), 12),
			)
			g.Assert(math.Round(seg_cd.SquareSegSegDistance(seg_ll), 12)).Equal(
				math.Round(sqr(d.Distance(&l)), 12),
			)

			g.Assert(math.Round(seg_dc.SquareSegSegDistance(seg_jk), 12)).Equal(
				math.Round(sqr(c.Distance(&j)), 12),
			)
			g.Assert(math.Round(seg_dc.SquareSegSegDistance(seg_jj), 12)).Equal(
				math.Round(sqr(c.Distance(&j)), 12),
			)
			g.Assert(math.Round(seg_jj.SquareSegSegDistance(seg_dc), 12)).Equal(
				math.Round(sqr(c.Distance(&j)), 12),
			)

			g.Assert(math.Round(seg_ab.SquareSegSegDistance(seg_no), 12)).Equal(
				math.Round(sqr(b.Distance(&n)), 12),
			)
			g.Assert(math.Round(seg_no.SquareSegSegDistance(seg_ab), 12)).Equal(
				math.Round(sqr(n.Distance(&b)), 12),
			)
			//no intersects tu
			g.Assert(math.Round(seg_no.SquareSegSegDistance(seg_tu), 12)).Equal(0.0)

			a = Point{16.82295, 10.44635}
			b = Point{28.99656, 15.76452}
			var on_ab = Point{25.32, 14.16}
			var tpoints = []Point{
				{30., 0.}, {15.78786, 25.26468}, {-2.61504, -3.09018}, {28.85125, 27.81773}, a, b, on_ab,
			}
			var t_dists = []float64{
				14.847874769195268, 13.993482735483202, 23.68698164986413,
				12.054085875760137, 0.00, 0.00, 0.0014882215157408923}
			var tvect = NewSegmentAB(a, b)
			var dists = make([]float64, len(tpoints))

			for i, tp := range tpoints {
				dists[i] = tvect.SquareDistanceToPoint(&tp)
			}

			for i := range tpoints {
				g.Assert(math.Round(dists[i], 2)).Equal(math.Round(sqr(t_dists[i]), 2))
			}
		})

	})

}
