package mbr

import (
	"testing"
	. "github.com/franela/goblin"
	"github.com/intdxdt/simplex/geom/point"
	"math"
)

func TestMBR(t *testing.T) {
	g := Goblin(t)

	m00 := New(point.Point{0, 0}, point.Point{0, 0})
	m00.Expand_xy(2, 2)

	n00 := New(point.Point{0, 0}, point.Point{0, 0})
	n00.Expand_xy(-2, -2)

	m0 := New(point.Point{1, 1}, point.Point{1, 1})
	m0.Expand_by(1, 1)

	m1 := New(point.Point{0, 0}, point.Point{2, 2})
	m2 := New(point.Point{4, 5}, point.Point{8, 9})
	m3 := New(point.Point{1.7, 1.5}, point.Point{5, 9})
	m4 := New(point.Point{5, 0}, point.Point{8, 2})
	m5 := New(point.Point{5, 11}, point.Point{8, 9})
	m6 := New(point.Point{0, 0}, point.Point{2, -2})
	m7 := New(point.Point{-2, 1}, point.Point{4, -2})
	m8 := New(point.Point{-1, 0}, point.Point{1, -1.5})

	p := point.Point{1.7, 1.5}  // POINT(1.7 1.5)

	g.Describe("minimum bounding box", func() {

		m0123 := New(point.Point{0, 2}, point.Point{1, 3})
		clone_m0123 := m0123.Clone()

		g.It("equals ", func() {
			g.Assert(m1.As_array()).Equal([]float64{0, 0, 2, 2})
			g.Assert(clone_m0123.Equals(m0123)).IsTrue()
			g.Assert(m0.Equals(m1)).IsTrue()
			g.Assert(m00.Equals(m1)).IsTrue()

		})
		g.It("intersects, distance", func() {
			g.Assert(m1.Intersects_point(p)).IsTrue()

			g.Assert(m00.Intersects(n00)).IsTrue()
			nm00, success := m00.Intersection(n00)
			g.Assert(success).IsTrue()

			g.Assert(nm00.ll.Equals(point.Point{0, 0})).IsTrue()
			g.Assert(nm00.ur.Equals(point.Point{0, 0})).IsTrue()

			g.Assert(m1.Intersects(m2)).IsFalse()
			_, success = m1.Intersection(m2)
			g.Assert(success).IsFalse()
			g.Assert(m1.Intersects(m3)).IsTrue()
			g.Assert(m2.Intersects(m3)).IsTrue()

			m13, _ := m1.Intersection(m3)
			m23, _ := m2.Intersection(m3)
			_m13 := []float64{1.7, 1.5, 2, 2}
			_m23 := []float64{4, 5, 5, 9}

			g.Assert(_m13).Equal(m13.As_array())
			g.Assert(_m23).Equal(m23.As_array())

			g.Assert(m3.Intersects(m4)).IsTrue()
			g.Assert(m2.Intersects(m5)).IsTrue()
			g.Assert(m7.Intersects(m6)).IsTrue()
			g.Assert(m6.Intersects(m7)).IsTrue()

			m67, _ := m6.Intersection(m7)
			m76, _ := m7.Intersection(m6)
			m78, _ := m7.Intersection(m8)

			g.Assert(m67.Equals(m6)).IsTrue()
			g.Assert(m67.Equals(m76)).IsTrue()
			g.Assert(m78.Equals(m8)).IsTrue()

			m25, _ := m2.Intersection(m5)
			m34, _ := m3.Intersection(m4)

			g.Assert(m25.Width()).Equal(m5.Width())
			g.Assert(m25.Height()).Equal(0.0)
			g.Assert(m34.Width()).Equal(0.0)
			g.Assert(m34.Height()).Equal(0.5)
			g.Assert(m3.Distance(m4)).Equal(0.0)

			g.Assert(m1.Distance(m2)).Equal(math.Hypot(2, 3))
			g.Assert(m1.Distance(m3)).Equal(0.0)

			a := New(point.Point{-7.703505430214746 ,3.0022503796012305},
				point.Point{-5.369812194018422, 5.231449888803689})
			g.Assert(m1.Distance(a)).Equal(math.Hypot(-5.369812194018422, 3.0022503796012305-2))

			b := New(point.Point{-4.742849832055231, -4.1033230559816065},
				point.Point{-1.9563504455521576, -2.292098454754609})
			g.Assert(m1.Distance(b)).Equal(math.Hypot(-1.9563504455521576, -2.292098454754609))

		})

		g.It("contains, disjoint , contains completely", func() {
			p1 := point.Point{-5.95, 9.28}
			p2 := point.Point{-0.11, 12.56}
			p3 := point.Point{3.58, 11.79}
			p4 := point.Point{-1.16, 14.71}

			mp12 := New(p1, p2)
			mp34 := New(p3, p4)

			// intersects but segment are disjoint
			g.Assert(mp12.Intersects(mp34)).IsTrue()
			g.Assert(mp12.Intersects_bounds(p3, p4)).IsTrue()
			g.Assert(mp12.Intersects_bounds(m1.ll, m1.ur)).IsFalse()
			g.Assert(mp12.Intersects_point(p3)).IsFalse()
			g.Assert(m1.Contains_xy(1, 1)).IsTrue()

			mbr11 := New(point.Point{1, 1}, point.Point{1.5, 1.5})
			mbr12 := New(point.Point{1, 1}, point.Point{2, 2})
			mbr13 := New(point.Point{1, 1}, point.Point{2.000045, 2.00001})
			mbr14 := New(point.Point{2.000045, 2.00001}, point.Point{4.000045, 4.00001})

			g.Assert(m1.Contains(mbr11)).IsTrue()
			g.Assert(m1.Contains(mbr12)).IsTrue()
			g.Assert(m1.Contains(mbr13)).IsFalse()
			g.Assert(m1.Disjoint(mbr13)).IsFalse()  // False
			g.Assert(m1.Disjoint(mbr14)).IsTrue()   // True disjoint


			g.Assert(m1.Contains_xy(1.5, 1.5)).IsTrue()
			g.Assert(m1.Contains_xy(2, 2)).IsTrue()

			g.Assert(m1.Completely_contains_mbr(mbr11)).IsTrue()
			g.Assert(m1.Completely_contains_xy(1.5, 1.5)).IsTrue()
			g.Assert(m1.Completely_contains_xy(1.5, 1.5)).IsTrue()
			g.Assert(m1.Completely_contains_xy(2, 2)).IsFalse()
			g.Assert(m1.Completely_contains_mbr(mbr12)).IsFalse()
			g.Assert(m1.Completely_contains_mbr(mbr13)).IsFalse()
		})

		g.It("translate, expand by, area", func() {

			ma := New(point.Point{0, 0}, point.Point{2, 2})
			mb := New(point.Point{-1, -1}, point.Point{1.5, 1.9})
			mc := New(point.Point{1.7, 1.5}, point.Point{5, 9})
			md := ma.Clone()
			ma.Expand(mc)
			md.Expand(mb)

			g.Assert(ma.As_array()).Equal([]float64{0, 0, 5, 9}) //ma modified by expand
			g.Assert(mc.As_array()).Equal([]float64{1.7, 1.5, 5, 9})//should not be touched
			g.Assert(md.As_array()).Equal([]float64{-1,-1, 2,2}) //ma modified by expand

			//mc area
			g.Assert(mc.Area()).Equal(24.75)

			mt := m1.Translate(1, 1)
			mby := m1.Clone()
			mby.Expand_by(-3, -3)

			m1c := m1.Center()
			mtc := mt.Center()

			g.Assert(m1c.Equals(point.Point{1, 1}))
			g.Assert(mtc.Equals(point.Point{2, 2}))
			g.Assert(mt.As_array()).Equal([]float64{1, 1, 3, 3})
			g.Assert(mby.As_array()).Equal([]float64{-1, -1, 3, 3})
		})

		g.It("is null, string", func() {
			mm := m1.Clone()
			mm.ur[1] = math.NaN()
			g.Assert(mm.Is_null()).IsTrue()
			g.Assert(m1.Is_null()).IsFalse()
			g.Assert(m1.String()).Equal("POLYGON ((0 0, 0 2, 2 2, 2 0, 0 0))")
		})

	})

}
