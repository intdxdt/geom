package geom

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/math"
	"testing"
)

func TestCart(t *testing.T) {
	g := goblin.Goblin(t)
	p0 := Pt(0.0, 0.0)
	pn := Pt(0.0, math.NaN())
	p1 := Pt(4, 5)
	p2 := Pt(4.0, 5.0)
	p3 := Point{4.0, math.NaN()}

	g.Describe("geom.point", func() {
		g.It("x, y access & null", func() {
			g.Assert(p0.IsZero()).IsTrue()
			g.Assert(pn.IsNull()).IsTrue()

			g.Assert(p1.IsZero()).IsFalse()
			g.Assert(p1.Equals2D(&p2)).IsTrue()

			g.Assert(p1[X]).Equal(4.0)
			g.Assert(p1[Y]).Equal(5.0)
			g.Assert(p0.IsNull()).IsFalse()
			g.Assert(p3.IsNull()).IsTrue()
		})

	})

	g.Describe("Point distance and to polygon ", func() {
		g.It("sqrt(3**2,4**2) ", func() {
			var pt = Point{3., 0.}
			g.Assert(pt.Magnitude(&Point{0., 4.})).Equal(5.0)
			g.Assert(pt.MagnitudeSquare(&Point{0., 4.})).Equal(25.0)
		})
		g.It("sqrt(2)", func() {
			var pt = Point{3, 4}
			g.Assert(pt.Magnitude(&Point{4, 5})).Equal(math.Sqrt2)
			g.Assert(pt.MagnitudeSquare(&Point{4, 5})).Equal(2.0)
		})
	})

	g.Describe("operators", func() {
		g.It("component ", func() {
			cx, cy := Component(5, math.Deg2rad(53.13010235415598))
			g.Assert(math.FloatEqual(cx, 3.0)).IsTrue()
			g.Assert(math.FloatEqual(cy, 4.0)).IsTrue()
		})
		g.It("add ", func() {
			var a, b = Point{3., 0.}, Point{0., 4.}
			var cx, cy = a.Add(b[X], b[Y])
			g.Assert(Point{cx, cy}).Equal(Point{3., 4.})
		})

		g.It("sub & neg ", func() {
			var a, b = Point{3., 4.}, Point{4, 5}
			var subpt = b.Neg()
			g.Assert(subpt).Equal(Point{-4, -5})
			var cx, cy = a.Sub(b[X], b[Y])
			g.Assert(Point{cx, cy}).Equal(Point{-1.0, -1.0})
		})
	})
}

//Test Neg
func TestNegCart(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Negate Vector", func() {
		g.It("should test vector negation", func() {
			a := []float64{10, 150, 6.5}
			e := []float64{280, 280, 12.8}
			A := Pt(a[X], a[Y])
			B := Pt(e[X], e[Y])

			var vx, vy = B.Sub(A[X], A[Y])
			var pv = Pt(vx, vy)
			var nv = pv.Neg()
			var negA = Pt(0, 0)
			for i, v := range A {
				negA[i] = -v
			}
			vx, vy = KProduct(pv[X], pv[Y], -1)
			g.Assert(nv).Eql(Pt(vx, vy))
			g.Assert(A.Neg()).Eql(negA)

		})
	})

}

func TestMagDistCart(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Point - Vector MagnitudeXY", func() {
		g.It("should test vector MagnitudeXY and distance", func() {
			var a = Point{0, 0}
			var b = Point{3, 4}
			var z = Pt(0, 0)
			var o = Pt(1, 1)
			g.Assert(o.Magnitude(&z)).Equal(math.Sqrt2)
			var x = Pt(-3, 2)
			g.Assert(math.Round(x.Magnitude(&z), 8)).Equal(
				math.Round(3.605551275463989, 8),
			)
			g.Assert(MagnitudeXY(3, 4)).Equal(5.0)
			x = Pt(3, 4)
			g.Assert(x.Magnitude(&z)).Equal(5.0)
			g.Assert(a.Magnitude(&b)).Equal(5.0)
			x = Pt(3, 4)

			g.Assert(x.MagnitudeSquare(&z)).Equal(25.0)
			g.Assert(MagnitudeSquareXY(3, 4)).Equal(25.0)
			g.Assert(a.MagnitudeSquare(&b)).Equal(25.0)
			x = Pt(4.587, 0.)
			g.Assert(x.Magnitude(&z)).Equal(4.587)
		})
	})

}

func TestDotProductCart(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Point - Vector Dot Product", func() {
		g.It("should test dot product", func() {
			var dotProd = DotProduct(1.2, -4.2, 1.2, -4.2)
			g.Assert(19.08).Equal(math.Round(dotProd, 8))
		})
	})

}

func TestSideOfCart(t *testing.T) {
	var g = goblin.Goblin(t)
	var a = Pt(237, 289)
	var b = Pt(404.25, 357.25)
	var c = Pt(460, 380)
	var d = Pt(297.13043478260863, 339.30434782608694)
	var e = Pt(445.8260869565217, 350.17391304347825)

	cx, cy := b.Sub(a[X], a[Y])
	ab := Pt(cx, cy)

	cx, cy = c.Sub(a[X], a[Y])
	ac := Pt(cx, cy)

	cx, cy = d.Sub(a[X], a[Y])
	ad := Pt(cx, cy)

	cx, cy = e.Sub(a[X], a[Y])
	ae := Pt(cx, cy)

	g.Describe("Orientation and cross product", func() {
		g.It("orientation", func() {
			g.Assert(Orientation2D(&a, &b, &c) == 0).IsTrue()
			g.Assert(Orientation2D(&a, &c, &d) < 0).IsTrue()
			g.Assert(Orientation2D(&a, &c, &e) > 0).IsTrue()

		})
		g.It("cross product", func() {
			g.Assert(CrossProduct(ab[X], ab[Y], ac[X], ac[Y]) == 0).IsTrue()
			g.Assert(CrossProduct(ac[X], ac[Y], ad[X], ad[Y]) > 0).IsTrue()
			g.Assert(CrossProduct(ac[X], ac[Y], ae[X], ae[Y]) < 0).IsTrue()
		})
	})

}

