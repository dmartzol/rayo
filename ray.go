package rayo

type Ray struct {
	A Vector // origin
	B Vector // direction
}

func NewRay(a, b Vector) Ray {
	return Ray{a, b}
}

func (r Ray) Origin() Vector {
	return r.A
}

func (r Ray) Direction() Vector {
	return r.B
}

func (r Ray) PointAt(t float64) Vector {
	return r.A.Add(r.B.Mul(t))
}

func Color(ray Ray) Vector {
	unitDirection := ray.Direction().UnitVector()
	t := 0.5 * (unitDirection.Y + 1.0) // -1 < y < 1 => 0 < (y+1)*0.5 < 1.0 => 0 < t < 1
	a := Vector{1.0, 1.0, 1.0}.Mul(1.0 - t)
	b := Vector{.5, .7, 1.0}.Mul(t)
	return a.Add(b)
}
