package geom

import (
	"bytes"
	"strconv"
	"strings"
	"sort"
)

var wktEmpty = []byte("empty")
var wktPolygon = []byte("polygon")
var wktLinestring = []byte("linestring")
var wktPoint = []byte("point")


type WKTParserObj struct {
	shell Coords
	holes []Coords
	gtype GeoType
}

//Coords
func (self *WKTParserObj) Shell() Coords {
	return self.shell
}

//Holes
func (self *WKTParserObj) Holes() []Coords {
	return self.holes
}

//Geometry Type
func (self *WKTParserObj) GeometryType() GeoType {
	return self.gtype
}

//To array of coodinates of wkt string
func (self *WKTParserObj) ToCoordinates() []Coords {
	var shells = make([]Coords, 0)
	var sh = self.shell
	if self.gtype == GeoTypePoint || self.gtype == GeoTypeLineString {
		shells = append(shells, sh)
	} else if self.gtype == GeoTypePolygon {
		shells = append(shells, sh)
		for _, sh = range self.holes {
			shells = append(shells, sh)
		}
	}
	return shells
}

//New WKT parser object
func NewWKTParserObj(gtype GeoType, coords ...Coords) *WKTParserObj {
	var shells = make([]Coords, 0, len(coords))
	for i := range coords {
		shells = append(shells, Coords(coords[i]))
	}

	var obj *WKTParserObj
	if len(shells) == 1 {
		obj = &WKTParserObj{shells[0], nil, gtype}
	} else {
		obj = &WKTParserObj{shells[0], shells[1:], gtype}
	}
	return obj
}

//Read wkt string
func ReadWKT(wkt string, typeId GeoType) *WKTParserObj {
	var wktBytes = bytes.ToLower([]byte(wkt_string(wkt)))
	if isEmptyWKT(wktBytes) {
		return &WKTParserObj{gtype: typeId}
	}

	var tokens = aggregateTokens(buildTokens(wktBytes))
	var obj = &WKTParserObj{gtype: GeoTypeUnknown}
	if typeId == GeoTypeUnknown {
		return obj
	}

	if typeId == GeoTypePolygon {
		obj = wktPolygonParser(typeId, wktBytes, tokens[0])
	} else if typeId == GeoTypeLineString {
		obj = wktLinestringParser(typeId, wktBytes, tokens[0])
	} else if typeId == GeoTypePoint {
		obj = wktPointParser(typeId, wktBytes, tokens[0])
	}

	return obj
}

func buildTokens(stream []byte) []*wktToken {
	var tokens []*wktToken
	var stack []*wktToken
	var s *wktToken

	for i, o := range stream {
		if o == '(' {
			stack = append(stack, &wktToken{i: i})
		} else if o == ')' {
			s = popToken(&stack)
			s.j = i
			tokens = append(tokens, s)
		}
	}
	return tokens
}

func WKTType(stream string) []byte {
	if strings.Index(stream, "empty") != -1 {
		var gtype = "unknown"
		var subs = strings.Split(stream, "empty")
		if len(subs) > 1 {
			gtype = strings.TrimSpace(subs[0])
		}
		return []byte(gtype)
	}

	var char byte
	var name = make([]byte, 0, 16)

	for i := range stream {
		char = stream[i]
		if char == '(' {
			break
		}

		if char != ' ' {
			name = append(name, stream[i])
		}
	}

	return bytes.ToLower(name)
}

func aggregateTokens(tokens []*wktToken) []*wktToken {
	//var bln = false
	//bln, tokens = aggregateSequence(tokens)
	_, tokens = aggregateSequence(tokens)
	//useful for muti-wkt types
	//if bln && len(tokens) > 1 {
	//	for _, tok := range tokens {
	//		_, tok.children = aggregateSequence(tok.children)
	//	}
	//}
	return tokens
}

//Aggregate sequence - single wkt type - Point, LineString, Polygon
//Does not support Multi-WKT types
func aggregateSequence(tokens []*wktToken) (bool, []*wktToken) {
	if len(tokens) <= 1 {
		return false, tokens
	}

	sort.Sort(wktTokens(tokens))
	var head *wktToken
	var heads []*wktToken
	var aggregate = false

	for _, tok := range tokens {
		if head == nil {
			head = tok
			heads = append(heads, head)
		} else {
			if head.i < tok.i && tok.j < head.j {
				head.children = append(head.children, tok)
				aggregate = true
			}
			//use this branch for multi-wkt type
			//else {head = tok;heads = append(heads, head)}
		}
	}
	return aggregate, heads
}

