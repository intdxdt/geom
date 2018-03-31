package geom

import (
	"testing"
	"github.com/franela/goblin"
)

func TestCoordinate(t *testing.T) {
	g := goblin.Goblin(t)
	coords := Coordinates{{0.0, 0.0}, {1.0, 0.0}, {2.0, 0.0}}
	coords2d := Coordinates{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}}
	coords2d1 := Coordinates{{0.0, 0.2}, {1.0, 0.9}, {1.0, 0.5}}
	coords2d2 := Coordinates{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.5}}
	coords2dclone := Coordinates{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}}
	coords_2d := append(Coordinates{}, coords2dclone...)
	expect_2d := Coordinates{{0.0, 0.2}, {1.0, 0.05}, {1.0, 0.1}}
	xycoords := Coordinates(coords)

	g.Describe("geom.point", func() {
		g.It("tests coordinates as an array of points", func() {
			coords2 := Coordinates{{0.0, 0.0}, {1.0, 0.0}, {2.0, 0.0}}
			c, coords2 := coords2.Pop()
			c[2] = 4.5
			g.Assert(c.Equals2D(&Point{2.0, 0.0, 4.5})).IsTrue()
			g.Assert(c.Equals3D(NewPointXY(2.0, 0.0))).IsFalse()
			g.Assert(c.Equals3D(&Point{2.0, 0.0, 4.5})).IsTrue()
			c, coords2 = coords2.Pop()
			g.Assert(c.Equals2D(NewPointXY(1.0, 0.0))).IsTrue()
			g.Assert(c.Equals3D(NewPointXY(1.0, 0.0))).IsTrue()
			c, coords2 = coords2.Pop()
			g.Assert(c.Equals2D(NewPointXY(0.0, 0.0))).IsTrue()
			g.Assert(len(coords2)).Equal(0)
			c, coords2 = coords2.Pop()
			g.Assert(c == nil).IsTrue()

			g.Assert(xycoords.Sort()).Eql(coords)
			g.Assert(coords_2d.Sort()).Eql(expect_2d)
			g.Assert(coords2d.Sort()).Eql(Coordinates{{0.0, 0.2}, {1.0, 0.05}, {1.0, 0.1}})
			g.Assert(coords2d1.Sort()).Eql(Coordinates{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.9}})
			g.Assert(coords2d2.Sort()).Eql(Coordinates{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.5}})
		})
	})

}
