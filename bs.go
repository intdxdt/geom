package geom

//compare points as items - x | y ordering
func ptCmp(oa, ob interface{}) int {
	var a, b = oa.(Point), ob.(Point)
	var d = a[X] - b[X]
	if feq(d, 0) {
		d = a[Y] - b[Y]
	}

	var r = 1
	if feq(d, 0) {
		r = 0
	} else if d < 0 {
		r = -1
	}
	return r
}
