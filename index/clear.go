package index

import (
	"github.com/intdxdt/mbr"
	"math"
)

func emptyMBR() mbr.MBR[float64] {
	return mbr.MBR[float64]{
		math.Inf(1), math.Inf(1),
		math.Inf(-1), math.Inf(-1),
	}
}

func (tree *Index) Clear() *Index {
	tree.data = createNode(nil, 1, true, []node{})
	return tree
}

// IsEmpty checks for empty tree
func (tree *Index) IsEmpty() bool {
	return len(tree.data.children) == 0
}
