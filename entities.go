package main

import "github.com/gdamore/tcell/v2"

type Sprite struct {
	Char rune
	X, Y int
	Color tcell.Style
}

func NewSprite(char rune, x, y int, color tcell.Style) *Sprite {
	return &Sprite{
		Char: char,
		X:    x,
		Y:    y,
		Color: color, 
	}
}

func (s *Sprite) Draw(screen tcell.Screen) {
	screen.SetContent(s.X, s.Y, s.Char, nil, s.Color)
}

type Projectile struct {
	Sprite *Sprite
	SpeedX, SpeedY int
}

func NewProjectile(x, y, sx, sy int, color tcell.Style) *Projectile {
	return &Projectile{
		Sprite: NewSprite('*', x, y, color),
		SpeedX: sx,
		SpeedY: sy,
	}
}

func (p *Projectile) Update() {
	p.Sprite.X += p.SpeedX
	p.Sprite.Y += p.SpeedY
}