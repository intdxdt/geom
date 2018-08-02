package geom

import (
	"testing"
	"github.com/intdxdt/math"
	"github.com/franela/goblin"
	"github.com/intdxdt/sset"
)

func TestCHull(t *testing.T) {
	g := goblin.Goblin(t)

	var empty_hull = Coordinates([]Point{})
	var hullEql = func(g *goblin.G, hull, expects Coords) {
		hs := sset.NewSSet(ptCmp)
		g.Assert(hull.Len()).Equal(expects.Len())
		for i := range hull.Idxs {
			hs.Add(*hull.Pt(i))
		}
		for i := range expects.Idxs {
			g.Assert(hs.Contains(*expects.Pt(i))).IsTrue()
		}
	}

	var data []Point
	for i := 0; i < 100; i++ {
		data = append(data, PointXY(math.Floor(float64(i)/10.0), float64(i%10)))
	}

	g.Describe("convex & simple hull", func() {
		var sqr = []Point{
			{33.52991674117594, 27.137460594059416},
			{33.52991674117594, 30.589750223527805},
			{36.44941148514852, 30.589750223527805},
			{36.44941148514852, 27.137460594059416},
			{33.52991674117594, 27.137460594059416},
		}

		g.It("should test convex hull", func() {
			var hull = ConvexHull(Coordinates(data))
			var ply = NewPolygon(Coordinates(data))
			var ln = NewLineString(Coordinates(data))
			var ch = []Point{{0, 0}, {9, 0}, {9, 9}, {0, 9}}
			var ch_array = [][]float64{{0, 0, 0}, {9, 0, 0}, {9, 9, 0}, {0, 9, 0}, {0, 0, 0}}

			g.Assert(ch).Eql(hull.Points())
			g.Assert(ply.ConvexHull().Shell.ToArray()).Eql(ch_array)
			g.Assert(ln.ConvexHull().Shell.ToArray()).Eql(ch_array)

			var pt = Coordinates([]Point{{33.52991674117594, 27.137460594059416}})
			g.Assert(ConvexHull(pt).Len()).Equal(1)
		})

		g.It("should test convex hull of sqr as sqr", func() {
			var hull = ConvexHull(Coordinates(sqr))
			var hpoly = NewPolygon(hull)
			var sqrpoly = NewPolygon(Coordinates(sqr))

			g.Assert(hull.Len()).Equal(len(sqr) - 1)
			g.Assert(hpoly.Area()).Equal(sqrpoly.Area())
		})

		g.It("chull contruction - empty, one, two, three", func() {
			hullEql(g, ConvexHull(Coordinates([]Point{})), empty_hull)
			hullEql(g, ConvexHull(Coordinates([]Point{{200, 200}})), Coordinates([]Point{{200, 200}}))
			hullEql(g, ConvexHull(Coordinates([]Point{{200, 200}, {760, 300}})), Coordinates([]Point{{200, 200}, {760, 300}}))
			var ch = ConvexHull(Coordinates([]Point{{200, 200}, {760, 300}, {500, 500}}))
			var exp = Coordinates([]Point{{760, 300}, {200, 200}, {500, 500}})
			hullEql(g, ch, exp)
		})

		g.It("chull for four points", func() {
			var ch = ConvexHull(Coordinates([]Point{{200, 200}, {760, 300}, {500, 500}, {400, 400}}))
			var exp = Coordinates([]Point{{760, 300}, {200, 200}, {500, 500}})
			hullEql(g, ch, exp)
		})
		g.It("chull returns a polygon", func() {
			var coords = Coordinates([]Point{{200, 200}, {760, 300}, {500, 500}, {400, 400}})
			var ply = NewPolygon(coords)
			var hull = NewPolygon(ConvexHull(coords))
			g.Assert(hull.Area() > 0).IsTrue()
			g.Assert(ply.Area() == hull.Area()).IsTrue()
		})

		g.It("handles points with duplicate ordinates", func() {
			var ch = ConvexHull(Coordinates([]Point{{-10, -10}, {10, 10}, {10, -10}, {-10, 10}}))
			var exp = Coordinates([]Point{{10, 10}, {10, -10}, {-10, -10}, {-10, 10}})
			hullEql(g, ch, exp)
		})

		g.It("handles overlapping upper and lower hulls", func() {
			var ch = ConvexHull(Coordinates([]Point{{0, -10}, {0, 10}, {0, 0}, {10, 0}, {-10, 0}}))
			var exp = Coordinates([]Point{{10, 0}, {0, -10}, {-10, 0}, {0, 10}})
			hullEql(g, ch, exp)
		})
		// Cases below taken from http://uva.onlinejudge.org/external/6/681.html
		g.It("computes chull for  a set of 6 points with non-trivial hull", func() {
			var poly = Coordinates([]Point{{60, 20}, {250, 140}, {180, 170}, {79, 140}, {50, 60}, {60, 20}})
			ch := ConvexHull(poly)
			var exp = Coordinates([]Point{{250, 140}, {60, 20}, {50, 60}, {79, 140}, {180, 170}})
			hullEql(g, ch, exp)
		})

		g.It("chull for  a set of 12 points with non-trivial hull", func() {
			var poly = []Point{{50, 60}, {60, 20}, {70, 45}, {100, 70},
				{125, 90}, {200, 113}, {250, 140}, {180, 170}, {105, 140},
				{79, 140}, {60, 85}, {50, 60}}
			var expectedHull = Coordinates([]Point{{250, 140}, {60, 20}, {50, 60}, {79, 140}, {180, 170}})
			var ch = ConvexHull(Coordinates(poly))
			hullEql(g, ch, expectedHull)
		})

		g.It("chull for a set of 15 points with non-trivial hull", func() {
			var poly = []Point{{30, 30}, {50, 60}, {60, 20}, {70, 45}, {86, 39},
				{112, 60}, {200, 113}, {250, 50}, {300, 200}, {130, 240}, {76, 150},
				{47, 76}, {36, 40}, {33, 35}, {30, 30}}
			var expectedHull = Coordinates([]Point{{300, 200}, {250, 50}, {60, 20}, {30, 30}, {47, 76}, {76, 150}, {130, 240}})
			hullEql(g, ConvexHull(Coordinates(poly)), expectedHull)
		})
	})

}
