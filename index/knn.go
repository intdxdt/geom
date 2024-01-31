package index

import (
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/heap"
	"github.com/intdxdt/mbr"
)

func predicate(_ *KObj) (bool, bool) {
	return true, false
}

func (tree *Index) Knn(
	query mbr.MBR[float64], limit int, score func(*mbr.MBR[float64], *KObj) float64,
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
			var o = &KObj{}
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
				break
			}

			if limit != 0 && len(result) == limit {
				return result
			}
		}

		if !stop {
			var q = queue.Pop()
			if q == nil {
				nd = nil
			} else {
				nd = q.(*KObj).node
			}
		}
	}
	return result
}

func (tree *Index) KnnMinDist(
	query *mono.MBR,
	distScore func(query *mono.MBR, dbitem *mono.MBR) float64,
	predicate func(*KObj, float64) bool, mindist float64) float64 {

	var nd = &tree.data
	//var result []*mono.MBR
	var child *node
	var stop bool
	var queue = heap.NewHeap(kobjCmp, heap.NewHeapType().AsMin())

	for !stop && (nd != nil) {
		for i := range nd.children {
			child = &nd.children[i]
			var box_dist = child.bbox.Distance(&query.MBR)
			if box_dist < mindist {
				var o = &KObj{}
				o.node = child
				o.MBR = &child.bbox
				o.IsItem = len(child.children) == 0
				o.Distance = box_dist
				if o.IsItem {
					o.Distance = distScore(query, child.item)
					if o.Distance < mindist {
						mindist = o.Distance
					}
				}
				queue.Push(o)
			}
		}

		for !queue.IsEmpty() && queue.Peek().(*KObj).IsItem {
			var candidate = queue.Pop().(*KObj)
			stop = predicate(candidate, mindist)
			if stop {
				break
			}
		}

		if !stop {
			var q = queue.Pop()
			if q == nil {
				nd = nil
			} else {
				nd = q.(*KObj).node
			}
		}
	}
	return mindist
}
