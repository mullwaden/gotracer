package scene

import (
	alg "gotracer/algebra"
)

type PointLight struct {
	Pos   *alg.Vec3
	Color *alg.Vec3
}

func NewPointLight(pos *alg.Vec3, color *alg.Vec3) Light {

	var out Light = &PointLight{Pos: pos, Color: color}

	return out
}

func (l *PointLight) GetPoint() *alg.Vec3 {
	return l.Pos
}

func (l *PointLight) GetColor() *alg.Vec3 {
	return l.Color
}
