package geom

import (
    . "github.com/intdxdt/simplex/geom/point"
    . "github.com/intdxdt/simplex/geom/mbr"
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
func pop_coords(a []*Point) (*Point, []*Point) {
    var v *Point
    var n int
    if len(a) == 0 {
        return nil, a
    }
    n = len(a) - 1
    v, a[n] = a[n], nil
    return v, a[:n]
}


//compare bbox, ptr
func is_bbox(a *MBR, b *MBR) bool {
    return a == b
}


//line is ring
func (self *LineString) is_ring() bool {
    return self.coordinates[0].Equals(
        self.coordinates[len(self.coordinates) - 1],
    )
}
