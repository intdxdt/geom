package mbr

import (
	"simplex/util/math"
)

type MBR [4]float64

const (
	x1 = iota
	y1
	x2
	y2
)

func NewMBR(minx, miny, maxx, maxy float64) *MBR {
	minx, maxx = math.MinF64(minx, maxx), math.MaxF64(minx, maxx)
	miny, maxy = math.MinF64(miny, maxy), math.MaxF64(miny, maxy)
	return &MBR{minx, miny, maxx, maxy}
}

func (self *MBR) Clone() *MBR {
	return &MBR{self[x1], self[y1], self[x2], self[y2]}
}

func (self *MBR) BBox() *MBR {
	return self
}

func (self *MBR) MinX() float64 {
	return self[x1]
}

func (self *MBR) MinY() float64 {
	return self[y1]
}

func (self *MBR) MaxX() float64 {
	return self[x2]
}

func (self *MBR) MaxY() float64 {
	return self[y2]
}
