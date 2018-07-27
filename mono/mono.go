package mono

import "github.com/intdxdt/mbr"

const null = -9

type MBR struct {
	mbr.MBR
	I int
	J int
}

//new monotone mbr
func CreateMonoMBR(box mbr.MBR) MBR {
	return MBR{box, null, null}
}

//clone  mono mbr
func (box *MBR) BBox() *mbr.MBR {
	return &box.MBR
}

//update mono chain index
func (box *MBR) UpdateIndex(i, j int) {
	box.I, box.J = i, j
}
