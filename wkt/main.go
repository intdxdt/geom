package main

import (
	"fmt"
	"github.com/intdxdt/stack"
	"sort"
)

type Token struct {
	i        int
	j        int
	children []*Token
}

type Tokens []*Token

//len of coordinates - sort interface
func (s Tokens) Len() int {
	return len(s)
}

//swap - sort interface
func (s Tokens) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//less - 2d compare - sort interface
func (s Tokens) Less(i, j int) bool {
	return s[i].i < s[j].i
}

func main() {
	//var wkt = "POLYGON ((35 10, 45 45, 15 40, 10 20, 35 10), (20 30, 35 35, 30 20, 20 30))"
	var wkt = "MULTIPOLYGON (((40 40, 20 45, 45 30, 40 40)), ((20 35, 10 30, 10 10, 30 5, 45 20, 20 35),(30 20, 20 15, 20 25, 30 20)))"
	//var wkt = "MULTIPOINT ((10 40), (40 30), (20 20), (30 10))"
	//var wkt = "POLYGON ((30 10, 40 40, 20 40, 10 20, 30 10))"
	//var wkt = "POINT (30 10)"
	var wktBytes = []byte(wkt)
	fmt.Println(wktBytes)
	fmt.Println(byte('('))
	fmt.Println(byte(')'))
	var name, tokens = buildTokens(wktBytes)
	tokens = aggregateTokens(tokens)
	fmt.Println(string(name))
	fmt.Println(string(wktBytes[tokens[0].i+1 : tokens[0].j]))
	var tcoords = tokenCoordinates(wktBytes, tokens)
	var coords = tcoords[0].([][]byte)
	for _, ch := range coords {
		fmt.Println(string(ch))
	}
}

func buildTokens(stream []byte) ([]byte, []*Token) {
	var tokens []*Token
	var stk = stack.NewStack()
	var namespace bool
	var name []byte

	for i, o := range stream {
		if o == '(' {
			namespace = true
			stk = stk.Push(&Token{i: i})
		} else if o == ')' {
			var s = stk.Pop().(*Token)
			s.j = i
			tokens = append(tokens, s)
		}

		if !namespace {
			if o != ' ' {
				name = append(name, o)
			}
		}
	}

	return name, tokens
}

func aggregateTokens(tokens []*Token) []*Token {
	var bln = false
	bln, tokens = aggregateSequence(tokens)
	if bln {
		for _, tok := range tokens {
			_, tok.children = aggregateSequence(tok.children)
		}
	}
	return tokens
}

func aggregateSequence(tokens []*Token) (bool, []*Token) {
	sort.Sort(Tokens(tokens))
	var heads []*Token
	var head *Token
	var aggregate = false
	for _, tok := range tokens {
		if head == nil {
			head = tok
			heads = append(heads, head)
		} else {
			if head.i < tok.i && tok.j < head.j {
				head.children = append(head.children, tok)
				aggregate = true
			} else {
				head = tok
				heads = append(heads, head)
			}
		}
	}
	return aggregate, heads
}

func tokenCoordinates(wktbytes []byte, tokens []*Token) []interface{} {
	var tokenCoords []interface{}
	for _, tok := range tokens {
		var chBytes [][]byte
		if len(tok.children) == 0 {
			chBytes = append(chBytes, wktbytes[tok.i+1:tok.j])
		} else {
			for _, ch := range tok.children {
				chBytes = append(chBytes, wktbytes[ch.i+1:ch.j])
			}
		}
		tokenCoords = append(tokenCoords, chBytes)
	}

	return tokenCoords
}
