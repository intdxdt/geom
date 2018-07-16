package geom

type wktToken struct {
	children []*wktToken
	i        int
	j        int
}

type wktTokens []*wktToken

//len of coordinates - sort interface
func (s wktTokens) Len() int {
	return len(s)
}

//swap - sort interface
func (s wktTokens) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

//less - 2d compare - sort interface
func (s wktTokens) Less(i, j int) bool {
	return s[i].i < s[j].i
}
