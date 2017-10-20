package geom

import (
	"bytes"
	"strconv"
	"github.com/intdxdt/math"
)

type InterPoint struct {
	Pt         *Point
	I, J, K, L int
}

func InterPointCmp(a, b interface{}) int {
	self   := a.(*InterPoint)
	other  := b.(*InterPoint)
	return self.Compare(other)
}

//compare points as items - x | y ordering
func (self *InterPoint) Compare(other *InterPoint) int {
	d := self.Pt[X] - other.Pt[X]
	if math.FloatEqual(d, 0.0) {
		d = self.Pt[Y] - other.Pt[Y]
	}

	if math.FloatEqual(d, 0.0) {
		//check if close enougth to zero
		dx := self.I - other.I
		if dx == 0 {
			dx = self.J - other.J
		}
		if dx < 0 {
			return -1
		} else if dx > 0 {
			return 1
		}
		return 0
	} else if d < 0 {
		return -1
	}
	return 1
}

//string
func (self *InterPoint) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	buf.WriteString(strconv.FormatFloat(self.Pt[X], 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(self.Pt[Y], 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(self.Pt[Z], 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(float64(self.I), 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(float64(self.J), 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(float64(self.K), 'f', -1, 64) + ", ")
	buf.WriteString(strconv.FormatFloat(float64(self.L), 'f', -1, 64))
	buf.WriteString("]")

	return buf.String()
}