//wkt string
func wkt_string(wkt string) string {
	var buffer bytes.Buffer
	var tokens = strings.Split(wkt, "\n")
	for _, token := range tokens {
		buffer.WriteString(strings.TrimSpace(token))
	}
	return buffer.String()
}

//checks for the emptiness of wkt string
func isEmptyWKT(wkt []byte) bool {
	return bytes.Index(wkt, wktEmpty) != -1
}

//parse float
func parseF64(str []byte) float64 {
	var x, err = strconv.ParseFloat(string(str), 64)
	if err != nil {
		panic("unable to convert to float")
	}
	return x
}

//Parse point
func wktPointParser(typeId GeoType, wkt []byte, tok *wktToken) *WKTParserObj {
	var wktStr = wkt[tok.i+1 : tok.j]
	var indices = numberIndices(wktStr)
	var dim = len(indices) / 2
	var lns = parseNums(wktStr, indices)
	var  pts = make([]Point, 0, len(lns)/dim)
	for i := 0; i < len(lns); i += dim {
		pts = append(pts, CreatePoint(lns[i:i+dim]))
	}
	return &WKTParserObj{gtype: typeId, shell: Coords(Coordinates(pts))}
}

//parse linestring
func wktLinestringParser(typeId GeoType, wkt []byte, tok *wktToken) *WKTParserObj {
	return &WKTParserObj{
		gtype: typeId,
		shell: parseString(wkt, tok),
	}
}

//parse polygon
func wktPolygonParser(typeId GeoType, wkt []byte, token *wktToken) *WKTParserObj {
	var shell Coords
	var obj = &WKTParserObj{gtype: typeId}
	var n = len(token.children)
	var holes = make([]Coords, 0, n-1)

	for i, tok := range token.children {
		if i == 0 {
			shell = parseString(wkt, tok)
		} else {
			holes = append(holes, parseString(wkt, tok))
		}
	}
	obj.shell, obj.holes = shell, holes
	return obj
}

//parse linestring
func parseString(wkt []byte, tok *wktToken) Coords {
	var wktStr = wkt[tok.i+1 : tok.j]
	var indices = numberIndices(wktStr)
	var dim = dimension(wktStr)
	var lns = parseNums(wktStr, indices)
	var pts = make([]Point, 0, len(lns)/dim)

	for i := 0; i < len(lns); i += dim {
		pts = append(pts, CreatePoint(lns[i:i+dim]))
	}
	return  Coords(Coordinates(pts))
}

func numberIndices(stream []byte) []int {
	var indices []int
	var idx, i, j = -1, -1, -1
	var n = len(stream)

	for idx < n {
		idx++
		if idx >= n {
			break
		}
		if stream[idx] == ' ' || stream[idx] == ',' {
			continue
		}
		if i == -1 {
			i, j = idx, idx
			for j < n && !(stream[j] == ' ' || stream[j] == ',') {
				j++
			}
			indices = append(indices, i, j)
			idx, i, j = j, -1, -1
		}
	}
	return indices
}

func parseNums(strBytes []byte, indices []int) []float64 {
	var coordinates = make([]float64, 0, len(indices)/2)
	for i := 0; i < len(indices); i += 2 {
		coordinates = append(coordinates, parseF64(strBytes[indices[i]:indices[i+1]]))
	}
	return coordinates
}

//Detects dimension 2&3 in wkt string
func dimension(stream []byte) int {
	var idx, i, j = -1, -1, -1
	var dim, n = 1, len(stream)
	for idx < n {
		idx++
		if idx >= n {
			break
		}

		//fmt.Println(string(stream[idx]))
		if stream[idx] == ' ' || stream[idx] == ',' {
			continue
		}

		if i == -1 {
			i, j = idx, idx
			for j < n && !(stream[j] == ' ' || stream[j] == ',') {
				j++
			}
			for j < n && stream[j] == ' ' {
				j++
			}

			if j >= n || stream[j] == ',' {
				break
			} else {
				j--
			}

			dim++
			idx, i, j = j, -1, -1
		}
	}
	if dim == 1 {
		dim = -1
	}
	return dim
}

func popToken(tokens *[]*wktToken) *wktToken {
	var v *wktToken
	var a = *tokens
	var n = len(a) - 1
	if n < 0 {
		return nil
	}
	v, a[n] = a[n], nil
	*tokens = a[:n]
	return v
}
