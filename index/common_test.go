package index

import (
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/math"
	"github.com/intdxdt/mbr"
)

// brute force distance
func mindistBruteforce(ln, ln2 [][]float64) float64 {
	var dist = math.MaxFloat64
	var bln = false
	var n1, n2 = len(ln) - 1, len(ln2) - 1
	var d float64
	for i := 0; !bln && i < n1; i++ {
		for j := 0; !bln && j < n2; j++ {
			d = segsegDistance(ln[i], ln[i+1], ln2[j], ln2[j+1])
			if d < dist {
				dist = d
			}
			//dist = minf64(dist, d)
			bln = dist == 0
		}
	}
	return dist
}

func knnMinLinearDistance(a, b [][]float64) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}
	var db = segmentDB(b)
	var queries = queryBounds(a)

	var dist = math.MaxFloat64
	var d float64
	for q := range queries {
		db.KnnMinDist(&queries[q],
			func(query *mono.MBR, item *mono.MBR) (float64, float64) {
				d = segsegDistance(
					a[query.I], a[query.J], b[item.I], b[item.J],
				)

				if d < dist {
					dist = d
				}
				return d, dist
			},
			func(o *KObj) bool {
				return o.Distance > dist || dist == 0 //add to neibs, stop
			})
	}

	return dist
}
func queryBounds(coords [][]float64) []mono.MBR {
	var n = len(coords) - 1
	var I, J int
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		I, J = i, i+1
		items = append(items,
			mono.MBR{
				MBR: mbr.CreateMBR(
					coords[I][0], coords[I][1],
					coords[J][0], coords[J][1],
				), I: I, J: J,
			})
	}
	return items
}

func segmentDB(coords [][]float64) *Index {
	var tree = NewIndex()
	var n = len(coords) - 1
	var I, J int
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		I, J = i, i+1
		items = append(items,
			mono.MBR{
				MBR: mbr.CreateMBR(
					coords[I][0], coords[I][1],
					coords[J][0], coords[J][1],
				), I: I, J: J,
			})
	}
	return tree.Load(items)
}

//Distance betwen two segments
func segsegDistance(sa, sb, oa, ob []float64) float64 {
	var dist = math.NaN()
	var x1, y1 = sa[0], sa[1]
	var x2, y2 = sb[0], sb[1]

	var x3, y3 = oa[0], oa[1]
	var x4, y4 = ob[0], ob[1]

	var pta, ptb []float64
	var mua, mub float64
	var is_aspt_a, is_aspt_b bool
	var denom, numera, numerb float64

	denom = (y4-y3)*(x2-x1) - (x4-x3)*(y2-y1)
	numera = (x4-x3)*(y1-y3) - (y4-y3)*(x1-x3)
	numerb = (x2-x1)*(y1-y3) - (y2-y1)*(x1-x3)

	if math.Abs(denom) < math.EPSILON {
		is_aspt_a = math.FloatEqual(sa[0], sb[0]) && math.FloatEqual(sa[1], sb[1])
		is_aspt_b = math.FloatEqual(oa[0], ob[0]) && math.FloatEqual(oa[1], ob[1])

		if is_aspt_a && is_aspt_b {
			dist = math.Hypot(x1-x4, y1-y4)
		} else if is_aspt_a || is_aspt_b {
			var lna, lnb []float64

			if is_aspt_a {
				pta = sa
				lna, lnb = oa, ob
			} else if is_aspt_b {
				pta = oa
				lna, lnb = sa, sb
			}

			dist = distanceToPoint(lna, lnb, pta)
		} else {
			dist = math.MinF64(
				math.MinF64(
					distanceToPoint(oa, ob, sa),
					distanceToPoint(oa, ob, sb)),
				math.MinF64(
					distanceToPoint(sa, sb, oa),
					distanceToPoint(sa, sb, ob),
				))
		}

	} else {
		//if close to zero or one , snap
		mua = numera / denom
		if (mua == 0) || math.Abs(mua) < math.EPSILON { //a == b || Abs(a - b) < EPSILON
			mua = 0
		} else if (mua == 1) || math.Abs(mua-1) < math.EPSILON {
			mua = 1
		}

		mub = numerb / denom
		if (mub == 0) || math.Abs(mub) < math.EPSILON {
			mub = 0
		} else if (mub == 1) || math.Abs(mub-1) < math.EPSILON {
			mub = 1
		}

		if mua < 0 || mua > 1 || mub < 0 || mub > 1 {
			//the is intersection along the the segments
			if mua < 0 {
				pta = sa
			} else if mua > 1 {
				pta = sb
			}

			if mub < 0 {
				ptb = oa
			} else if mub > 1 {
				ptb = ob
			}

			if pta != nil && ptb != nil {
				dist = math.MinF64(
					distanceToPoint(oa, ob, pta),
					distanceToPoint(sa, sb, ptb),
				)
			} else if pta != nil {
				dist = distanceToPoint(oa, ob, pta)
			} else {
				dist = distanceToPoint(sa, sb, ptb)
			}
		} else {
			//lines intersect
			dist = 0
		}
	}

	return dist
}

func distanceToPoint(sa, sb, pt []float64) float64 {
	var dist = math.NaN()
	//var cPt *Point
	var ax, ay = sa[0], sa[1]
	var bx, by = sb[0], sb[1]
	var px, py = pt[0], pt[1]
	//var dab = sb.Sub(sa)
	var dx, dy = bx - ax, by - ay
	//a == b || Abs(a - b) < EPSILON
	var isz_x = (dx == 0) || math.Abs(dx) < math.EPSILON
	var isz_y = (dy == 0) || math.Abs(dy) < math.EPSILON

	if isz_x && isz_y {
		//line with zero length
		dist = math.Hypot(px-ax, py-ay)
	} else {
		var cPtx, cPty float64
		var u = (((px - ax) * dx) + ((py - ay) * dy)) / (dx*dx + dy*dy)

		if u < 0 {
			cPtx, cPty = ax, ay
		} else if u > 1 {
			cPtx, cPty = bx, by
		} else {
			cPtx, cPty = ax+u*dx, ay+u*dy
		}
		dist = math.Hypot(px-cPtx, py-cPty)
	}

	return dist
}
