package geom

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/robust"
)

//Component vector
func Component(m, d float64) (float64, float64) {
	return m * math.Cos(d), m * math.Sin(d)
}

//Dot Product of two points as vectors
func DotProduct(v, o Point) float64 {
	return DotProductXY(v[X], v[Y], o[X], o[Y])
}

//Dot Product of two points as vectors
func DotProductXY(vx, vy, ox, oy float64) float64 {
	return (vx * ox) + (vy * oy)
}

//Unit vector of point
func Unit(v *Point) (float64, float64) {
	return UnitXY(v[X], v[Y])
}

//Unit vector of point
func UnitXY(x, y float64) (float64, float64) {
	var m = MagnitudeXY(x, y)
	if feq(m, 0) {
		m = math.EPSILON
	}
	return x / m, y / m
}

//Projects  u on to v
func Project(u, onv *Point) float64 {
	return ProjectXY(u[X], u[Y], onv[X], onv[Y])
}

//Projects  u on to v, using x and y compoents of u and v
func ProjectXY(ux, uy, onvX, onvY float64) float64 {
	var cx, cy = UnitXY(onvX, onvY)
	return DotProductXY(ux, uy, cx, cy)
}

//2D cross product of AB and AC vectors given A, B, and C as points,
//i.e. z-component of their 3D cross product.
//Returns a positive value, if ABC makes a counter-clockwise turn,
//negative for clockwise turn, and zero if the points are collinear.
func Orientation2D(a, b, c *Point) float64 {
	return robust.Orientation2D(a[:Z], b[:Z], c[:Z])
}

//2D cross product of AB and AC vectors,
//i.e. z-component of their 3D cross product.
//negative cw and positive if ccw
func CrossProduct(ab, ac Point) float64 {
	return (ab[X] * ac[Y]) - (ab[Y] * ac[X])
}

//Computes vector magnitude given x an dy component
func MagnitudeXY(dx, dy float64) float64 {
	return math.Hypot(dx, dy)
}

//Computes vector magnitude given x an dy component
func MagnitudeSquareXY(dx, dy float64) float64 {
	return (dx * dx) + (dy * dy)
}

//Checks if catesian coordinate is null ( has NaN )
func IsNull(x, y float64) bool {
	return math.IsNaN(x) || math.IsNaN(y)
}

//Checks if x and y components are zero
func IsZero(x, y float64) bool {
	return feq(x, 0) && feq(y, 0)
}

//Dir computes direction in radians - counter clockwise from x-axis.
func Direction(x, y float64) float64 {
	var d = math.Atan2(y, x)
	if d < 0 {
		d += math.Tau
	}
	return d
}

//Revdir computes the reversed direction from a foward direction
func ReverseDirection(d float64) float64 {
	var r = d - math.Pi
	if d < math.Pi {
		r = d + math.Pi
	}
	return r
}

func DeflectionAngle(bearing1, bearing2 float64) float64 {
	var a = bearing2 - ReverseDirection(bearing1)
	if a < 0.0 {
		a = a + math.Tau
	}
	return math.Pi - a
}

//Extvect extends vector from the from end or from begin of vector
func Extend(x, y, magnitude, angle float64, fromEnd bool) (float64, float64) {
	//from a of v back direction initiates as fwd v direction anticlockwise
	//bb - back bearing
	//fb - forward bearing
	var bb = Direction(x, y)
	if fromEnd {
		bb += math.Pi
	}
	var fb = bb + angle
	if fb > math.Tau {
		fb -= math.Tau
	}
	return Component(magnitude, fb)
}

//Deflect_vector computes vector deflection given deflection angle and
// side of vector to deflect from (from_end)
func Deflect(v *Point, mag, deflAngle float64, fromEnd bool) (float64, float64) {
	return Extend(v[X], v[Y], mag, math.Pi - deflAngle, fromEnd)
}

//KProduct scales x and y components by constant  k
func KProduct(x, y, k float64) (float64, float64) {
	return k * x, k * y
}
