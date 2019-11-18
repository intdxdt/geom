package geom

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/math"
	"github.com/intdxdt/sset"
	"testing"
	"time"
)

func createHulls(indxs [][]int, coords Coords) []Geometry {
	hulls := make([]Geometry, 0, len(indxs))
	for _, o := range indxs {
		hulls = append(hulls, hullGeom(subCoordinates(coords, o[0], o[1])))
	}
	return hulls
}

//generates sub polyline from generator indices
func subCoordinates(coordinates Coords, i, j int) Coords {
	coordinates.Idxs = make([]int, 0, j-i+1)
	for v := i; v <= j; v++ {
		coordinates.Idxs = append(coordinates.Idxs, v)
	}
	return coordinates
}

//hull geom
func hullGeom(coords Coords) Geometry {
	var g Geometry
	if coords.Len() > 2 {
		g = NewPolygon(coords)
	} else if coords.Len() == 2 {
		g = NewLineString(coords)
	} else {
		g = coords.Pt(0)
	}
	return g
}

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

		g.It("chull intersects", func() {
			g.Timeout(1 * time.Hour)
			var wkt = "LINESTRING ( 168.27827643756547 144.97549489763233, 167.9558767112785 146.58749352906722, 169.3529421918554 147.0173598307832, 169.67534191814238 147.76962585878613, 169.3529421918554 148.1994921605021, 168.17080986213648 147.9845590096441, 167.8484101358495 149.48909106565, 169.67534191814238 150, 170.7500076724323 152.28322202680383, 170.7500076724323 153.03548805480676, 170.64254109700332 153.8952206582387, 172.2545397284382 157.0117513456795, 177.62786849988782 157.97895052454044, 180.7443991873286 159.6984157314043, 180.7443991873286 163.88961217313502, 182.14146466790552 164.53441162570897, 184.72066247820135 161.09548121198122, 186.33266110963623 156.47441846853454, 186.5475942604942 154.3250869599547, 187.72972659021312 153.03548805480676, 189.9191578972495 154.30638596574983, 192 155, 192.13585618280183 158.08641709996942, 192.5657224845178 159.26854942968833, 193.74785481423672 161.52534751369717, 198.04651783139641 161.84774723998416, 199.9809161891183 160.9880146365522, 203.52731317827505 161.84774723998416, 204.92437865885194 165.39414422914092, 204.70944550799396 167.65094231314976, 204.06464605542 170.76747300059054, 205.5691781114259 172.16453848116745, 208.3633090725797 171.19733930230652, 210.19024085487257 167.86587546400776, 211.90970606173647 166.25387683257284, 213.19930496688437 165.71654395542788, 215.24116990003523 166.03894368171487, 217 168, 215.86524749172676 169.13475250827324, 216.42330222975414 171.6272056040225, 216.20836907889617 174.4213365651763, 216.9606351068991 177.4304006771881, 214.4889038720323 180.1170650629129, 215.9934359280382 183.98586177835662, 216.53076880518316 184.95306095721756, 220.2920989451979 185.5978604097915, 222.22649730291977 185.49039383436252, 227.5998260743694 183.0186625994957, 229.21182470580428 179.25733245948095, 231.8984890915291 174.63626971603426 )"
			var coords = NewLineStringFromWKT(wkt).Coordinates
			var coordsA, coordsB, coordsC = coords, coords, coords
			coordsA.Idxs, coordsB.Idxs, coordsC.Idxs = []int{}, []int{}, []int{}
			for _, idx := range coords.Idxs {
				if idx < 30 {
					coordsA.Idxs = append(coordsA.Idxs, idx)
				}
				if idx >= 29 {
					coordsB.Idxs = append(coordsB.Idxs, idx)
				}
				if idx >= 27 {
					coordsC.Idxs = append(coordsC.Idxs, idx)
				}
			}
			var cvxA = ConvexHull(coordsA)
			var cvxB = ConvexHull(coordsB)
			var cvxC = ConvexHull(coordsC)

			var plyA = NewPolygon(cvxA)
			var plyB = NewPolygon(cvxB)
			var plyC = NewPolygon(cvxC)

			var inters = plyA.Intersection(plyB)
			g.Assert(len(inters)).Equal(1)
			g.Assert(inters[0].Equals2D(&coords.Pnts[29])).IsTrue()
			inters = plyA.Intersection(plyC)
			g.Assert(len(inters)).Equal(4)

			wkt = "LINESTRING ( 960 840, 980 840, 980 880, 1020 900, 1080 880, 1120 860, 1160 800, 1160 760, 1140 700, 1080 700, 1040 720, 1060 760, 1120 800, 1080 840, 1020 820, 940 760 )"
			coords = NewLineStringFromWKT(wkt).Coordinates
			var hulls = createHulls([][]int{{0, 2}, {2, 6}, {6, 8}, {8, 10}, {10, 12}, {12, coords.Len() - 1}}, coords)
			inters = hulls[0].Intersection(hulls[1])
			g.Assert(len(inters)).Equal(1)

		})
	})

}
