package geom

import (
	"testing"
	"github.com/franela/goblin"
)

func TestWKT(t *testing.T) {
	g := goblin.Goblin(t)

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
			obj := ReadWKT(pt)
			g.Assert(obj.gtype).Eql(GeoType_Point)
			g.Assert(obj.GeometryType()).Eql(GeoType_Point)
			g.Assert(obj.Shell() == nil).IsFalse()
			g.Assert(len(*obj.Shell())).Eql(1)
			g.Assert((*obj.shell)[0]).Eql([2]float64{30, 10})

			obj = ReadWKT(ept)
			g.Assert(obj.gtype).Eql(GeoType_Point)
			g.Assert(obj.GeometryType()).Eql(GeoType_Point)
			g.Assert(obj.Shell() == nil).IsTrue()
			g.Assert(obj.Holes() == nil).IsTrue()

			obj = ReadWKT(cpoly)
			g.Assert(obj.gtype).Eql(GeoType_Polygon)
			g.Assert(obj.GeometryType()).Eql(GeoType_Polygon)
			g.Assert(obj.Shell() == nil).Eql(false)
			g.Assert(len(*obj.Shell())).Eql(5)
			g.Assert(len(*obj.Holes())).Eql(1)
			g.Assert(len(*((*obj.Holes())[0]))).Eql(4)

			obj = ReadWKT(poly)
			g.Assert(obj.gtype).Eql(GeoType_Polygon)
			g.Assert(obj.GeometryType()).Eql(GeoType_Polygon)
			g.Assert(obj.shell == nil).Eql(false)
			g.Assert(len(*obj.shell)).Eql(5)
			g.Assert(obj.holes == nil).Eql(false)
			g.Assert(len(*obj.holes)).Eql(0)

			obj = ReadWKT(epoly)
			g.Assert(obj.gtype).Eql(GeoType_Polygon)
			g.Assert(obj.GeometryType()).Eql(GeoType_Polygon)
			g.Assert(obj.shell == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)

			obj = ReadWKT(ln)
			g.Assert(obj.gtype).Eql(GeoType_LineString)
			g.Assert(obj.GeometryType()).Eql(GeoType_LineString)
			g.Assert(obj.shell == nil).Eql(false)
			g.Assert(len(*obj.shell)).Eql(3)
			g.Assert(obj.holes == nil).Eql(true)

			obj = ReadWKT(eln)
			g.Assert(obj.gtype).Eql(GeoType_LineString)
			g.Assert(obj.GeometryType()).Eql(GeoType_LineString)
			g.Assert(obj.shell == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)
		})

		g.It("should throw", func(done goblin.Done) {
			defer func() {
				r := recover()
				if r != nil {
					g.Assert(r != nil).Equal(true)
				} else {
					g.Fail("did not throw")
				}
				done()
			}()
			ReadWKT(tln)
		})

	})

	g.Describe("WKT Write", func() {

		sh := [][2]float64{{35, 10}, {45, 45}, {15, 40}, {10, 20}, {35, 10}}
		h1 := [][2]float64{{20, 30}, {35, 35}, {30, 20}, {20, 30}}
		wkt_sh := "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10))"

		g.It("tests wkt writer", func() {
			g.Assert(WriteWKT(ReadWKT(pt))).Eql("POINT (30 10)")
			ept := ReadWKT(ept)

			g.Assert(WriteWKT(ept)).Eql("POINT EMPTY")

			g.Assert(WriteWKT(ReadWKT(ln))).Eql("LINESTRING (30 10, 10 30, 40 40)")
			g.Assert(WriteWKT(ReadWKT(eln))).Eql("LINESTRING EMPTY")

			g.Assert(WriteWKT(ReadWKT(poly))).Eql(poly)
			g.Assert(WriteWKT(ReadWKT(cpoly))).Eql(cpoly)
			g.Assert(WriteWKT(NewWKTParserObj(GeoType_Polygon, sh))).Eql(wkt_sh)
			g.Assert(WriteWKT(NewWKTParserObj(GeoType_Polygon, sh, h1))).Eql(cpoly)
			g.Assert(WriteWKT(ReadWKT(epoly))).Eql(epoly)
		})
	})

	g.Describe("WKT ToArray", func() {
		ln := "LINESTRING (2.28 3.7, 2.98 5.36, 3.92 4.8, 3.9 3.64, 2.28 3.7)"
		sh := [][2]float64{{35, 10}, {45, 45}, {15, 40}, {10, 20}, {35, 10}}
		h1 := [][2]float64{{20, 30}, {35, 35}, {30, 20}, {20, 30}}
		poly_array := [][][2]float64{sh, h1}
		ln_array := [][2]float64{{2.28, 3.7}, {2.98, 5.36}, {3.92, 4.8}, {3.9, 3.64}, {2.28, 3.7}}

		g.It("tests wkt to array", func() {
			ln_obj := ReadWKT(ln)
			poly_obj := ReadWKT(cpoly)
			g.Assert(ln_obj.ToArray()[0]).Eql(ln_array)
			g.Assert(poly_obj.ToArray()).Eql(poly_array)
		})
	})
}
