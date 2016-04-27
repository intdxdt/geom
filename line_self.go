package geom

import (
    . "github.com/intdxdt/simplex/geom/mbr"
    . "github.com/intdxdt/simplex/struct/sset"
    "github.com/intdxdt/simplex/struct/item"
    "strconv"
)

//computes whether linestring is simple
func (self *LineString) IsSimple() bool{
    var cache   = make(map[string]bool)
    var bln     = true//, bcomplx, chain, inters, jbox, qbox
    var ln1     = make([]*Segment, 0)
    var ln2     = make([]*Segment, 0)
    var ptlist  = make([]*Point, 0)
    var query *MBR

    for i := 0; bln && i < len(self.chains); i++ {
        chain := self.chains[i]
        query = chain.BBox()
        inters := self.index.Search(query)

        for j := 0; bln && j < len(inters); j++ {
            jbox := (*inters[j].GetItem()).(*MonoMBR)

            ckey := self.cashe_key(chain, jbox)

            if cache[ckey] || jbox.MBR.Equals(chain.MBR) {
                continue//already checked || already monotone
            }

            self.cashe_ij(cache, chain, jbox, true)
            qbox, ok := chain.MBR.Intersection(jbox.MBR)

            if ok == false && chain.j == jbox.i {
                continue //non overlapping && contiguous
            }

            ln1 = self.segs_inrange(
                ln1, qbox, chain.i, chain.j, false, false,
            )
            ln2 = self.segs_inrange(
                ln2, qbox, jbox.i, jbox.j, false, false,
            )
            ptlist = self.segseg_intersection(
                ln1, ln2, ptlist, false,
            )

            bcomplx := (
                (chain.j != jbox.i && len(ptlist) > 0) ||
                (chain.j == jbox.i && len(ptlist) > 1))
            if bcomplx {
                bln = false
            }
        }
    }
    return bln
}



//cache box ij keys
func (self *LineString) cashe_ij(cashe map[string]bool, box1, box2 *MonoMBR, rev bool) {
    var a = strconv.Itoa(box1.i) + "_" + strconv.Itoa(box1.j)
    var b = strconv.Itoa(box2.i) + "_" + strconv.Itoa(box2.j)

    cashe[a + "-" + b] = true
    if rev {
        cashe[b + "-" + a] = true
    }
}


//cache key
func (self *LineString)  cashe_key(box1, box2 *MonoMBR) string {
    var a = strconv.Itoa(box1.i) + "_" + strconv.Itoa(box1.j)
    var b = strconv.Itoa(box2.i) + "_" + strconv.Itoa(box2.j)

    return a + "-" + b
}


// self intersection coordinates
func (self *LineString) SelfIntersection() []*Point {

    var ckey string
    var cache = make(map[string]bool)
    var ln1 = make([]*Segment, 0)
    var ln2 = make([]*Segment, 0)
    var ptlist = make([]*Point, 0)

    var query *MBR
    var chain *MonoMBR
    var bcomplx bool

    var selfinters = NewSSet()

    for i := 0; i < len(self.chains); i++ {
        chain = self.chains[i]
        query = chain.BBox()
        inters := self.index.Search(query)

        for j := 0; j < len(inters); j++ {
            jbox := (*inters[j].GetItem()).(*MonoMBR)
            ckey = self.cashe_key(chain, jbox)

            if cache[ckey] || jbox.MBR.Equals(chain.MBR) {
                continue//already checked || already monotone
            }

            self.cashe_ij(cache, chain, jbox, true)
            qbox, bln := chain.MBR.Intersection(jbox.MBR)
            if bln == false && chain.j == jbox.i {
                continue//non overlapping && contiguous
            }

            ln1=self.segs_inrange(ln1, qbox, chain.i, chain.j, false, false)
            ln2=self.segs_inrange(ln2, qbox, jbox.i, jbox.j, false, false)
            ptlist = self.segseg_intersection(ln1, ln2, ptlist, false)

            bcomplx = (chain.j != jbox.i && len(ptlist) > 0) ||
                (chain.j == jbox.i && len(ptlist) > 1)

            if bcomplx {
                for _, p := range ptlist {
                    selfinters.Add(p)
                }
            }
        }
    }

    var ptinters = make([]*Point, 0)
    selfinters.Each(func(o item.Item) {
        ptinters = append(ptinters, o.(*Point))
    })
    return ptinters
}






