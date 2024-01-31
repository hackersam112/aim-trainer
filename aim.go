package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/leedenison/gologo"
	"github.com/leedenison/gologo/obj"
)

type particle struct {
	radius float64
	object *gologo.Object
}

var (
	p    *particle
	last time.Time
)

func squareDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2)
}

func mouseButtonPressed(w *glfw.Window, button glfw.MouseButton, action glfw.Action, mod glfw.ModifierKey) {
	if button == glfw.MouseButtonLeft && action == glfw.Press {
		cx, cy := w.GetCursorPos()
		px, py := p.object.GetPosition()

		// y axis needs to be inverted between screen coordinates and opengl coordinates
		sx, sy := w.GetSize()
		cy = float64(sy) - cy

		sqDist := squareDistance(cx, cy, float64(px), float64(py))

		if sqDist <= math.Pow(p.radius, 2) {
			now := time.Now()
			fmt.Printf("Target clicked after: %v\n", now.Sub(last))
			last = now

			p.object.SetPosition(float32(rand.Intn(sx)), float32(rand.Intn(sy)))
		}
	}
}

func main() {
	g := gologo.Init()
	defer g.Close()

	p = &particle{
		radius: 50.0,
		object: obj.Polygon(
			mgl32.Vec2{200.0, 300.0},
			10,
			50.0,
			mgl32.Vec4{1.0, 1.0, 1.0, 1.0}),
	}

	g.Window.SetMouseButtonCallback(mouseButtonPressed)

	last = time.Now()

	for !g.Window.ShouldClose() {
		g.ClearBackBuffer()
		p.object.Draw()
		g.Window.SwapBuffers()
		g.CheckForEvents()
	}
}
