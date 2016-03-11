package linestring

import (
    "github.com/intdxdt/simplex/geom/point"
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
    coordinates []*point.Point
    monosize    int
    bucketsize  int
    index       *rtree.RTree
    bbox        *MonoMBR
    length      float64
}

//New LineString from a given coordinates {Array} [[x,y], ....[x,y]]
func NewLineString(coordinates []*point.Point) *LineString {
    self := &LineString{}
    self.chains = make([]*MonoMBR, 0)

    if len(coordinates) == 1 {
        coordinates = append(coordinates, coordinates[0].Clone())
    }

    self.coordinates = make([]*point.Point, len(coordinates))
    copy(self.coordinates, coordinates)

    if len(coordinates) == 0 {
        //at least a segment a ring p1----p2----p1
        panic("a linestring must have at least 1 coordinate")
    }

    //init
    self.monosize = int(math.Log2(float64(len(coordinates)) + 1.0))
    self.bucketsize = 9
    self.index = rtree.NewRTree(self.bucketsize)

    self.process_chains(0, 0)
    self.build_index()
    return self
}

//Clone linestring
func (self *LineString) Clone() *LineString {
    return NewLineString(self.coordinates)

}

//envelope of linestring
func (self *LineString) Envelope() string {
    return self.bbox.MBR.String()
}

func (self *LineString)  PrintChains() {
    for i := 0; i < len(self.chains); i++ {
        fmt.Println(self.chains[i].String())
    }
}