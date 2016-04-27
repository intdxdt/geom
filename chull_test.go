package geom

import (
    "testing"
    . "github.com/franela/goblin"
    "math"
)

func TestCHull(t *testing.T) {
    g := Goblin(t)
    var data = make([]*Point, 0);
    for i := 0; i < 100; i++ {
        data = append(data, NewPointXY(math.Floor(float64(i) / 10.0), float64(i % 10)))
    }

    g.Describe("Convex Hull", func() {
        g.It("it should test convex hull", func() {
            var hull = ConvexHull(data)
            var ch = []*Point{{0, 0}, {9, 0}, {9, 9}, {0, 9}, {0, 0}}
            g.Assert(ch).Eql(hull)

            var pt = []*Point{ {33.52991674117594, 27.137460594059416},}
            g.Assert(len(ConvexHull(pt))).Equal(1)
        })
        var sqr = []*Point{
            {33.52991674117594, 27.137460594059416},
            {33.52991674117594, 30.589750223527805},
            {36.44941148514852, 30.589750223527805},
            {36.44941148514852, 27.137460594059416},
            {33.52991674117594, 27.137460594059416},
        }
        g.It("it should test convex hull of sqr as sqr", func() {
            var hull = ConvexHull(sqr)
            var shull = SimpleHull(sqr)
            var hpoly = NewPolygon(hull)
            var shpoly = NewPolygon(shull)
            var sqrpoly = NewPolygon(sqr)
            g.Assert(len(hull)).Equal(len(sqr))
            g.Assert(hpoly.Area()).Equal(sqrpoly.Area())
            g.Assert(shpoly.Area()).Equal(sqrpoly.Area())
        })
    })

}


