package geom

//intersection of self linestring with other
func (self *LineString) Intersection(other *LineString) []*Point {
    var ptlist = make([]*Point, 0)

    if self.bbox.Disjoint(other.bbox.MBR) {
        return ptlist //disjoint
    }

    //if root mbrs intersect
    //var i, q, lnrange, ibox, qbox, qrng
    var othersegs = make([]*Segment, 0)
    var selfsegs = make([]*Segment, 0)

    var query = other.bbox
    var inrange = self.index.Search(query.MBR)
    for i := 0; i < len(inrange); i++ {
        //cur self box
        ibox := (*inrange[i].GetItem()).(*MonoMBR)
        //search ln using ibox

        lnrange := other.index.Search(query.MBR)
        for q := 0; q < len(lnrange); q++ {
            qbox := (*lnrange[q].GetItem()).(*MonoMBR)
            qrng, ok := ibox.BBox().Intersection(qbox.BBox())

            if ok {
                selfsegs = self.segs_inrange(selfsegs, qrng, ibox.i, ibox.j, false, false)
                othersegs = other.segs_inrange(othersegs, qrng, qbox.i, qbox.j, false, false)
                ptlist = self.segseg_intersection(selfsegs, othersegs, ptlist, true)
            }
        }
    }
    return ptlist  //debug
}


//Checks if line intersects other line
//other{LineString} - geometry types and array as Point
func (self *LineString) intersects_linestring(other *LineString) bool {

    var bln = false
    //if root mbrs intersect
    var othersegs = make([]*Segment, 0)
    var selfsegs = make([]*Segment, 0)

    var query = other.bbox.MBR
    var inrange = self.index.Search(query)

    for i := 0; !bln && i < len(inrange); i++ {
        //search ln using ibox
        ibox := (*inrange[i].GetItem()).(*MonoMBR)
        query = ibox.BBox()
        lnrange := other.index.Search(query)

        for q := 0; !bln && q < len(lnrange); q++ {
            qbox := (*lnrange[q].GetItem()).(*MonoMBR)
            qrng, _ := ibox.Intersection(qbox.MBR)
            selfsegs = self.segs_inrange(
                selfsegs, qrng, ibox.i, ibox.j, false, false,
            )
            othersegs = other.segs_inrange(
                othersegs, qrng, qbox.i, qbox.j, false, false,
            )
            if len(othersegs) > 0 && len(selfsegs) > 0 {
                bln = self.segseg_intersects(selfsegs, othersegs)
            }
        }
    }
    return bln
}


//line intersect polygon rings
func (self *LineString) intersects_polygon(lns []*LineString) bool {
    var bln , intersects_hole, in_hole bool
    var rings = make([]*LinearRing, len(lns))
    for i, ln := range lns {
        rings[i] = &LinearRing{ln}
    }
    var shell = rings[0]
    bln = self.Intersects(shell.LineString)
    //if false, check if shell contains line
    if !bln {

        bln = shell.contains_line(self)
        //inside shell, does it touch hole boundary ?
        for i := 1; bln && !intersects_hole && i < len(rings); i++ {
            intersects_hole = self.Intersects(rings[i].LineString)
        }
        //inside shell but does not touch the boundary of holes
        if bln && !intersects_hole {
            //check if completely contained in hole
            for i := 1; !in_hole && i < len(rings); i++ {
                in_hole = rings[i].contains_line(self)
            }
        }
        bln = bln && !in_hole
    }
    return bln
}


////test intersects of self line string with point
//func (self *LineString) IntersectsPoint(other *Point) bool {
//    if other == nil {
//        return false
//    }
//    return self.Intersects(other.AsLineString())
//}


// Tests whether a collection of segments from line a and line b intersects
// TODO:Improve from O(n2) - although expects few number of segs from index selection
func (self *LineString)segseg_intersects(segsa []*Segment, segsb []*Segment) bool {
    var bln = false
    for a := 0; !bln && a < len(segsa); a++ {
        for b := 0; !bln && b < len(segsb); b++ {
            bln = segsa[a].Intersects(segsb[b], false)
        }
    }
    return bln
}

//
///*
// description list of self intersection coordinates
// */
//func (self *LineString)self_intersection () {
//
//  var cache map[string]string
//    var ckey string
//  var bcomplx, chain, inters, jbox, qbox
//  var ln1 = [], ln2 = [], ptlist = [], i, j
//  var cmp = func (a, b) {
//    return a[0] - b[0] || a[1] - b[1]
//  }
//  var selfinters = struct.sset(cmp)
//
//  for (i = 0 i < self.len(chains) ++i) {
//    chain = self.chains[i]
//    inters = self.index.search(self._searchbox(chain))
//
//    for (j = 0 j < len(inters) ++j) {
//      jbox = inters[j]
//      ckey = self._cashe_key(chain, jbox)
//
//      if cache[ckey] || jbox.equals(chain) {
//        continue//already checked || already monotone
//      }
//
//      self._cashe_ij(cache, chain, jbox, true)
//      qbox = chain.intersection(jbox)
//      if qbox.isnil() && chain.j == jbox.i {
//        continue//non overlapping && contiguous
//      }
//      self._segsinrange(ln1, qbox, chain.i, chain.j)
//      self._segsinrange(ln2, qbox, jbox.i, jbox.j)
//      self._segseg_intersection(ln1, ln2, ptlist)
//
//      bcomplx = (chain.j != jbox.i && len(ptlist) > 0) ||
//                (chain.j == jbox.i && len(ptlist) > 1)
//      if bcomplx {
//        _.each(ptlist, func (pt) {
//          selfinters.append(Point(pt))
//        })
//      }
//    }
//  }
//  return selfinters.slice(0)
//}
//






