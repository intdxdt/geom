package main

import (
	"fmt"
	. "simplex/geom"
)

func main() {
	//awkt := "LINESTRING ( 350 710, 400 770, 450 770, 480 810, 570 820, 670 730, 720 760, 930 800 )"

	plywkt := "POLYGON (( 720 760, 860 770, 950 700, 930 640, 800 600, 740 580, 730 500, 760 440, 720 360, 620 390, 510 480, 460 570, 440 630, 450 740, 480 810, 570 820, 570 770, 580 740, 670 730, 720 760 ), ( 630 670, 580 650, 590 600, 650 580, 710 600, 710 670, 630 670 ), ( 780 650, 800 640, 850 710, 830 720, 780 650 ))"
	//plywkt2 :="POLYGON (( 860 920, 950 880, 860 800, 930 720, 880 690, 830 700, 810 730, 790 790, 820 840, 810 870, 860 920 ), ( 840 750, 860 750, 850 800, 830 800, 840 750 ))"

	//ptAwkt := "POINT ( 630 650 )"
	//ptBwkt := "POINT ( 710 600 )"
	//ptCwkt := "POINT ( 722.1298042987639 582.0334837046336 )"
	ptDwkt := "POINT ( 720 360 )"

	//g := NewLineStringFromWKT(awkt)
	ply := NewPolygonFromWKT(plywkt)
	//ply2 := NewPolygonFromWKT(plywkt2)

	//ptA := NewPointFromWKT(ptAwkt)
	ptB := NewPointFromWKT(ptDwkt)
	//ptC := NewPointFromWKT(ptCwkt)
	//ptD := NewPointFromWKT(ptDwkt)

	//g.PrivLnIntersection(ply)
	pts := ply.Intersection(ptB)
	for _, pt := range pts {
		fmt.Println(pt.WKT())
	}

	a := NewPointXY(0, 0)
	b := NewPointXY(-3, 4)
	k := &Point{2, 2}
	n := &Point{1, 5}

	seg_ab := NewSegment(a, b)
	seg_kn := &Segment{A: k, B: n}
	bln := seg_ab.Intersects(seg_kn)
	fmt.Println(bln)
}
