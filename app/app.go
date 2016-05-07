package main

import (
	"fmt"

	. "simplex/geom"
)

func main() {
	//pt1_out := NewPointFromWKT("POINT ( 49.8322373906287 49.1670033843562 )")
	pt2_out := NewPointFromWKT("POINT (  26.70508112717612 29.46609249326697 )")
	//pnt3_in := NewPointFromWKT("POINT ( 27.439276564111122 38.76590136111034 )")
	poly := NewPolygonFromWKT("" +
        "POLYGON (( 35 10, 45 45, 15 40, 10 20, 35 10 ), " +
		"( 20 30, 35 35, 30 20, 20 30 )" +
        ")")
	//a := pt1_out.Distance(poly)
	b := pt2_out.Distance(poly)
	//c := pnt3_in.Distance(poly)
	//fmt.Println(a)
	fmt.Println(b)
	//fmt.Println(c)
}
