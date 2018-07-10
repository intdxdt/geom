package geom

import "github.com/intdxdt/math"

//Int cmp
func intCmp(a, b interface{}) int {
	return a.(int) - b.(int)
}

//compare points as items - x | y ordering
func ptCmp(oa, ob interface{}) int {
	var a, b = oa.(Point), ob.(Point)
	var d = a[X] - b[X]
	if math.FloatEqual(d, 0) {
		d = a[Y] - b[Y]
	}
	var r = 1
	if math.FloatEqual(d, 0) {
		r = 0
	} else if d < 0 {
		r = -1
	}
	return r
}

//binary search assumes the array is sorted
func bsInts(array []int, key int, ij ...int) int {
	low := 0
	high := len(array) - 1
	if len(ij) > 0 {
		low = ij[0]
		high = ij[1]
	}

	var mid, v int
	for low <= high {
		mid = low + ((high - low) >> 1)
		v   = intCmp(array[mid], key)
		if v < 0 {
			low = mid + 1 //too low
		} else if v > 0 {
			high = mid - 1 //too high
		} else {
			return mid //found
		}
	}
	return -(low + 1) //key not found
}

//binary search assumes the array is sorted
func bsPts(array []Point, key *Point, ij ...int) int {
	low := 0
	high := len(array) - 1
	if len(ij) > 0 {
		low = ij[0]
		high = ij[1]
	}

	var mid, v int
	for low <= high {
		mid = low + ((high - low) >> 1)
		v   = ptCmp(&array[mid], key)
		if v < 0 {
			low = mid + 1 //too low
		} else if v > 0 {
			high = mid - 1 //too high
		} else {
			return mid //found
		}
	}
	return -(low + 1) //key not found
}
