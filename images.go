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

//go:embed assets/ranking-waiting.png
var rank1Bytes []byte
var rank1Image *ebiten.Image

//go:embed assets/ranking-display.png
var rank2Bytes []byte
var rank2Image *ebiten.Image

//go:embed assets/ranking-display-normal.png
var rank3Bytes []byte
var rank3Image *ebiten.Image

//go:embed assets/alphabet.png
var alphabetBytes []byte
var alphabetImage *ebiten.Image

//go:embed assets/score.png
var scoreBytes []byte
var scoreImage *ebiten.Image

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

	decoded, _, err = image.Decode(bytes.NewReader(rank1Bytes))
	if err != nil {
		log.Fatal(err)
	}
	rank1Image = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(rank2Bytes))
	if err != nil {
		log.Fatal(err)
	}
	rank2Image = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(rank3Bytes))
	if err != nil {
		log.Fatal(err)
	}
	rank3Image = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(alphabetBytes))
	if err != nil {
		log.Fatal(err)
	}
	alphabetImage = ebiten.NewImageFromImage(decoded)

	decoded, _, err = image.Decode(bytes.NewReader(scoreBytes))
	if err != nil {
		log.Fatal(err)
	}
	scoreImage = ebiten.NewImageFromImage(decoded)
}
