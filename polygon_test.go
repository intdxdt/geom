package geom

import (
    "testing"
    . "github.com/franela/goblin"
    "fmt"
)

func TestPolygon(t *testing.T) {
    g := Goblin(t)

    sh := []*Point{{35, 10}, {45, 45}, {15, 40}, {10, 20}, {35, 10}, }
    h1 := []*Point{{20, 30}, {35, 35}, {30, 20}, {20, 30}, }
    poly0 := NewPolygon(sh)
    poly := NewPolygon(sh, h1)
    //poly := NewPolygon(sh)

    g.Describe("Polygon", func() {
        g.It("should test polygon relates", func() {
            wkt := "POLYGON (( 33.52991674117594 27.137460594059416, 33.52991674117594 30.589750223527805, 36.44941148514852 30.589750223527805, 36.44941148514852 27.137460594059416, 33.52991674117594 27.137460594059416 ))"
            ply_inpoly := NewPolygonFromWKT(wkt)
            ply_inpoly_clone := ply_inpoly.Clone()
            g.Assert(poly.Intersects(ply_inpoly)).IsTrue()
            g.Assert(poly.Intersects(ply_inpoly_clone)).IsTrue()
            //g.Assert(poly.Contains(ply_inpoly)).IsTrue()
        })
        g.It("should test polygon string", func() {
            g.Assert(poly.String()).Eql(
                "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))",
            )
            g.Assert(fmt.Sprintf("%v", poly0)).Eql(
                "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10))",
            )
            g.Assert(fmt.Sprintf("%v", poly)).Eql(
                "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10),(20 30, 35 35, 30 20, 20 30))",
            )
        })

    })
}
