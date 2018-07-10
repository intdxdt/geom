package geom

import (
	"testing"
	"github.com/franela/goblin"
)

func TestPointToPolyTangent(t *testing.T) {
	g := goblin.Goblin(t)
	wkt := "LINESTRING ( 240 200, 240 260, 280 250, 290 220, 350 210, 340 280, 290 300, 290 340, 350 370, 390 360, 470 380, 510 320, 470 260, 550 220, 590 310, 650 200, 540 170, 470 190, 440 140, 370 120, 300 160, 370 170, 410 240, 390 270 )"
	ln := NewLineStringFromWKT(wkt)

	coords := ln.coordinates
	//hull := ConvexHull(coords)
	//ply := NewPolygon(hull)

	var i, j int
	pt0 := PointXY(570, 60)
	pt1 := PointXY(780, 320)
	pt2 := PointXY(190, 410)
	pt3 := PointXY(120, 210)

	g.Describe("TangentPointToPoly", func() {
		g.It("should tangent point to polygon", func() {
			i, j = TangentPointToPoly(&pt0, coords)

			g.Assert([]int{i, j}).Eql([]int{15, 19})
			i, j = TangentPointToPoly(&pt1, coords)

			g.Assert([]int{i, j}).Eql([]int{10, 15})
			i, j = TangentPointToPoly(&pt2, coords)

			g.Assert([]int{i, j}).Eql([]int{0, 10})

			i, j = TangentPointToPoly(&pt3, coords)
			g.Assert([]int{i, j}).Eql([]int{19, 7})
		})

	})
}
