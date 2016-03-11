package linestring

import (
    "testing"
    . "github.com/franela/goblin"
    pt "github.com/intdxdt/simplex/geom/point"
    "github.com/intdxdt/simplex/util/math"
)

func TestLineStringEdit(t *testing.T) {
    g := Goblin(t)

    a := &pt.Point{-2, -4}
    b := &pt.Point{1, -1}
    c := &pt.Point{-1, 4}

    pts := []*pt.Point{a, b, c}
    d := &pt.Point{5.6, 7.9}

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
    })
}
