package geom

import . "github.com/intdxdt/simplex/geom/mbr"

const (
    x = iota
    y
    null = -9
)

const (
    Unkown = iota - 1
    GeoType_Point
    GeoType_LineString
    GeoType_Polygon
)

type Geometry interface {
    Envelope() *MBR
    AsLinear() []*LineString
    Intersects(Geometry) bool
}
