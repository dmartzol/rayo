package rayo

import "math"

type Vector struct {
	X, Y, Z float64
}

func NewVector(x, y, z float64) Vector {
	return Vector{X: x, Y: y, Z: z}
}

func (v Vector) Mul(t float64) Vector {
	return v.Scale(t)
}

func (v Vector) Scale(t float64) Vector {
	return Vector{v.X * t, v.Y * t, v.Z * t}
}

func (v Vector) Add(b Vector) Vector {
	return Vector{v.X + b.X, v.Y + b.Y, v.Z + b.Z}
}

func (v Vector) Modulo() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) UnitVector() Vector {
	modulo := v.Modulo()
	return Vector{X: v.X / modulo, Y: v.Y / modulo, Z: v.Z / modulo}
}
