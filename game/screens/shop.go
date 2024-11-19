package screens

import (
	"go-game/game"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
)

type Image struct {
	Path string
	Width int
	Height int
}

func Shop(screen tcell.Screen, g *game.Game, coins int) {
	fps := 0
	frameCount := 0
	lastFPSUpdate := time.Now()
	ticker := time.NewTicker(time.Second / 30)
	defer ticker.Stop()

	SHOP_TEXT := `
	███████  ██   ██  ███████   ██████ 
	██       ██   ██  ██   ██   ██  ██
	███████  ███████  ██   ██   ██████
	     ██  ██   ██  ██   ██   ██
	███████  ██   ██  ███████   ██
	----------------------------------
	`
	MAP_HEIGHT := 30
	MAP_WIDTH := 78
	LEFT_BORDER_X := 40
	TOP_BORDER_Y := 10

	X_OFFSET := 5
	Y_OFFSET := 5
	BACKGROUND_POINT := game.Point{X: LEFT_BORDER_X + X_OFFSET, Y: TOP_BORDER_Y + Y_OFFSET}
	CHILDHOOD_POINT := game.Point{X: LEFT_BORDER_X + X_OFFSET, Y: TOP_BORDER_Y + MAP_HEIGHT - Y_OFFSET}
	NOW_POINT := game.Point{X: LEFT_BORDER_X + MAP_WIDTH - X_OFFSET, Y: TOP_BORDER_Y + Y_OFFSET}
	FUTURE_POINT := game.Point{X: LEFT_BORDER_X + MAP_WIDTH - X_OFFSET, Y: TOP_BORDER_Y + MAP_HEIGHT - Y_OFFSET}

	ITEMS := make(map[string]struct{}, 0)

	if game.CheckInventory(g.Player.Name, "background") == 1 {
		ITEMS["background"] = struct{}{}
	}
	if game.CheckInventory(g.Player.Name, "childhood") == 1 {
		ITEMS["childhood"] = struct{}{}
	}
	if game.CheckInventory(g.Player.Name, "now") == 1 {
		ITEMS["now"] = struct{}{}
	}
	if game.CheckInventory(g.Player.Name, "future") == 1 {
		ITEMS["future"] = struct{}{}
	}

	for {	
		// draw logic
		screen.Clear()

		// draw ui
		game.DrawString(screen, 0, 0, "Coins: " + strconv.Itoa(coins))
		game.DrawString(screen, 147, 0, strconv.Itoa(fps))

		g.Player.Sprite.Draw(screen)

		// draw map
		game.DrawString(screen, 63, 0, SHOP_TEXT)
		game.DrawRect(screen, LEFT_BORDER_X, TOP_BORDER_Y, MAP_WIDTH, MAP_HEIGHT, tcell.StyleDefault)

		// draw items
		game.DrawString(screen, BACKGROUND_POINT.X-30, BACKGROUND_POINT.Y, "BACKGROUND (3 COINS) ->")
		game.DrawString(screen, BACKGROUND_POINT.X, BACKGROUND_POINT.Y, "🌳")

		game.DrawString(screen, CHILDHOOD_POINT.X-29, CHILDHOOD_POINT.Y, "CHILDHOOD (5 COINS) ->")
		game.DrawString(screen, CHILDHOOD_POINT.X, CHILDHOOD_POINT.Y, "🧸")

		game.DrawString(screen, NOW_POINT.X+9, NOW_POINT.Y, "<- NOW (7 COINS)")
		game.DrawString(screen, NOW_POINT.X, NOW_POINT.Y, "😎")

		game.DrawString(screen, FUTURE_POINT.X+9, FUTURE_POINT.Y, "<- FUTURE (10 COINS)") 
		game.DrawString(screen, FUTURE_POINT.X, FUTURE_POINT.Y, "🏙️")


		screen.Show()

		// movement
		if screen.HasPendingEvent() {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return 
				}
				switch ev.Rune() {
				case 'k', 'w':
					if g.Player.Sprite.Y > TOP_BORDER_Y + 1{
						g.Player.Sprite.Y--
					}
				case 'j', 's':
					if g.Player.Sprite.Y < TOP_BORDER_Y + MAP_HEIGHT - 1 {
						g.Player.Sprite.Y++
					}
				case 'h', 'a':
					if g.Player.Sprite.X > LEFT_BORDER_X + 1 {
						g.Player.Sprite.X--
					}
				case 'l', 'd':
					if g.Player.Sprite.X < LEFT_BORDER_X + MAP_WIDTH - 1 {
						g.Player.Sprite.X++
					}
				}
			}
		}	

		// check collisions with items
		if g.Player.Sprite.X == BACKGROUND_POINT.X && g.Player.Sprite.Y == BACKGROUND_POINT.Y {
			if ITEMS["background"] == struct{}{} {
				g.Player.Sprite.X = 75
				g.Player.Sprite.Y = 20
				Slides(
					screen, 
					&Image{
						Path: "assets/background/sister.JPG",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/background/boat.JPG",
						Width: 150,
						Height: 50,
					},
				)	
			} else if coins >= 3 {
				coins -= 3
				game.AddItem(g.Player.Name, "coin", -3)
				game.PurchaseItem(g.Player.Name, "background", 1)
				
				// Render item screen
				g.Player.Sprite.X = 75
				g.Player.Sprite.Y = 20
				Slides(
					screen, 
					&Image{
						Path: "assets/background/sister.JPG",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/background/boat.JPG",
						Width: 150,
						Height: 50,
					},
				)
			}
		}
		if g.Player.Sprite.X == CHILDHOOD_POINT.X && g.Player.Sprite.Y == CHILDHOOD_POINT.Y {
			if ITEMS["childhood"] == struct{}{} {
				g.Player.Sprite.X = 75
				g.Player.Sprite.Y = 20
				Slides(
					screen, 
					&Image{
						Path: "assets/childhood/nerd.jpg",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/childhood/catnerd.jpg",
						Width: 80,
						Height: 45,
					},
					&Image{
						Path: "assets/childhood/steam.png",
						Width: 80,
						Height: 27,
					},
					&Image{
						Path: "assets/childhood/band.JPG",
						Width: 70,
						Height: 50,
					},
				)	
			} else if coins >= 5 {
				coins -= 5
				game.AddItem(g.Player.Name, "coin", -5)
				game.PurchaseItem(g.Player.Name, "childhood", 1)
				
				// Render item screen
				g.Player.Sprite.X = 75
				g.Player.Sprite.Y = 20
				Slides(
					screen, 
					&Image{
						Path: "assets/childhood/nerd.jpg",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/childhood/catnerd.jpg",
						Width: 80,
						Height: 45,
					},
					&Image{
						Path: "assets/childhood/steam.png",
						Width: 80,
						Height: 27,
					},
					&Image{
						Path: "assets/childhood/band.JPG",
						Width: 70,
						Height: 50,
					},
				)
			}
		}
		if g.Player.Sprite.X == NOW_POINT.X && g.Player.Sprite.Y == NOW_POINT.Y {
			if ITEMS["now"] == struct{}{} {
				g.Player.Sprite.X = 75
				g.Player.Sprite.Y = 20
				Slides(
					screen, 
					&Image{
						Path: "assets/now/sledset.jpg",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/now/nyc.JPG",
						Width: 90,
						Height: 49,
					},
					&Image{
						Path: "assets/now/comedy.JPG",
						Width: 150,
						Height: 50,
					},
					&Image{
						Path: "assets/now/nyc.JPG",
						Width: 90,
						Height: 49,
					},
					&Image{
						Path: "assets/now/comedy.JPG",
						Width: 150,
						Height: 50,
					},
					&Image{
						Path: "assets/now/cave.JPG",
						Width: 100,
						Height: 49,
					},
				)	
			} else if coins >= 7 {
				coins -= 7
				game.AddItem(g.Player.Name, "coin", -7)
				game.PurchaseItem(g.Player.Name, "now", 1)
				
				// Render item screen
				g.Player.Sprite.X = 75
				g.Player.Sprite.Y = 20
				Slides(
					screen, 
					&Image{
						Path: "assets/now/sledset.jpg",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/now/nyc.JPG",
						Width: 90,
						Height: 49,
					},
					&Image{
						Path: "assets/now/comedy.JPG",
						Width: 150,
						Height: 50,
					},
					&Image{
						Path: "assets/now/cave.JPG",
						Width: 100,
						Height: 49,
					},
				)
			}
		}
		if g.Player.Sprite.X == FUTURE_POINT.X && g.Player.Sprite.Y == FUTURE_POINT.Y {
			if ITEMS["future"] == struct{}{} {
				g.Player.Sprite.X = 75
				g.Player.Sprite.Y = 20
				Slides(
					screen, 
					&Image{
						Path: "assets/future/seattle.jpg",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/future/gopher.png",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/future/dart.png",
						Width: 100,
						Height: 50,
					},
				)	
			} else if coins >= 10 {
				coins -= 10
				game.AddItem(g.Player.Name, "coin", -10)
				game.PurchaseItem(g.Player.Name, "future", 1)
				
				// Render item screen
				g.Player.Sprite.X = 75
				g.Player.Sprite.Y = 20
				Slides(
					screen, 
					&Image{
						Path: "assets/future/seattle.jpg",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/future/gopher.png",
						Width: 100,
						Height: 50,
					},
					&Image{
						Path: "assets/future/dart.png",
						Width: 100,
						Height: 50,
					},
				)
			}
		}
			
		// fps counter logic
		frameCount++
		if time.Since(lastFPSUpdate) >= time.Second {
			fps = frameCount
			frameCount = 0
			lastFPSUpdate = time.Now()
		}
	
		<-ticker.C
	}
}