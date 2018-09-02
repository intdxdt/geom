package geom

import (
	"bytes"
	"github.com/intdxdt/mbr"
	"strings"
)


const (
	GeoTypeUnknown GeoType = iota - 1
	GeoTypePoint
	GeoTypeSegment
	GeoTypeLineString
	GeoTypePolygon
)

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

// Read geometry from WKT
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
	return other.Geometry().(*LineString)
}

func CastAsPolygon(other Geometry) *Polygon {
	return other.Geometry().(*Polygon)
}

func CastAsSegment(other Geometry) *Segment {
	return other.Geometry().(*Segment)
}

// is polygon
func (gt GeoType) IsPolygon() bool {
	return gt == GeoTypePolygon
}

// is linestring
func (gt GeoType) IsLineString() bool {
	return gt == GeoTypeLineString
}

// is linestring
func (gt GeoType) IsSegment() bool {
	return gt == GeoTypeSegment
}

// is point
func (gt GeoType) IsPoint() bool {
	return gt == GeoTypePoint
}
