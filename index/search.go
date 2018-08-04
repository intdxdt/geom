package index

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
)

//Search item
func (tree *Index) Search(query mbr.MBR) []*mono.MBR {
	var bbox = &query
	var result []*mono.MBR
	var nd = &tree.data

	if !intersects(bbox, &nd.bbox) {
		return []*mono.MBR{}
	}

	var nodesToSearch []*idxNode
	var child *idxNode

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

//All items from  root idxNode
func (tree *Index) All() []*mono.MBR {
	return all(&tree.data, []*mono.MBR{})
}

//all - fetch all items from idxNode
func all(nd *idxNode, result []*mono.MBR) []*mono.MBR {
	var nodesToSearch []*idxNode
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
