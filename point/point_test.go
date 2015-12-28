package point

import (
	"testing"
	. "github.com/franela/goblin"
	"math"
)

func TestFloatEqual(t *testing.T) {
	g := Goblin(t)
	g.Describe("geom.point", func() {
		p1 := Point{4, 5}
		//p2 := Point{4.0, 5.0}
		//p3 := Point{4, 5}
		//p4 := New([]float64{4, 5})
		//p5 := New([]float64{4.0, 2.5})
		//p6 := p5.Clone()


		g.It("x, y access", func() {
			g.Assert(p1.X()).Equal(4.)
			g.Assert(p1.Y()).Equal(5.)
		})
		g.It("clone equals", func() {
			g.Assert(p1.Equals(p1.Clone())).IsTrue()
		})
		g.It("as array", func() {
			g.Assert(p1.As_array()).Equal([2]float64{p1[x], p1[y]})
		})

	})

	g.Describe("Point distance", func() {
		g.It("sqrt(3**2,4**2) ", func() {
			pt := &Point{3., 0.}
			g.Assert(pt.Distance(Point{0., 4.})).Equal(5.0)
		})
		g.It("sqrt(2)", func() {
			pt := &Point{3, 4}
			g.Assert(pt.Distance(Point{4, 5})).Equal(math.Sqrt2)
		})
	})

	g.Describe("Point operators", func() {
		g.It("add ", func() {
			a, b := &Point{3., 0.}, &Point{0., 4.}
			g.Assert(a.Add(b)).Equal(Point{3., 4.})
		})

		g.It("sub ", func() {
			a, b := &Point{3., 4.}, &Point{4, 5}
			g.Assert(a.Sub(b)).Equal(Point{-1.0, -1.0})
		})
	})

	g.Describe("Point operators", func() {
		g.It("wkt ", func() {

			a := &Point{3.87, 7.45}
			g.Assert(a.Wkt()).Equal("POINT (3.87 7.45)")
		})
	})

}
