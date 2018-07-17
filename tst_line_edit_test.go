package geom

import (
	"testing"
	"github.com/intdxdt/math"
	"github.com/franela/goblin"
	"fmt"
)

func TestLineStringEdit(t *testing.T) {
	var g = goblin.Goblin(t)
	var a = Point{-2, -4}
	var b = Point{1, -1}
	var c = Point{-1, 4}

	var pts = []Point{a, b, c}

	g.Describe("Linestring", func() {
		g.It("should test length on append", func() {
			var ln  = NewLineString(pts)

			g.Assert(
				math.Round(ln.length, 10)).Equal(
				math.Round(9.62780549425, 10),
			)
			g.Assert(len(ln.chains)).Equal(2)



			//test util pop_coords
			g.Assert(len(pts)).Equal(3)

			v, pts := pop_coords(pts)
			g.Assert(len(pts)).Equal(2)
			g.Assert(v).Equal(c)

			v, pts = pop_coords(pts)
			g.Assert(len(pts)).Equal(1)
			g.Assert(v).Equal(b)

			v, pts = pop_coords(pts)
			g.Assert(len(pts)).Equal(0)
			g.Assert(v).Equal(a)

			v, pts = pop_coords(pts)
			g.Assert(len(pts)).Equal(0)
			g.Assert(v.IsNull()).IsTrue()
		})
		g.It("should test intersection", func() {
			var a = Point{-4.975454545454546, 0.2551515151515151}
			var b = Point{-3.9389015151515148, 1.156155303030303}
			var c = Point{1.5, -2}
			var d = Point{-1.5, 2}
			var h_prime = Point{0.4841546875717521, -0.6455395491824757}

			var h = Point{0.484154648492778, -0.645539531323704}
			var i = Point{0.925118053504632, -1.233490738006176}
			var ln_e *LineString
			fmt.Println(">? ln_e >> ", ln_e == nil)
			var pt_e Point
			var ln_ab       = NewLineString([]Point{a, b})
			var ln_cd       = NewLineString([]Point{c, d})
			var ln_cd_clone = ln_cd.Clone()
			var ln_hi       = NewLineString([]Point{h, i})

			var ok = ln_cd.Intersects(ln_ab)
			g.Assert(ok).IsFalse()
			g.Assert(ln_cd_clone.Intersects(ln_ab)).IsFalse()

			g.Assert(ln_cd.Intersects(ln_e)).IsFalse()

			ok = ln_cd.Intersects(ln_hi)
			g.Assert(ok).IsTrue() //at h, i

			var pts = ln_cd.Intersection(ln_hi)
			g.Assert(len(pts)).Equal(2) //at h, i

			pts = ln_cd.Intersection(ln_ab)
			g.Assert(len(pts)).Equal(0) //disjoint

			g.Assert(ln_cd.Intersects(pt_e)).IsTrue()
			g.Assert(ln_cd.Intersects(h)).IsTrue()        //at h
			g.Assert(ln_cd.Intersects(h_prime)).IsFalse() //disjoint
		})
	})
}
