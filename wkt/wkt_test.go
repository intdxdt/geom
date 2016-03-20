package wkt

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/geom/point"
)

func TestWKT(t *testing.T) {
    g := Goblin(t)

    var pt = " \n\rPOINT (30 10)\n\r "
    var ept = " \n\rPOINT EMPTY\n\r "
    var ln = " \n\rLINESTRING (30 10, 10 30, 40 40)\n\r "
    var tln = " \n\rLINESTRING (30 1$0.$, 10 v, 40 40)\n\r "
    var eln = "LINESTRING EMPTY"

    var poly = "POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10))"
    var cpoly = "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))"
    var epoly = "POLYGON EMPTY"

    g.Describe("WKT Read", func() {
        g.It("test wkt parser", func() {
            obj := Read(pt)
            g.Assert(obj.gtype).Eql("point")
            g.Assert(obj.shell == nil).Eql(false)
            g.Assert(len(*obj.shell)).Eql(1)
            g.Assert((*obj.shell)[0]).Eql(&Point{30, 10})

            obj = Read(ept)
            g.Assert(obj.gtype).Eql("point")
            g.Assert(obj.shell == nil).Eql(true)
            g.Assert(obj.holes == nil).Eql(true)

            obj = Read(cpoly)
            g.Assert(obj.gtype).Eql("polygon")
            g.Assert(obj.shell == nil).Eql(false)
            g.Assert(len(*obj.shell)).Eql(5)
            g.Assert(len(*obj.holes)).Eql(1)
            g.Assert(len(*((*obj.holes)[0]))).Eql(4)

            obj = Read(poly)
            g.Assert(obj.gtype).Eql("polygon")
            g.Assert(obj.shell == nil).Eql(false)
            g.Assert(len(*obj.shell)).Eql(5)
            g.Assert(obj.holes == nil).Eql(false)
            g.Assert(len(*obj.holes)).Eql(0)

            obj = Read(epoly)
            g.Assert(obj.gtype).Eql("polygon")
            g.Assert(obj.shell == nil).Eql(true)
            g.Assert(obj.holes == nil).Eql(true)

            obj = Read(ln)
            g.Assert(obj.gtype).Eql("linestring")
            g.Assert(obj.shell == nil).Eql(false)
            g.Assert(len(*obj.shell)).Eql(3)
            g.Assert(obj.holes == nil).Eql(true)

            obj = Read(eln)
            g.Assert(obj.gtype).Eql("linestring")
            g.Assert(obj.shell == nil).Eql(true)
            g.Assert(obj.holes == nil).Eql(true)
        })

        g.It("should throw", func(done Done) {
            defer func() {
                r := recover()
                if r != nil {
                    g.Assert(r != nil).Equal(true)
                } else {
                    g.Fail("did not throw")
                }
                done()
            }()
            Read(tln)
        })

    })

    g.Describe("WKT Write", func() {
        g.It("tests wkt writer", func() {
            g.Assert(Write(Read(pt))).Eql("POINT (30 10)")
            g.Assert(Write(Read(ept))).Eql("POINT EMPTY")
            g.Assert(Write(Read(ln))).Eql("LINESTRING (30 10, 10 30, 40 40)")
            g.Assert(Write(Read(ln))).Eql("LINESTRING (30 10, 10 30, 40 40)")
            g.Assert(Write(Read(eln))).Eql("LINESTRING EMPTY")
            g.Assert(Write(Read(poly))).Eql("POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10))")
            g.Assert(Write(Read(epoly))).Eql("POLYGON EMPTY")
            g.Assert(Write(Read(cpoly))).Eql("POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))")
        })
    })
}

