package linestring

import (
    "testing"
    . "github.com/franela/goblin"
    pt "github.com/intdxdt/simplex/geom/point"
)

func TestLineString(t *testing.T) {
    g := Goblin(t)

    pts := []pt.Point{
        {5.6, 7.9}, {5.6, 8.9}, {6.6, 8.9},
        {6.6, 7.9}, {5.6, 7.9},
    }
    pts_1   := []pt.Point{{5.6, 7.9}}
    ln      := New(pts)
    cln     := ln.Clone()
    ln_1    := New(pts_1)

    g.Describe("Linestring", func() {
        g.It("should test length", func() {
            g.Assert(ln.Length() == 4.0).IsTrue()
            g.Assert(cln.Length() == 4.0).IsTrue()
            g.Assert(ln_1.Length() == 0.0).IsTrue()
        })

        g.It("should throw if empty coordinates", func(done Done) {
            defer func() {
                r := recover()
                if r != nil {
                    g.Assert(r != nil).Equal(true)
                } else {
                    g.Fail("did not throw")
                }
                done()
            }()
            pts := make([]pt.Point, 0)
            New(pts)
        })

        g.It("should be array of points", func() {
            ln.build_index()
            g.Assert(ln.ToArray()).Eql(pts)
            g.Assert(cln.ToArray()).Eql(pts)
            ln.build_index()
            g.Assert(ln.ToArray()).Eql(pts)
            g.Assert(ln_1.ToArray()).Eql([]pt.Point{pts_1[0], pts_1[0]})
        })

    })
}
