package main

import (
	"fmt"
	"purple-rain/internals/models"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	backgroundColor = rl.NewColor(230, 230, 250, 255)
)

func main() {
	var drops [500]models.Drop
	rl.InitWindow(640, 360, "purple-rain")
	rl.InitAudioDevice()
	song := rl.LoadMusicStream("assets/rain.mp3")
	rl.PlayMusicStream(song)
	for i := 0; i < len(drops); i++ {
		drops[i] = models.NewDrop()
	}
	rl.SetTargetFPS(60)
	var paused bool
	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(song)
		if !rl.IsMusicStreamPlaying(song) && !paused {
			rl.PlayMusicStream(song)
		}
		if rl.IsKeyPressed(rl.KeySpace) {
			paused = !paused
			if paused {
				rl.PauseMusicStream(song)
			} else {
				rl.PauseMusicStream(song)
			}
		}
		fmt.Println("time: ", rl.GetMusicTimePlayed(song)/rl.GetMusicTimeLength(song))
		rl.BeginDrawing()
		rl.ClearBackground(backgroundColor)
		for i := 0; i < len(drops); i++ {
			if !paused {
				drops[i].Fall()
			}
			drops[i].Show()
		}
		rl.EndDrawing()
	}
	rl.UnloadMusicStream(song)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

