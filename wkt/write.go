package wkt

import (
    //"strconv"
    . "github.com/intdxdt/simplex/geom/point"
    "fmt"
    "strings"
)

func Write(obj *WKTParserObj) string {
    var s string
    if obj.gtype == "point" {
        s = fmt.Sprintf("POINT %s", str_point(obj.shell))
    }else if obj.gtype == "linestring" {
        s = fmt.Sprintf("LINESTRING %s", str_polyline(obj.shell))
    }else if obj.gtype == "polygon" {
        wkt := str_polygon(obj)
        if is_empty(wkt) {
            s = fmt.Sprintf("POLYGON %s", wkt)
        } else {
            s = fmt.Sprintf("POLYGON (%s)", wkt)
        }
    }
    return s
}

func coord(pt *Point) string {
    return fmt.Sprintf("%v %v", pt.X(), pt.Y())
}

func str_point(shell *Shell) string {
    var s string = "EMPTY"
    if shell != nil && len(*shell) > 0 {
        s = fmt.Sprintf("(%s)", coord((*shell)[0]))
    }
    return s
}

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
            lnstr[i] = coord(sh[i])
        }
        s = fmt.Sprintf("(%s)", strings.Join(lnstr, ", "))
    }
    return s
}

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
