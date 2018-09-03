package geom

import (
	"testing"
	"github.com/franela/goblin"
)

func TestCoordinate(t *testing.T) {
	var g = goblin.Goblin(t)
	var cds = Coordinates([]Point{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}})
	var c0 = cds.ShallowClone(0, 1)
	g.Assert(cds.DataView()).Equal([]Point{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}})
	g.Assert(cds.Clone().DataView()).Equal([]Point{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}})
	g.Assert(c0.Points()).Equal([]Point{{0.0, 0.2}})
	var c1 = cds.ShallowClone(0, 2)
	g.Assert(c1.Points()).Equal([]Point{{0.0, 0.2}, {1.0, 0.1}})
	g.Assert(c1).Equal(cds.ShallowClone(2))

	cds = Coordinates([]Point{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}})
	var cds_1 = cds.Clone()
	cds_1.Append(Point{2, 2})
	cds_1.Append(Point{3, 3})
	g.Assert(cds.DataView()).Equal([]Point{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}})
	g.Assert(cds_1.DataView()).Equal([]Point{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}, {2, 2}, {3, 3}})

	var c2 = cds.ShallowClone(1, 3)
	g.Assert(c2.Points()).Equal([]Point{{1.0, 0.1}, {1.0, 0.05}})
	g.Assert(c2.Clone().Points()).Equal([]Point{{1.0, 0.1}, {1.0, 0.05}})

	g.Assert(cds.Slice(1, 3).Points()).Equal(c2.Points())
	var c4 = cds.ShallowClone(2, 3)
	g.Assert(c4.Points()).Equal([]Point{{1.0, 0.05}})
	g.Assert(cds.Slice(2, 3).Points()).Equal(c4.Points())
	var c5 = cds.ShallowClone(2, 2)
	g.Assert(c5.Points()).Equal([]Point{})

	var coords = []Point{{0.0, 0.0}, {1.0, 0.0}, {2.0, 0.0}}
	var coords2d = Coordinates([]Point{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}})
	var coords2d1 = Coordinates([]Point{{0.0, 0.2}, {1.0, 0.9}, {1.0, 0.5}})
	var coords2d2 = Coordinates([]Point{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.5}})
	var coords2dclone = []Point{{0.0, 0.2}, {1.0, 0.1}, {1.0, 0.05}}
	var coords_2d = Coordinates(append([]Point{}, coords2dclone...))
	var expect_2d = []Point{{0.0, 0.2}, {1.0, 0.05}, {1.0, 0.1}}
	var xycoords = Coordinates(coords)

	g.Describe("geom.point", func() {
		g.It("tests Coords as an array of points", func() {
			var coords2 = Coordinates([]Point{{0.0, 0.0}, {1.0, 0.0}, {2.0, 0.0}})
			g.Assert(coords2.FirstIndex()).Equal(0)
			g.Assert(coords2.LastIndex()).Equal(2)
			var bln, c = coords2.Pop()
			g.Assert(coords2.FirstIndex()).Equal(0)
			g.Assert(coords2.LastIndex()).Equal(1)

			c[2] = 4.5
			g.Assert(bln).IsTrue()
			g.Assert(c.Equals2D(&Point{2.0, 0.0, 4.5})).IsTrue()
			var pt = PointXY(2.0, 0.0)
			g.Assert(c.Equals3D(&pt)).IsFalse()
			g.Assert(c.Equals3D(&Point{2.0, 0.0, 4.5})).IsTrue()

			bln, c = coords2.Pop()
			g.Assert(bln).IsTrue()
			pt = PointXY(1.0, 0.0)
			g.Assert(c.Equals2D(&pt)).IsTrue()
			pt = PointXY(1.0, 0.0)
			g.Assert(c.Equals3D(&pt)).IsTrue()

			bln, c = coords2.Pop()
			g.Assert(bln).IsTrue()
			pt = PointXY(0.0, 0.0)
			g.Assert(c.Equals2D(&pt)).IsTrue()
			g.Assert(coords2.Len()).Equal(0)
			bln, c = coords2.Pop()
			g.Assert(bln).IsFalse()
			g.Assert(c.IsNull()).IsTrue()

			g.Assert(xycoords.Sort().Points()).Eql(coords)
			g.Assert(coords_2d.Sort().Points()).Eql(expect_2d)
			g.Assert(coords2d.Sort().Points()).Eql([]Point{{0.0, 0.2}, {1.0, 0.05}, {1.0, 0.1}})
			g.Assert(coords2d1.Sort().Points()).Eql([]Point{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.9}})
			g.Assert(coords2d2.Sort().Points()).Eql([]Point{{0.0, 0.2}, {1.0, 0.5}, {1.0, 0.5}})
		})
	})

}
