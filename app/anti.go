package main

import (
    "fmt"
    . "simplex/geom"
    "simplex/cart2d"
    "simplex/util/math"
)

func mainx() {
    var wkt = "LINESTRING ( 460 580, 472 597, 506 613, 532 597, 529 576, 522 549, 532 538, 563 531, 578 544, 584 561, 590 588, 590 605, 603 615, 607 650, 578 660, 558 679, 589 695, 611 696, 642 681, 644 658, 652 649, 679 641, 693 622, 679 606, 686 582, 746 581, 755 577, 764 534, 760 512, 737 488, 709 496, 685 471, 680 446, 671 429, 638 429, 630 442, 623 453, 585 429, 539 414, 510 433, 471 439, 433 431, 419 454, 428 476, 442 480, 473 480, 468 514, 440 521, 421 521, 412 534, 425 542, 449 545, 459 549 )"
    g := ReadGeometry(wkt);
    coords := g.AsLinear()[0].Coordinates();
    //coords := []*Point{A, B, C, D,I, F}
    h := ConvexHull(coords)
    fmt.Println(h[0].WKT())
    fmt.Println(h[1].WKT())
    fmt.Println(h[7].WKT())
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
