package index

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
)

//loadBoxes loads bounding boxes
func (tree *Index) loadBoxes(data []mbr.MBR) *Index {
	var items = make([]mono.MBR, 0, len(data))
	for i := range data {
		items = append(items, mono.MBR{MBR: data[i]})
	}
	return tree.Load(items)
}

//Load implements bulk loading
func (tree *Index) Load(items []mono.MBR) *Index {
	var n = len(items)
	if n < tree.minEntries {
		for i := range items {
			tree.insert(&items[i])
		}
		return tree
	}

	var data = make([]*mono.MBR,  n)
	for i := range items {
		data[i] = &items[i]
	}

	// recursively build the tree with the given data from stratch using OMT algorithm
	var nd = tree.buildTree(data, 0, n-1, 0)

	if len(tree.data.children) == 0 {
		// save as is if tree is empty
		tree.data = nd
	} else if tree.data.height == nd.height {
		// split root if trees have the same height
		tree.splitRoot(tree.data, nd)
	} else {
		if tree.data.height < nd.height {
			// swap trees if inserted one is bigger
			tree.data, nd = nd, tree.data
		}

		// insert the small tree into the large tree at appropriate level
		tree.insertNode(nd, tree.data.height-nd.height-1)
	}

	return tree
}
