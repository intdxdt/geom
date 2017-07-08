package geom

import (
	"github.com/franela/goblin"
	"sort"
	"testing"
)

func TestCoordinate(t *testing.T) {
	g := goblin.Goblin(t)
	coords := Coordinates{{0.0, 0.0}, {1.0, 0.0}, {2.0, 0.0}}
	_2d := Coordinates{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}}
	_2d1 := Coordinates{{0.0, 0.2}, {1.0, 0.9}, {1.0, 0.5}}
	_2d2 := Coordinates{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.5}}
	_2dclone := Coordinates{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}}
	_2dcoords := XYCoordinates{_2dclone}
	_2dexpect := Coordinates{{0.0, 0.2}, {1.0, 0.05}, {1.0, 0.1}}
	xcoords := XCoordinates{coords}
	ycoords := YCoordinates{coords}
	xycoords := XYCoordinates{coords}

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

			sort.Sort(&xcoords)
			g.Assert(xcoords.Coordinates).Eql(coords)
			xcoords.Sort()
			g.Assert(xcoords.Coordinates).Eql(coords)

			sort.Sort(&ycoords)
			g.Assert(ycoords.Coordinates).Eql(coords)
			ycoords.Sort()
			g.Assert(ycoords.Coordinates).Eql(coords)

			sort.Sort(xycoords)
			g.Assert(xycoords.Coordinates).Eql(coords)

			sort.Sort(_2dcoords)
			g.Assert(_2dcoords.Coordinates).Eql(_2dexpect)

			_2dy := YCoordinates{_2d}
			_2dy.Sort()
			g.Assert(_2dy.Coordinates).Eql(Coordinates{{1.0, 0.05}, {1.0, 0.1}, {0.0, 0.2}})

			_2dx := XCoordinates{_2d}
			_2dx.Sort()
			g.Assert(_2dx.Coordinates).Eql(Coordinates{{0.0, 0.2}, {1.0, 0.05}, {1.0, 0.1}})

			_2dxy := XYCoordinates{_2d}
			_2dxy.Sort()
			g.Assert(_2dxy.Coordinates).Eql(Coordinates{{0.0, 0.2}, {1.0, 0.05}, {1.0, 0.1}})

			_2dxy = XYCoordinates{_2d1}
			_2dxy.Sort()
			g.Assert(_2dxy.Coordinates).Eql(Coordinates{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.9}})

			_2dxy = XYCoordinates{_2d2}
			_2dxy.Sort()
			g.Assert(_2dxy.Coordinates).Eql(Coordinates{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.5}})
		})
	})

}
