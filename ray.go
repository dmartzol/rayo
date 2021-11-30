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
	unitDirection := ray.Direction.UnitVector()
	t := 0.5 * (unitDirection.Y + 1.0) // -1 < y < 1 => 0 < (y+1)*0.5 < 1.0 => 0 < t < 1
	a := Vector{1.0, 1.0, 1.0}.Mul(1.0 - t)
	b := Vector{.5, .7, 1.0}.Mul(t)
	return a.Add(b)
}
