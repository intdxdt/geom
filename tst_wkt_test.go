package geom

import (
	"testing"
	"github.com/franela/goblin"
	"time"
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
			g.Timeout(1 * time.Hour)
			obj := readWKT(pt, GeoTypePoint)
			g.Assert(obj.gtype).Eql(GeoTypePoint)
			g.Assert(obj.GeometryType()).Eql(GeoTypePoint)
			g.Assert(obj.Shell() == nil).IsFalse()
			g.Assert(len(obj.Shell())).Eql(1)
			g.Assert((obj.shell)[0]).Eql([]float64{30, 10})

			obj = readWKT(ept, GeoTypePoint)
			g.Assert(obj.gtype).Eql(GeoTypePoint)
			g.Assert(obj.GeometryType()).Eql(GeoTypePoint)
			g.Assert(obj.Shell() == nil).IsTrue()
			g.Assert(obj.Holes() == nil).IsTrue()

			obj = readWKT(cpoly, GeoTypePolygon)
			g.Assert(obj.gtype).Eql(GeoTypePolygon)
			g.Assert(obj.GeometryType()).Eql(GeoTypePolygon)
			g.Assert(obj.Shell() == nil).Eql(false)
			g.Assert(len(obj.Shell())).Eql(5)
			g.Assert(len(obj.Holes())).Eql(1)
			g.Assert(len(obj.Holes()[0])).Eql(4)

			obj = readWKT(poly, GeoTypePolygon)
			g.Assert(obj.gtype).Eql(GeoTypePolygon)
			g.Assert(obj.GeometryType()).Eql(GeoTypePolygon)
			g.Assert(obj.shell == nil).Eql(false)
			g.Assert(len(obj.shell)).Eql(5)
			g.Assert(obj.holes == nil).Eql(false)
			g.Assert(len(obj.holes)).Eql(0)

			obj = readWKT(epoly, GeoTypePolygon)
			g.Assert(obj.gtype).Eql(GeoTypePolygon)
			g.Assert(obj.GeometryType()).Eql(GeoTypePolygon)
			g.Assert(obj.shell == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)

			obj = readWKT(ln, GeoTypeLineString)
			g.Assert(obj.gtype).Eql(GeoTypeLineString)
			g.Assert(obj.GeometryType()).Eql(GeoTypeLineString)
			g.Assert(obj.shell == nil).Eql(false)
			g.Assert(len(obj.shell)).Eql(3)
			g.Assert(obj.holes == nil).Eql(true)

			obj = readWKT(eln, GeoTypeLineString)
			g.Assert(obj.gtype).Eql(GeoTypeLineString)
			g.Assert(obj.GeometryType()).Eql(GeoTypeLineString)
			g.Assert(obj.shell == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)

			var unknownLn = "unknown empty"
			obj = readWKT(unknownLn, GeoTypeUnknown)
			g.Assert(obj.gtype).Eql(GeoTypeUnknown)
			g.Assert(obj.GeometryType()).Eql(GeoTypeUnknown)
			g.Assert(obj.shell == nil).Eql(true)
			g.Assert(obj.holes == nil).Eql(true)

			var gtype = wktType("polygon empty")
			g.Assert(string(gtype) == "polygon").IsTrue()
			gtype = wktType(unknownLn)
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
			readWKT(tln, GeoTypeLineString)
		})

	})

	g.Describe("WKT Write", func() {
		var sh = [][]float64{{35, 10}, {45, 45}, {15, 40}, {10, 20}, {35, 10}}
		var h1 = [][]float64{{20, 30}, {35, 35}, {30, 20}, {20, 30}}
		var wkt_sh = "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10))"
		g.It("tests wkt writer", func() {
			g.Timeout(1 * time.Hour)
			g.Assert(WriteWKT(readWKT(pt, GeoTypePoint))).Eql("POINT (30 10)")
			ept := readWKT(ept, GeoTypePoint)

			g.Assert(WriteWKT(ept)).Eql("POINT EMPTY")

			g.Assert(WriteWKT(readWKT(ln, GeoTypeLineString))).Eql("LINESTRING (30 10, 10 30, 40 40)")
			g.Assert(WriteWKT(readWKT(eln, GeoTypeLineString))).Eql("LINESTRING EMPTY")

			g.Assert(WriteWKT(readWKT(poly, GeoTypePolygon))).Eql(poly)
			g.Assert(WriteWKT(readWKT(cpoly, GeoTypePolygon))).Eql(cpoly)
			g.Assert(WriteWKT(NewWKTParserObj(GeoTypePolygon, sh))).Eql(wkt_sh)
			g.Assert(WriteWKT(NewWKTParserObj(GeoTypePolygon, sh, h1))).Eql(cpoly)
			g.Assert(WriteWKT(readWKT(epoly, GeoTypePolygon))).Eql(epoly)
		})
	})

	g.Describe("WKT ToArray", func() {
		var ln = "LINESTRING (2.28 3.7, 2.98 5.36, 3.92 4.8, 3.9 3.64, 2.28 3.7)"
		var sh = [][]float64{{35, 10}, {45, 45}, {15, 40}, {10, 20}, {35, 10}}
		var h1 = [][]float64{{20, 30}, {35, 35}, {30, 20}, {20, 30}}
		var poly_array = [][][]float64{sh, h1}
		var ln_array = [][]float64{{2.28, 3.7}, {2.98, 5.36}, {3.92, 4.8}, {3.9, 3.64}, {2.28, 3.7}}

		g.It("tests wkt to array", func() {
			ln_obj := readWKT(ln, GeoTypeLineString)
			poly_obj := readWKT(cpoly, GeoTypePolygon)
			g.Assert(ln_obj.ToArray()[0]).Eql(ln_array)
			g.Assert(poly_obj.ToArray()).Eql(poly_array)
		})
	})
}
