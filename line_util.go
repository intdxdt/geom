package geom

import (
    . "simplex/geom/mbr"
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
