package main

import (
	"regexp"
	"strings"
	"fmt"
)

func main() {
	wkt := "POLYGON (( 160 340, 160 380, 180 380, 180 340, 160 340 ))"
	var o = regexp.MustCompile(`^\s*(?P<type>[A-Za-z]+)\s*\(\s*(?P<coords>.*)\s*\)\s*$`)
	match := o.FindStringSubmatch(wkt)
	fmt.Println(strings.ToLower(match[1]))
}
