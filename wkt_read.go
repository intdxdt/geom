package geom

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
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

//wkt type and coordiantes
func (wktreg *wktRegex) wkt_type_coords(wkt string) map[string]string {
	wkt = strings.TrimSpace(wkt)
	var captures = make(map[string]string)
	captures["wkt"], captures["type"], captures["coords"] = "", "", ""

	if is_empty_wkt(wkt) {
		wktreg = &re_emptyTypeStr
	}
	var match = wktreg.FindStringSubmatch(wkt)
	if match != nil {
		for i, name := range wktreg.SubexpNames() {
			if i == 0 || name == "" {
				if i == 0 {
					captures["wkt"] = match[i]
				}
				continue
			}
			var val = match[i]
			if name == "type" {
				val = strings.ToLower(val)
			}
			captures[name] = val
		}
	}
	return captures
}

type Shell [][]float64
type Holes []Shell

type WKTParserObj struct {
	shell Shell
	holes Holes
	gtype int
}

//Shell
func (self *WKTParserObj) Shell() Shell {
	return self.shell
}

//Holes
func (self *WKTParserObj) Holes() Holes {
	return self.holes
}

//Geometry Type
func (self *WKTParserObj) GeometryType() int {
	return self.gtype
}

//To array of coodinates of wkt string
func (self *WKTParserObj) ToArray() [][][]float64 {
	coords := make([][][]float64, 0)
	sh := self.shell
	if self.gtype == GeoType_Point || self.gtype == GeoType_LineString {
		coords = append(coords, sh)
	} else if self.gtype == GeoType_Polygon {
		coords = append(coords, sh)
		for _, sh = range self.holes {
			coords = append(coords, sh)
		}
	}
	return coords
}

//New WKT parser object
func NewWKTParserObj(gtype int, coords ...[][]float64) *WKTParserObj {
	var shells = make([]Shell, 0, len(coords))
	for i := range coords {
		shells = append(shells, Shell(coords[i]))
	}
	var obj *WKTParserObj
	if len(shells) == 1 {
		obj = &WKTParserObj{
			shells[0],
			nil,
			gtype,
		}
	} else {
		obj = &WKTParserObj{
			shells[0],
			Holes(shells[1:]),
			gtype,
		}
	}
	return obj
}

//Read wkt string
func ReadWKT(wkt string) *WKTParserObj {
	wkt = wkt_string(wkt)
	var parser func(string, *WKTParserObj)
	var matches = re_typeStr.wkt_type_coords(wkt)
	var obj = &WKTParserObj{nil, nil, GeoType_Unkown}

	mtype, coords := matches["type"], matches["coords"]

	if mtype == "polygon" {
		obj.gtype, parser = GeoType_Polygon, wkt_polygon_parser
	} else if mtype == "linestring" {
		obj.gtype, parser = GeoType_LineString, wkt_linestring_parser
	} else if mtype == "point" {
		obj.gtype, parser = GeoType_Point, wkt_point_parser
	}

	if coords != "" && obj.gtype != GeoType_Unkown {
		parser(coords, obj)
	}
	return obj
}

//Read wkt as geometry
func ReadGeometry(wkt string) Geometry {
	var g Geometry
	var obj = ReadWKT(wkt)

	if obj.gtype == GeoType_Polygon {
		var pts [][]Point
		for _, v := range obj.ToArray() {
			pts = append(pts, AsPointArray(v))
		}
		g = NewPolygon(pts...)
	} else if obj.gtype == GeoType_LineString {
		g = NewLineStringFromArray(obj.ToArray()[0])
	} else if obj.gtype == GeoType_Point {
		var pt = CreatePoint(obj.ToArray()[0][0][:])
		g = &pt
	}

	return g
}

//wkt string
func wkt_string(wkt string) string {
	var buffer bytes.Buffer
	tokens := strings.Split(wkt, "\n")
	for _, token := range tokens {
		buffer.WriteString(strings.TrimSpace(token))
	}
	return buffer.String()
}

//Parse point
func wkt_point_parser(wkt_coords string, obj *WKTParserObj) {
	//var coords = str.trim().split(this.regExes.spaces)
	var coords = strings.TrimSpace(wkt_coords)
	var coord = re_spaces.Split(coords, -1)
	var pt = []float64{
		wkt_parse_float(coord[X]),
		wkt_parse_float(coord[Y]),
	}
	obj.shell, obj.holes = Shell{pt}, nil
}

//parse linestring
func wkt_linestring_parser(wkt_coords string, obj *WKTParserObj) {
	var coords = strings.TrimSpace(wkt_coords)
	var shell = wkt_string_coords(coords)
	obj.shell, obj.holes = shell, nil
}

//parse polygon
func wkt_polygon_parser(wkt_coords string, obj *WKTParserObj) {
	var coords = strings.TrimSpace(wkt_coords)
	var rings = re_parenComma.Split(coords, -1)
	var n = len(rings)
	var holes = make(Holes, n-1)
	var shell Shell

	for i := 0; i < n; i++ {
		ring := re_trimParens.ReplaceAllString(rings[i], "$1")
		comps := wkt_string_coords(ring)
		if i == 0 {
			shell = comps
		}
		if i > 0 {
			holes[i-1] = comps
		}
	}
	obj.shell, obj.holes = shell, holes
}

//string coords
func wkt_string_coords(str string) Shell {
	var points = strings.Split(strings.TrimSpace(str), ",")
	var n = len(points)
	var comps = make(Shell, n)

	for i := 0; i < n; i++ {
		coords := re_spaces.Split(strings.TrimSpace(points[i]), -1)
		pt := []float64{
			wkt_parse_float(coords[X]),
			wkt_parse_float(coords[Y]),
		}
		comps[i] = pt
	}
	return comps
}

//parse float
func wkt_parse_float(str string) float64 {
	x, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic("unable to convert to float")
	}
	return x
}

//checks for the emptiness of wkt string
func is_empty_wkt(wkt string) bool {
	return strings.Index(wkt, "EMPTY") != -1
}
