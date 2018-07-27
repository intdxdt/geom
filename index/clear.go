package index

import (
	"math"
	"github.com/intdxdt/mbr"
)

func emptyMBR() mbr.MBR {
	return mbr.MBR{
		math.Inf(1), math.Inf(1),
		math.Inf(-1), math.Inf(-1),
	}
}

func (tree *Index) Clear() *Index {
	tree.data = createIdxNode(nil, 1, true, []idxNode{})
	return tree
}

//IsEmpty checks for empty tree
func (tree *Index) IsEmpty() bool {
	return len(tree.data.children) == 0
}
