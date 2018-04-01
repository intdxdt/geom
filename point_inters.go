package geom

import (
	"bytes"
	"github.com/intdxdt/math"
	"fmt"
)

type InterPoint struct {
	*Point
	Inter VBits
}

type IntPts []*InterPoint

func (s IntPts) Len() int {
	return len(s)
}
func (s IntPts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s IntPts) Less(i, j int) bool {
	return (s[i].Point[X] < s[j].Point[X]) || (
		math.FloatEqual(s[i].Point[X], s[j].Point[X]) && s[i].Point[Y] < s[j].Point[Y])
}

func (p *InterPoint) IsIntersection() bool {
	return p.Inter == 0
}

func (p *InterPoint) IsVertex() bool {
	var mask = SelfMask | OtherMask
	return p.Inter&mask > 0
}

func (p *InterPoint) IsVertexSelf() bool {
	return p.Inter&SelfMask > 0
}

func (p *InterPoint) IsVertexOther() bool {
	return p.Inter&OtherMask > 0
}

func (p *InterPoint) IsVerteXOR() bool {
	return ( p.IsVertexSelf() &&  !p.IsVertexOther()) ||
			(!p.IsVertexSelf() &&   p.IsVertexOther())
}

//string
func (self *InterPoint) String() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	buf.WriteString(math.FloatToString(self.Point[X]) + ", ")
	buf.WriteString(math.FloatToString(self.Point[Y]) + ", ")
	buf.WriteString(math.FloatToString(self.Point[Z]) + ", ")
	buf.WriteString(fmt.Sprintf("%04b", self.Inter))
	buf.WriteString("]")
	return buf.String()
}
