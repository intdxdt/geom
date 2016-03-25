package linestring

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/geom/point"
    "github.com/intdxdt/simplex/util/math"
)

func TestLineStringEdit(t *testing.T) {
    g := Goblin(t)

    a := &Point{-2, -4}
    b := &Point{1, -1}
    c := &Point{-1, 4}

    pts := []*Point{a, b, c}
    d := &Point{5.6, 7.9}

    g.Describe("Linestring", func() {
        g.It("should test length on append", func() {
            ln := NewLineString(pts)
            g.Assert(math.Round(ln.length, 10)).Equal(math.Round(9.62780549425, 10))
            g.Assert(len(ln.chains)).Equal(2)

            ln.Append(d)
            g.Assert(math.Round(ln.Length(), 10)).Equal(math.Round(17.2939648978, 10))
            g.Assert(len(ln.chains)).Equal(3)

            ln.Pop()
            g.Assert(math.Round(ln.length, 10)).Equal(math.Round(9.62780549425, 10))
            g.Assert(len(ln.chains)).Equal(2)
        })
        g.It("should test intersection", func() {
             c := &Point{1.5, -2}
             d := &Point{-1.5, 2}
             h_prime := &Point{0.4841546875717521, -0.6455395491824757}

             h := &Point{0.484154648492778, -0.645539531323704}
             i := &Point{0.925118053504632, -1.233490738006176}

            ln_cd := NewLineString([]*Point{c, d})
            ln_hi := NewLineString([]*Point{h, i})

            ok  := ln_cd.Intersects(ln_hi)
            g.Assert(ok).Equal(true) //at h, i

            pts  := ln_cd.Intersection(ln_hi)
            g.Assert(len(pts)).Equal(2) //at h, i

            g.Assert(ln_cd.IntersectsPoint(h)).Equal(true) //at h
            g.Assert(ln_cd.IntersectsPoint(h_prime)).Equal(false) //disjoint

        })
    })
}
