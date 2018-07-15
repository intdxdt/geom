package geom

import (
	"bytes"
)

const (
	minx = iota
	miny
	maxx
	maxy
)

func QuadrantRelation(g, other Geometry) string {
	//Expand oject intersections with extended bouding box intersections
	var buffer bytes.Buffer
	for _, q := range quadrants(g, other) {
		if g.Intersects(q) {
			buffer.WriteString("T")
		} else {
			buffer.WriteString("F")
		}
	}
	return buffer.String()
}

func quadrants(g, other Geometry) []Geometry {
	var nw, nn, ne, ww, ii, ee, sw, ss, se Geometry
	var box = g.BBox()
	var other_mbr = other.BBox()

	box.ExpandIncludeMBR(other_mbr)
	box.ExpandByDelta(1e2, 1e2)

	var mat [][]Point
	var xs = []float64{box[minx], other_mbr[minx], other_mbr[maxx], box[maxx]}
	var ys = []float64{box[miny], other_mbr[miny], other_mbr[maxy], box[maxy]}

	for j := 0; j < len(ys); j++ {
		var row_mat []Point
		for i := 0; i < len(xs); i++ {
			row_mat = append(row_mat, PointXY(xs[i], ys[j]))
		}
		mat = append(mat, row_mat)
	}

	/*
	       .(3,0).|.(3,1).|.(3,2).|.(3,3).
	              nw      nn      ne
	       .(2,0).|.(2,1).|.(2,2).|.(2,3).
	              ww      ii      ee
	       .(1,0).|.(1,1).|.(1,2).|.(1,3).
	              sw      ss      se
	       .(0,0).|.(0,1).|.(0,2).|.(0,3).
	   //TODO: ii can be improved by changing ii to convex hull
	*/

	var _, ispoint = other.(*Point)

	nw = NewPolygon([]Point{mat[2][0], mat[3][0], mat[3][1], mat[2][1], mat[2][0]})
	ne = NewPolygon([]Point{mat[2][2], mat[3][2], mat[3][3], mat[2][3], mat[2][2]})
	sw = NewPolygon([]Point{mat[0][0], mat[1][0], mat[1][1], mat[0][1], mat[0][0]})
	se = NewPolygon([]Point{mat[0][2], mat[1][2], mat[1][3], mat[0][3], mat[0][2]})

	if ispoint {
		nn = NewLineString([]Point{mat[2][1], mat[3][1]})
		ww = NewLineString([]Point{mat[1][0], mat[1][1]})
		ii = &mat[1][1]
		ee = NewLineString([]Point{mat[1][2], mat[1][3]})
		ss = NewLineString([]Point{mat[0][1], mat[1][1]})
	} else {
		nn = NewPolygon([]Point{mat[2][1], mat[3][1], mat[3][2], mat[2][2], mat[2][1]})
		ww = NewPolygon([]Point{mat[1][0], mat[2][0], mat[2][1], mat[1][1], mat[1][0]})
		ii = NewPolygon([]Point{mat[1][1], mat[2][1], mat[2][2], mat[1][2], mat[1][1]})
		ee = NewPolygon([]Point{mat[1][2], mat[2][2], mat[2][3], mat[1][3], mat[1][2]})
		ss = NewPolygon([]Point{mat[0][1], mat[1][1], mat[1][2], mat[0][2], mat[0][1]})
	}

	return []Geometry{
		nw, nn, ne,
		ww, ii, ee,
		sw, ss, se,
	}
}
