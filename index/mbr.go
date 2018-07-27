package index

import (
	"sort"
	"github.com/intdxdt/mbr"
)

//calcBBox calculates its bbox from bboxes of its children.
func calcBBox(nd *idxNode) {
	nd.bbox = distBBox(nd, 0, len(nd.children))
}

//distBBox computes min bounding rectangle of idxNode children from k to p-1.
func distBBox(nd *idxNode, k, p int) mbr.MBR {
	var bbox = emptyMBR()
	for i := k; i < p; i++ {
		extend(&bbox, &nd.children[i].bbox)
	}
	return bbox
}

//allDistMargin computes total margin of all possible split distributions.
//Each idxNode is at least m full.
func (tree *Index) allDistMargin(nd *idxNode, m, M int, sortBy sortBy) float64 {
	if sortBy == byX {
		sort.Sort(xNodePath{nd.children})
		//bubbleAxis(*idxNode.getChildren(), byX, byY)
	} else if sortBy == byY {
		sort.Sort(yNodePath{nd.children})
		//bubbleAxis(*idxNode.getChildren(), byY, byX)
	}

	var i int
	var leftBBox = distBBox(nd, 0, m)
	var rightBBox = distBBox(nd, M-m, M)
	var margin = bboxMargin(&leftBBox) + bboxMargin(&rightBBox)

	for i = m; i < M-m; i++ {
		extend(&leftBBox, &nd.children[i].bbox)
		margin += bboxMargin(&leftBBox)
	}

	for i = M - m - 1; i >= m; i-- {
		extend(&rightBBox, &nd.children[i].bbox)
		margin += bboxMargin(&rightBBox)
	}
	return margin
}
