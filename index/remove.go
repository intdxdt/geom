package index

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
)

func nodeAtIndex(a []*idxNode, i int) *idxNode {
	var n = len(a)
	if i > n-1 || i < 0 || n == 0 {
		return nil
	}
	return a[i]
}

func nodeSiblingAtIndex(a []idxNode, i int) *idxNode {
	var n = len(a)
	if i > n-1 || i < 0 || n == 0 {
		return nil
	}
	return &a[i]
}

func popNode(a []*idxNode) (*idxNode, []*idxNode) {
	var v *idxNode
	var n int
	if len(a) == 0 {
		return nil, a
	}
	n = len(a) - 1
	v, a[n] = a[n], nil
	return v, a[:n]
}

func popIndex(indxs *[]int) int {
	var n, v int
	a := *indxs
	n = len(a) - 1
	if n < 0 {
		return 0
	}
	v, a[n] = a[n], 0
	*indxs = a[:n]
	return v
}

//remove idxNode at given index from idxNode slice.
func removeNode(a []idxNode, i int) []idxNode {
	var n = len(a) - 1
	if i > n {
		return a
	}
	a, a[n] = append(a[:i], a[i+1:]...), idxNode{}
	return a
}

//condense idxNode and its path from the root
func (tree *Index) condense(path []*idxNode) {
	var parent *idxNode
	var i = len(path) - 1
	// go through the path, removing empty nodes and updating bboxes
	for i >= 0 {
		if len(path[i].children) == 0 {
			//go through the path, removing empty nodes and updating bboxes
			if i > 0 {
				parent = path[i-1]
				index := sliceIndex(len(parent.children), func(s int) bool {
					return path[i] == &parent.children[s]
				})
				if index != -1 {
					parent.children = removeNode(parent.children, index)
				}
			} else {
				//root is empty, rest root
				tree.Clear()
			}
		} else {
			calcBBox(path[i])
		}
		i--
	}
}

//Remove Item from Index
//NOTE: remove node
func (tree *Index) Remove(item *mono.MBR) *Index {
	if item == nil {
		return tree
	}
	tree.removeItem(&item.MBR, func(nd *idxNode, i int) bool {
		return nd.children[i].item == item
	})
	return tree
}

//Remove Item from Index
//NOTE:if item is a bbox , then first found bbox is removed
func (tree *Index) removeMBR(item *mbr.MBR) *Index {
	tree.removeItem(item,
		func(nd *idxNode, i int) bool {
			return nd.children[i].bbox.Equals(item)
		})
	return tree
}

func (tree *Index) removeItem(item *mbr.MBR, predicate func(*idxNode, int) bool) *Index {
	var nd = &tree.data
	var parent *idxNode
	var bbox = item.BBox()
	var path = make([]*idxNode, 0)
	var indexes = make([]int, 0)
	var i, index int
	var goingUp bool

	// depth-first iterative self traversal
	for (nd != nil) || (len(path) > 0) {
		if nd == nil {
			// go up
			nd, path = popNode(path)
			parent = nodeAtIndex(path, len(path)-1)
			i = popIndex(&indexes)
			goingUp = true
		}

		if nd.leaf {
			// check current idxNode
			//index = idxNode.children.indexOf(item)
			index = sliceIndex(len(nd.children), func(i int) bool {
				return predicate(nd, i)
			})

			//if found
			if index != -1 {
				//item found, remove the item and condense self upwards
				//idxNode.children.splice(index, 1)
				nd.children = removeNode(nd.children, index)
				path = append(path, nd)
				tree.condense(path)
				return tree
			}
		}

		if !goingUp && !nd.leaf && contains(&nd.bbox, bbox) {
			// go down
			path = append(path, nd)
			indexes = append(indexes, i)
			i = 0
			parent = nd
			nd = &nd.children[0]
		} else if parent != nil {
			// go right
			i++
			nd = nodeSiblingAtIndex(parent.children, i)
			goingUp = false
		} else {
			nd = nil
		} // nothing found
	}
	return tree
}
