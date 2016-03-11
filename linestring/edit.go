package linestring

import (
    p "github.com/intdxdt/simplex/geom/point"
//"github.com/intdxdt/simplex/geom/mbr"
    "github.com/intdxdt/simplex/struct/rtree"
//"math"
)

/*
 description append coordinate, by extending line
 param pnt{Point}
 */
func (self *LineString) Append(pnt *p.Point) *LineString {
    //chain index
    var idx = len(self.chains) - 1
    var chain *MonoMBR
    chain, self.chains = pop_mono_mbr(self.chains)
    //remove chain from index
    node := self.find_chain(chain)
    self.index.Remove(node)
    //subtract length of newly poped chain
    self.length -= self.chain_length(chain)
    var coord = pnt.Clone()
    //push coord
    self.coordinates = append(self.coordinates, coord)
    i := chain.i
    j := len(self.coordinates) - 1
    self.process_chains(i, j)
    //add newly pushed chains to index
    for ; idx < len(self.chains); idx++ {
        self.index.Insert(self.chains[idx])
    }
    self.update_rootmbr()
    return self
}

func (self *LineString) find_chain(ch *MonoMBR) *rtree.Node {
    res := self.index.Search(ch.MBR)
    var node *rtree.Node
    bbox := &ch.MBR
    node = nil
    if len(res) >= 1 {
        for i := 0; i < len(res); i++ {
            if res[i].BBox().Equals(bbox) {
                node = res[i]
                break
            }
        }
    }
    return node
}


// update root mbr
func (self *LineString)  update_rootmbr() {
    self.bbox = self.chains[0].Clone()
    self.bbox.i = 0
    self.bbox.j = len(self.coordinates) - 1
    for i := 1; i < len(self.chains); i++ {
        self.bbox.MBR.ExpandIncludeMBR(&self.chains[i].MBR)
    }
}
