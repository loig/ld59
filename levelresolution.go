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
func (l level) draw(trail trail, screen *ebiten.Image) {

	opacity := float32(1)
	if len(l.moveAnimations) <= 0 {
		if l.finished || l.dead {
			opacity = 0
			if l.levelAppearsDone {
				opacity = 1
				if l.levelAppearsFrames <= globalLevelOpacityFrames {
					opacity = float32(l.levelAppearsFrames) / float32(globalLevelOpacityFrames)
				}
			}
		} else {
			if !l.levelAppearsReady {
				opacity = 0
			}
			if l.levelAppearsReady && !l.levelAppearsDone {
				opacity = float32(globalLevelOpacityFrames-l.levelAppearsFrames) / float32(globalLevelOpacityFrames)
			}
		}
	}

	drawBack(screen)

	//lBack := true
	for y, line := range l.grid {
		//back := lBack
		for x, s := range line {
			//drawBack(x, y, back, screen)
			s.draw(x, y, l.cursorX == x && l.cursorY == y, opacity, screen)
			//back = !back
		}
		//lBack = !lBack
	}

	trail.draw(screen)
	if len(l.moveAnimations) > 0 {
		l.moveAnimations[len(l.moveAnimations)-1].draw(screen)
	} else {
		drawPlayer(l.playerX, l.playerY, opacity, screen)
	}

	drawCursor(l.cursorX, l.cursorY, xor(l.cursorX == l.playerX, l.cursorY == l.playerY), opacity, screen)

	drawCountDown(l.framesLeft, l.frames, screen)
}

func drawCountDown(left, total int, screen *ebiten.Image) {
	ySize := (globalScreenHeight * left) / total

	op := &ebiten.DrawImageOptions{}
	imX := 0
	imY := 0
	screen.DrawImage(countdownImage.SubImage(image.Rect(imX, imY, imX+globalScreenWidth, imY+ySize)).(*ebiten.Image), op)

}

/*
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
*/

func drawBack(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(fondImage, op)
}

func (s signalElement) draw(x, y int, big bool, opacity float32, screen *ebiten.Image) {
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
	op.ColorScale.ScaleAlpha(opacity)
	imX := int(globalCellSize * s)
	imY := 0
	screen.DrawImage(images.SubImage(image.Rect(imX, imY, imX+globalCellSize, imY+globalCellSize)).(*ebiten.Image), op)
}

func drawPlayer(x, y int, opacity float32, screen *ebiten.Image) {

	xx := globalGridX + x*globalCellSize
	yy := globalGridY + y*globalCellSize

	freelyDrawPlayer(xx, yy, opacity, screen)
}

func freelyDrawPlayer(xx, yy int, opacity float32, screen *ebiten.Image) {
	if xx >= globalGridX-5 && yy >= globalGridY-5 {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(xx), float64(yy))
		op.ColorScale.ScaleAlpha(opacity)
		imX := 0
		imY := globalCellSize
		screen.DrawImage(images.SubImage(image.Rect(imX, imY, imX+globalCellSize, imY+globalCellSize)).(*ebiten.Image), op)
	}
}

func drawCursor(x, y int, reachable bool, opacity float32, screen *ebiten.Image) {
	if opacity >= 1 && reachable && x >= 0 && x < globalLevelSizeX && y >= 0 && y < globalLevelSizeY {
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

// Updating level
func (l *level) update(mouseX, mouseY int) (levelFinished, dead, playMove, playMiss, playNote bool, note int, xTrail, yTrail int) {

	xTrail = globalGridX + l.playerX*globalCellSize
	yTrail = globalGridY + l.playerY*globalCellSize

	if len(l.moveAnimations) > 0 {
		current := l.moveAnimations[len(l.moveAnimations)-1]
		x, y := current.xTo, current.yTo
		xTrail = current.posX
		yTrail = current.posY

		if !current.playedMove {
			playMove = true
			l.moveAnimations[len(l.moveAnimations)-1].playedMove = true
		}

		if l.moveAnimations[len(l.moveAnimations)-1].update() {

			if current.success {
				if l.grid[y][x] != signalElementNone {
					playNote = true
					note = int(l.grid[y][x])
				}
				l.grid[y][x] = signalElementNone
			} else {
				playMiss = true
			}

			l.moveAnimations = l.moveAnimations[:len(l.moveAnimations)-1]
		}
		if !l.finished && !l.dead {
			oldPlayMiss := playMiss
			l.finished, playMiss = l.updatePlayer(mouseX, mouseY)
			playMiss = playMiss || oldPlayMiss
			if l.finished {
				l.levelAppearsFrames = globalLevelOpacityFrames + globalEndLevelWaitFrames
			}
		}
		return false, false, playMove, playMiss, playNote, note, xTrail, yTrail
	}

	if l.finished || l.dead {
		if l.levelAppearsDone {
			l.levelAppearsFrames--
			if l.levelAppearsFrames <= 0 {
				l.levelAppearsDone = false
				l.levelAppearsFrames = globalFramesBeforeLevel
			}
			return false, false, playMove, playMiss, playNote, note, xTrail, yTrail
		}
		if l.levelAppearsReady {
			l.levelAppearsFrames--
			if l.levelAppearsFrames <= 0 {
				l.levelAppearsReady = false
			}
			return false, false, playMove, playMiss, playNote, note, xTrail, yTrail
		}
		return l.finished, l.dead, playMove, playMiss, playNote, note, xTrail, yTrail
	}

	if !l.levelAppearsReady {
		l.levelAppearsFrames--
		if l.levelAppearsFrames <= 0 {
			l.levelAppearsFrames = globalLevelOpacityFrames
			l.levelAppearsReady = true
		}
		return false, false, playMove, playMiss, playNote, note, xTrail, yTrail
	}

	if !l.levelAppearsDone {
		l.levelAppearsFrames--
		if l.levelAppearsFrames <= 0 {
			l.levelAppearsDone = true
			l.levelAppearsFrames = globalLevelOpacityFrames
		}
		return false, false, playMove, playMiss, playNote, note, xTrail, yTrail
	}

	l.framesLeft--
	if l.framesLeft == 0 {
		l.dead = true
		l.levelAppearsFrames = globalLevelOpacityFrames + globalEndLevelWaitFrames
		return false, false, playMove, playMiss, playNote, note, xTrail, yTrail
	}

	oldPlayMiss := playMiss
	l.finished, playMiss = l.updatePlayer(mouseX, mouseY)
	playMiss = playMiss || oldPlayMiss
	if l.finished {
		l.levelAppearsFrames = globalLevelOpacityFrames + globalEndLevelWaitFrames
	}

	return false, false, playMove, playMiss, playNote, note, xTrail, yTrail
}

// Updating player position
func (l *level) updatePlayer(mouseX, mouseY int) (levelFinished, playMiss bool) {

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
		return l.movePlayer(l.cursorX, l.cursorY), false
	case inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight):
		l.revertMove()
		playMiss = true
	}

	return false, playMiss
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
		l.moveAnimations = addAnimation(l.moveAnimations, newMove(l.playerX, l.playerY, newPosX, newPosY, false))
		l.moveAnimations = addAnimation(l.moveAnimations, newMove(newPosX, newPosY, l.playerX, l.playerY, true))
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

	// setup animation
	l.moveAnimations = addAnimation(l.moveAnimations, newMove(l.playerX, l.playerY, newPosX, newPosY, true))

	// apply move
	l.playerProgress++
	l.playerX = newPosX
	l.playerY = newPosY
	//l.moveDistance = 0
	//l.grid[l.playerY][l.playerX] = signalElementNone

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
