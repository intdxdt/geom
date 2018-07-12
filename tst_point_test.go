package geom

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/math"
	"testing"
	. "github.com/intdxdt/mbr"
)

func TestPoint(t *testing.T) {
	g := goblin.Goblin(t)
	p0 := PointXY(0.0, 0.0)
	p1 := PointXY(4, 5)
	p2 := PointXY(4.0, 5.0)
	p3 := CreatePoint([]float64{4, 5})
	p4 := CreatePoint([]float64{4, 5.01})
	p5 := CreatePoint([]float64{4})
	p6 := Point{4.0, math.NaN()}
	p7 := PointXY(4.0, 4.9)
	p8 := PointXY(3.9, 4.9)

	p9 := PointXYZ(3.9, 4.9, 9.8)
	p10 := CreatePoint([]float64{3.9, 4.9, 9.8})
	p11 := CreatePoint([]float64{3.9, 4.9, 9.8, 7.9})

	g.Describe("geom.point", func() {
		g.It("loads wkt as point", func() {
			g.Assert(p1.Type().IsPoint()).IsTrue()
			var p1Box = p1.BBox()
			g.Assert(p1Box.Area()).Equal(0.0)

			g.Assert(p1.Area()).Equal(0.0)
			g.Assert(PointFromWKT(p1.WKT())).Eql(p1)
			g.Assert(PointFromWKT(p4.WKT())).Eql(p4)
		})

		g.It("x, y, z access", func() {
			g.Assert(p0.IsZero()).IsTrue()
			g.Assert(p1.IsZero()).IsFalse()
			g.Assert(p1.Equals2D(&p2)).IsTrue()
			g.Assert(p1.Equals2D(&p1)).IsTrue()
			//g.Assert(p1.Coordinates().Equals2D(p2)).IsTrue()

			g.Assert(p1[X]).Equal(4.0)
			g.Assert(p1[Y]).Equal(5.0)
			g.Assert(p1[Z]).Equal(0.0)

			g.Assert(p5[X]).Equal(4.0)
			g.Assert(p5[Y]).Equal(0.0)

			g.Assert(p3[X]).Equal(p1[X])
			g.Assert(p3[Y]).Equal(p1[Y])

			//3.9, 4.9, 9.8
			g.Assert(p9[X]).Equal(3.9)
			g.Assert(p9[Y]).Equal(4.9)
			g.Assert(p9[Z]).Equal(9.8)
			g.Assert(p9.Equals2D(&p10)).IsTrue()
			g.Assert(p9.Equals2D(&p11)).IsTrue()

		})

		g.It("point relate", func() {
			pc := p1
			g.Assert(p1.Equals2D(&pc)).IsTrue()
			g.Assert(p1.Compare(&pc)).Equal(0)
			g.Assert(p1.Compare(&p2)).Equal(0)
			g.Assert(p1.Compare(&p4)).Equal(-1)
			g.Assert(p1.Compare(&p0)).Equal(1)
			g.Assert(p1.Compare(&p8)).Equal(1)
			g.Assert(p8.Compare(&p1)).Equal(-1)
			g.Assert(p1.Compare(&p7)).Equal(1)
			g.Assert(p7.Compare(&p1)).Equal(-1)

		})

		g.It("as array", func() {
			g.Assert(p1.ToArray()).Equal([]float64{p1[X], p1[Y], 0})
		})

	})

	g.Describe("Point distance and to polygon ", func() {
		g.It("sqrt(3**2,4**2) ", func() {

			pt := &Point{3., 0.}
			g.Assert(pt.Distance(&Point{0., 4.})).Equal(5.0)
			g.Assert(pt.SquareDistance(&Point{0., 4.})).Equal(25.0)

		})
		g.It("sqrt(2)", func() {
			pt := &Point{3, 4}
			g.Assert(pt.Distance(&Point{4, 5})).Equal(math.Sqrt2)
			g.Assert(pt.SquareDistance(&Point{4, 5})).Equal(2.0)
		})
	})

	g.Describe("Point operators", func() {
		g.It("add ", func() {
			a, b := &Point{3., 0.}, &Point{0., 4.}
			g.Assert(a.Add(b[0], b[1])).Equal(Point{3., 4.})
		})

		g.It("sub & neg ", func() {
			a, b := &Point{3., 4.}, &Point{4, 5}
			nb := b.Neg()
			g.Assert(a.Sub(b[0], b[1])).Equal(Point{-1.0, -1.0})
			g.Assert(nb).Equal(Point{-4, -5})
		})
	})

	g.Describe("type conversion & util", func() {
		g.It("wkt string", func() {
			a := Point{3.87, 7.45}
			g.Assert(a.WKT()).Equal("POINT (3.87 7.45)")
			g.Assert(a.BBox()).Equal(CreateMBR(3.87, 7.45, 3.87, 7.45))
			g.Assert(a.ConvexHull().Shell.ToArray()).Equal([][]float64{{3.87, 7.45, 0}, {3.87, 7.45, 0}, {3.87, 7.45, 0}})
		})
	})

	g.Describe("point relates", func() {
		g.It("intersect , equals, isnull ", func() {
			var p0 *Point
			g.Assert(p3.Equals2D(&p1)).IsTrue()
			g.Assert(p3.Intersects(&p1)).IsTrue()
			g.Assert(p3.Intersects(p1.Geometry())).IsTrue()
			g.Assert(p3.Intersects(p0)).IsFalse()
			g.Assert(p3.Geometry().Intersects(p0)).IsFalse()
			g.Assert(p3.Disjoint(&p1)).IsFalse()
			g.Assert(p3.Disjoint(&p4)).IsTrue()
			g.Assert(p6.IsNull()).IsTrue()
		})
	})

}