func TestCCW(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Sidedness", func() {
		g.It("should test side of point to vector", func() {
			var k = Point{-0.887, -1.6128}
			var u = Point{4.55309, 1.42996}
			var testpoints = []Point{{2, 2}, {0, 2}, {0, -2}, {2, -2}, {0, 0}, {2, 0}, u, k}

			left, right, on := func(x float64) bool {
				return x < 0
			}, func(x float64) bool {
				return x > 0
			}, func(x float64) bool {
				return math.FloatEqual(x, 0)
			}

			var sides = make([]float64, len(testpoints))
			for i, pnt := range testpoints {
				sides[i] = Orientation2D(&k, &u, &pnt)
			}
			g.Assert(Orientation2D(&k, &u, &Point{2, 2}) < 0).IsTrue()

			sideOut := []func(x float64) bool{
				left, left, right, right, left,
				right, on, on,
			}

			for i := range sideOut {
				g.Assert(sideOut[i](sides[i])).IsTrue()
			}
		})
	})

}

func TestProj(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - unit & Project", func() {
		var A = Point{0.88682, -1.06102}
		var B = Point{3.5, 1.0}
		g.It("should test projection", func() {
			g.Assert(math.Round(Project(&A, &B), 5)).Equal(0.56121)
		})
		g.It("should test Unit", func() {
			Z := Point{0., 0.}
			cx, cy := UnitVector(Z[X], Z[Y])
			g.Assert(math.FloatEqual(cx, 0)).IsTrue()
			g.Assert(math.FloatEqual(cy, 0)).IsTrue()
		})
	})
}

func TestDirection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Direction", func() {
		g.It("should test vector direction", func() {
			var A = Point{0, 0}
			var B = Point{-1, 0}
			var cx, cy = B.Sub(A[X], A[Y])
			var v = Pt(cx, cy)
			g.Assert(Direction(1, 1)).Equal(0.7853981633974483)
			g.Assert(Direction(-1, 0)).Equal(math.Pi)
			g.Assert(Direction(v[X], v[Y])).Equal(math.Pi)
			g.Assert(Direction(1, math.Sqrt(3))).Equal(math.Deg2rad(60))
			g.Assert(Direction(0, -1)).Equal(math.Deg2rad(270))
		})
	})

}

func TestReverseDirection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector RevDirection", func() {
		g.It("should test reverse vector direction", func() {
			var A = Point{0, 0}
			var B = Point{-1, 0}
			var cx, cy = B.Sub(A[X], A[Y])
			var v = Pt(cx, cy)
			g.Assert(ReverseDirection(Direction(v[X], v[Y]))).Equal(0.0)
			g.Assert(ReverseDirection(0.7853981633974483)).Equal(0.7853981633974483 + math.Pi)
			g.Assert(ReverseDirection(0.7853981633974483 + math.Pi)).Equal(0.7853981633974483)
		})
	})

}

func TestDeflection(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector Deflection", func() {
		g.It("should test reverse vector direction", func() {
			var ln0 = []Point{{0, 0}, {20, 30}}
			var ln1 = []Point{{20, 30}, {40, 15}}
			var cx, cy = ln0[1].Sub(ln0[0][X], ln0[0][Y])
			var v0 = Point{cx, cy}
			cx, cy = ln1[1].Sub(ln1[0][X], ln1[0][Y])
			var v1 = Point{cx, cy}

			g.Assert(math.Round(DeflectionAngle(
				Direction(v0[X], v0[Y]),
				Direction(v1[X], v1[Y]),
			), 10)).Equal(math.Round(math.Deg2rad(93.17983011986422), 10))
			g.Assert(math.Round(DeflectionAngle(
				Direction(v0[X], v0[Y]),
				Direction(v0[X], v0[Y]),
			), 10)).Equal(math.Deg2rad(0.0))

			ln1 = []Point{{20, 30}, {20, 60}}
			cx, cy = ln1[1].Sub(ln1[0][X], ln1[0][Y])
			v1 = Point{cx, cy}
			g.Assert(math.Round(DeflectionAngle(
				Direction(v0[X], v0[Y]),
				Direction(v1[X], v1[Y]),
			), 10)).Equal(
				math.Round(math.Deg2rad(-33.690067525979806), 10),
			)
		})
	})

}

func TestDistanceToPoint(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Dist2Vect", func() {
		g.It("should test distance vector", func() {
			var a       = Point{16.82295, 10.44635}
			var b       = Point{28.99656, 15.76452}
			var onAb    = Point{25.32, 14.16}

			tpoints := []Point{
				{30., 0.},
				{15.78786, 25.26468},
				{-2.61504, -3.09018},
				{28.85125, 27.81773},
				a,
				b,
				onAb,
			}

			var tDists = []float64{14.85, 13.99, 23.69, 12.05, 0.00, 0.00, 0.00}
			var dists  = make([]float64, len(tpoints))

			for i, tp := range tpoints {
				dists[i] = DistanceToPoint(&a, &b, &tp)
			}

			for i := range tpoints {
				g.Assert(math.Round(dists[i], 2)).Equal(math.Round(tDists[i], 2))
			}
		})
	})

}
