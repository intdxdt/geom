package linestring

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/geom/point"
    "github.com/intdxdt/simplex/util/math"
    "fmt"
)

func TestLineStringEdit(t *testing.T) {
    g := Goblin(t)

    a := &Point{-2, -4}
    b := &Point{1, -1}
    c := &Point{-1, 4}

    pts := []*Point{a, b, c}
    d := &Point{5.6, 7.9}
    var wkt = "LINESTRING ( 1.8550969696969695 4.406890909090909, 2.030969696969697 4.634490909090909, 2.1999454545454546 4.717254545454545, 2.2930545454545452 4.858642424242424, 2.499963636363636 4.920715151515152, 2.6482484848484846 5.096587878787879, 2.9862 5.179351515151516, 3.2008705964930875 5.313330045424207, 3.458420073920762 5.392514658801376, 3.61096174774788 5.153923835635884, 3.6539863224170674 4.794081938402683, 3.8378185960035944 4.590693039966525 )"

    g.Describe("Linestring", func() {
        g.It("should test length on append", func() {
            ln := NewLineString(pts)
            ln2 := NewLineStringFromWKT(wkt)

            g.Assert(
                math.Round(ln.length, 10)).Equal(
                math.Round(9.62780549425, 10),
            )
            g.Assert(len(ln.chains)).Equal(2)

            ln.Append(d)
            g.Assert(math.Round(ln.Length(), 10)).Equal(
                math.Round(17.2939648978, 10),
            )
            g.Assert(len(ln.chains)).Equal(3)

            ln.Pop()
            fmt.Println("number of coords:", len(ln.coordinates))
            g.Assert(math.Round(ln.length, 10)).Equal(
                math.Round(9.62780549425, 10),
            )
            g.Assert(len(ln.chains)).Equal(2)

            ln.Pop()
            fmt.Println("number of coords:", len(ln.coordinates))
            g.Assert(math.Round(ln.length, 10)).Equal(
                math.Round(4.242640687119285, 10),
            )
            g.Assert(len(ln.chains)).Equal(1)

            ln.Pop()
            g.Assert(math.Round(ln.length, 10)).Equal(
                math.Round(4.242640687119285, 10),
            )
            g.Assert(len(ln.chains)).Equal(1)
            //------------------------------------------
            g.Assert(len(ln2.chains)).Equal(4)
            ln2.Pop()
            g.Assert(len(ln2.chains)).Equal(4)
            ln2.Pop()
            g.Assert(len(ln2.chains)).Equal(4)
            g.Assert(len(ln2.MonoChains())).Equal(4)
            ln2.Pop()
            g.Assert(len(ln2.chains)).Equal(3)
            g.Assert(len(ln2.MonoChains())).Equal(3)
            ln2.Pop()
            g.Assert(len(ln2.chains)).Equal(3)

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
            g.Assert(v==nil).IsTrue()
        })
        g.It("should test intersection", func() {
            a := &Point{-4.975454545454546, 0.2551515151515151, }
            b := &Point{-3.9389015151515148, 1.156155303030303}
            c := &Point{1.5, -2}
            d := &Point{-1.5, 2}
            h_prime := &Point{0.4841546875717521, -0.6455395491824757}

            h := &Point{0.484154648492778, -0.645539531323704}
            i := &Point{0.925118053504632, -1.233490738006176}
            var ln_e *LineString
             var pt_e *Point
            ln_ab := NewLineString([]*Point{a, b})
            ln_cd := NewLineString([]*Point{c, d})
            fmt.Println(ln_cd)
            fmt.Println(ln_ab)

            ln_hi := NewLineString([]*Point{h, i})



            ok := ln_cd.Intersects(ln_ab)
            g.Assert(ok).IsFalse()

            ok = ln_cd.Intersects(ln_e)
            g.Assert(ok).IsFalse()

            ok = ln_cd.Intersects(ln_hi)
            g.Assert(ok).IsTrue() //at h, i

            pts := ln_cd.Intersection(ln_hi)
            g.Assert(len(pts)).Equal(2) //at h, i

            pts = ln_cd.Intersection(ln_ab)
            g.Assert(len(pts)).Equal(0) //disjoint

            g.Assert(ln_cd.IntersectsPoint(pt_e)).IsFalse()
            g.Assert(ln_cd.IntersectsPoint(h)).IsTrue() //at h
            g.Assert(ln_cd.IntersectsPoint(h_prime)).IsFalse() //disjoint

        })
    })
}
