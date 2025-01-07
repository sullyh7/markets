package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Panel struct {
	position      rl.Vector2
	width, height float32
	title         string
}

func NewPanel(title string, pos rl.Vector2, width, height float32) *Panel {
	return &Panel{
		title:    title,
		position: pos,
		width:    width,
		height:   height,
	}
}

func (p *Panel) update() {

}

func (p *Panel) getDrawPos(x, y float32) rl.Vector2 {
	return rl.NewVector2(p.position.X+x, p.position.Y-y+p.height)
}

func (p *Panel) render() {
	gui.Panel(rl.NewRectangle(p.position.X, p.position.Y, p.width, p.height), p.title)
	pos := p.getDrawPos(10, 10)
	rl.DrawCircle(int32(pos.X), int32(pos.Y), 2, rl.Red)
}
