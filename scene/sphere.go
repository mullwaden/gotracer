package scene

import (
	alg "gotracer/algebra"
	"math"
)

type Sphere struct {
	Pos    *alg.Vec3
	Radius float64
}

func NewSphere(pos *alg.Vec3, radius float64) SceneItem {

	var out SceneItem = &Sphere{Pos: pos, Radius: radius}

	return out
}

func (s *Sphere) GetColor() *alg.Vec3 {
	return alg.NewVec3(.7, .7, 0)
}

func (s *Sphere) Intersect(ray *alg.Ray) *Intersection {
	oc := ray.Pos.Sub(s.Pos)
	a := ray.Dir.Dot(ray.Dir)
	b := 2.0 * oc.Dot(ray.Dir)
	c := oc.Dot(oc) - s.Radius*s.Radius

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	dist := (-b - math.Sqrt(discriminant)) / (2.0 * a)

	if dist < 0 {
		return nil
	}

	point := ray.ToPoint(dist)
	normal := point.Sub(s.Pos).Norm()

	return &Intersection{
		Point: point,
		Norm:  normal,
		Dist:  dist,
	}
}
