package index

//split at index
func splitAtIndex(nodes []node, index int) ([]node, []node) {
	var ln = len(nodes)
	var ext = make([]node, 0, ln-index)
	for i := index; i < ln; i++ {
		ext = append(ext, nodes[i])
		nodes[i] = node{}
	}
	return nodes[:index], ext
}

//slice index
func sliceIndex(limit int, predicate func(i int) bool) int {
	var index = -1
	for i := 0; i < limit; i++ {
		if predicate(i) {
			index = i
			break
		}
	}
	return index
}

//minimum float
func min(a, b float64) float64 {
	if b < a {
		return b
	}
	return a
}

//maximum float
func max(a, b float64) float64 {
	if b > a {
		return b
	}
	return a
}

//minint
func minInt(a, b int) int {
	if b < a {
		return b
	}
	return a
}

//maximum integer
func maxInt(a, b int) int {
	if b > a {
		return b
	}
	return a
}
