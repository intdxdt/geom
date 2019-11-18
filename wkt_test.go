package geom

import (
	"github.com/franela/goblin"
	"testing"
	"time"
)

func TestWKT(t *testing.T) {
	var g = goblin.Goblin(t)

	var pt = " \n\rPOINT (30 10 2.5)\n\r "
	var ept = " \n\rPOINT EMPTY\n\r "
	var ln = " \n\rLINESTRING (30 10 1, 10 30 2, 40 40 3)\n\r "
	var tln = " \n\rLINESTRING (30 1$0.$, 10 v, 40 40)\n\r "
	var eln = "LINESTRING EMPTY"

	var poly = "POLYGON ((30 10 1, 40 40 2, 20 40 3, 10 20 4, 30 10 5))"
	var cpoly = "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))"
	var epoly = "POLYGON EMPTY"

	g.Describe("WKT Read", func() {
		g.It("test wkt parser", func() {
			g.Timeout(1 * time.Hour)
			obj := ReadWKT(pt, GeoTypePoint)
			g.Assert(obj.gtype).Eql(GeoTypePoint)
			g.Assert(obj.GeometryType()).Eql(GeoTypePoint)
			g.Assert(obj.Shell().Pnts == nil).IsFalse()
			g.Assert(len(obj.Shell().Pnts)).Eql(1)
			g.Assert(obj.shell.Pnts[0][:]).Eql([]float64{30, 10, 2.5})

			obj = ReadWKT(ept, GeoTypePoint)
			g.Assert(obj.gtype).Eql(GeoTypePoint)
			g.Assert(obj.GeometryType()).Eql(GeoTypePoint)
			g.Assert(obj.Shell().Pnts == nil).IsTrue()
			g.Assert(obj.Holes() == nil).IsTrue()

			obj = ReadWKT(cpoly, GeoTypePolygon)
			g.Assert(obj.gtype).Eql(GeoTypePolygon)
			g.Assert(obj.GeometryType()).Eql(GeoTypePolygon)
			g.Assert(obj.Shell().Pnts == nil).Eql(false)
			g.Assert(len(obj.Shell().Pnts)).Eql(5)
			g.Assert(len(obj.Holes())).Eql(1)
			g.Assert(len(obj.Holes()[0].Pnts)).Eql(4)

			obj = ReadWKT(poly, GeoTypePolygon)
			g.Assert(obj.gtype).Eql(GeoTypePolygon)
			g.Assert(obj.GeometryType()).Eql(GeoTypePolygon)
			g.Assert(obj.shell.Pnts == nil).Eql(false)
			g.Assert(len(obj.shell.Pnts)).Eql(5)
			g.Assert(obj.holes == nil).Eql(false)
			g.Assert(len(obj.holes)).Eql(0)

			obj = ReadWKT(epoly, GeoTypePolygon)
			g.Assert(obj.gtype).Eql(GeoTypePolygon)
			g.Assert(obj.GeometryType()).Eql(GeoTypePolygon)
			g.Assert(obj.shell.Pnts == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)

			obj = ReadWKT(ln, GeoTypeLineString)
			g.Assert(obj.gtype).Eql(GeoTypeLineString)
			g.Assert(obj.GeometryType()).Eql(GeoTypeLineString)
			g.Assert(obj.shell.Pnts == nil).Eql(false)
			g.Assert(len(obj.shell.Pnts)).Eql(3)
			g.Assert(obj.holes == nil).Eql(true)

			obj = ReadWKT(eln, GeoTypeLineString)
			g.Assert(obj.gtype).Eql(GeoTypeLineString)
			g.Assert(obj.GeometryType()).Eql(GeoTypeLineString)
			g.Assert(obj.shell.Pnts == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)

			var unknownLn = "unknown empty"
			obj = ReadWKT(unknownLn, GeoTypeUnknown)
			g.Assert(obj.gtype).Eql(GeoTypeUnknown)
			g.Assert(obj.GeometryType()).Eql(GeoTypeUnknown)
			g.Assert(obj.shell.Pnts == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)

			var notImplemented = "MultiPoint ((3 4), (5 6))"
			obj = ReadWKT(notImplemented, GeoTypeUnknown)
			g.Assert(obj.gtype).Eql(GeoTypeUnknown)
			g.Assert(obj.GeometryType()).Eql(GeoTypeUnknown)
			g.Assert(obj.shell.Pnts == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)

			var gtype = WKTType("polygon empty")
			g.Assert(string(gtype) == "polygon").IsTrue()
			gtype = WKTType(unknownLn)
			g.Assert(string(gtype) == "unknown").IsTrue()
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
			ReadWKT(tln, GeoTypeLineString)
		})

	})

	g.Describe("WKT Write", func() {
		var sh = AsCoordinates([][]float64{{35, 10}, {45, 45}, {15, 40}, {10, 20}, {35, 10}})
		var h1 = AsCoordinates([][]float64{{20, 30}, {35, 35}, {30, 20}, {20, 30}})
		var wkt_sh = "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10))"
		g.It("tests wkt writer", func() {
			g.Timeout(1 * time.Hour)
			g.Assert(WriteWKT(ReadWKT(pt, GeoTypePoint))).Eql("POINT (30 10)")
			g.Assert(WriteWKT3D(ReadWKT(pt, GeoTypePoint))).Eql("POINT (30 10 2.5)")
			ept := ReadWKT(ept, GeoTypePoint)

			g.Assert(WriteWKT(ept)).Eql("POINT EMPTY")

			g.Assert(WriteWKT(ReadWKT(ln, GeoTypeLineString))).Eql("LINESTRING (30 10, 10 30, 40 40)")
			g.Assert(WriteWKT3D(ReadWKT(ln, GeoTypeLineString))).Eql("LINESTRING (30 10 1, 10 30 2, 40 40 3)")
			g.Assert(WriteWKT(ReadWKT(eln, GeoTypeLineString))).Eql("LINESTRING EMPTY")

			g.Assert(WriteWKT(ReadWKT(poly, GeoTypePolygon))).Eql("POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10))")
			g.Assert(WriteWKT3D(ReadWKT(poly, GeoTypePolygon))).Eql(poly)
			g.Assert(WriteWKT(ReadWKT(cpoly, GeoTypePolygon))).Eql(cpoly)
			g.Assert(WriteWKT(NewWKTParserObj(GeoTypePolygon, sh))).Eql(wkt_sh)
			g.Assert(WriteWKT(NewWKTParserObj(GeoTypePolygon, sh, h1))).Eql(cpoly)
			g.Assert(WriteWKT(ReadWKT(epoly, GeoTypePolygon))).Eql(epoly)
			g.Assert(WriteWKT3D(ReadWKT(epoly, GeoTypePolygon))).Eql(epoly)
		})
	})

	g.Describe("WKT ToCoordinates", func() {
		var ln = "LINESTRING (2.28 3.7, 2.98 5.36, 3.92 4.8, 3.9 3.64, 2.28 3.7)"

		var sh = AsCoordinates([][]float64{{35, 10}, {45, 45}, {15, 40}, {10, 20}, {35, 10}})
		var h1 = AsCoordinates([][]float64{{20, 30}, {35, 35}, {30, 20}, {20, 30}})
		var poly_array = []Coords{sh, h1}
		var ln_array = Coordinates([]Point{{2.28, 3.7}, {2.98, 5.36}, {3.92, 4.8}, {3.9, 3.64}, {2.28, 3.7}})

		g.It("tests wkt to array", func() {
			ln_obj := ReadWKT(ln, GeoTypeLineString)
			poly_obj := ReadWKT(cpoly, GeoTypePolygon)
			g.Assert(ln_obj.ToCoordinates()[0]).Eql(ln_array)
			g.Assert(poly_obj.ToCoordinates()).Eql(poly_array)
		})
	})

	g.Describe("WKT utils", func() {
		g.It("tests wkt to array", func() {
			g.Timeout(1 * time.Hour)
			var tokens []*wktToken
			var v = popToken(&tokens)
			g.Assert(v == nil).IsTrue()
			g.Assert(dimension([]byte(" ")) == -1).IsTrue()
			g.Assert(dimension([]byte(" 3.142 ")) == -1).IsTrue()
			g.Assert(dimension([]byte("3.78   4.17 ")) == 2).IsTrue()
			g.Assert(dimension([]byte("   3.78    4.17   ")) == 2).IsTrue()
			g.Assert(dimension([]byte("   3.78    4.17   ,    3.78    4.17   ")) == 2).IsTrue()
			g.Assert(dimension([]byte(" 3.142 4.45 5.78 ")) == 3).IsTrue()
			g.Assert(dimension([]byte("3.36 4.78 5.67 , 1.12 1.34 2.47")) == 3).IsTrue()
			g.Assert(dimension([]byte("3.112 4.27 3.35, 5.12 6.14 2.57")) == 3).IsTrue()
			g.Assert(dimension([]byte(" 3.78 4.17   ,  3.18 4.11 ")) == 2).IsTrue()
		})
	})
}
