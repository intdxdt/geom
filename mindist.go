package geom

import (
	"time"
	"github.com/TopoSimplify/pln"
	"github.com/intdxdt/geom/index"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/mbr"
	"math"
)

func BigMinDist(a, b pln.Polyline) (float64, float64) {
	var duration float64
	const N = 1000
	var dists = [N]float64{}

	for i := 0; i < N; i++ {
		var t0 = time.Now()
		if a.Coordinates.Len() > b.Coordinates.Len() {
			a, b = b, a
		}
		var db, _ = segmentDB(b)
		var queries = queryBounds(a)
		var dist = math.MaxFloat64
		var d float64
		var item *mono.MBR
		for q := range queries {
			db.KnnMinDist(queries[q].MBR, 1, func(query *mbr.MBR, boxer *index.KObj) (float64, float64) {
				item = boxer.GetNode()
				d = SegSegDistance(
					a.Coordinates.Pt(queries[q].I),
					a.Coordinates.Pt(queries[q].J),
					b.Coordinates.Pt(item.I),
					b.Coordinates.Pt(item.J),
				)
				if d < dist {
					dist = d
				}
				return d, dist
			}, func(_ *index.KObj) (bool, bool) {
				return false, d > dist || dist == 0 //add to neibs, stop
			})
		}
		var t1 = time.Now()
		dists[i] = dist
		duration += t1.Sub(t0).Seconds()

	}
	return dists[0], (duration / float64(N)) * 1000
}

func queryBounds(ln pln.Polyline) []mono.MBR {
	var n = ln.Len() - 1
	var a, b *Point
	var I, J int
	var items = make([]mono.MBR, 0, n)
	for i := 0; i < n; i++ {
		a, b = ln.Coordinates.Pt(i), ln.Coordinates.Pt(i+1)
		//items = append(items, mbr.CreateMBR(a[geom.X], a[geom.Y], b[geom.X], b[geom.Y]))
		I, J = ln.Coordinates.Idxs[i], ln.Coordinates.Idxs[i+1]
		items = append(items, mono.MBR{
			MBR: mbr.CreateMBR(a[X], a[Y], b[X], b[Y]),
			I:   I, J: J,
		})
	}
	return items
}

func segmentDB(polyline pln.Polyline) (*index.Index, []mono.MBR) {
	var tree = index.NewIndex()
	var data = polyline.SegmentBounds()
	tree.Load(data)
	return tree, data
}

