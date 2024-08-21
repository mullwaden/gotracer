package main

import (
	alg "gotracer/algebra"
	scene "gotracer/scene"
)

func CreateSceneA() *scene.Scene {
	a := scene.NewSphere(alg.NewVec3(0, 0, -100), 5)
	b := scene.NewSphere(alg.NewVec3(0, 6, -120), 3)
	c := scene.NewSphere(alg.NewVec3(-25, 5, -100), 8)
	d := scene.NewSphere(alg.NewVec3(5, -25, -100), 1)
	e := scene.NewPlane(alg.NewVec3(0, 30, 0), alg.NewVec3(0, -1, 0))

	l1 := scene.NewPointLight(alg.NewVec3(0, -100, 0), alg.NewVec3(1, 1, 1))
	l2 := scene.NewPointLight(alg.NewVec3(50, -50, -50), alg.NewVec3(1, 1, 1))

	return &scene.Scene{
		Items:  &[]scene.SceneItem{b, a, c, d, e},
		Lights: &[]scene.Light{l1, l2},
	}

}
