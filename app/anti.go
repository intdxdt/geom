package main

import (
    "fmt"
    . "simplex/geom"
    "simplex/cart2d"
    "simplex/util/math"
)

func main() {
    var A = NewPointXY(70, 55)
    var B = NewPointXY(75, 55)
    var C = NewPointXY(80, 55)
    var D = NewPointXY(85, 55)

    coords := []*Point{A, B, C, D}
    //coords := []*Point{A, B, C, D,I, F}
    h := ConvexHull(coords)
    fmt.Println(h)
    fmt.Println(NewPolygon(h).WKT())

    hull := NewHull(h)
    idx := hull.Antipodal(1, 0)
    fmt.Println(idx)

    fmt.Println(math.FloatEqual(30.0000000000056, 30.))

    //v_bc := vect.NewVect(&vect.Options{A:B, B:C})
    //v_cd := vect.NewVect(&vect.Options{A:C, B:D})
    //
    //ov := v_bc.Extvect(30, math.Deg2rad(270), true)
    //d := v_cd.Project(ov)
    //fmt.Println(d)
    //
    //fmt.Println("ov->v", ov.Vector())
    //
    ////==================================================
    //
    //
    //
    //
    //i, j := 2,3
    //var ptI, ptJ = h[i], h[j]
   	//var cmpIJ = ptJ.Sub(ptI)
    //var oth = orthvector(cmpIJ, math.Deg2rad(270))
    //fmt.Println(oth)
    //
    //var uvect = D.Sub(ptJ)
    //
    //
    //dist := cart2d.Project(uvect, oth)
    //fmt.Println(dist)

}


func orthvector(v *Point, angle float64) *Point {
	m := 30.0
	cx, cy := cart2d.Extend(v, m, angle, true)
	return NewPointXY(cx, cy)
}
