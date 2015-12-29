package point

import (
	"testing"
	. "github.com/franela/goblin"
	"math"
)

func TestPoint(t *testing.T) {
	g := Goblin(t)
	p1 := Point{4, 5}
	p2 := Point{4.0, 5.0}
	p3 := New([]float64{4, 5})
	p4 := New([]float64{4, 5.01})
	p5 := New([]float64{4})
	p6 := Point{4.0, math.NaN()}

	g.Describe("geom.point", func() {

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
			g.Assert(p1.As_array()).Equal([2]float64{p1[x], p1[y]})
		})

	})

	g.Describe("Point distance", func() {
		g.It("sqrt(3**2,4**2) ", func() {
			pt := Point{3., 0.}
			g.Assert(pt.Distance(Point{0., 4.})).Equal(5.0)
		})
		g.It("sqrt(2)", func() {
			pt := Point{3, 4}
			g.Assert(pt.Distance(Point{4, 5})).Equal(math.Sqrt2)
		})
	})

	g.Describe("Point operators", func() {
		g.It("add ", func() {
			a, b := Point{3., 0.}, Point{0., 4.}
			g.Assert(a.Add(b)).Equal(Point{3., 4.})
		})

		g.It("sub & neg ", func() {
			a, b := Point{3., 4.}, Point{4, 5}
			nb := b.Neg()
			g.Assert(a.Sub(b)).Equal(Point{-1.0, -1.0})
			g.Assert(nb).Equal(Point{-4, -5})
		})
	})

	g.Describe("type conversion & util", func() {
		g.It("wkt string", func() {
			a := Point{3.87, 7.45}
			g.Assert(a.String()).Equal("3.87 7.45")
			g.Assert(a.Wkt()).Equal("POINT (3.87 7.45)")
			g.Assert(a.Bbox()).Equal([]Point{{3.87, 7.45}, {3.87, 7.45}})
			g.Assert(a.Convex_hull()).Equal([]Point{{3.87, 7.45}, {3.87, 7.45}, {3.87, 7.45}, {3.87, 7.45}})
		})
	})

	g.Describe("type conversion & util", func() {
		g.It("string, wkt , bbox, chull", func() {
			a := Point{3.87, 7.45}
			g.Assert(a.String()).Equal("3.87 7.45")
			g.Assert(a.Wkt()).Equal("POINT (3.87 7.45)")
			g.Assert(a.Bbox()).Equal([]Point{{3.87, 7.45}, {3.87, 7.45}})
			g.Assert(a.Convex_hull()).Equal([]Point{{3.87, 7.45}, {3.87, 7.45}, {3.87, 7.45}, {3.87, 7.45}})
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
