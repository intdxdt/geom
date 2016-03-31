package geom

import (
    "math"
)

//Distance computes distance between two points
func (self *Point ) Distance(other Geometry) float64 {
    var dist = math.NaN()
    if IsNullGeometry(other) {
        return dist
    }
    pt, ispoint := IsPoint(other)
    ln, isline := IsLineString(other)
    poly, ispolygon := IsPolygon(other)
    if ispoint {
        dist = self.Sub(pt).Magnitude()
    } else if isline {
        lnear := self.AsLinear()[0]
        dist = ln.Distance(lnear)
    }  else if ispolygon {
        lnear := self.AsLinear()[0]
        dist = lnear.Distance(poly.Shell.LineString)
    }
    return dist
}

// Computes the distance between self and another linestring
// the distance between intersecting linestrings is 0.  Otherwise, the
// distance is the Euclidean distance between the closest points.
func (self *Polygon) Distance(other Geometry) float64 {
    var dist = math.NaN()
    if IsNullGeometry(other) {
        return dist
    }
    _, isline := IsLineString(other)
    _, ispoint := IsPoint(other)
    poly, ispolygon := IsPolygon(other)
    //reverse intersect line inter poly
    if isline || ispoint {
        other_lns := other.AsLinear()
        if len(other_lns) > 0 {
            dist = self.Shell.Distance(other_lns[0])
        }
    } else if ispolygon {
        dist = self.Shell.Distance(poly.Shell.LineString)
    }
    return dist
}



//description  Computes the distance between self and another linestring
//the distance between intersecting linestrings is 0.  Otherwise, the
//distance is the Euclidean distance between the closest points.
func (self *LineString) Distance(other *LineString) float64 {
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
        if len(lnrange) == 0 {
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
        dist = math.NaN()
    }
    return dist
}


// brute force distance
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


//minimum distance
func segseg_mindist(segsa, segsb []*Segment) float64 {
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




