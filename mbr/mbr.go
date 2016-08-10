package mbr

import (
    "math"
)

type MBR [4]float64

const (
    x1 = iota
    y1
    x2
    y2
)

func NewMBR(minx, miny, maxx, maxy float64) *MBR {
    minx, maxx = math.Min(minx, maxx), math.Max(minx, maxx)
    miny, maxy = math.Min(miny, maxy), math.Max(miny, maxy)
    return &MBR{minx, miny, maxx, maxy}
}

func (self *MBR ) Clone() *MBR {
    return &MBR{self[x1], self[y1], self[x2], self[y2]}
}

func (self *MBR ) BBox() *MBR {
    return self
}
