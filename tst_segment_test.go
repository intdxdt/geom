package geom

import (
	"testing"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/math"
	"github.com/franela/goblin"
	"fmt"
)

func TestSegment(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("Segment", func() {
		g.It("should test segment intersection", func() {
			wkt := "POLYGON (( -0.3604422185430426 -10, -0.3604422185430426 0.5291138245033155, 10 0.5291138245033155, 10 -10, -0.3604422185430426 -10 ))"
			ply := NewPolygonFromWKT(wkt)
			a := NewPointXY(0, 0)
			b := NewPointXY(-3, 4)
			c := NewPointXY(1.5, -2)
			d := NewPointXY(-1.5, 2)
			e := NewPointXY(0.5, 3)
			//f := &Point{-2, -2}
			gk := &Point{-1.5, -2.5}
			h := &Point{0.484154648492778, -0.645539531323704}
			i := &Point{0.925118053504632, -1.233490738006176}
			k := &Point{2, 2}
			n := &Point{1, 5}

			seg_ab := NewSegment(a, b)
			ln_ab := NewLineString([]*Point{a, b})
			seg_de := &Segment{A: d, B: e}

			seg_cd := &Segment{c, d, nil}
			seg_gkh := &Segment{gk, h, nil}
			seg_hi := &Segment{A: h, B: i}
			seg_ak := &Segment{A: a, B: k}
			seg_kn := &Segment{A: k, B: n}

			g.Assert(seg_ab.Type().IsSegment()).IsTrue()
			g.Assert(seg_ab.IsSimple()).IsTrue()
			g.Assert(seg_ab.Type().IsLineString()).IsFalse()
			g.Assert(seg_ab.BBox().Equals(mbr.NewMBR(0, 0, -3, 4))).IsTrue()
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

			g.Assert(seg_ab.Intersection(seg_ak)).Eql([]*Point{a})
			g.Assert(seg_ab.Distance(seg_ak)).Equal(0.0)
			fmt.Println(seg_ab.Distance(seg_kn))
			g.Assert(math.FloatEqual(seg_ab.Distance(seg_kn), 2.8)).IsTrue()

			pts, ok := seg_ab.SegSegIntersection(seg_de, false)
			g.Assert(ok).IsTrue()
			g.Assert(pts[0]).Equal(&Point{-1.5, 2})

			pts, ok = seg_ab.SegSegIntersection(seg_cd, false)
			g.Assert(ok).IsTrue()
			g.Assert(pts[0]).Equal(&Point{-1.5, 2})
			g.Assert(pts[1]).Equal(&Point{0.0, 0.0})

			pts, ok = seg_gkh.SegSegIntersection(seg_cd, false)
			g.Assert(ok).IsTrue()
			g.Assert(len(pts)).Equal(1) //at h

			pts, ok = seg_hi.SegSegIntersection(seg_cd, false)
			g.Assert(ok).IsTrue()
			g.Assert(len(pts)).Equal(2) //at h, i

			pts, ok = seg_hi.SegSegIntersection(seg_ab, false)
			g.Assert(seg_hi.SegSegIntersects(seg_ab, false)).Equal(ok)
			g.Assert(ok).IsFalse()
			g.Assert(len(pts)).Equal(0) //empty

			pts, ok = seg_ak.SegSegIntersection(seg_kn, false)
			g.Assert(seg_ak.SegSegIntersects(seg_kn, false)).Equal(ok)
			g.Assert(ok).IsTrue()
			g.Assert(len(pts)).Equal(1) //at k
			g.Assert(pts[0]).Equal(k)   //k
		})
	})

}

