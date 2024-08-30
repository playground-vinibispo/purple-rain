package main

import (
	"purple-rain/internals/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	backgroundColor = rl.NewColor(230, 230, 250, 255)
)

func main() {
	var drops [500]models.Drop
	rl.InitWindow(640, 360, "purple-rain")
	for i := 0; i < len(drops); i++ {
		drops[i] = models.NewDrop()
	}
	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)
		for i := 0; i < len(drops); i++ {
			drops[i].Fall()
			drops[i].Show()
		}
		rl.EndDrawing()
	}
	rl.CloseWindow()
}

