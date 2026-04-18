package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	var g game

	ebiten.SetWindowTitle("Ludum Dare 59")
	ebiten.SetWindowSize(2*globalScreenWidth, 2*globalScreenHeight)

	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
