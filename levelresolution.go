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

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Drawing a level
func (l level) draw(screen *ebiten.Image) {
	lBack := true
	for y, line := range l.grid {
		back := lBack
		for x, s := range line {
			drawBack(x, y, back, screen)
			s.draw(x, y, l.cursorX == x && l.cursorY == y, screen)
			back = !back
		}
		lBack = !lBack
	}

	drawPlayer(l.playerX, l.playerY, screen)
	drawCursor(l.cursorX, l.cursorY, xor(l.cursorX == l.playerX, l.cursorY == l.playerY), screen)
}

func drawBack(x, y int, back bool, screen *ebiten.Image) {
	xx := float64(globalGridX + x*globalCellSize)
	yy := float64(globalGridY + y*globalCellSize)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(xx, yy)
	imX := 2 * globalCellSize
	if !back {
		imX += globalCellSize
	}
	imY := globalCellSize
	screen.DrawImage(images.SubImage(image.Rect(imX, imY, imX+globalCellSize, imY+globalCellSize)).(*ebiten.Image), op)
}

func (s signalElement) draw(x, y int, big bool, screen *ebiten.Image) {
	xx := float64(globalGridX + x*globalCellSize)
	yy := float64(globalGridY + y*globalCellSize)

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

func drawPlayer(x, y int, screen *ebiten.Image) {

	xx := globalGridX + x*globalCellSize
	yy := globalGridY + y*globalCellSize

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xx), float64(yy))
	imX := 0
	imY := globalCellSize
	screen.DrawImage(images.SubImage(image.Rect(imX, imY, imX+globalCellSize, imY+globalCellSize)).(*ebiten.Image), op)
}

func drawCursor(x, y int, reachable bool, screen *ebiten.Image) {
	if reachable {
		xx := globalGridX + x*globalCellSize
		yy := globalGridY + y*globalCellSize

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(xx), float64(yy))
		imX := globalCellSize
		imY := globalCellSize
		screen.DrawImage(images.SubImage(image.Rect(imX, imY, imX+globalCellSize, imY+globalCellSize)).(*ebiten.Image), op)
	}
}

/*
func (l level) drawMoveDistances(x, y, distance int, screen *ebiten.Image) {

	// up
	if y-distance >= 0 {
		drawMoveDistance(x, y-distance, 0, screen)
	}

	// down
	if y+distance < len(l.grid) {
		drawMoveDistance(x, y+distance, 2, screen)
	}

	// left
	if x-distance >= 0 {
		drawMoveDistance(x-distance, y, 3, screen)
	}

	// right
	if x+distance < len(l.grid[l.playerY]) {
		drawMoveDistance(x+distance, y, 1, screen)
	}

}

func drawMoveDistance(x, y, direction int, screen *ebiten.Image) {

	xx := globalGridX + x*globalCellSize
	yy := globalGridY + y*globalCellSize

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(xx), float64(yy))
	imX := globalCellSize
	imY := globalCellSize
	screen.DrawImage(images.SubImage(image.Rect(imX, imY, imX+globalCellSize, imY+globalCellSize)).(*ebiten.Image), op)
	imX += globalCellSize + direction*globalCellSize
	screen.DrawImage(images.SubImage(image.Rect(imX, imY, imX+globalCellSize, imY+globalCellSize)).(*ebiten.Image), op)
}
*/

// Updating player position
func (l *level) updatePlayer(mouseX, mouseY int) (levelFinished bool) {

	l.cursorX = mouseToGridX(mouseX)
	l.cursorY = mouseToGridY(mouseY)

	/*
		switch {
		case inpututil.IsKeyJustPressed(ebiten.KeyUp):
			return l.movePlayer(l.playerX, l.playerY-1-l.moveDistance)
		case inpututil.IsKeyJustPressed(ebiten.KeyDown):
			return l.movePlayer(l.playerX, l.playerY+1+l.moveDistance)
		case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
			return l.movePlayer(l.playerX-1-l.moveDistance, l.playerY)
		case inpututil.IsKeyJustPressed(ebiten.KeyRight):
			return l.movePlayer(l.playerX+1+l.moveDistance, l.playerY)
		case inpututil.IsKeyJustPressed(ebiten.KeyBackspace) ||
			inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight):
			l.revertMove()
		case inpututil.IsKeyJustPressed(ebiten.KeyTab) ||
			isWheelUpJustUsed():
			l.updateMoveDistance(true)
		case inpututil.IsKeyJustPressed(ebiten.KeyShiftLeft) ||
			inpututil.IsKeyJustPressed(ebiten.KeyShiftRight) ||
			isWheelDownJustUsed():
			l.updateMoveDistance(false)
		}
	*/

	switch {
	case inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft):
		return l.movePlayer(l.cursorX, l.cursorY)
	case inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight):
		l.revertMove()
	}

	return false
}

func (l *level) movePlayer(newPosX, newPosY int) (levelFinished bool) {
	possibleMove := newPosX == l.playerX || newPosY == l.playerY

	// move is not possible
	if !possibleMove {
		return
	}

	// check if move is in grid and to colored tile
	if newPosY < 0 || newPosY >= len(l.grid) ||
		newPosX < 0 || newPosX >= len(l.grid[newPosY]) ||
		l.grid[newPosY][newPosX] > signalElementMax {
		possibleMove = false
	}

	// move is not possible
	if !possibleMove {
		return
	}

	// check actual signal progress
	possibleMove = l.grid[newPosY][newPosX] == l.signal[l.playerProgress]

	// move is not possible
	if !possibleMove {
		return
	}

	// record move
	l.moveRecord = append(l.moveRecord, playerState{
		lastX: l.playerX, lastY: l.playerY,
		x: newPosX, y: newPosY,
		lastProgress: l.playerProgress,
		progress:     l.playerProgress + 1,
		gridErased:   l.grid[newPosY][newPosX],
	})

	// apply move
	l.playerProgress++
	l.playerX = newPosX
	l.playerY = newPosY
	//l.moveDistance = 0
	l.grid[l.playerY][l.playerX] = signalElementNone

	// end of signal: level completed
	return l.playerProgress >= len(l.signal)
}

func (l *level) revertMove() {
	if len(l.moveRecord) > 0 {
		record := l.moveRecord[len(l.moveRecord)-1]
		l.moveRecord = l.moveRecord[:len(l.moveRecord)-1]
		l.grid[record.y][record.x] = record.gridErased
		l.playerX = record.lastX
		l.playerY = record.lastY
		l.playerProgress = record.lastProgress
		//l.moveDistance = 0
	}
}

/*
func (l *level) updateMoveDistance(inc bool) {
	step := 1
	if !inc {
		step = -1
	}

	l.moveDistance += step

	if l.moveDistance < 0 {
		l.moveDistance = 0
	}

	if l.moveDistance >= l.maxPossibleMove() {
		l.moveDistance = l.maxPossibleMove() - 1
	}
}

func (l level) maxPossibleMove() int {

	upY := l.playerY
	downY := len(l.grid) - l.playerY - 1

	leftX := l.playerX
	rightX := len(l.grid[l.playerY]) - l.playerX - 1

	return max(upY, downY, leftX, rightX)
}
*/
