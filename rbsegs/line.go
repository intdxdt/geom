package main

import (
	"fmt"
	"github.com/intdxdt/geom"
)

func main() {
	var awkt = "LINESTRING ( 800 1600, 1000 1800, 1200 1600, 1400 2000, 1000 2200 )"
	var bwkt = "LINESTRING ( 1000 2000, 1600 1600, 1200 2200, 1600 2000, 1402.6469565217394 1490.912173913043, 875.6904347826086 1716.3034782608693 )"

	//var c = geom.Point{1.5, -2}
	//var d = geom.Point{-1.5, 2}
	//var h = geom.Point{0.484154648492778, -0.645539531323704}
	//var i = geom.Point{0.925118053504632, -1.233490738006176}
	//var ln_cd = NewLineString(geom.Coordinates([]geom.Point{c, d}))
	//var ln_hi = NewLineString(geom.Coordinates([]geom.Point{h, i}))

	var aln = geom.NewLineString(geom.NewLineStringFromWKT(awkt).Coordinates)
	var bln = geom.NewLineString(geom.NewLineStringFromWKT(bwkt).Coordinates)
	for _, o := range aln.Intersection(bln){
		fmt.Println(o.WKT())
	}
}
