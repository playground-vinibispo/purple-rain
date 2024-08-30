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
	speed  float32
	length float32
	z      float32
}

func random(minimum, maximum float32) float32 {
	return minimum + rand.Float32()*(maximum-minimum)
}

func NewDrop() Drop {
	x := random(0, float32(rl.GetScreenWidth()))
	y := random(-500, -50)
	z := random(0, 20)
	speed := rl.Remap(z, 0, 20, 1, 20)
	length := rl.Remap(z, 0, 20, 10, 20)
	return Drop{
		Vector2: rl.NewVector2(x, y),
		speed:   speed,
		length:  length,
		z:       z,
	}
}

func (d *Drop) Fall() {
	d.Y += d.speed
	gravity := rl.Remap(d.z, 0, 20, 0, 0.2)
	d.speed += gravity
	if d.Y > float32(rl.GetScreenHeight()) {
		d.Y = random(-200, -100)
		d.speed = rl.Remap(d.z, 0, 20, 4, 10)
	}
}

func (d *Drop) Show() {
	thick := rl.Remap(d.z, 0, 20, 1, 3)
	// rl.DrawLineV(d.Vector2, rl.NewVector2(d.X, d.Y+d.length), purple)
	rl.DrawLineEx(d.Vector2, rl.NewVector2(d.X, d.Y+10), thick, purple)
}
