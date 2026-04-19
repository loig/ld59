/*
LD59, a game for Ludum Dare 59
Copyright (C) 2026 Loïg Jezequel

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"bytes"
	_ "embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/images.png
var imagesBytes []byte
var images *ebiten.Image

//go:embed assets/fond.png
var fondBytes []byte
var fondImage *ebiten.Image

//go:embed assets/digits.png
var digitsBytes []byte
var digitsImage *ebiten.Image

//go:embed assets/countdown.png
var countdownBytes []byte
var countdownImage *ebiten.Image

//go:embed assets/title.png
var titleBytes []byte
var titleImage *ebiten.Image

//go:embed assets/gameover.png
var gameoverBytes []byte
var gameoverImage *ebiten.Image

//go:embed assets/howto.png
var howtoBytes []byte
var howtoImage *ebiten.Image

// load images
func loadImages() {
	decoded, _, err := image.Decode(bytes.NewReader(imagesBytes))
	if err != nil {
		log.Fatal(err)
	}
	images = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(fondBytes))
	if err != nil {
		log.Fatal(err)
	}
	fondImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(digitsBytes))
	if err != nil {
		log.Fatal(err)
	}
	digitsImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(countdownBytes))
	if err != nil {
		log.Fatal(err)
	}
	countdownImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(titleBytes))
	if err != nil {
		log.Fatal(err)
	}
	titleImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(gameoverBytes))
	if err != nil {
		log.Fatal(err)
	}
	gameoverImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(howtoBytes))
	if err != nil {
		log.Fatal(err)
	}
	howtoImage = ebiten.NewImageFromImage(decoded)
}
