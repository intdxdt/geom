package geom

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/mbr"
	"bytes"
	"strings"
)

var nan = math.NaN()

const (
	X    = iota
	Y
	Z
	null = -9
)

const (
	GeoTypeUnknown    = iota - 1
	GeoTypePoint
	GeoTypeSegment
	GeoTypeLineString
	GeoTypePolygon
)

var feq = math.FloatEqual

//geometry constructor
type GeometryFn func([]*Point) Geometry

type Geometry interface {
	BBox() *mbr.MBR
	AsLinear() []*LineString
	Intersects(Geometry) bool
	Intersection(Geometry) []Point
	Distance(Geometry) float64
	Type() *geoType
	WKT() string
}

type geoType struct {
	gtype int
}

//New geometry
func NewGeometry(wkt string) Geometry {
	var g Geometry
	wkt = strings.ToLower(wkt_string(wkt))
	var gtype = wktType(wkt)

	if bytes.Equal(gtype, wktPolygon) {
		g = NewPolygonFromWKT(wkt)
	} else if bytes.Equal(gtype, wktPoint) {
		var pt = PointFromWKT(wkt)
		g = &pt
	} else if bytes.Equal(gtype, wktLinestring) {
		g = NewLineStringFromWKT(wkt)
	}
	return g
}

//Read wkt as geometry
//func ReadGeometry(wkt string) Geometry {
//	var g Geometry
//	var typeId = wktType([]byte(wkt))
//	var obj = readWKT(wkt, typeId)
//
//	if obj.gtype == GeoTypePolygon {
//		var pts [][]Point
//		for _, v := range obj.ToArray() {
//			pts = append(pts, AsPointArray(v))
//		}
//		g = NewPolygon(pts...)
//	} else if obj.gtype == GeoTypeLineString {
//		g = NewLineStringFromArray(obj.ToArray()[0])
//	} else if obj.gtype == GeoTypePoint {
//		var pt = CreatePoint(obj.ToArray()[0][0][:])
//		g = &pt
//	}
//
//	return g
//}

//New geoType
func new_geoType(gtype int) *geoType {
	return &geoType{gtype}
}

//Value
func (gt *geoType) Value() int {
	return gt.gtype
}

//is polygon
func (gt *geoType) IsPolygon() bool {
	return gt.gtype == GeoTypePolygon
}

//is linestring
func (gt *geoType) IsLineString() bool {
	return gt.gtype == GeoTypeLineString
}

//is linestring
func (gt *geoType) IsSegment() bool {
	return gt.gtype == GeoTypeSegment
}

//is point
func (gt *geoType) IsPoint() bool {
	return gt.gtype == GeoTypePoint
}
