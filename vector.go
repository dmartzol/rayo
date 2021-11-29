package rayo

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vec3 {
	return Vec3{X: x, Y: y, Z: z}
}

func (v Vec3) Mul(t float64) Vec3 {
	return v.Scale(t)
}

func (v Vec3) Scale(t float64) Vec3 {
	return Vec3{v.X * t, v.Y * t, v.Z * t}
}

func (v Vec3) Add(b Vec3) Vec3 {
	return Vec3{v.X + b.X, v.Y + b.Y, v.Z + b.Z}
}

func (v Vec3) Modulo() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vec3) UnitVector() Vec3 {
	modulo := v.Modulo()
	return Vec3{X: v.X / modulo, Y: v.Y / modulo, Z: v.Z / modulo}
}
