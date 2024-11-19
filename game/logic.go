package game

import (
	"math/rand"

	"github.com/gdamore/tcell/v2"
)

type Player struct {
	Sprite     *Sprite
	Name       string
	Coins      int
	Score      int
	NearMisses int
}

type Game struct {
	Player      Player
	Level       int
	Coins       []*Sprite
	Projectiles []*Projectile
	Alive       bool
	CoinCount   int
}

type Point struct {
	X int
	Y int
}

func InitGame(playerName string) *Game {
	return &Game{
		Player: Player{
			Sprite:     NewSprite('🤓', 70, 20, tcell.StyleDefault.Foreground(tcell.ColorTeal)),
			Name:       playerName,
			Score:      0,
			NearMisses: 0,
			Coins:      0,
		},
		Level:       1,
		Coins:       GenerateCoins(1, tcell.StyleDefault.Foreground(tcell.ColorYellow)),
		Projectiles: GenerateProjectiles(1, tcell.StyleDefault.Foreground(tcell.ColorRed)),
		Alive:       true,
		CoinCount:   0,
	}
}

func GenerateCoins(level int, color tcell.Style) []*Sprite {
	coins := make([]*Sprite, level+2)

	for i := range level + 2 {
		coins[i] = NewSprite('o', rand.Intn(MAP_WIDTH-10)+6, rand.Intn(MAP_HEIGHT-10)+4, color)
	}

	return coins
}

func GenerateProjectile(color tcell.Style) *Projectile {
	spawn := rand.Intn(4)
	var x, y, sx, sy int
	switch spawn {
	// top of map
	case 0:
		x = rand.Intn(MAP_WIDTH-10) + 5
		y = rand.Intn(5) - 5
		sx = 0
		sy = 1
	// right side of map
	case 1:
		x = rand.Intn(5) + MAP_WIDTH
		y = rand.Intn(MAP_HEIGHT-10) + 5
		sx = -1
		sy = 0
	// bottom of map
	case 2:
		x = rand.Intn(MAP_WIDTH-10) + 5
		y = rand.Intn(5) + MAP_HEIGHT
		sx = 0
		sy = -1
	// left side of map
	case 3:
		x = rand.Intn(5) - 5
		y = rand.Intn(MAP_HEIGHT-10) + 5
		sx = 1
		sy = 0
	}
	return NewProjectile(x, y, sx, sy, color)
}

func GenerateProjectiles(level int, color tcell.Style) []*Projectile {
	projectiles := make([]*Projectile, level*4)

	for i := range level * 4 {
		projectiles[i] = GenerateProjectile(color)
	}

	return projectiles
}
