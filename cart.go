package geom

import (
	"github.com/intdxdt/math"
	"github.com/intdxdt/robust"
)

// Component vector
func Component(m, d float64) (float64, float64) {
	return m * math.Cos(d), m * math.Sin(d)
}

// DotProduct - Dot Product of two points as vectors
func DotProduct(vx, vy, ox, oy float64) float64 {
	return (vx * ox) + (vy * oy)
}

// UnitVector of point
func UnitVector(x, y float64) (float64, float64) {
	var m = MagnitudeXY(x, y)
	if feq(m, 0) {
		m = math.EPSILON
	}
	return x / m, y / m
}

// Project u on to v
func Project(u, onv *Point) float64 {
	return ProjectXY(u[X], u[Y], onv[X], onv[Y])
}

// ProjectXY - Projects  u on to v, using x and y compoents of u and v
func ProjectXY(ux, uy, onvX, onvY float64) float64 {
	var cx, cy = UnitVector(onvX, onvY)
	return DotProduct(ux, uy, cx, cy)
}

// Orientation2D - 2D cross product of AB and AC vectors given A, B, and Pnts as points,
// i.e. z-component of their 3D cross product.
// Returns a positive value, if ABC makes a counter-clockwise turn,
// negative for clockwise turn, and zero if the points are collinear.
func Orientation2D(a, b, c *Point) float64 {
	return robust.Orientation2D(a[:Z], b[:Z], c[:Z])
}

// CrossProduct - 2D cross product of AB and AC vectors,
// i.e. z-component of their 3D cross product.
// negative cw and positive if ccw
func CrossProduct(ax, ay, bx, by float64) float64 {
	return (ax * by) - (ay * bx)
}

// MagnitudeXY - Computes vector magnitude given x an dy component
func MagnitudeXY(dx, dy float64) float64 {
	return math.Hypot(dx, dy)
}

// MagnitudeSquareXY - Computes vector magnitude given x an dy component
func MagnitudeSquareXY(dx, dy float64) float64 {
	return math.Hypot2(dx, dy)
}

// IsNull - Checks if Cartesian coordinate is null ( has NaN )
func IsNull(x, y float64) bool {
	return math.IsNaN(x) || math.IsNaN(y)
}

// IsZero - Checks if x and y components are zero
func IsZero(x, y float64) bool {
	return feq(x, 0) && feq(y, 0)
}

// Direction - computes direction in radians - counter clockwise from x-axis.
func Direction(x, y float64) float64 {
	var d = math.Atan2(y, x)
	if d < 0 {
		d += math.Tau
	}
	return d
}

// ReverseDirection - computes the reversed direction from a foward direction
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

// Extend - extends vector from the from end or from begin of vector
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

// Deflect - computes vector deflection given deflection angle and
// side of vector to deflect from (from_end)
func Deflect(vx, vy, mag, deflAngle float64, fromEnd bool) (float64, float64) {
	return Extend(vx, vy, mag, math.Pi-deflAngle, fromEnd)
}

// KProduct - scales x and y components by constant  k
func KProduct(x, y, k float64) (float64, float64) {
	return k * x, k * y
}
