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
type GeometryFn func([]Point) Geometry

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

//Read geometry from WKT
func ReadGeometry(wkt string) Geometry {
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

//New geoType
func newGeoType(gtype int) *geoType {
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
