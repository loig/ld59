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
	"image"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

func isWheelUpJustUsed() bool {
	_, dy := ebiten.Wheel()
	return dy > 0
}

func isWheelDownJustUsed() bool {
	_, dy := ebiten.Wheel()
	return dy < 0
}

func mouseToGridX(mousePos int) int {
	if mousePos < globalGridX {
		return -1
	}
	return (mousePos - globalGridX) / globalCellSize
}

func mouseToGridY(mousePos int) int {
	if mousePos < globalGridY {
		return -1
	}
	return (mousePos - globalGridY) / globalCellSize
}

func xor(b1, b2 bool) bool {
	return (b1 || b2) && !(b1 && b2)
}

func getNumberAsArray(n int) (res []int) {

	if n == 0 {
		return []int{0}
	}

	for n > 0 {
		res = append(res, n%10)
		n = n / 10
	}

	slices.Reverse(res)

	return
}

func drawNumberAt(x, y, n int, screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))

	nTab := getNumberAsArray(n)

	for _, digit := range nTab {
		imX := digit * 32
		imY := 0
		screen.DrawImage(digitsImage.SubImage(image.Rect(imX, imY, imX+32, imY+32)).(*ebiten.Image), op)
		op.GeoM.Translate(20, 0)
	}

}

func getFrames(lastFrames, levelNumber int) int {
	if levelNumber == 0 {
		return globalMaxFramePerLevel
	}

	if levelNumber%globalFrameStep == 0 {
		newFrames1 := lastFrames / 2
		newFrames2 := lastFrames - globalMaxFrameDecrease
		return max(newFrames1, newFrames2)
	}

	return lastFrames
}
