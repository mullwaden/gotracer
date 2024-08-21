package main

import (
	alg "gotracer/algebra"
	scene "gotracer/scene"
	"image"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

const fov float64 = 90.0

var cameraPoint = 0.5 / math.Tan(fov*math.Pi/(4.0*180.0))

func TracePixel(sceneData *scene.Scene, x float64, y float64) *alg.Vec3 {
	pos := alg.NewVec3(x, y, 0)
	dir := pos.Sub(alg.NewVec3(0, 0, cameraPoint))

	ray := alg.NewRay(pos, dir)

	var intersection *scene.Intersection
	var sceneItem scene.SceneItem

	lastDist := math.MaxFloat64

	for _, item := range *sceneData.Items {
		ii := item.Intersect(ray)

		if ii != nil && ii.Dist < lastDist {
			intersection = ii
			lastDist = intersection.Dist
			sceneItem = item
		}
	}

	if intersection == nil {
		return alg.NewVec3(0, 0, 0)
	}

	colorOut := alg.NewVec3(.1, .1, .1)

	for _, light := range *sceneData.Lights {
		isBlocked, hitAngle := scene.TraceLight(sceneData, intersection, light.GetPoint())

		if !isBlocked {
			colorOut = colorOut.Add(light.GetColor().Scale(hitAngle))
		}
	}

	colorOut = colorOut.Mul(sceneItem.GetColor())

	// for each light, sample random points in light
	// trace from intersection to light, ignore if hit on the way

	return colorOut
}

func RenderScene(scene *scene.Scene) image.Image {
	w := 640
	h := 480

	// largest dimension goes from -0.5 to 0.5
	scale := math.Max(float64(w), float64(h))

	m := image.NewRGBA(image.Rect(0, 0, w, h))

	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			dx := (float64(x) - float64(w)*0.5) / scale
			dy := (float64(y) - float64(h)*0.5) / scale

			m.Set(x, y, TracePixel(scene, dx, dy).ToRGBA())

		}
	}

	//TracePixel(scene, 0, 0)

	return m
}

func ShowImage(img image.Image) {

	dx := img.Bounds().Bounds().Dx()
	dy := img.Bounds().Bounds().Dy()

	a := app.New()
	w := a.NewWindow("Images")

	content := canvas.NewImageFromImage(img)
	w.SetContent(content)
	w.Resize(fyne.NewSize(float32(dx), float32(dy)))

	w.ShowAndRun()
}
