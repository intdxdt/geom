package wkt

import (
    "regexp"
    "strings"
    "strconv"
    p "github.com/intdxdt/simplex/geom/point"
)

var re_typeStr = wktRegex{
    regexp.MustCompile(`^\s*(?P<type>[A-Za-z]+)\s*\(\s*(?P<coords>.*)\s*\)\s*$`),
}
var re_emptyTypeStr = wktRegex{
    regexp.MustCompile(`^\s*(?P<type>[A-Za-z]+)\s*EMPTY\s*$`),
}

var re_spaces = wktRegex{regexp.MustCompile(`\s+`)}
var re_parenComma = wktRegex{regexp.MustCompile(`\)\s*,\s*\(`)}
var re_trimParens = wktRegex{regexp.MustCompile(`^\s*\(?(.*?)\)?\s*$`)}

type wktRegex struct {
    *regexp.Regexp
}

const (
    x = iota
    y
)

type WKTParserObj struct {
    shell *[]p.Point
    holes *[][]p.Point
    gtype string
}

func Read(wkt string) *WKTParserObj {
    matches := re_typeStr.type_coords(wkt);
    obj := &WKTParserObj{nil, nil, ""}
    wkt = *matches["wkt"]
    gtype, coords := matches["type"], matches["coords"]

    if *gtype == "polygon" {
        if coords != nil {
            polygon(*coords, obj)
        }
    } else if *gtype == "linestring" {
        if coords != nil {
            linestring(*coords, obj)
        }
    } else if *gtype == "point" {
        if coords != nil {
            point(*coords, obj)
        }
    }
    obj.gtype = *gtype

    return obj
}

//parse point
func point(wkt_coords string, obj *WKTParserObj) {
    //var coords = str.trim().split(this.regExes.spaces)
    var coords = strings.TrimSpace(wkt_coords)
    var coord = re_spaces.Split(coords, -1)
    pt := p.Point{
        parse_float(coord[x]),
        parse_float(coord[y]),
    }
    var shell = []p.Point{pt}
    obj.shell, obj.holes = &shell, nil
}

//parse linestring
func linestring(wkt_coords string, obj *WKTParserObj) {
    var coords = strings.TrimSpace(wkt_coords)
    shell := string_coords(coords)
    obj.shell, obj.holes = &shell, nil
}

//parse polygon
func polygon(wkt_coords string, obj *WKTParserObj) {
    var coords = strings.TrimSpace(wkt_coords)
    var rings = re_parenComma.Split(coords, -1)

    var n = len(rings)
    holes := make([][]p.Point, n - 1)

    var shell []p.Point

    for i := 0; i < n; i++ {
        ring := re_trimParens.ReplaceAllString(rings[i], "$1")
        components := string_coords(ring)
        if i == 0 {
            shell = components
        };
        if i > 0 {
            holes[i - 1] = components
        };
    }
    obj.shell, obj.holes = &shell, &holes
}

//string coords
func string_coords(str string) []p.Point {
    var points = strings.Split(strings.TrimSpace(str), ",")
    var n = len(points)
    var components = make([]p.Point, n)

    for i := 0; i < n; i++ {
        coords := re_spaces.Split(strings.TrimSpace(points[i]), -1)
        components[i][x] = parse_float(coords[x])
        components[i][y] = parse_float(coords[y])
    }
    return components;
}

//parse float
func parse_float(str string) float64 {
    x, err := strconv.ParseFloat(str, 64)
    if err != nil {
        panic("unable to convert to float")
    }
    return x
}

//wkt type and coordiantes
func (self *wktRegex) type_coords(wkt string) map[string]*string {
    wkt = strings.TrimSpace(wkt)
    captures := make(map[string]*string)
    captures["wkt"], captures["type"], captures["coords"] = nil, nil, nil

    if (strings.Index(wkt, "EMPTY") != -1) {
        self = &re_emptyTypeStr
    }
    match := self.FindStringSubmatch(wkt)
    if match != nil {
        for i, name := range self.SubexpNames() {
            if i == 0 || name == "" {
                if i == 0 {
                    captures["wkt"] = &match[i]
                }
                continue
            }
            val := match[i]
            if name == "type" {
                val = strings.ToLower(val)
            }
            captures[name] = &val
        }
    }
    return captures
}


