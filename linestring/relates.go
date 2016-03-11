package linestring

import (
    "github.com/intdxdt/simplex/geom/point"
//"github.com/intdxdt/simplex/geom/mbr"
//    "fmt"
)

//intersection of self linestring with other
func (self *LineString) Intersection(other *LineString) []*point.Point {
    var ptlist = make([]*point.Point, 0)
    if !self.bbox.Intersects(other.bbox.MBR) {
        return ptlist //disjoint
    }

    //if root mbrs intersect
    //var i, q, lnrange, ibox, qbox, qrng
    var othersegs = make([]*LineString, 0)
    var selfsegs = make([]*LineString, 0)

    var query = other.bbox
    var inrange = self.index.Search(query.MBR)
    for i := 0; i < len(inrange); i++ {
        //cur self box
        ibox := (*inrange[i].GetItem()).(*MonoMBR)
        //search ln using ibox

        lnrange := other.index.Search(query.MBR)
        for q := 0; q < len(lnrange); q++ {
            qbox := (*lnrange[q].GetItem()).(*MonoMBR)
            qrng, _ := ibox.BBox().Intersection(qbox.BBox())
            self.segs_inrange(selfsegs, qrng, ibox.i, ibox.j, false, false)
            other.segs_inrange(othersegs, qrng, qbox.i, qbox.j, false, false)
            //self.segseg_intersection(selfsegs, othersegs, ptlist, true)
        }
    }
    //return _.map(ptlist, func (pt) {
    //  return Point(pt)
    //})
    return ptlist  //debug
}




