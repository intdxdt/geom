package index

import (
	"math"
	"sort"
)

// _split overflowed idxNode into two
func (tree *Index) split(insertPath []*idxNode, level int) {
	var nd = insertPath[level]
	var newNode = createIdxNode(nil, nd.height, nd.leaf, []idxNode{})
	var M = len(nd.children)
	var m = tree.minEntries

	tree.chooseSplitAxis(nd, m, M)
	var at = tree.chooseSplitIndex(nd, m, M)
	//perform split at index
	nd.children, newNode.children = splitAtIndex(nd.children, at)

	calcBBox(nd)
	calcBBox(&newNode)

	if level > 0 {
		insertPath[level-1].addChild(newNode)
	} else {
		tree.splitRoot(*nd, newNode)
	}
}

//_splitRoot splits the root of tree.
func (tree *Index) splitRoot(nd, other idxNode) {
	// split root idxNode
	tree.data = createIdxNode(nil, nd.height+1, false, []idxNode{nd, other})
	calcBBox(&tree.data)
}

//_chooseSplitIndex selects split index.
func (tree *Index) chooseSplitIndex(nd *idxNode, m, M int) int {
	var i, index int
	var overlap, area, minOverlap, minArea float64

	minOverlap, minArea = math.Inf(1), math.Inf(1)

	for i = m; i <= M-m; i++ {
		var bbox1 = distBBox(nd, 0, i)
		var bbox2 = distBBox(nd, i, M)

		overlap = intersectionArea(&bbox1, &bbox2)
		area = bboxArea(&bbox1) + bboxArea(&bbox2)

		// choose distribution with minimum overlap
		if overlap < minOverlap {
			minOverlap = overlap
			index = i

			if area < minArea {
				minArea = area
			}
		} else if overlap == minOverlap {
			// otherwise choose distribution with minimum area
			if area < minArea {
				minArea = area
				index = i
			}
		}
	}

	return index
}

//_chooseSplitAxis selects split axis : sorts idxNode children
//by the best axis for split.
func (tree *Index) chooseSplitAxis(nd *idxNode, m, M int) {
	var xMargin = tree.allDistMargin(nd, m, M, byX)
	var yMargin = tree.allDistMargin(nd, m, M, byY)

	// if total distributions margin value is minimal for x, sort by minX,
	// otherwise it's already sorted by minY
	if xMargin < yMargin {
		sort.Sort(&xNodePath{nd.children})
	}
}
