package geom

import (
    "testing"
    . "github.com/franela/goblin"
    . "github.com/intdxdt/simplex/geom/point"
    . "github.com/intdxdt/simplex/util/math"
)

func TestLinearRing(t *testing.T) {
    g := Goblin(t)
    pts_closed := []*Point{
        {5.6, 7.9}, {5.6, 8.9}, {6.6, 8.9}, {6.6, 7.9}, {5.6, 7.9},
    }

    pts_open := []*Point{
        {5.6, 7.9}, {5.6, 8.9}, {6.6, 8.9}, {6.6, 7.9},
    }


    ln1      := NewLinearRing(pts_closed)
    ln2      := NewLinearRing(pts_open)

    rng0 := NewLinearRing([]*Point{{2.28, 3.7}, {2.98, 5.36}, {3.92, 4.8}, {3.9, 3.64}, {2.28, 3.7}})

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

            g.Assert(rng0.IntersectsPoint(pt0)).Equal(false)
            g.Assert(rng0.IntersectsPoint(pt1)).Equal(true)
            g.Assert(rng0.IntersectsPoint(pt2)).Equal(true)
            g.Assert(rng0.IntersectsPoint(pt3)).Equal(false)
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
            g.Assert(Round(rng0.Area(),12) == 1.9164).IsTrue()
            g.Assert(Round(rng1.Area(),12) == 0.12).IsTrue()
        })

    })
}
