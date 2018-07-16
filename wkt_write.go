package geom

import (
	"strings"
)

//write wkt
func WriteWKT(obj *WKTParserObj) string {
	var s string
	if obj.gtype == GeoTypePoint {
		s = "POINT " + str_point(obj.shell)
	} else if obj.gtype == GeoTypeLineString {
		s = "LINESTRING " + str_polyline(obj.shell)
	} else if obj.gtype == GeoTypePolygon {
		var wkt = str_polygon(obj)
		if wkt == "EMPTY" {
			s = "POLYGON " + wkt
		} else {
			s = "POLYGON (" + wkt + ")"
		}
	}
	return s
}

//str point
func str_point(shell Shell) string {
	var s = "EMPTY"
	if shell != nil && len(shell) > 0 {
		s = "(" + coord_str(shell[0]) + ")"
	}
	return s
}

//str polyline
func str_polyline(shell Shell) string {
	var s = "EMPTY"
	if shell == nil {
		return s
	}

	var n = len(shell)
	if n > 0 {
		var lnstr = make([]string, n)
		for i := 0; i < n; i++ {
			lnstr[i] = coord_str(shell[i])
		}
		s = "(" + strings.Join(lnstr, ", ") + ")"
	}
	return s
}

//str polygon
func str_polygon(obj *WKTParserObj) string {
	var n int
	var shell = str_polyline(obj.shell)
	if len(obj.holes) > 0 {
		n = len(obj.holes)
	}
	var rings = make([]string, n+1)
	rings[0] = shell
	if n > 0 {
		for i := 0; i < n; i++ {
			rings[i+1] = str_polyline(obj.holes[i])
		}
	}
	return strings.Join(rings, ",")
}