func TestSegDist(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("SegSeg and SegToPoint", func() {
		g.It("should test segment distance to point and seg", func() {
			var a = NewPointXY(-0.8, -2.6)
			var b = NewPointXY(-1, 1)
			var c = NewPointXY(-2, 3)
			var d = NewPointXY(7, -3)
			var e = NewPointXY(1.6, 0.6)
			var f = NewPointXY(-8, 4)
			var gi = NewPointXY(10, -8)
			var j = NewPointXY(-3.5, 4)
			var k = NewPointXY(-5, 5)
			var l = NewPointXY(8.5, -4)
			var m = NewPointXY(10, -5)
			var t = NewPointXY(1, 6)
			var u = NewPointXY(6, 4)

			var n = NewPointXY(1, 3)
			var o = NewPointXY(6, 5)
			var expects = math.Round(1.1094003924504583, 12)

			seg_ab := NewSegment(a, b)
			seg_ba := NewSegment(b, a)
			seg_cd := NewSegment(c, d)
			seg_dc := NewSegment(d, c)
			seg_dd := NewSegment(d, d)
			seg_ff := NewSegment(f, f)
			seg_ef := NewSegment(e, f)
			seg_fg := NewSegment(f, gi)
			seg_jk := NewSegment(j, k)
			seg_jj := NewSegment(j, j)
			seg_lm := NewSegment(l, m)
			seg_ll := NewSegment(l, l)
			seg_no := NewSegment(n, o)
			seg_tu := NewSegment(t, u)

			g.Assert(math.Round(seg_ab.SegSegDistance(seg_ab), 12)).Equal(0.0)
			g.Assert(math.Round(seg_ab.SegSegDistance(seg_cd), 12)).Equal(expects)
			g.Assert(math.Round(seg_ab.SegSegDistance(seg_dc), 12)).Equal(expects)
			g.Assert(math.Round(seg_ba.SegSegDistance(seg_cd), 12)).Equal(expects)
			g.Assert(math.Round(seg_cd.SegSegDistance(seg_ab), 12)).Equal(expects)

			g.Assert(math.Round(seg_dc.SegSegDistance(seg_ef), 12)).Equal(0.0)
			g.Assert(seg_dd.SegSegDistance(seg_ff)).Equal(d.Distance(f))
			g.Assert(math.Round(seg_dc.SegSegDistance(seg_fg), 12)).Equal(
				math.Round(2.496150883013531, 12),
			)

			g.Assert(math.Round(seg_dc.SegSegDistance(seg_lm), 12)).Equal(
				math.Round(d.Distance(l), 12),
			)

			g.Assert(math.Round(seg_cd.SegSegDistance(seg_lm), 12)).Equal(
				math.Round(d.Distance(l), 12),
			)
			g.Assert(math.Round(seg_dc.SegSegDistance(seg_ll), 12)).Equal(
				math.Round(d.Distance(l), 12),
			)
			g.Assert(math.Round(seg_cd.SegSegDistance(seg_ll), 12)).Equal(
				math.Round(d.Distance(l), 12),
			)

			g.Assert(math.Round(seg_dc.SegSegDistance(seg_jk), 12)).Equal(
				math.Round(c.Distance(j), 12),
			)
			g.Assert(math.Round(seg_dc.SegSegDistance(seg_jj), 12)).Equal(
				math.Round(c.Distance(j), 12),
			)
			g.Assert(math.Round(seg_jj.SegSegDistance(seg_dc), 12)).Equal(
				math.Round(c.Distance(j), 12),
			)

			g.Assert(math.Round(seg_ab.SegSegDistance(seg_no), 12)).Equal(
				math.Round(b.Distance(n), 12),
			)
			g.Assert(math.Round(seg_no.SegSegDistance(seg_ab), 12)).Equal(
				math.Round(n.Distance(b), 12),
			)
			//no intersects tu
			g.Assert(math.Round(seg_no.SegSegDistance(seg_tu), 12)).Equal(0.0)

			a = &Point{16.82295, 10.44635}
			b = &Point{28.99656, 15.76452}
			on_ab := &Point{25.32, 14.16}

			tpoints := []*Point{
				{30., 0.},
				{15.78786, 25.26468},
				{-2.61504, -3.09018},
				{28.85125, 27.81773},
				a, b, on_ab,
			}

			t_dists := []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
			tvect := NewSegment(a, b)
			dists := make([]float64, len(tpoints))

			for i, tp := range tpoints {
				dists[i] = tvect.DistanceToPoint(tp)
			}

			pt1_out := NewPointFromWKT("POINT ( 49.8322373906287 49.1670033843562 )")
			pt2_out := NewPointFromWKT("POINT (  26.70508112717612 29.46609249326697 )")
			pnt3_in := NewPointFromWKT("POINT ( 27.439276564111122 38.76590136111034 )")
			poly := NewPolygonFromWKT("POLYGON (( 35 10, 45 45, 15 40, 10 20, 35 10 ), ( 20 30, 35 35, 30 20, 20 30 ))")
			var pt_online = NewPointXY(45.00000, 45.000000000000000000000000001)
			ln := poly.Shell.AsLinear()[0]

			g.Assert(math.Round(ln.Distance(poly), 12)).Equal(math.Round(0, 12))
			g.Assert(math.Round(ln.Distance(pt_online), 12)).Equal(math.Round(0, 12))
			g.Assert(math.Round(pt1_out.Distance(ln), 12)).Equal(math.Round(6.380786425247758, 12))
			g.Assert(math.Round(ln.Distance(pt1_out), 12)).Equal(math.Round(6.380786425247758, 12))
			g.Assert(math.Round(pt1_out.Distance(poly), 12)).Equal(math.Round(6.380786425247758, 12))
			g.Assert(math.Round(poly.Distance(pt1_out), 12)).Equal(math.Round(6.380786425247758, 12))
			g.Assert(math.Round(pt2_out.Distance(poly), 12)).Equal(math.Round(2.626841960149983, 12))
			g.Assert(math.Round(poly.Distance(pt2_out), 12)).Equal(math.Round(2.626841960149983, 12))
			g.Assert(math.Round(pnt3_in.Distance(poly), 12)).Equal(math.Round(0.0, 12))
			g.Assert(math.Round(poly.Distance(pnt3_in), 12)).Equal(math.Round(0.0, 12))

			var null_pt *Point
			var null_poly *Polygon
			var null_ln *LineString
			g.Assert(math.IsNaN(poly.Distance(null_pt))).IsTrue()
			g.Assert(math.IsNaN(pt2_out.Distance(null_pt))).IsTrue()
			g.Assert(math.IsNaN(ln.Distance(null_pt))).IsTrue()
			g.Assert(math.IsNaN(pt2_out.Distance(null_poly))).IsTrue()
			g.Assert(math.IsNaN(poly.Distance(null_poly))).IsTrue()
			g.Assert(math.IsNaN(poly.Distance(null_ln))).IsTrue()

			var seg_aa = NewSegment(a, a)
			g.Assert(seg_aa.DistanceToPoint(a)).Equal(0.0)
			g.Assert(a.SideOf(a, b).IsOn()).IsTrue()
			g.Assert(b.SideOf(a, b).IsOn()).IsTrue()

			seg_ab = NewSegment(a, b)
			g.Assert(seg_ab.SideOf(a).IsOn()).IsTrue()
			g.Assert(seg_ab.SideOf(b).IsOn()).IsTrue()

			for i := range tpoints {
				g.Assert(math.Round(dists[i], 2)).Equal(math.Round(t_dists[i], 2))
			}
		})

	})

}
