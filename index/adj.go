package index

import "github.com/intdxdt/mbr"

// adjust bboxes along the given tree path
func (tree *Index) adjustParentBBoxes(bbox *mbr.MBR, path []*node, level int) {
	for i := level; i >= 0; i-- {
		path[i].bbox.ExpandIncludeMBR(bbox)
	}
}
