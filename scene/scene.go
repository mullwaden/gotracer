package scene

import (
	alg "gotracer/algebra"
	"math"
)

type Intersection struct {
	Point *alg.Vec3
	Norm  *alg.Vec3
	Dist  float64
}

type SceneItem interface {
	Intersect(ray *alg.Ray) *Intersection
	GetColor() *alg.Vec3
}

type Light interface {
	GetPoint() *alg.Vec3
	GetColor() *alg.Vec3
}

type Scene struct {
	Items  *[]SceneItem
	Lights *[]Light
}

func TraceLight(
	sceneData *Scene,
	intersection *Intersection,
	point *alg.Vec3,
) (bool, float64) {
	lightRay := alg.NewRay(intersection.Point, point.Sub(intersection.Point))

	// if light ray direction faces away from surface normal
	// it is blocking itself
	if lightRay.Dir.Dot(intersection.Norm) < 0 {
		return true, 0.0
	} else {
		for _, item := range *sceneData.Items {

			intersection := item.Intersect(lightRay)

			if intersection != nil {
				return true, 0.0
			}
		}
	}

	return false, math.Abs(lightRay.Dir.Dot(intersection.Norm))
}
