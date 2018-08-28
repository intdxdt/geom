package mono

import "github.com/intdxdt/mbr"

type MBR struct {
	mbr.MBR
	I int
	J int
}

//clone  mono mbr
func (box *MBR) BBox() *mbr.MBR {
	return &box.MBR
}

//update mono chain index
func (box *MBR) UpdateIndex(i, j int) {
	box.I, box.J = i, j
}
