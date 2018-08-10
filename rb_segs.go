package geom

const (
	CreateRED = iota
	CreateBLUE
	RemoveRED
	RemoveBLUE
)

func addSegment(
	index int,
	red *LineString, redList *bfList,
	blue *LineString, blueList *bfList,
	visit func(int, int) bool, flip bool) bool {

	//Look up segment
	var ra = red.Coordinates.Pt(index)
	var rb = red.Coordinates.Pt(index + 1)
	var ba, bb *Point

	//Read out components
	var l0 = minf64(ra[Y], rb[Y])
	var h0 = maxf64(ra[Y], rb[Y])

	//Scan over blue intervals for point
	var intervals = blueList.intervals
	var blueIndex = blueList.index
	var count = blueList.count
	var ptr = 2 * count
	var h1, l1 float64
	var bindex int
	var ret bool

	for i := count - 1; !ret && i >= 0; i-- {
		ptr += -1
		h1 = intervals[ptr]
		ptr += -1
		l1 = intervals[ptr]

		//Test if intervals overlap
		if l0 <= h1 && l1 <= h0 {
			bindex = blueIndex[i]
			ba = blue.Coordinates.Pt(bindex)
			bb = blue.Coordinates.Pt(bindex + 1)

			//Test if segments intersect
			if SegSegIntersects(ra, rb, ba, bb) {
				if flip {
					ret = visit(bindex, index)
				} else {
					ret = visit(index, bindex)
				}
			}
		}
	}
	redList.insert(l0, h0, index)
	return ret
}
