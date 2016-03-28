package geom

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/geom/mbr"
)

func TestLineString(t *testing.T) {
    g := Goblin(t)
    pts := []*Point{
        {5.6, 7.9}, {5.6, 8.9}, {6.6, 8.9}, {6.6, 7.9}, {5.6, 7.9},
    }
    pt_array := [][2]float64{
        {5.6, 7.9}, {5.6, 8.9}, {6.6, 8.9}, {6.6, 7.9}, {5.6, 7.9},
    }

    ln := NewLineString(pts)
    cln := ln.Clone()
    pt_lnstr := NewLineStringFromPoint(pts[0])

    g.Describe("Linestring", func() {
        g.It("should test length", func() {
            g.Assert(ln.Length() == 4.0).IsTrue()
            g.Assert(pt_lnstr.Length() == 0.0).IsTrue()
            g.Assert(ln.IsRing()).IsTrue()
            g.Assert(ln.len(len(ln.coordinates) - 1, 0) == ln.Length()).IsTrue()
            g.Assert(ln.chain_length(ln.chains[0])).Equal(ln.chain_length(ln.chains[1]))
            g.Assert(ln.chain_length(ln.chains[2])).Equal(ln.chain_length(ln.chains[3]))
            g.Assert(cln.Length() == 4.0).IsTrue()
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
            pts := make([]*Point, 0)
            NewLineString(pts)
        })

        g.It("should be slice of array", func() {
            ln.build_index()
            g.Assert(ln.ToArray()).Eql(pt_array)
            g.Assert(cln.ToArray()).Eql(pt_array)
            ln.build_index()
            g.Assert(ln.ToArray()).Eql(pt_array)
        })

    })

    g.Describe("Linestring - Coords", func() {

        g.It("should be slice of points", func() {
            ln.build_index()
            g.Assert(ln.Coordinates()).Eql(pts)
            g.Assert(cln.Coordinates()).Eql(pts)
            ln.build_index()
            g.Assert(ln.Coordinates()).Eql(pts)
        })

        g.It("should test coords indexing", func() {
            g.Assert(ln.VertexAt(0)).Eql(pts[0])
            g.Assert(ln.VertexAt(ln.LenVertices() - 1)).Eql(pts[len(pts) - 1])
            g.Assert(ln.LenVertices()).Eql(len(pts))
        })

        g.It("should test envelope", func() {
            box := NewMBR(pts[0][x], pts[0][y], pts[0][x], pts[0][y], )
            for _, v := range pts[1:] {
                box.ExpandIncludeXY(v.X(), v.Y())
            }
            g.Assert(ln.Envelope()).Eql(box)
        })

    })

    g.Describe("Linestring - WKT", func() {
        g.It("should test wkt string", func() {
            lnstr := "LINESTRING (5.6 7.9, 5.6 8.9, 6.6 8.9, 6.6 7.9, 5.6 7.9)"
            g.Assert(ln.String()).Eql(lnstr)
        })
    })
}
