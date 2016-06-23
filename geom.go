package geom

import (
    . "simplex/geom/mbr"
)

const (
    x = iota
    y
    z
    null = -9
)

const (
    GeoType_Unkown = iota - 1
    GeoType_Point
    GeoType_LineString
    GeoType_Polygon
)

type Geometry interface {
    BBox() *MBR
    AsLinear() []*LineString
    Intersects(Geometry) bool
    Distance(Geometry) float64
    Type() *geoType
    WKT() string 
}

type geoType struct {
    gtype int
}

//New Side
func new_geoType(gtype int) *geoType {
    return &geoType{gtype}
}

//is polygon
func (self *geoType) IsPolygon() bool {
    return self.gtype == GeoType_Polygon
}

//is linestring
func (self *geoType) IsLineString() bool {
    return self.gtype == GeoType_LineString
}

//is point
func (self *geoType) IsPoint() bool {
    return self.gtype == GeoType_Point
}


