package rayo

type Ray struct {
	Origin    Vector
	Direction Vector
}

func NewRay(a, b Vector) Ray {
	return Ray{a, b}
}

func (r Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Direction.Mul(t))
}

func Color(ray Ray) Vector {
	if hitSphere(Vector{0, 0, -1}, 0.5, ray) {
		return Vector{1, 0, 0}
	}
	unitDirection := ray.Direction.UnitVector()
	t := 0.5 * (unitDirection.Y + 1.0) // -1 < y < 1 => 0 < (y+1)*0.5 < 1.0 => 0 < t < 1
	a := Vector{1.0, 1.0, 1.0}.Mul(1.0 - t)
	b := Vector{.5, .7, 1.0}.Mul(t)
	return a.Add(b)
}

func hitSphere(center Vector, radius float64, ray Ray) bool {
	oc := ray.Origin.Sub(center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * oc.Dot(ray.Direction)
	c := oc.Dot(oc) - (radius * radius)
	discriminant := b*b - 4*a*c
	return discriminant > 0
}
