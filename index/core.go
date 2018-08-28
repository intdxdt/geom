package index

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/geom/mono"
)

var inf = math.Inf(1)
var feq =  math.FloatEqual
type sortBy int

const (
	byX sortBy = iota
	byY
)
const (
	cmpMinX = iota
	cmpMinY
)

func maxEntries(x int) int {
	return maxInt(4, x)
}

func minEntries(x int) int {
	return maxInt(2, int(math.Ceil(float64(x)*0.4)))
}

func swapItem(arr []*mono.MBR, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func popInt(a []int) (int, []int) {
	var n, v int
	n = len(a) - 1
	v, a[n] = a[n], 0
	a = a[:n]
	return v, a
}

func appendInts(a []int, v ...int) []int {
	for i := range v {
		a = append(a, v[i])
	}
	return a
}
