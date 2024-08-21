package algebra

import "fmt"

type Ray struct {
	Pos *Vec3
	Dir *Vec3
}

func NewRay(pos *Vec3, dir *Vec3) *Ray {

	return &Ray{Pos: pos, Dir: dir.Norm()}
}

func (ray *Ray) ToPoint(d float64) *Vec3 {

	return ray.Pos.Add(ray.Dir.Scale(d))
}

func (ray *Ray) String() string {
	return fmt.Sprintf("pos: %v dir: %v", ray.Pos, ray.Dir)
}
