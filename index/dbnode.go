package index

import (
	"github.com/intdxdt/mbr"
	"github.com/intdxdt/geom/mono"
)

//idxNode type for internal idxNode
type idxNode struct {
	children []idxNode
	item     *mono.MBR
	height   int
	leaf     bool
	bbox     mbr.MBR
}

//createIdxNode creates a new idxNode
func createIdxNode(item *mono.MBR, height int, leaf bool, children []idxNode) idxNode {
	var box mbr.MBR
	if item == nil {
		box = emptyMBR()
	} else {
		box = item.MBR
	}

	return idxNode{
		children: children,
		item:     item,
		height:   height,
		leaf:     leaf,
		bbox:     box,
	}
}

//idxNode type for internal idxNode
func newLeafNode(item *mono.MBR) idxNode {
	return idxNode{
		children: []idxNode{},
		item:     item,
		height:   1,
		leaf:     true,
		bbox:     item.MBR,
	}
}

//MBR returns bbox property
func (nd *idxNode) BBox() *mbr.MBR {
	return &nd.bbox
}

//add child
func (nd *idxNode) addChild(child idxNode) {
	nd.children = append(nd.children, child)
}

//Constructs children of idxNode
func makeChildren(items []*mono.MBR) []idxNode {
	var chs = make([]idxNode, 0, len(items))
	for i := range items {
		chs = append(chs, newLeafNode(items[i]))
	}
	return chs
}
