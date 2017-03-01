package main

import (
	p "github.com/bregydoc/PGoJs/Processing"
)

func setup() {
	p.CreateCanvas(600, 600)
	p.Background(230)
}

func draw() {
	p.StrokeWeight(4)
	if !p.MouseIsPressed {
		p.NoStroke()
	} else {
		p.Stroke("rgba(139,195,74 ,1)")
	}
	p.Line(p.PmouseX, p.PmouseY, p.MouseX, p.MouseY)
}

func main() {
	p.Setup = setup
	p.Draw = draw

	p.LaunchApp()
}
