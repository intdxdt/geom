package linestring

import (
    "github.com/intdxdt/simplex/geom/mbr"
    "github.com/intdxdt/simplex/geom/point"
)

//segments in range
//xor - altenate segments if nothing is in range of box
func (self *LineString) segs_inrange(seglist []*LineString, box *mbr.MBR, i, j int, extend, xor bool) []*LineString{

    //extend or refresh list
    if !extend {
        seglist = make([]*LineString, 0)
    }

    altsegs := make([]*LineString, 0)//, bool, seg
    for ; i < j; i++ {
        inters := box.IntersectsBounds(self.coordinates[i], self.coordinates[i + 1])
        coords := []*point.Point{self.coordinates[i], self.coordinates[i + 1]}
        var seg = NewLineString(coords)

        if inters {
            seglist = append(seglist, seg)
        } else {
            altsegs = append(altsegs, seg)
        }
    }
    if xor && len(seglist) == 0 {
        for i := range altsegs {
            seglist = append(seglist, altsegs[i])
        }
    }
    return seglist
}

