package main

import (
	"github.com/ddomurad/goCraft/core"
	"github.com/ddomurad/goCraft/resource"
	"github.com/ddomurad/goCraft/simple2d"
	"github.com/go-gl/gl/v3.3-core/gl"
)

const (
	CUSTOM_SHADER = "custom_shader"
)

type CustomShader2dScene struct {
	fillColor   core.Color
	borderColor core.Color
	camera      *simple2d.Camera2d
	animation   float32
}

func (s *CustomShader2dScene) Render(dt float64, r *simple2d.Renderer2d, app *core.App) {
	r.SetShader(CUSTOM_SHADER)
	shader := r.GetShader()
	ul := shader.GetUniformLocation("animation")
	gl.Uniform1f(ul, s.animation)

	r.ApplyCamera(s.camera)

	r.DrawElipse(0, 0, 1, 1, s.animation, s.fillColor)
	r.DrawElipseBorder(0, -0, 1, 1, s.animation, 2, s.borderColor)

	s.animation += float32(dt)
}

func main() {
	app := core.InitApp("goCraft - custom shader", 800, 600, true, false)

	scene2d := &CustomShader2dScene{
		fillColor:   core.Color{0.1, 0.1, 0.1, 1.0},
		borderColor: core.Color{0.4, 0.4, 1, 1.0},
		camera:      simple2d.NewCamera(),
	}

	scene2d.camera.SetZoom(1)

	renderer := simple2d.NewRenderer2d(app, scene2d)

	renderer.Init()
	renderer.SetClearColor([4]float32{0.2, 0.2, 0.2, 1.0})
	renderer.SetCrictleSegments(100)

	app.ResourceManager.PreloadReource(resource.RT_SHADER, CUSTOM_SHADER, resource.ShaderFileSource{
		VertexShaderPath:   "./cmd/custom_shader/custom.vs",
		FragmentShaderPath: "./cmd/custom_shader/custom.fs",
	})

	for app.Run() {
		app.Render(renderer)
	}
}
