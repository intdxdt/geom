package geom

import (
    "strings"
)

//write wkt
func WriteWKT(obj *WKTParserObj) string {
    var s string
    if obj.gtype == GeoType_Point {
        s = "POINT " + str_point(obj.shell)
    } else if obj.gtype == GeoType_LineString {
        s = "LINESTRING " + str_polyline(obj.shell)
    } else if obj.gtype == GeoType_Polygon {
        wkt := str_polygon(obj)
        if is_empty_wkt(wkt) {
            s = "POLYGON " + wkt
        } else {
            s = "POLYGON (" + wkt + ")"
        }
    }
    return s
}

//str point
func str_point(shell *Shell) string {
    var s string = "EMPTY"
    if shell != nil && len(*shell) > 0 {
        sh := *shell
        s = "(" + coord_str(&sh[0]) + ")"
    }
    return s
}

//str polyline
func str_polyline(shell *Shell) string {
    var s string = "EMPTY"
    if shell == nil {
        return s
    }
    var sh = *shell
    n := len(sh)
    if n > 0 {
        lnstr := make([]string, n)
        for i := 0; i < n; i++ {
            lnstr[i] = coord_str(&sh[i])
        }
        s = "(" +  strings.Join(lnstr, ", ") + ")"
    }
    return s
}

//str polygon
func str_polygon(obj *WKTParserObj) string {
    shell := str_polyline(obj.shell)
    var n int
    var holes Holes
    if obj.holes != nil {
        holes = *obj.holes
        n = len(holes)
    }
    rings := make([]string, n + 1)
    rings[0] = shell
    if n > 0 {
        for i := 0; i < n; i++ {
            rings[i + 1] = str_polyline(holes[i])
        }
    }
    return strings.Join(rings, ",")
}
