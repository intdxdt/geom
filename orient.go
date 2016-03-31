package geom

type Orient struct {
    side int8
}
//new orientation
func NewOrientation() *Orient{
    return &Orient{'o'}
}
//is left
func (self *Orient) IsLeft() bool {
    return self.side == 'l'
}
//is on
func (self *Orient) IsOn() bool {
    return self.side == 'o'
}
//is right
func (self *Orient) IsRight() bool {
    return self.side == 'r'
}
//is on or left
func (self *Orient) IsOnOrLeft() bool {
    return self.IsOn() || self.IsLeft()
}

//is on or right
func (self *Orient) IsOnOrRight() bool {
    return self.IsOn() || self.IsRight()
}

