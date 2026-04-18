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
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type signalElement int

const (
	signalElementMax  = 5
	signalElementNone = signalElementMax + 1
)

type playerState struct {
	lastX, lastY int
	x, y         int
	lastProgress int
	progress     int
	gridErased   signalElement
}

type level struct {
	signal                           []signalElement
	grid                             [][]signalElement
	playerX, playerY, playerProgress int
	moveDistance                     int
	moveRecord                       []playerState
}

// Creating a level
func createLevel() level {
	l := level{
		signal: []signalElement{0, 1, 0},
		grid: [][]signalElement{
			{0, 1, 2, 0, 0, 0, 2},
			{0, 1, 2, 0, 0, 0, 2},
			{0, 1, 2, 0, 0, 0, 2},
			{1, 2, 2, signalElementNone, 0, 1, 1},
			{2, 0, 1, 1, 1, 2, 1},
			{0, 1, 2, 0, 0, 0, 1},
			{0, 1, 2, 0, 0, 0, 1},
		},
		playerProgress: 0,
	}

	l.playerY = len(l.grid) / 2
	if len(l.grid) > 0 {
		l.playerX = len(l.grid[0]) / 2
	}

	return l
}

// Drawing a level
func (l level) draw(screen *ebiten.Image) {
	for y, line := range l.grid {
		for x, s := range line {
			s.draw(x, y, screen)
		}
	}

	drawPlayer(l.playerX, l.playerY, screen)
	l.drawMoveDistances(l.playerX, l.playerY, l.moveDistance+1, screen)
}

func (s signalElement) draw(x, y int, screen *ebiten.Image) {
	xx := globalGridX + x*globalCellSize
	yy := globalGridY + y*globalCellSize
	col := color.RGBA{R: 0, G: 0, B: 0, A: 255}
	switch s {
	case 0:
		col.R = 255
	case 1:
		col.B = 255
	case 2:
		col.G = 255
	}

	vector.FillRect(screen, float32(xx), float32(yy), globalCellSize, globalCellSize, col, false)
}

func drawPlayer(x, y int, screen *ebiten.Image) {

	xx := globalGridX + x*globalCellSize + globalCellSize/4
	yy := globalGridY + y*globalCellSize + globalCellSize/4

	vector.FillRect(screen, float32(xx), float32(yy), globalCellSize/2, globalCellSize/2, color.White, false)
}

func (l level) drawMoveDistances(x, y, distance int, screen *ebiten.Image) {

	// up
	if y-distance >= 0 {
		drawMoveDistance(x, y-distance, screen)
	}

	// down
	if y+distance < len(l.grid) {
		drawMoveDistance(x, y+distance, screen)
	}

	// left
	if x-distance >= 0 {
		drawMoveDistance(x-distance, y, screen)
	}

	// right
	if x+distance < len(l.grid[l.playerY]) {
		drawMoveDistance(x+distance, y, screen)
	}

}

func drawMoveDistance(x, y int, screen *ebiten.Image) {

	xx := globalGridX + x*globalCellSize + globalCellSize/4
	yy := globalGridY + y*globalCellSize + globalCellSize/4

	vector.FillRect(screen, float32(xx), float32(yy), globalCellSize/2, globalCellSize/2, color.Black, false)
}

// Updating player position
func (l *level) updatePlayer() (levelFinished bool) {

	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyUp):
		return l.movePlayer(l.playerX, l.playerY-1-l.moveDistance)
	case inpututil.IsKeyJustPressed(ebiten.KeyDown):
		return l.movePlayer(l.playerX, l.playerY+1+l.moveDistance)
	case inpututil.IsKeyJustPressed(ebiten.KeyLeft):
		return l.movePlayer(l.playerX-1-l.moveDistance, l.playerY)
	case inpututil.IsKeyJustPressed(ebiten.KeyRight):
		return l.movePlayer(l.playerX+1+l.moveDistance, l.playerY)
	case inpututil.IsKeyJustPressed(ebiten.KeyBackspace):
		l.revertMove()
	case inpututil.IsKeyJustPressed(ebiten.KeyTab):
		l.updateMoveDistance(true)
	case inpututil.IsKeyJustPressed(ebiten.KeyShiftLeft) || inpututil.IsKeyJustPressed(ebiten.KeyShiftRight):
		l.updateMoveDistance(false)
	}

	return false
}

func (l *level) movePlayer(newPosX, newPosY int) (levelFinished bool) {
	possibleMove := true

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
	l.moveDistance = 0
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
	}
}

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
