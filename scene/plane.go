package scene

import (
	alg "gotracer/algebra"
)

type Plane struct {
	Pos  *alg.Vec3
	Norm *alg.Vec3
}

func NewPlane(pos *alg.Vec3, norm *alg.Vec3) SceneItem {

	var out SceneItem = &Plane{Pos: pos, Norm: norm.Norm()}

	return out
}

func (p *Plane) GetColor() *alg.Vec3 {
	return alg.NewVec3(0, 0, 1)
}

func (p *Plane) Intersect(ray *alg.Ray) *Intersection {
	denom := p.Norm.Dot(ray.Dir)

	if denom > -0.0001 {
		return nil
	}

	dist := p.Pos.Sub(ray.Pos).Dot(p.Norm) / denom

	if dist < 0 {
		return nil
	}

	point := ray.ToPoint(dist)

	return &Intersection{
		Point: point,
		Norm:  p.Norm,
		Dist:  dist,
	}

}
