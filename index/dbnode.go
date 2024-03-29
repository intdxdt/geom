package index

import (
	"github.com/intdxdt/geom/mono"
	"github.com/intdxdt/mbr"
)

// node type for internal node
type node struct {
	children []node
	item     *mono.MBR
	height   int
	leaf     bool
	bbox     mbr.MBR[float64]
}

// createNode creates a new node
func createNode(item *mono.MBR, height int, leaf bool, children []node) node {
	var box mbr.MBR[float64]
	if item == nil {
		box = emptyMBR()
	} else {
		box = item.MBR
	}

	return node{
		children: children,
		item:     item,
		height:   height,
		leaf:     leaf,
		bbox:     box,
	}
}

// node type for internal node
func newLeafNode(item *mono.MBR) node {
	return node{
		children: []node{},
		item:     item,
		height:   1,
		leaf:     true,
		bbox:     item.MBR,
	}
}

// MBR returns bbox property
func (nd *node) BBox() *mbr.MBR[float64] {
	return &nd.bbox
}

// Add child
func (nd *node) addChild(child node) {
	nd.children = append(nd.children, child)
}

// Constructs children of node
func makeChildren(items []*mono.MBR) []node {
	var chs = make([]node, 0, len(items))
	for i := range items {
		chs = append(chs, newLeafNode(items[i]))
	}
	return chs
}
