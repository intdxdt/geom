package geom

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/util/math"
    . "github.com/intdxdt/simplex/geom/mbr"
    "math"
)

func TestPoint(t *testing.T) {
    g := Goblin(t)
    p1 := NewPointXY(4, 5)
    p2 := NewPointXY(4.0, 5.0)
    p3 := NewPoint([]float64{4, 5})
    p4 := NewPoint([]float64{4, 5.01})
    p5 := NewPoint([]float64{4})
    p6 := &Point{4.0, math.NaN()}

    g.Describe("geom.point", func() {
        g.It("loads wkt as point", func() {
            g.Assert(NewPointFromWKT(p1.String())).Eql(p1)
            g.Assert(NewPointFromWKT(p4.String())).Eql(p4)
        })

        g.It("x, y access", func() {
            g.Assert(p1.Equals(p2)).IsTrue()
            g.Assert(p5.X()).Equal(4.0)
            g.Assert(p5.Y()).Equal(0.0)
            g.Assert(p3.X()).Equal(p1.X())
            g.Assert(p3.Y()).Equal(p1.Y())
        })

        g.It("clone equals", func() {
            pc := p1.Clone()
            g.Assert(p1.Equals(pc)).IsTrue()
        })

        g.It("as array", func() {
            g.Assert(p1.ToArray()).Equal([2]float64{p1[x], p1[y]})
        })

    })

    g.Describe("Point distance", func() {
        g.It("sqrt(3**2,4**2) ", func() {
            pt := &Point{3., 0.}
            g.Assert(pt.Distance(&Point{0., 4.})).Equal(5.0)
            g.Assert(pt.SquareDistance(&Point{0., 4.})).Equal(25.0)
        })
        g.It("sqrt(2)", func() {
            pt := &Point{3, 4}
            g.Assert(pt.Distance(&Point{4, 5})).Equal(math.Sqrt2)
            g.Assert(pt.SquareDistance(&Point{4, 5})).Equal(2.0)
        })
    })

    g.Describe("Point operators", func() {
        g.It("add ", func() {
            a, b := &Point{3., 0.}, &Point{0., 4.}
            g.Assert(a.Add(b)).Equal(&Point{3., 4.})
        })

        g.It("sub & neg ", func() {
            a, b := &Point{3., 4.}, &Point{4, 5}
            nb := b.Neg()
            g.Assert(a.Sub(b)).Equal(&Point{-1.0, -1.0})
            g.Assert(nb).Equal(&Point{-4, -5})
        })
    })

    g.Describe("type conversion & util", func() {
        g.It("wkt string", func() {
            a := Point{3.87, 7.45}
            g.Assert(a.String()).Equal("POINT (3.87 7.45)")
            g.Assert(a.BBox()).Equal(NewMBR(3.87, 7.45, 3.87, 7.45))
            g.Assert(a.ConvexHull()).Equal([]*Point{{3.87, 7.45}, {3.87, 7.45}, {3.87, 7.45}, {3.87, 7.45}})
        })
    })

    g.Describe("type conversion & util", func() {
        g.It("string, wkt , bbox, chull", func() {
            a := Point{3.87, 7.45}
            g.Assert(a.String()).Equal("POINT (3.87 7.45)")
            g.Assert(a.BBox()).Equal(NewMBR(3.87, 7.45, 3.87, 7.45))
            g.Assert(a.ConvexHull()).Equal([]*Point{{3.87, 7.45}, {3.87, 7.45}, {3.87, 7.45}, {3.87, 7.45}})
        })
    })

    g.Describe("point relates", func() {
        g.It("intersect , equals, isnull ", func() {
            g.Assert(p3.Equals(p1)).IsTrue()
            g.Assert(p3.Intersects(p1)).IsTrue()
            g.Assert(p3.Disjoint(p1)).IsFalse()
            g.Assert(p3.Disjoint(p4)).IsTrue()
            g.Assert(p6.IsNull()).IsTrue()
        })
    })

}

func TestMagDist(t *testing.T) {
    g := Goblin(t)
    g.Describe("Point - Vector Magnitude", func() {
        g.It("should test vector magnitude and distance", func() {
            a := &Point{0, 0 }
            b := &Point{3, 4 }

            g.Assert(NewPointXY(1, 1).Magnitude()).Equal(math.Sqrt2)
            g.Assert(Round(NewPointXY(-3, 2).Magnitude(), 8)).Equal(
                Round(3.605551275463989, 8),
            )

            g.Assert(NewPointXY(3, 4).Magnitude()).Equal(5.0)
            g.Assert(a.Distance(b)).Equal(5.0)

            g.Assert(NewPointXY(3, 4).SquareMagnitude()).Equal(25.0)
            g.Assert(a.SquareDistance(b)).Equal(25.0)

            g.Assert(NewPointXY(4.587, 0.).Magnitude()).Equal(4.587)
        })
    })

}

func TestDotProduct(t *testing.T) {
    g := Goblin(t)
    g.Describe("Point - Vector Dot Product", func() {
        g.It("should test dot product", func() {
            dot_prod := NewPointXY(1.2, -4.2).DotProduct(NewPointXY(1.2, -4.2))
            g.Assert(19.08).Equal(Round(dot_prod, 8))
        })
    })

}

func TestUnit(t *testing.T) {
    g := Goblin(t)
    g.Describe("Point -  Unit Vector", func() {
        g.It("should test unit vector", func() {
            v := &Point{-3, 2}
            unit_v := v.UnitVector()
            for i, v := range *unit_v {
                (*unit_v)[i] = Round(v, 6)
            }
            g.Assert(unit_v).Equal(&Point{-0.83205, 0.5547})
        })
    })

}

func TestAngleAtPnt(t *testing.T) {
    g := Goblin(t)
    g.Describe("Point - Angle at Point", func() {
        g.It("should test angle formed at point by vector", func() {
            a := &Point{-1.28, 0.74}
            b := &Point{1.9, 4.2}
            c := &Point{3.16, -0.84}
            g.Assert(Round(a.AngleAtPoint(b, c), 8)).Equal(Round(1.1694239325184717, 8), )
            g.Assert(Round(b.AngleAtPoint(a, c), 8)).Equal(Round(0.9882331199311394, 8), )
        })
    })

}
