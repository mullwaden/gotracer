package algebra

import (
	"fmt"
	"image/color"
	"math"
)

type Vec3 struct {
	X float64
	Y float64
	Z float64
}

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v *Vec3) Dot(other *Vec3) float64 {
	return v.X*other.X + v.Y*other.Y + v.Z*other.Z
}

func (v *Vec3) Len() float64 {
	return math.Sqrt(v.Dot(v))
}

func (v *Vec3) Mul(other *Vec3) *Vec3 {
	return NewVec3(
		v.X*other.X,
		v.Y*other.Y,
		v.Z*other.Z,
	)
}

func (v *Vec3) Add(other *Vec3) *Vec3 {
	return NewVec3(
		v.X+other.X,
		v.Y+other.Y,
		v.Z+other.Z,
	)
}

func (v *Vec3) Sub(other *Vec3) *Vec3 {

	return NewVec3(
		v.X-other.X,
		v.Y-other.Y,
		v.Z-other.Z,
	)
}

func (v *Vec3) Scale(s float64) *Vec3 {
	return NewVec3(
		v.X*s,
		v.Y*s,
		v.Z*s,
	)
}

func (v *Vec3) Norm() *Vec3 {
	len := v.Len()

	return v.Scale(1 / len)
}

func (v *Vec3) String() string {
	return fmt.Sprintf("[%v,%v,%v]", v.X, v.Y, v.Z)
}

func clampedUint(v float64) uint8 {
	return uint8(math.Max(math.Min(v, 1.0), 0.0) * 255.0)
}

func (v *Vec3) ToRGBA() color.RGBA {
	return color.RGBA{clampedUint(v.X), clampedUint(v.Y), clampedUint(v.Z), 255}
}
