package mbr

import (
    "fmt"
)

//String converts mbr to wkt string
func (self *MBR) String() string {
    lx, ly := self[x1], self[y1]
    ux, uy := self[x2], self[y2]

    return fmt.Sprintf(
        "POLYGON ((%v %v, %v %v, %v %v, %v %v, %v %v))",
        lx, ly, lx, uy, ux, uy, ux, ly, lx, ly,
    )
}