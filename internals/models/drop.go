package models

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	purple = rl.NewColor(138, 43, 226, 255)
)

type Drop struct {
	rl.Vector2
	Speed float32
}

func random(minimum, maximum float32) float32 {
	return minimum + rand.Float32()*(maximum-minimum)
}

func NewDrop() Drop {
	x := random(0, float32(rl.GetScreenWidth()))
	y := random(-200, -100)
	speed := random(4, 10)
	return Drop{rl.NewVector2(x, float32(y)), float32(speed)}
}

func (d *Drop) Fall() {
	d.Y += d.Speed
	if d.Y > float32(rl.GetScreenHeight()) {
		d.Y = random(-200, -100)
	}
}

func (d *Drop) Show() {
	rl.DrawLineV(d.Vector2, rl.NewVector2(d.X, d.Y+10), purple)
	// rl.DrawLineEx(d.Vector2, rl.NewVector2(d.X, d.Y+10), 2, purple)
}
