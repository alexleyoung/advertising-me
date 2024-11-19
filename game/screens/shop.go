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

type ShopItem struct {
	Name     string
	Cost     int
	Symbol   string
	Position game.Point
	LabelX   int
	Label    string
	Slides   []*Image
}

func handleItemPurchase(screen tcell.Screen, g *game.Game, items map[string]struct{}, item ShopItem, coins int) {
	if _, exists := items[item.Name]; exists {
		g.Player.Sprite.X = 75
		g.Player.Sprite.Y = 20
		Slides(screen, item.Slides)
		return
	}
	
	if coins >= item.Cost {
		coins -= item.Cost
		game.AddItem(g.Player.Name, "coin", -item.Cost)
		game.PurchaseItem(g.Player.Name, item.Name, 1)
		
		g.Player.Sprite.X = 75
		g.Player.Sprite.Y = 20
		Slides(screen, item.Slides)
	}
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

	shopItems := []ShopItem{
		{
			Name:     "background",
			Cost:     3,
			Symbol:   "🌳",
			Position: BACKGROUND_POINT,
			LabelX:   BACKGROUND_POINT.X - 30,
			Label:    "BACKGROUND (3 COINS) ->",
			Slides: []*Image{
				{Path: "assets/background/sister.JPG", Width: 100, Height: 50},
				{Path: "assets/background/boat.JPG", Width: 150, Height: 50},
			},
		},
		{
			Name:     "childhood",
			Cost:     5,
			Symbol:   "🧸",
			Position: CHILDHOOD_POINT,
			LabelX:   CHILDHOOD_POINT.X - 29,
			Label:    "CHILDHOOD (5 COINS) ->",
			Slides: []*Image{
				{Path: "assets/childhood/nerd.jpg", Width: 100, Height: 50},
				{Path: "assets/childhood/catnerd.jpg", Width: 80, Height: 45},
				{Path: "assets/childhood/steam.png", Width: 80, Height: 27},
				{Path: "assets/childhood/band.JPG", Width: 70, Height: 50},
			},
		},
		{
			Name:     "now",
			Cost:     7,
			Symbol:   "😎",
			Position: NOW_POINT,
			LabelX:   NOW_POINT.X + 9,
			Label:    "<- NOW (7 COINS)",
			Slides: []*Image{
				{Path: "assets/now/sledset.jpg", Width: 100, Height: 50},
				{Path: "assets/now/nyc.JPG", Width: 90, Height: 49},
				{Path: "assets/now/comedy.JPG", Width: 150, Height: 50},
				{Path: "assets/now/cave.JPG", Width: 100, Height: 49},
			},
		},
		{
			Name:     "future",
			Cost:     10,
			Symbol:   "🏙️",
			Position: FUTURE_POINT,
			LabelX:   FUTURE_POINT.X + 9,
			Label:    "<- FUTURE (10 COINS)",
			Slides: []*Image{
				{Path: "assets/future/seattle.jpg", Width: 100, Height: 50},
				{Path: "assets/future/gopher.png", Width: 100, Height: 50},
				{Path: "assets/future/dart.png", Width: 100, Height: 50},
			},
		},
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
		for _, item := range shopItems {
			game.DrawString(screen, item.LabelX, item.Position.Y, item.Label)
			game.DrawString(screen, item.Position.X, item.Position.Y, item.Symbol)
		}

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
		for _, item := range shopItems {
			if g.Player.Sprite.X == item.Position.X && g.Player.Sprite.Y == item.Position.Y {
				handleItemPurchase(screen, g, ITEMS, item, coins)
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