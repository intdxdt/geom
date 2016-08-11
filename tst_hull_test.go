package geom

import (
    "testing"
    . "github.com/franela/goblin"
)

func TestHull(t *testing.T) {

    g := Goblin(t)
    var A = NewPointXY(172.0, 224.0)
    var B = NewPointXY(180.0, 158.0)
    var C = NewPointXY(266.0, 46.0)
    var D = NewPointXY(374.0, 38.0)
    var E = NewPointXY(480.0, 100.0)
    var F = NewPointXY(500.0, 200.0)
    var G = NewPointXY(440.0, 300.0)
    var H = NewPointXY(340.0, 280.0)
    var I = NewPointXY(200.0, 240.0)

    coords := []*Point{A, B, C, D, E, F, G, H, I}
    hull := NewHull(coords)
    g.Describe("Geometry", func() {
        g.It("it should test Hull Antipodal", func() {
            g.Assert(hull.Antipodal(2, 3)).Equal(6)
            g.Assert(hull.Antipodal(0, 1)).Equal(5)
            g.Assert(hull.Antipodal(4, 5)).Equal(0)
        })
    })

}

