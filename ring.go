package geom

type LinearRing struct {
	*LineString
}

//new linear ring
func NewLinearRing(coordinates []*Point) *LinearRing {
	coords := CloneCoordinates(coordinates)
	if len(coordinates) > 1 {
		if !IsRing(coordinates) {
			coords = append(coords, coordinates[0].Clone())
		}
	}
	return &LinearRing{NewLineString(coords, false)}
}
