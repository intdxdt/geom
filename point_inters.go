package geom

import (
    "bytes"
    "strconv"
    . "simplex/struct/item"
    . "simplex/util/math"
)

type InterPoint struct {
    Pt         *Point
    I, J, K, L int
}

//compare points as items - x | y ordering
func (self *InterPoint) Compare(o Item) int {
    other := o.(*InterPoint)
    d := self.Pt[x] - other.Pt[x]
    if FloatEqual(d, 0.0) {
        //x's are close enougth to each other
        d = self.Pt[y] - other.Pt[y]
    }

    if FloatEqual(d, 0.0) {
        //check if close enougth ot zero
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
    buf.WriteString(strconv.FormatFloat(self.Pt[x], 'f', -1, 64) + ", ")
    buf.WriteString(strconv.FormatFloat(self.Pt[y], 'f', -1, 64) + ", ")
    buf.WriteString(strconv.FormatFloat(self.Pt[z], 'f', -1, 64) + ", ")
    buf.WriteString(strconv.FormatFloat(float64(self.I), 'f', -1, 64) + ", ")
    buf.WriteString(strconv.FormatFloat(float64(self.J), 'f', -1, 64) + ", ")
    buf.WriteString(strconv.FormatFloat(float64(self.K), 'f', -1, 64) + ", ")
    buf.WriteString(strconv.FormatFloat(float64(self.L), 'f', -1, 64))
    buf.WriteString("]")

    return buf.String()
}

