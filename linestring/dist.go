package geom

import (
    . "github.com/intdxdt/simplex/geom/segment"
    //"github.com/intdxdt/simplex/geom/mbr"
    "math"
)

//length of line
func (self *LineString) Length() float64 {
    return self.len(0, len(self.coordinates) - 1)
}

//compute length of chain
func (self *LineString) chain_length(chain *MonoMBR) float64 {
    return self.len(chain.i, chain.j)
}

//length of line from index i to j
func (self *LineString) len(i, j int) float64 {
    var dist float64
    if j < i {
        i, j = j, i
    }
    for ; i < j; i++ {
        dist += self.coordinates[i].Distance(self.coordinates[i + 1])
    }
    return dist;
}


//description  Computes the distance between self and another linestring
//the distance between intersecting linestrings is 0.  Otherwise, the
//distance is the Euclidean distance between the closest points.
//param other{LineString}
func (self *LineString) Distance(other *LineString) float64{
    var othersegs = make([]*Segment, 0)
    var selfsegs = make([]*Segment, 0)

    if self.bbox.Disjoint(other.bbox.MBR) {
        return self.mindist_bruteforce(other)
    }
    //if root mbrs intersect
    var bln = false
    var dist float64 = -1
    var query, _ = self.bbox.Intersection(other.bbox.MBR)
    var inrange = self.index.Search(query)
    if len(inrange) == 0 {
        //go bruteforce
        dist = self.mindist_bruteforce(other)
        bln = true
    }
    for i := 0; !bln && i < len(inrange); i++ {
        //cur self box
        ibox := (*inrange[i].GetItem()).(*MonoMBR)
        //search ln using ibox
        query = ibox.MBR
        lnrange := other.index.Search(query)
        if len(lnrange) == 0  {
            //go bruteforce
            dist = self.mindist_bruteforce(other)
            bln = true
        }
        for q := 0; !bln && q < len(lnrange); q++ {
            qbox := (*lnrange[q].GetItem()).(*MonoMBR)
            qrng, ok := ibox.BBox().Intersection(qbox.BBox())
            if ok {
                var xor_segs = true //segments when nothing is in range of qrng
                self.segs_inrange(selfsegs, qrng, ibox.i, ibox.j, false, xor_segs)
                other.segs_inrange(othersegs, qrng, qbox.i, qbox.j, false, xor_segs)

                _dist := segseg_mindist(selfsegs, othersegs)
                if dist < 0 {
                    dist = _dist
                }else {
                    dist = math.Min(_dist, dist)
                }
                if dist == 0.0 {
                    bln = true
                }
            }

        }
    }

    if dist < 0 {
        panic("invalid distance")
    }
    return dist
}


// bruteforce dist,
func (self *LineString) mindist_bruteforce(other *LineString) float64 {
    var bln = false
    var ln = self.coordinates
    var ln2 = other.coordinates
    var dist float64 = -1
    for i := 0; !bln && i < len(ln) - 1; i++ {
        for j := 0; !bln && j < len(ln2) - 1; j++ {
            segA := &Segment{ln[i], ln[i + 1]}
            segB := &Segment{ln2[j], ln2[j + 1]}
            _dist := segA.Distance(segB)
            if dist < 0 {
                dist = _dist
            } else {
                dist = math.Min(_dist, dist)
            }
            if _dist == 0.0 {
                bln = true
            }
        }
    }
    return dist
}


/*
 description minimum distance
 param segsa
 param segsb
 returns {number}
 private
 */
func segseg_mindist(segsa,  segsb []*Segment) float64 {
    var bln = false
    var dist = -1.0
    var _dist float64
    for a := 0; !bln && a < len(segsa); a++ {
        for b := 0; !bln && b < len(segsb); b++ {
            bln = segsa[a].Intersects(segsb[b], false)
            if bln {
                dist = 0.0
            } else {
                _dist = segsa[a].Distance(segsb[b])
                if dist < 0 {
                    dist = _dist
                }else {
                    dist = math.Min(_dist, dist)
                }
            }
        }
    }
    return dist
}




