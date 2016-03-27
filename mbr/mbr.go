package geom

import (
    "math"
    . "github.com/intdxdt/simplex/geom/point"
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

//ExpandXY expands mbr to include x and y
func NewMBRFromPoints(pts ...*Point) *MBR {
    if len(pts) == 0 {
        return NewMBR(0, 0, 0, 0)
    }
    i := 0
    x, y := pts[i].X(), pts[i].Y()
    mbr := NewMBR(x, y, x, y)
    for i = 1; i < len(pts); i++ {
        x, y = pts[i].X(), pts[i].Y()
        mbr.ExpandIncludeXY(x, y)
    }
    return mbr
}

func (self *MBR ) Clone() *MBR {
    return &MBR{self[x1], self[y1], self[x2], self[y2]}
}

func (self *MBR ) BBox() *MBR {
    return self
}