func TestMagDist(t *testing.T) {
	var g = goblin.Goblin(t)
	g.Describe("Point - Vector Magnitude", func() {
		g.It("should test vector magnitude and distance", func() {
			var a = Point{0, 0}
			var b = Point{3, 4}
			var z = PointXY(0, 0)
			var pt = PointXY(1, 1)
			g.Assert((&pt).Magnitude(&z)).Equal(math.Sqrt2)
			pt = PointXY(-3, 2)
			g.Assert(math.Round((&pt).Magnitude(&z), 8)).Equal(
				math.Round(3.605551275463989, 8),
			)
			pt = PointXY(3, 4)
			g.Assert((&pt).Magnitude(&z)).Equal(5.0)
			g.Assert(a.Distance(&b)).Equal(5.0)
			pt = PointXY(3, 4)
			g.Assert((&pt).SquareMagnitude(&z)).Equal(25.0)
			g.Assert(a.SquareDistance(&b)).Equal(25.0)

			pt = PointXY(4.587, 0.)
			g.Assert((&pt).Magnitude(&z)).Equal(4.587)
		})
	})

}

func TestDotProduct(t *testing.T) {
	var g = goblin.Goblin(t)
	g.Describe("Point - Vector Dot Product", func() {
		g.It("should test dot product", func() {
			pt := PointXY(1.2, -4.2)
			pt_a := PointXY(1.2, -4.2)
			dot_prod := pt_a.DotProduct(&pt)
			g.Assert(19.08).Equal(math.Round(dot_prod, 8))
		})
	})

}

func TestAngleAtPnt(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Point - Angle at Point", func() {
		g.It("should test angle formed at point by vector", func() {
			a := &Point{-1.28, 0.74}
			b := &Point{1.9, 4.2}
			c := &Point{3.16, -0.84}
			g.Assert(math.Round(a.AngleAtPoint(b, c), 8)).Equal(math.Round(1.1694239325184717, 8))
			g.Assert(math.Round(b.AngleAtPoint(a, c), 8)).Equal(math.Round(0.9882331199311394, 8))
		})
	})

}

func TestSideOf(t *testing.T) {
	var g = goblin.Goblin(t)
	var a = PointXY(237, 289)
	var b = PointXY(404.25, 357.25)
	var c = PointXY(460, 380)
	var d = PointXY(297.13043478260863, 339.30434782608694)
	var e = PointXY(445.8260869565217, 350.17391304347825)

	g.Describe("side of point", func() {
		g.It("side of line a, c", func() {
			g.Assert(b.SideOf(&a, &c).IsOn()).IsTrue()
			g.Assert(b.SideOf(&a, &c).IsOnOrLeft()).IsTrue()
			g.Assert(b.SideOf(&a, &c).IsOnOrRight()).IsTrue()

			g.Assert(d.SideOf(&a, &c).IsLeft()).IsTrue()
			g.Assert(d.SideOf(&a, &c).IsOnOrLeft()).IsTrue()
			g.Assert(d.SideOf(&a, &c).IsOnOrRight()).IsFalse()

			g.Assert(e.SideOf(&a, &c).IsRight()).IsTrue()
			g.Assert(e.SideOf(&a, &c).IsOnOrRight()).IsTrue()
			g.Assert(e.SideOf(&a, &c).IsOnOrLeft()).IsFalse()
		})
	})

}
