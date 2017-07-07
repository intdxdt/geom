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
	Intersection(Geometry) []*Point
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
func (gt *geoType) IsPolygon() bool {
	return gt.gtype == GeoType_Polygon
}

//is linestring
func (gt *geoType) IsLineString() bool {
	return gt.gtype == GeoType_LineString
}

//is point
func (gt *geoType) IsPoint() bool {
	return gt.gtype == GeoType_Point
}
