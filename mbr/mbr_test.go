package mbr

import (
    "testing"
    "math"
    . "github.com/franela/goblin"
    util "github.com/intdxdt/simplex/util/math"
)

func TestMBR(t *testing.T) {
    g := Goblin(t)

    m00 := NewMBR(0, 0, 0, 0)
    m00.ExpandIncludeXY(2, 2)

    n00 := NewMBR(0, 0, 0, 0)
    n00.ExpandIncludeXY(-2, -2)

    m0 := NewMBR(1, 1, 1, 1)
    m0.ExpandByDelta(1, 1)

    m1 := NewMBR(0, 0, 2, 2)
    m2 := NewMBR(4, 5, 8, 9)
    m3 := NewMBR(1.7, 1.5, 5, 9)
    m4 := NewMBR(5, 0, 8, 2)
    m5 := NewMBR(5, 11, 8, 9)
    m6 := NewMBR(0, 0, 2, -2)
    m7 := NewMBR(-2, 1, 4, -2)
    m8 := NewMBR(-1, 0, 1, -1.5)

    p := []float64{1.7, 1.5, 3.4}  // POINT(1.7 1.5, 3.4)
    p0 := []float64{1.7}  // POINT(1.7 1.5)

    g.Describe("minimum bounding box", func() {

        m0123 := NewMBR(0, 2, 1, 3)
        clone_m0123 := m0123.Clone()

        g.It("equals ", func() {
            g.Assert(m1.AsArray()).Equal([]float64{0, 0, 2, 2})
            g.Assert(clone_m0123.Equals(m0123)).IsTrue()
            g.Assert(m0.Equals(m1)).IsTrue()
            g.Assert(m0.BBox() == m0).IsTrue()
            g.Assert(m00.Equals(m1)).IsTrue()
        })

        g.It("intersects, distance", func() {
            g.Assert(m1.IntersectsPoint(p)).IsTrue()
            g.Assert(m1.IntersectsPoint(p0)).IsFalse()

            g.Assert(m00.Intersects(n00)).IsTrue()
            nm00, success := m00.Intersection(n00)
            g.Assert(success).IsTrue()

            g.Assert(nm00[x1] == 0.0 && nm00[y1] == 0.0).IsTrue()
            g.Assert(nm00[x2] == 0.0 && nm00[y2] == 0.0).IsTrue()

            g.Assert(m1.Intersects(m2)).IsFalse()
            _, success = m1.Intersection(m2)
            g.Assert(success).IsFalse()
            g.Assert(m1.Intersects(m3)).IsTrue()
            g.Assert(m2.Intersects(m3)).IsTrue()

            m13, _ := m1.Intersection(m3)
            m23, _ := m2.Intersection(m3)
            _m13 := []float64{1.7, 1.5, 2, 2}
            _m23 := []float64{4, 5, 5, 9}

            g.Assert(_m13).Equal(m13.AsArray())
            g.Assert(_m23).Equal(m23.AsArray())

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

            d := math.Hypot(2, 3)
            g.Assert(m1.Distance(m2)).Equal(d)
            g.Assert(m1.DistanceSquare(m2)).Equal(util.Round(d * d, 12))
            g.Assert(m1.Distance(m3)).Equal(0.0)
            g.Assert(m1.DistanceSquare(m3)).Equal(0.0)

            a := NewMBR(
                -7.703505430214746, 3.0022503796012305,
                -5.369812194018422, 5.231449888803689)
            g.Assert(m1.Distance(a)).Equal(math.Hypot(-5.369812194018422, 3.0022503796012305 - 2))

            b := NewMBR(-4.742849832055231, -4.1033230559816065,
                -1.9563504455521576, -2.292098454754609)
            g.Assert(m1.Distance(b)).Equal(math.Hypot(-1.9563504455521576, -2.292098454754609))

        })

        g.It("contains, disjoint , contains completely", func() {
            p1 := []float64{-5.95, 9.28}
            p2 := []float64{-0.11, 12.56}
            p3 := []float64{3.58, 11.79}
            p4 := []float64{-1.16, 14.71}
            p4x := []float64{-1.16}

            mp12 := NewMBR(p1[x1], p1[y1], p2[x1], p2[y1])
            mp34 := NewMBR(p3[x1], p3[y1], p4[x1], p4[y1])

            // intersects but segment are disjoint
            g.Assert(mp12.Intersects(mp34)).IsTrue()
            g.Assert(mp12.IntersectsBounds(p3, p4)).IsTrue()
            g.Assert(mp12.IntersectsBounds(p3, p4x)).IsFalse()
            g.Assert(mp12.IntersectsBounds(
                []float64{m1[x1], m1[y1]},
                []float64{m1[x2], m1[y2]},
            )).IsFalse()
            g.Assert(mp12.IntersectsPoint(p3)).IsFalse()
            g.Assert(m1.ContainsXY(1, 1)).IsTrue()

            mbr11 := NewMBR(1, 1, 1.5, 1.5)
            mbr12 := NewMBR(1, 1, 2, 2)
            mbr13 := NewMBR(1, 1, 2.000045, 2.00001)
            mbr14 := NewMBR(2.000045, 2.00001, 4.000045, 4.00001)

            g.Assert(m1.Contains(mbr11)).IsTrue()
            g.Assert(m1.Contains(mbr12)).IsTrue()
            g.Assert(m1.Contains(mbr13)).IsFalse()
            g.Assert(m1.Disjoint(mbr13)).IsFalse()  // False
            g.Assert(m1.Disjoint(mbr14)).IsTrue()   // True disjoint

            g.Assert(m1.ContainsXY(1.5, 1.5)).IsTrue()
            g.Assert(m1.ContainsXY(2, 2)).IsTrue()

            g.Assert(m1.CompletelyContainsMBR(mbr11)).IsTrue()
            g.Assert(m1.CompletelyContainsXY(1.5, 1.5)).IsTrue()
            g.Assert(m1.CompletelyContainsXY(1.5, 1.5)).IsTrue()
            g.Assert(m1.CompletelyContainsXY(2, 2)).IsFalse()
            g.Assert(m1.CompletelyContainsMBR(mbr12)).IsFalse()
            g.Assert(m1.CompletelyContainsMBR(mbr13)).IsFalse()
        })

        g.It("translate, expand by, area", func() {

            ma := NewMBR(0, 0, 2, 2)
            mb := NewMBR(-1, -1, 1.5, 1.9)
            mc := NewMBR(1.7, 1.5, 5, 9)
            md := ma.Clone()
            ma.ExpandIncludeMBR(mc)
            md.ExpandIncludeMBR(mb)

            g.Assert(ma.AsArray()).Equal([]float64{0, 0, 5, 9}) //ma modified by expand
            g.Assert(mc.AsArray()).Equal([]float64{1.7, 1.5, 5, 9})//should not be touched
            g.Assert(md.AsArray()).Equal([]float64{-1, -1, 2, 2}) //ma modified by expand

            //mc area
            g.Assert(mc.Area()).Equal(24.75)

            mt := m1.Translate(1, 1)
            mby := m1.Clone()
            mby.ExpandByDelta(-3, -3)

            m1c := m1.Center()
            mtc := mt.Center()

            g.Assert(m1c).Eql([]float64{1, 1})
            g.Assert(mtc).Eql([]float64{2, 2})
            g.Assert(mt.AsArray()).Equal([]float64{1, 1, 3, 3})
            g.Assert(mby.AsArray()).Equal([]float64{-1, -1, 3, 3})
        })

        g.It("is string", func() {
            g.Assert(m1.String()).Equal("POLYGON ((0 0, 0 2, 2 2, 2 0, 0 0))")
        })


    })

}
