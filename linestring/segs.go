package linestring

import (
    . "github.com/intdxdt/simplex/geom/mbr"
    . "github.com/intdxdt/simplex/geom/segment"
    . "github.com/intdxdt/simplex/geom/point"
)

//segments in range
//xor - altenate segments if nothing is in range of box
func (self *LineString) segs_inrange(seglist []*Segment, box *MBR, i, j int, extend, xor bool) []*Segment {
    //extend or refresh list
    if !extend {
        seglist = make([]*Segment, 0)
    }

    altsegs := make([]*Segment, 0)//, bool, seg
    for ; i < j; i++ {
        inters := box.IntersectsBounds(self.coordinates[i], self.coordinates[i + 1])
        var seg = &Segment{self.coordinates[i], self.coordinates[i + 1]}
        if inters {
            seglist = append(seglist, seg)
        } else {
            altsegs = append(altsegs, seg)
        }
    }
    if xor && len(seglist) == 0 {
        seglist = append(seglist, altsegs...)
    }
    return seglist
}

/*
 description segment ptlist
 param segsa{[]}
 param segsb{[]}
 param ptlist{[]}
 param [append]{boolean}
 returns {Array}
 private
 */
func (self *LineString) segseg_intersection(segsa, segsb []*Segment, ptlist []*Point, extend bool) []*Point {

    if !extend {
        ptlist = make([]*Point, 0)
    }
    for a := 0; a < len(segsa); a++ {
        for b := 0; b < len(segsb); b++ {
            coord, ok := segsa[a].Intersection(segsb[b], false)
            if !ok {
                continue
            }
            for _, pt := range coord {
                if !contains_point(ptlist, pt) {
                    ptlist = append(ptlist, coord...)
                }
            }
        }
    }
    return ptlist
}



//linear search if point is a member of list of points
func contains_point(coords []*Point, pt *Point) bool {
    bln := false
    n := len(coords)
    for i := 0; !bln && i < n; i++ {
        bln = pt.Equals(coords[i])
    }
    return bln
}


