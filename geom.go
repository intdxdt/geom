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
	GeoTypeUnknown    GeoType = iota - 1
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
	Bounds() mbr.MBR
	AsLinear() []*LineString
	Intersects(Geometry) bool
	Intersection(Geometry) []Point
	Distance(Geometry) float64
	Type() GeoType
	WKT() string
	Geometry() Geometry
}

type GeoType int

//Read geometry from WKT
func ReadGeometry(wkt string) Geometry {
	wkt = strings.ToLower(wkt_string(wkt))
	var g Geometry
	var gtype = wktType(wkt)

	if bytes.Equal(gtype, wktPolygon) {
		g = NewPolygonFromWKT(wkt)
	} else if bytes.Equal(gtype, wktPoint) {
		g = PointFromWKT(wkt)
	} else if bytes.Equal(gtype, wktLinestring) {
		g = NewLineStringFromWKT(wkt)
	}
	return g
}

func CastAsPoint(other Geometry) Point {
	var pt, ok = other.(Point)
	if !ok {
		pt = *(other.(*Point))
	}
	return pt
}

func CastAsLineString(other Geometry) *LineString {
	return other.(*LineString)
}

func CastAsPolygon(other Geometry) *Polygon {
	return other.(*Polygon)
}

func CastAsSegment(other Geometry) *Segment {
	return other.(*Segment)
}

//is polygon
func (gt GeoType) IsPolygon() bool {
	return gt == GeoTypePolygon
}

//is linestring
func (gt GeoType) IsLineString() bool {
	return gt == GeoTypeLineString
}

//is linestring
func (gt GeoType) IsSegment() bool {
	return gt == GeoTypeSegment
}

//is point
func (gt GeoType) IsPoint() bool {
	return gt == GeoTypePoint
}
