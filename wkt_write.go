package geom

import "strings"

//write wkt
func WriteWKT(obj *WKTParserObj) string {
	var s string
	if obj.gtype == GeoTypePoint {
		s = "POINT " + strPoint(obj.shell, coordStr)
	} else if obj.gtype == GeoTypeLineString {
		s = "LINESTRING " + strPolyline(obj.shell, coordStr)
	} else if obj.gtype == GeoTypePolygon {
		var wkt = strPolygon(obj, coordStr)
		if wkt == "EMPTY" {
			s = "POLYGON " + wkt
		} else {
			s = "POLYGON (" + wkt + ")"
		}
	}
	return s
}

//write wkt 3d
func WriteWKT3D(obj *WKTParserObj) string {
	var s string
	if obj.gtype == GeoTypePoint {
		s = "POINT " + strPoint(obj.shell, coordStr3D)
	} else if obj.gtype == GeoTypeLineString {
		s = "LINESTRING " + strPolyline(obj.shell, coordStr3D)
	} else if obj.gtype == GeoTypePolygon {
		var wkt = strPolygon(obj, coordStr3D)
		if wkt == "EMPTY" {
			s = "POLYGON " + wkt
		} else {
			s = "POLYGON (" + wkt + ")"
		}
	}
	return s
}

//str point
func strPoint(shell Coords, fnCoordStr func([]float64) string) string {
	var s = "EMPTY"
	if shell.Pnts != nil && len(shell.Pnts) > 0 {
		s = "(" + fnCoordStr(shell.Pnts[0][:]) + ")"
	}
	return s
}

//str polyline
func strPolyline(shell Coords, fnCoordStr func([]float64) string) string {
	var s = "EMPTY"
	if shell.Pnts == nil {
		return s
	}

	var n = shell.Len()
	if n > 0 {
		var lnstr = make([]string, n)
		for i := 0; i < n; i++ {
			lnstr[i] = fnCoordStr(shell.Pt(i)[:])
		}
		s = "(" + strings.Join(lnstr, ", ") + ")"
	}
	return s
}

//str polygon
func strPolygon(obj *WKTParserObj, fnCoordStr func([]float64) string) string {
	var n int
	var shell = strPolyline(obj.shell, fnCoordStr)
	if len(obj.holes) > 0 {
		n = len(obj.holes)
	}
	var rings = make([]string, n+1)
	rings[0] = shell
	if n > 0 {
		for i := 0; i < n; i++ {
			rings[i+1] = strPolyline(obj.holes[i], fnCoordStr)
		}
	}
	return strings.Join(rings, ",")
}
