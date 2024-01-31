package index

import (
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/mbr"
)

// Search item
func (tree *Index) Search(query mbr.MBR[float64]) []*mono.MBR {
	var bbox = &query
	var result []*mono.MBR
	var nd = &tree.data

	if !intersects(bbox, &nd.bbox) {
		return nil
	}

	var nodesToSearch []*node
	var child *node

	for {
		for i := range nd.children {
			child = &nd.children[i]

			if intersects(bbox, &child.bbox) {
				if nd.leaf {
					result = append(result, child.item)
				} else if contains(bbox, &child.bbox) {
					result = all(child, result)
				} else {
					nodesToSearch = append(nodesToSearch, child)
				}
			}
		}

		nd, nodesToSearch = popNode(nodesToSearch)
		if nd == nil {
			break
		}
	}
	return result
}

// All items from  root node
func (tree *Index) All() []*mono.MBR {
	return all(&tree.data, []*mono.MBR{})
}

// all - fetch all items from node
func all(nd *node, result []*mono.MBR) []*mono.MBR {
	var nodesToSearch []*node
	for {
		if nd.leaf {
			for i := range nd.children {
				result = append(result, nd.children[i].item)
			}
		} else {
			for i := range nd.children {
				nodesToSearch = append(nodesToSearch, &nd.children[i])
			}
		}

		nd, nodesToSearch = popNode(nodesToSearch)
		if nd == nil {
			break
		}
	}

	return result
}
