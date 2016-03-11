package linestring

import (
    "github.com/intdxdt/simplex/geom/point"
    "github.com/intdxdt/simplex/geom/mbr"
)

//pop chain from chainl list
func pop_mono_mbr(a []*MonoMBR) (*MonoMBR, []*MonoMBR) {
    var v *MonoMBR
    var n int
    if len(a) == 0 {
        return nil, a
    }
    n = len(a) - 1
    v, a[n] = a[n], nil
    return v, a[:n]
}


//pop chain from chainl list
func pop_coords(a []*point.Point) (*point.Point, []*point.Point) {
    var v *point.Point
    var n int
    if len(a) == 0 {
        return nil, a
    }
    n = len(a) - 1
    v, a[n] = a[n], nil
    return v, a[:n]
}


//compare bbox, ptr
func is_bbox(a *mbr.MBR, b *mbr.MBR) bool {
    return a == b
}

//convert point to line
func (self *LineString) conv_pt_linestring(other *point.Point) []*LineString {
    coords, lines := make([]*point.Point, 2), make([]*LineString, 1)
    coords[0], coords[1] = other.Clone(), other.Clone()
    lines[0] = NewLineString(coords)
    return lines
}

 //line is ring
func (self *LineString) is_ring() bool{
  var n  = len(self.coordinates)
  var p0 = self.coordinates[0]
  var pn = self.coordinates[n - 1]
  return (n > 2) && (p0[x] == pn[x]) && (p0[y] == pn[y])
}
