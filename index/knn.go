package index

import (
	"github.com/intdxdt/heap"
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/math"
)

func predicate(_ *KObj) (bool, bool) {
	return true, false
}

func (tree *Index) Knn(
	query mbr.MBR, limit int, score func(*mbr.MBR, *KObj) float64,
	predicates ...func(*KObj) (bool, bool)) []*mono.MBR {

	var predFn = predicate
	if len(predicates) > 0 {
		predFn = predicates[0]
	}

	var nd = &tree.data
	var result []*mono.MBR
	var child *node
	var stop, pred bool
	var queue = heap.NewHeap(kobjCmp, heap.NewHeapType().AsMin())

	for !stop && (nd != nil) {
		for i := range nd.children {
			child = &nd.children[i]
			var o = kobjPool.Get().(*KObj)
			o.node = child
			o.MBR = &child.bbox
			o.IsItem = len(child.children) == 0
			o.Distance = score(&query, o)
			queue.Push(o)
		}

		for !queue.IsEmpty() && queue.Peek().(*KObj).IsItem {
			var candidate = queue.Pop().(*KObj)
			pred, stop = predFn(candidate)
			if pred {
				result = append(result, candidate.GetNode())
			}

			if stop {
				kobjPool.Put(candidate)
				break
			}

			if limit != 0 && len(result) == limit {
				kobjPool.Put(candidate)
				return result
			}
			kobjPool.Put(candidate)
		}

		if !stop {
			var q = queue.Pop()
			if q == nil {
				nd = nil
			} else {
				var candidate = q.(*KObj)
				nd = candidate.node
				kobjPool.Put(candidate)
			}
		}
	}
	for _, o := range queue.View() {
		kobjPool.Put(o.(*KObj))
	}
	return result
}

func (tree *Index) KnnMinDist(
	query *mono.MBR,
	distScore func(query *mono.MBR, dbitem *mono.MBR) (float64, float64),
	predicate func(*KObj) bool,
) {

	var nd = &tree.data
	//var result []*mono.MBR
	var child *node
	var stop bool
	var queue = heap.NewHeap(kobjCmp, heap.NewHeapType().AsMin())
	var mindist = math.MaxFloat64

	for !stop && (nd != nil) {
		for i := range nd.children {
			child = &nd.children[i]
			var box_dist = child.bbox.Distance(&query.MBR)
			if box_dist < mindist {
				var o = kobjPool.Get().(*KObj)
				o.node = child
				o.MBR = &child.bbox
				o.IsItem = len(child.children) == 0
				o.Distance = box_dist
				if o.IsItem {
					o.Distance, mindist = distScore(query, child.item)
				}
				queue.Push(o)
			}
		}

		for !queue.IsEmpty() && queue.Peek().(*KObj).IsItem {
			var candidate = queue.Pop().(*KObj)
			stop = predicate(candidate)
			if stop {
				kobjPool.Put(candidate)
				break
			}
			kobjPool.Put(candidate)
		}

		if !stop {
			var q = queue.Pop()
			if q == nil {
				nd = nil
			} else {
				var candidate = q.(*KObj)
				nd = candidate.node
				kobjPool.Put(candidate)
			}
		}
	}
	for _, o := range queue.View() {
		kobjPool.Put(o.(*KObj))
	}
}
