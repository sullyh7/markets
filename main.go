package main

import (
	"fmt"
	"time"

	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	latestPrice int
)

func main() {

	rl.InitWindow(1200, 800, "Markets")
	defer rl.CloseWindow()

	rl.SetTargetFPS(30)
	gui.SetStyle(0, gui.BACKGROUND_COLOR, 0x2d2d2dff)
	// panel := NewPanel("test", rl.NewVector2(100, 100), 800, 600)
	latestPrice = 100

	// testing
	go func() {
		for {
			latestPrice += 1
			time.Sleep(time.Second * 1)
		}

	}()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.GetColor(0x2d2d2dff))
		// panel.render()
		fontSize := 400
		textWidth := rl.MeasureText(fmt.Sprint(latestPrice), int32(fontSize))
		rl.DrawText(fmt.Sprint((latestPrice)), int32(rl.GetScreenWidth())/2-textWidth/2, int32(rl.GetScreenHeight())/2-int32(fontSize)/2, int32(fontSize), rl.Blue)
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
