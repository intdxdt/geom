package geom

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/util/math"
    "fmt"
)

func TestLinearRing(t *testing.T) {
    g := Goblin(t)
    pts_closed := []*Point{
        {5.6, 7.9}, {5.6, 8.9}, {6.6, 8.9}, {6.6, 7.9}, {5.6, 7.9},
    }

    pts_open := []*Point{
        {5.6, 7.9}, {5.6, 8.9}, {6.6, 8.9}, {6.6, 7.9},
    }

    wkt := "POLYGON ((2.9326 4.3639, 3.1457 4.6164, 3.3706 4.5257, 3.5245 4.2849, 3.2167 4.1547, 3.0944 4.0718, 2.9918 4.1902, 2.9326 4.3639))"
    wkt0 := "POLYGON (( 3.705786620853205 4.927616481193835, 4.025775008326719 5.030841563599844, 4.183630587978802 5.015056005634635, 4.187576977470104 4.81379014157823, 4.136273914083177 4.6756665093826575, 3.9586863869745836 4.632256224978335, 3.87581220765724 4.770379857173907, 3.7376885754616676 4.774326246665209, 3.705786620853205 4.927616481193835 ))"
    wkt1 := "LINESTRING ( 4.601947874056822 5.334713554430103, 4.519073694739478 5.007163226652032, 4.8308384645523414 4.849307646999948, 5.142603234365206 4.963752942247709, 5.241262971647757 4.545435656169689 )"
    wkt2 := "LINESTRING ( 4.234933651365729 4.09160086466995, 3.765313301900782 4.0166194643352116, 3.694278291057345 3.8548174951918264, 3.852133870709428 3.4522857670790152 )"

    var ply = NewPolygonFromWKT(wkt)
    var ply0 = NewPolygonFromWKT(wkt0)

    var ln1 = NewLinearRing(pts_closed)
    var ln2 = NewLinearRing(pts_open)
    var ln3 = NewLineStringFromWKT(wkt1)
    var ln4 = NewLineStringFromWKT(wkt2)

    rng0 := NewLinearRing([]*Point{{2.28, 3.7}, {2.98, 5.36}, {3.92, 4.8}, {3.9, 3.64}, {2.28, 3.7}})

    fmt.Println(rng0.String())

    //points in relation to ln1
    pt0 := &Point{2.42747717129387, 4.4873795209295695} //outside
    pt1 := &Point{3.92, 4.8} //POINT(3.92 4.8) at vertex
    pt2 := &Point{2.724860976740761, 4.754956030556663, } //online
    pt3 := &Point{2.908711515786465, 4.440556719843803 } //inside

    g.Describe("LinearRing", func() {
        g.It("should test length", func() {
            g.Assert(ln1.Length() == 4.0).IsTrue()
            g.Assert(ln2.Length() == 4.0).IsTrue()
        })
    })

    g.Describe("LinearRing - Relation to Point", func() {
        g.It("should test point in ring & intersect", func() {
            g.Assert(ln1.PointCompletelyInRing(pt0)).Equal(false)
            g.Assert(ln1.PointCompletelyInRing(pt1)).Equal(false)
            g.Assert(ln1.PointCompletelyInRing(pt2)).Equal(false)
            g.Assert(ln1.PointCompletelyInRing(pt3)).Equal(false)

            g.Assert(rng0.PointCompletelyInRing(pt0)).Equal(false)
            g.Assert(rng0.PointCompletelyInRing(pt1)).Equal(false)
            g.Assert(rng0.PointCompletelyInRing(pt2)).Equal(false)
            g.Assert(rng0.PointCompletelyInRing(pt3)).Equal(true)

            fmt.Println("pt1  :", pt1)
            fmt.Println("rng0 :", rng0)

            g.Assert(rng0.Intersects(pt0)).Equal(false)
            g.Assert(rng0.Intersects(pt1)).Equal(true)
            g.Assert(rng0.Intersects(pt2)).Equal(true)
            g.Assert(rng0.Intersects(pt3)).Equal(false)
        })
    })
    g.Describe("LinearRing - Relation to Polygon", func() {
        g.It("should test polygon in ring & intersect", func() {
            g.Assert(rng0.contains_polygon(ply)).IsTrue()
            g.Assert(rng0.contains_polygon(ply0)).IsFalse()
            g.Assert(rng0.contains_polygon(ply0)).IsFalse()

        })
    })
    g.Describe("LinearRing - Relation to Linestring", func() {
        g.It("should test polygon in ring & intersect", func() {
            var rngc = rng0.Clone()
            g.Assert(rng0.intersects_linestring(ln3)).IsFalse()
            g.Assert(rngc.intersects_linestring(ln3)).IsFalse()
            g.Assert(rngc.Intersects(ln3)).IsFalse()
            g.Assert(rng0.contains_line(ln3)).IsFalse()

            g.Assert(rng0.contains_line(ln4)).IsFalse()
            g.Assert(rng0.intersects_linestring(ln4)).IsTrue()
            g.Assert(rngc.intersects_linestring(ln4)).IsTrue()
            g.Assert(rngc.Intersects(ln4)).IsTrue()
        })
    })
}

func TestLinearRingArea(t *testing.T) {
    g := Goblin(t)

    rng0 := NewLinearRing([]*Point{
        {2.28, 3.7}, {2.98, 5.36}, {3.92, 4.8}, {3.9, 3.64}, {2.28, 3.7},
    })

    rng1 := NewLinearRing([]*Point{{3, 1.6}, {3, 2}, {2.4, 2.8}});

    g.Describe("LinearRing", func() {
        g.It("should test area", func() {
            g.Assert(Round(rng0.Area(), 12) == 1.9164).IsTrue()
            g.Assert(Round(rng1.Area(), 12) == 0.12).IsTrue()
        })

    })
}
