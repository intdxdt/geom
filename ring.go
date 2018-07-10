package geom

type LinearRing struct {
	*LineString
}

//new linear ring
func NewLinearRing(coordinates []Point) *LinearRing {
	var n = len(coordinates)
	var coords = coordinates[:n:n]
	if n > 1 {
		if !IsRing(coords) {
			coords = append(coords, coords[0])
		}
	}
	return &LinearRing{NewLineString(coords)}
}
