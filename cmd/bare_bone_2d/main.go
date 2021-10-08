package main

import (
	"github.com/ddomurad/goCraft/core"
	"github.com/ddomurad/goCraft/simple2d"
)

type BareBone2dScene struct {
	fillColor   core.Color
	borderColor core.Color
	camera      *simple2d.Camera2d
	animation   float32
}

func (s *BareBone2dScene) Render(dt float64, r *simple2d.Renderer2d, app *core.App) {
	r.SetShader(simple2d.DRI_SHADER_SIMPLE)
	r.ApplyCamera(s.camera)

	r.DrawRect(0, 0, 1, 2, s.animation, s.fillColor)
	r.DrawRectBorder(0, -0, 1, 2, s.animation, 2, s.borderColor)

	s.animation += float32(dt)
}

func main() {
	app := core.InitApp("goCraft - bare bone 2d", 800, 600, true, false)

	scene2d := &BareBone2dScene{
		fillColor:   core.Color{0.1, 0.1, 0.1, 1.0},
		borderColor: core.Color{0.4, 0.4, 1, 1.0},
		camera:      simple2d.NewCamera(),
	}

	scene2d.camera.SetZoom(0.5)

	renderer := simple2d.NewRenderer2d(app, scene2d)

	renderer.Init()
	renderer.SetClearColor([4]float32{0.2, 0.2, 0.2, 1.0})

	for app.Run() {
		app.Render(renderer)
	}
}
