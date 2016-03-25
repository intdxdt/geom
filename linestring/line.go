package linestring

import (
    . "github.com/intdxdt/simplex/geom/point"
    . "github.com/intdxdt/simplex/geom/mbr"
    "github.com/intdxdt/simplex/struct/rtree"
    "math"
    "fmt"
)

const (
    x = iota
    y
    null = -9
)

type LineString struct {
    chains      []*MonoMBR
    coordinates []*Point
    monosize    int
    bucketsize  int
    index       *rtree.RTree
    bbox        *MonoMBR
    length      float64
}

//New LineString from a given coordinates {Array} [[x,y], ....[x,y]]
func NewLineString(coordinates []*Point) *LineString {
    if len(coordinates) < 2 {
        panic("a linestring must have at least 2 coordinate")
    }
    self := &LineString{}
    self.chains = make([]*MonoMBR, 0)

    self.coordinates = make([]*Point, len(coordinates))
    copy(self.coordinates, coordinates)

    //init
    self.monosize = int(math.Log2(float64(len(coordinates)) + 1.0))
    self.bucketsize = 9
    self.index = rtree.NewRTree(self.bucketsize)

    self.process_chains(0, 0)
    self.build_index()
    return self
}

//clone linestring
func (self *LineString) Clone() *LineString {
    return NewLineString(self.coordinates)
}

//envelope of linestring
func (self *LineString) Envelope() *MBR {
    return self.bbox.MBR
}

//print chains
func (self *LineString)  PrintChains() {
    for i := 0; i < len(self.chains); i++ {
        //fmt.Println(self.chains[i].String())
        fmt.Printf("%p\n", self.chains[i])
    }
}

