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
	speed := 1
	return Drop{rl.NewVector2(x, float32(y)), float32(speed)}
}

func (d *Drop) Fall() {
	d.Y += d.Speed
}

func (d *Drop) Show() {
	rl.DrawLineEx(d.Vector2, rl.NewVector2(d.X, d.Y+10), 2, purple)
}
