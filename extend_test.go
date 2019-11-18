package geom

import (
	"github.com/franela/goblin"
	"github.com/intdxdt/math"
	"testing"
)

const precision = 8

var A2 = Point{0.88682, -1.06102}
var B2 = Point{3.5, 1}
var C2 = Point{-3, 1}
var D2 = Point{-1.5, -3}

func TestExtVect(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("Vector - Extend", func() {
		g.It("should test extending a vector", func() {

			var va = A2
			var vb = B2
			var vc = C2
			var vd = D2
			var cx, cy = B2.Sub(D2[X], D2[Y])
			var vdb = Point{cx, cy}
			cx, cy = C2.Sub(B2[X], B2[Y])
			var vbc = Point{cx, cy}

			g.Assert(math.Round(Direction(va[X], va[Y]), precision)).Equal(
				math.Round(math.Deg2rad(309.889497029295), precision),
			)
			g.Assert(math.Round(Direction(vb[X], vb[Y]), precision)).Equal(
				math.Round(math.Deg2rad(15.945395900922854), precision),
			)
			g.Assert(math.Round(Direction(vc[X], vc[Y]), precision)).Equal(
				math.Round(math.Deg2rad(161.565051177078), precision),
			)
			g.Assert(math.Round(Direction(vd[X], vd[Y]), precision)).Equal(
				math.Round(math.Deg2rad(243.43494882292202), precision),
			)
			g.Assert(math.Round(MagnitudeXY(vdb[X], vdb[Y]), 4)).Equal(
				math.Round(6.4031242374328485, 4),
			)
			g.Assert(math.Round(Direction(vdb[X], vdb[Y]), precision)).Equal(
				math.Round(math.Deg2rad(38.65980825409009), precision),
			)
			deflangle := 157.2855876468
			cx, cy = Extend(
				vdb[X], vdb[Y],
				3.64005494464026,
				math.Deg2rad(180+deflangle),
				true,
			)
			vo := Point{cx, cy}

			g.Assert(math.Round(vo[0], precision)).Equal(
				math.Round(-vb[0], precision),
			)
			g.Assert(math.Round(vo[1], precision)).Equal(
				math.Round(-vb[1], precision),
			)

			// "vo by extending vdb by angle to origin"
			// "vo by extending vdb by angle to origin"
			deflangleB := 141.34019174590992

			// extend to c from end
			cx, cy = Extend(vdb[X], vdb[Y], 6.5, math.Deg2rad(180+deflangleB), true)
			vextc := Point{cx, cy}
			g.Assert(math.Round(vbc[0], precision)).Equal(math.Round(vextc[0], precision))
			g.Assert(math.Round(vbc[1], precision)).Equal(math.Round(vextc[1], precision))

			// "vextc with magnitudie extension from vdb Pnts"
			g.Assert(math.Round(vextc[0], precision)).Equal(-MagnitudeXY(cx, cy))
			// "vextc horizontal vector test:  extension from vdb Pnts"
			g.Assert(math.Round(vextc[1], precision)).Equal(0.)

			cx, cy = Deflect(5, 0, 2, math.Deg2rad(90), true)
			//deflection is the right hand angle
			g.Assert(math.Round(cx, precision)).Equal(math.Round(0.0, precision))
			g.Assert(math.Round(cy, precision)).Equal(math.Round(-2, precision))

			cx, cy = Deflect(5, 0, 2, math.Deg2rad(90), false)
			g.Assert(math.Round(cx, precision)).Equal(math.Round(0.0, precision))
			g.Assert(math.Round(cy, precision)).Equal(math.Round(2, precision))
		})
	})

}
