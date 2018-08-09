package geom

type bfList struct {
	intervals []float64
	index     []int
	count     int
}

//It is silly, but this is faster than doing the right thing for up to a
//few thousand segments, which hardly occurs in practice.
func createBruteForceList(capacity int) bfList {
	return bfList{
		intervals: make([]float64, 2*capacity), //pool.mallocDouble(2 * capacity)
		index:     make([]int, capacity),       //pool.mallocInt32(capacity)
		count:     0,
	}
}

func (brt *bfList) reset() {
	for i := range brt.intervals {
		brt.intervals[i] = 0
	}
	for i := range brt.index {
		brt.index[i] = 0
	}
}

func (brt *bfList) insert(lo, hi float64, index int) {
	var count = brt.count
	brt.index[count] = index
	brt.intervals[2*count] = lo
	brt.intervals[2*count+1] = hi
	brt.count += 1
}

func (brt *bfList) remove(index int) {
	var count = brt.count
	var rindex = brt.index
	var intervals = brt.intervals
	for i := count - 1; i >= 0; i-- {
		if rindex[i] == index {
			rindex[i] = rindex[count-1]
			intervals[2*i] = intervals[2*(count-1)]
			intervals[2*i+1] = intervals[2*count-1]
			brt.count += -1
			break
		}
	}
}
