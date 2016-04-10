package geom

const (
    Left = -1 + iota
    On
    Right
)

//type Side of zero
type Side struct {
    s int
}

//New Side
func NewSide() *Side {
    return &Side{On}
}

//is left
func (self *Side) IsLeft() bool {
    return self.s == Left
}
//as left
func (self *Side) AsLeft() *Side {
    self.s = Left
    return  self
}

//Is on
func (self *Side) IsOn() bool {
    return self.s == On
}

//As on
func (self *Side) AsOn() *Side {
    self.s = On
    return self
}

//Is right
func (self *Side) IsRight() bool {
    return self.s == Right
}

//As right
func (self *Side) AsRight() *Side {
    self.s = Right
    return self
}

//is on or left
func (self *Side) IsOnOrLeft() bool {
    return self.IsOn() || self.IsLeft()
}

//is on or right
func (self *Side) IsOnOrRight() bool {
    return self.IsOn() || self.IsRight()
}

