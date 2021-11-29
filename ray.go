package rayo

type Ray struct {
	A Vec3 // origin
	B Vec3 // direction
}

func NewRay(a, b Vec3) Ray {
	return Ray{a, b}
}

func (r Ray) Origin() Vec3 {
	return r.A
}

func (r Ray) Direction() Vec3 {
	return r.B
}

func (r Ray) PointAt(t float64) Vec3 {
	return r.A.Add(r.B.Mul(t))
}

func Color(ray Ray) Vec3 {
	unitDirection := ray.Direction().UnitVector()
	t := 0.5 * (unitDirection.Y + 1.0) // -1 < y < 1 => 0 < (y+1)*0.5 < 1.0 => 0 < t < 1
	a := Vec3{1.0, 1.0, 1.0}.Mul(1.0 - t)
	b := Vec3{.5, .7, 1.0}.Mul(t)
	return a.Add(b)
}
