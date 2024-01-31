package mono

import "github.com/intdxdt/mbr"

const null = -9

type MBR struct {
	mbr.MBR[float64]
	I int
	J int
}

// CreateMonoMBR - new monotone mbr
func CreateMonoMBR(box mbr.MBR[float64]) MBR {
	return MBR{box, null, null}
}

// BBox - clone  mono mbr
func (box *MBR) BBox() *mbr.MBR[float64] {
	return &box.MBR
}

// UpdateIndex - update mono chain index
func (box *MBR) UpdateIndex(i, j int) {
	box.I, box.J = i, j
}
