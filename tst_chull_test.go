package geom

import (
    "testing"
    "math"
    "simplex/struct/sset"
    . "github.com/franela/goblin"
)

func TestCHull(t *testing.T) {
    g := Goblin(t)

    var empty_hull = make([]*Point, 0)
    hullEql := func(g *G, hull, expects Coordinates) {
        hs := sset.NewSSet()
        g.Assert(hull.Len()).Equal(expects.Len())
        for _, pt := range hull {
            hs.Add(pt)
        }
        for _, pt := range expects {
            g.Assert(hs.Contains(pt)).IsTrue()
        }
    }

    var data = make([]*Point, 0);
    for i := 0; i < 100; i++ {
        data = append(data, NewPointXY(math.Floor(float64(i) / 10.0), float64(i % 10)))
    }

    g.Describe("convex & simple hull", func() {
        var sqr = []*Point {
            {33.52991674117594, 27.137460594059416},
            {33.52991674117594, 30.589750223527805},
            {36.44941148514852, 30.589750223527805},
            {36.44941148514852, 27.137460594059416},
            {33.52991674117594, 27.137460594059416},
        }

        g.It("should test convex hull", func() {
            var hull = ConvexHull(data)
            var ch = []*Point{{0, 0}, {9, 0}, {9, 9}, {0, 9}}
            g.Assert(ch).Eql(hull)
            var pt = []*Point{{33.52991674117594, 27.137460594059416}, }
            g.Assert(len(ConvexHull(pt))).Equal(0)
        })

        g.It("should test convex hull of sqr as sqr", func() {
            var hull    = ConvexHull(sqr)
            var shull   = SimpleHull(sqr)
            var hpoly   = NewPolygon(hull)
            var shpoly  = NewPolygon(shull)
            var sqrpoly = NewPolygon(sqr)

            g.Assert(len(hull)).Equal(len(sqr) - 1)
            g.Assert(hpoly.Area()).Equal(sqrpoly.Area())
            g.Assert(shpoly.Area()).Equal(sqrpoly.Area())
        })

        g.It("simple hull - init 0, 1, 2", func() {
            wktL := "POLYGON (( 409 189, 429 235, 395 289, 366.14493191082227 293.7217384145927, 340 298, 354.5470100598068 275.01994063016036, 381.63571099302817 269.17705887748156, 379.0628213878595 250.73801670710623, 405.6493473079356 240.4464582864316, 409 189 ))"
            wktR := "POLYGON (( 479 184, 504 231, 601 254, 638 223, 594.3279183536367 219.00571157669285, 572.887171643898 225.43793558961448, 552.4235912992449 207.74149982270123, 517.9988600669667 217.71926677410852, 506.4208568437078 198.8514096695384, 479 184 ))"

            var coords0 = NewPolygonFromWKT(wktL).Shell.Coordinates()
            var coords1 = NewPolygonFromWKT(wktR).Shell.Coordinates()

            ch0 := SimpleHull(coords0)
            ch1 := SimpleHull(coords1)

            exp0 := []*Point{{409, 189}, {429, 235}, {395, 289}, {340, 298}, {409, 189} }
            exp1 := []*Point{{479, 184}, {504, 231}, {601, 254}, {638, 223}, {479, 184} }

            hullEql(g, ch0, exp0)
            hullEql(g, ch1, exp1)
        })

        g.It("chull of zero point is empty", func() {
            hullEql(g, ConvexHull([]*Point{}), empty_hull, )
            hullEql(g, SimpleHull([]*Point{}), empty_hull, )
        })
        g.It("chull of one point is empty", func() {
            hullEql(g, ConvexHull([]*Point{{200, 200}}), empty_hull, )
            hullEql(g, SimpleHull([]*Point{{200, 200}}), empty_hull, )
        })
        g.It("chull of  two points is empty", func() {
            hullEql(g, ConvexHull([]*Point{{200, 200}, {760, 300}}), empty_hull, )
            hullEql(g, SimpleHull([]*Point{{200, 200}, {760, 300}}), empty_hull, )
        })
        g.It("chull for three points", func() {
            ch := ConvexHull([]*Point{{200, 200}, {760, 300}, {500, 500}})
            exp := []*Point{{760, 300}, {200, 200}, {500, 500}}
            hullEql(g, ch, exp)
        })
        g.It("chull for four points", func() {
            ch := ConvexHull([]*Point{{200, 200}, {760, 300}, {500, 500}, {400, 400}})
            exp := []*Point{{760, 300}, {200, 200}, {500, 500}}
            hullEql(g, ch, exp)
        })
        g.It("chull returns a polygon", func() {
            coords := []*Point{{200, 200}, {760, 300}, {500, 500}, {400, 400}}
            ply := NewPolygon(coords)
            hull := NewPolygon(ConvexHull(coords))
            g.Assert(hull.Area() > 0).IsTrue()
            g.Assert(ply.Area() == hull.Area()).IsTrue()
        })

        g.It("handles points with duplicate ordinates", func() {
            ch := ConvexHull([]*Point{{-10, -10}, {10, 10}, {10, -10}, {-10, 10}})
            exp := []*Point{{10, 10}, {10, -10}, {-10, -10}, {-10, 10}}
            hullEql(g, ch, exp)
        })

        g.It("handles overlapping upper and lower hulls", func() {
            ch := ConvexHull([]*Point{{0, -10}, {0, 10}, {0, 0}, {10, 0}, {-10, 0}})
            exp := []*Point{{10, 0}, {0, -10}, {-10, 0}, {0, 10}}
            hullEql(g, ch, exp)
        })
        // Cases below taken from http://uva.onlinejudge.org/external/6/681.html
        g.It("computes chull for  a set of 6 points with non-trivial hull", func() {
            var poly = []*Point{{60, 20}, {250, 140}, {180, 170}, {79, 140}, {50, 60}, {60, 20}};
            ch := ConvexHull(poly)
            var exp = []*Point{{250, 140}, {60, 20}, {50, 60}, {79, 140}, {180, 170}};
            hullEql(g, ch, exp)
        })

        g.It("chull for  a set of 12 points with non-trivial hull", func() {
            var poly = []*Point{{50, 60}, {60, 20}, {70, 45}, {100, 70},
                {125, 90}, {200, 113}, {250, 140}, {180, 170}, {105, 140},
                {79, 140}, {60, 85}, {50, 60}};
            var expectedHull = []*Point{{250, 140}, {60, 20}, {50, 60},
                {79, 140}, {180, 170}};
            ch := ConvexHull(poly);
            hullEql(g, ch, expectedHull)
        })


        g.It("chull for a set of 15 points with non-trivial hull", func() {
            var poly =  []*Point{{30,30}, {50,60}, {60,20}, {70,45}, {86,39},
                {112,60}, {200,113}, {250,50}, {300,200}, {130,240}, {76,150},
                {47,76}, {36,40}, {33,35}, {30,30}};
            var expectedHull =  []*Point{{300,200}, {250,50}, {60,20}, {30,30},
                {47,76}, {76,150}, {130,240}};
            hullEql(g, ConvexHull(poly), expectedHull);
        })
    })

}


