package index

/*
 (c) 2015, Titus Tienaah
 A library for 2D spatial indexing of points and rectangles.
 https://github.com/mourner/rbush
 @after  (c) 2015, Vladimir Agafonkin
*/

//Index type
type Index struct {
	data       node
	maxEntries int
	minEntries int
}

func NewIndex(nodeCap ...int) *Index {
	var bucketSize = 8
	var tree = &Index{}
	tree.Clear()
	if len(nodeCap) > 0 {
		bucketSize = nodeCap[0]
	}
	// bucket size(node) == 8 by default
	tree.maxEntries = maxEntries(bucketSize)
	tree.minEntries = minEntries(tree.maxEntries)
	return tree
}
