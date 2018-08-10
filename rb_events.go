package geom

import (
	"sort"
	"github.com/intdxdt/math"
)

type event struct {
	val float64
	ev  int
	idx int
}

//coordinates iterable of points
type events []event

//Lexicographic sort
func (o events) Sort() {
	sort.Sort(o)
}

//Len for sort interface
func (o events) Len() int {
	return len(o)
}

//Swap for sort interface
func (o events) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

//lexical sort of x and y coordinates
func (o events) Less(i, j int) bool {
	var a, b = o[i], o[j]
	var d = a.val - b.val
	var id int
	//x's are close enough to each other
	if math.FloatEqual(d, 0) {
		id = a.ev - b.ev
	} else {
		return d < 0
	}

	//y's are close enough to each other
	if id == 0 {
		id = a.idx - b.idx
	}
	return id < 0
}

func prepareEvents(red, blue *LineString) []event {
	var nr = red.Coordinates.Len() - 1
	var nb = blue.Coordinates.Len() - 1
	var i, ptr, idx int
	var n = nr + nb
	var data = make([]event, 0, 2*n)

	for i, idx = 0, 0; i < len(red.rbEvents); i += 2 {
		//red.rbEvents[i].ev, red.rbEvents[i].idx = CreateRED, idx
		data = append(data, event{red.rbEvents[i], CreateRED, idx})
		ptr++

		//red.rbEvents[i+1].ev, red.rbEvents[i+1].idx = RemoveRED, idx
		data = append(data, event{red.rbEvents[i+1], RemoveRED, idx})
		ptr++
		idx++
	}

	for i, idx = 0, 0; i < len(blue.rbEvents); i += 2 {
		//blue.rbEvents[i].ev, blue.rbEvents[i].idx = CreateBLUE, idx
		data = append(data, event{blue.rbEvents[i], CreateBLUE, idx})
		ptr++

		//blue.rbEvents[i+1].ev, blue.rbEvents[i+1].idx = RemoveBLUE, idx
		data = append(data, event{blue.rbEvents[i+1], RemoveBLUE, idx})
		//data[ptr] = &blue.rbEvents[i+1]
		ptr++
		idx++
	}

	events(data).Sort()

	return data
}
