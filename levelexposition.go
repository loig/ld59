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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func (l level) drawExposition(screen *ebiten.Image) {

	if l.signalPosition < len(l.signal) {
		l.signal[l.signalPosition].drawFreely(l.signalX, l.signalY, false, screen)
	}

}

func (l *level) updateExposition() (done bool) {

	if !l.expositionDone {
		l.signalElementFramesLeft--
		if l.signalElementFramesLeft <= 0 {
			l.signalElementFramesLeft = l.signalElementFrames
			l.signalPosition++
			l.signalX, l.signalY = getNewSignalPlace(l.signalX, l.signalY)
			l.expositionDone = l.signalPosition >= len(l.signal)
		}
	}

	return l.expositionDone
}

func (s signalElement) drawFreely(x, y int, big bool, screen *ebiten.Image) {
	xx := float64(x)
	yy := float64(y)

	op := &ebiten.DrawImageOptions{}
	if big {
		scaleFactor := 1.2
		xx = xx - globalCellSize*scaleFactor/2 + globalCellSize/2
		yy = yy - globalCellSize*scaleFactor/2 + globalCellSize/2
		op.GeoM.Scale(scaleFactor, scaleFactor)
	}
	op.GeoM.Translate(xx, yy)
	imX := int(globalCellSize * s)
	imY := 0
	screen.DrawImage(images.SubImage(image.Rect(imX, imY, imX+globalCellSize, imY+globalCellSize)).(*ebiten.Image), op)
}

func getNewSignalPlace(x, y int) (xx, yy int) {
	return farButNotMuch(x, globalScreenWidth),
		farButNotMuch(y, globalScreenHeight)
}

func farButNotMuch(pos, posMax int) (newPos int) {
	shift := rand.Intn(posMax/2) + posMax/4
	newPos = (pos + shift) % posMax
	if newPos < 2*globalCellSize {
		newPos = 2 * globalCellSize
	}
	if newPos > posMax-2*globalCellSize {
		newPos = posMax - 2*globalCellSize
	}
	return
}
