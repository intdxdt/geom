package geom

import (
    "regexp"
    "strings"
    "strconv"
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


type Shell [][2]float64
type Holes []*Shell

type WKTParserObj struct {
    shell *Shell
    holes *Holes
    gtype int
}

func (self *WKTParserObj) GeometryType() int {
    return self.gtype
}

//To array of coodinates of wkt string
func (self *WKTParserObj) ToArray() [][][2]float64 {
    coords := make([][][2]float64, 0)
    sh := self.shell
    if self.gtype == GeoType_Point || self.gtype == GeoType_LineString {
        coords = append(coords, *sh)
    } else if (self.gtype == GeoType_Polygon) {
        coords = append(coords, *sh)
        for _, sh = range *self.holes {
            coords = append(coords, *sh)
        }
    }
    return coords
}


func NewWKTParserObj(gtype int, coords ...[][2]float64) *WKTParserObj {
    shells := make([]*Shell, len(coords))
    for i := 0; i < len(coords); i++ {
        sh := Shell(coords[i])
        shells[i] = &sh
    }
    //shell := shells[0]
    if len(shells) == 1 {
        return &WKTParserObj{shells[0], nil, gtype }
    }
    holes := Holes(shells[1:])
    return &WKTParserObj{shells[0], &holes, gtype}
}

func ReadWKT(wkt string) *WKTParserObj {
    var parser func(*string, *WKTParserObj)
    matches := re_typeStr.wkt_type_coords(wkt);
    obj := &WKTParserObj{nil, nil, Unkown}

    mtype, coords := *matches["type"], matches["coords"]

    if mtype == "polygon" {
        obj.gtype, parser = GeoType_Polygon, wkt_polygon_parser
    } else if mtype == "linestring" {
        obj.gtype, parser = GeoType_LineString, wkt_linestring_parser
    } else if mtype == "point" {
        obj.gtype, parser = GeoType_Point, wkt_point_parser
    }

    if coords != nil  && obj.gtype != Unkown {
        parser(coords, obj)
    }
    return obj
}

//parse point
func wkt_point_parser(wkt_coords *string, obj *WKTParserObj) {
    //var coords = str.trim().split(this.regExes.spaces)
    var coords = strings.TrimSpace(*wkt_coords)
    var coord = re_spaces.Split(coords, -1)
    pt := [2]float64{
        wkt_parse_float(coord[x]),
        wkt_parse_float(coord[y]),
    }
    obj.shell, obj.holes = &Shell{pt}, nil
}

//parse linestring
func wkt_linestring_parser(wkt_coords *string, obj *WKTParserObj) {
    var coords = strings.TrimSpace(*wkt_coords)
    shell := wkt_string_coords(&coords)
    obj.shell, obj.holes = shell, nil
}

//parse polygon
func wkt_polygon_parser(wkt_coords *string, obj *WKTParserObj) {
    var coords = strings.TrimSpace(*wkt_coords)
    var rings = re_parenComma.Split(coords, -1)

    var n = len(rings)
    holes := make(Holes, n - 1)

    var shell *Shell

    for i := 0; i < n; i++ {
        ring := re_trimParens.ReplaceAllString(rings[i], "$1")
        comps := wkt_string_coords(&ring)
        if i == 0 {
            shell = comps
        };
        if i > 0 {
            holes[i - 1] = comps
        };
    }
    obj.shell, obj.holes = shell, &holes
}

//string coords
func wkt_string_coords(str *string) *Shell {
    var points = strings.Split(strings.TrimSpace(*str), ",")
    var n = len(points)
    var comps = make(Shell, n)

    for i := 0; i < n; i++ {
        coords := re_spaces.Split(strings.TrimSpace(points[i]), -1)
        pt := [2]float64{
            wkt_parse_float(coords[x]),
            wkt_parse_float(coords[y]),
        }
        comps[i] = pt
    }
    return &comps
}

//parse float
func wkt_parse_float(str string) float64 {
    x, err := strconv.ParseFloat(str, 64)
    if err != nil {
        panic("unable to convert to float")
    }
    return x
}

//wkt type and coordiantes
func (self *wktRegex) wkt_type_coords(wkt string) map[string]*string {
    wkt = strings.TrimSpace(wkt)
    captures := make(map[string]*string)
    captures["wkt"], captures["type"], captures["coords"] = nil, nil, nil

    if is_empty_wkt(wkt) {
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

//checks for the emptiness of wkt string
func is_empty_wkt(wkt string) bool {
    return strings.Index(wkt, "EMPTY") != -1
}

