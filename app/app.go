package main

import (
    "fmt"
    . "simplex/geom"
)

func main() {
    coords := []*Point{{1, 2, 3}, {3, 4, 5}, {6, 7, 8}}
    ln := NewLineString(coords)
    fmt.Println(ln)
    ln.Clone()
    //fmt.Println(c)
}
