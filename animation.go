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
	"github.com/hajimehoshi/ebiten/v2"
)

type move struct {
	frames, totalFrames int
	xFrom, yFrom        int
	xTo, yTo            int
	playedMove          bool
	success             bool
	posX, posY          int
}

func newMove(xF, yF, xT, yT int, success bool) (m move) {
	m.frames = 0
	m.xFrom = xF
	m.xTo = xT
	m.yFrom = yF
	m.yTo = yT
	distance := max(xT-xF, xF-xT, yT-yF, yF-yT)
	m.totalFrames = globalMoveFrames*distance - (globalMoveFrames*(distance-1))/2
	m.playedMove = false
	m.success = success
	return
}

func (m *move) update() (done bool) {
	m.frames++

	timeProportion := float64(m.frames) / float64(m.totalFrames)

	m.posX = getPositionFromTime(m.xFrom, m.xTo, globalGridX, timeProportion)
	m.posY = getPositionFromTime(m.yFrom, m.yTo, globalGridY, timeProportion)

	return m.frames >= m.totalFrames
}

func (m move) draw(screen *ebiten.Image) {
	//timeProportion := float64(m.frames) / float64(m.totalFrames)

	//x := getPositionFromTime(m.xFrom, m.xTo, globalGridX, timeProportion)
	//y := getPositionFromTime(m.yFrom, m.yTo, globalGridY, timeProportion)

	freelyDrawPlayer(m.posX, m.posY, 1, screen)
}

func getPositionFromTime(from, to, shift int, time float64) int {
	realFrom := shift + from*globalCellSize
	realTo := shift + to*globalCellSize

	distance := realTo - realFrom

	return realFrom + int(timeFunc(time)*float64(distance))
}

// should return a result between 0 and 1 if time is between 0 and 1
func timeFunc(time float64) float64 {
	//return time * time
	return -((time - 1) * (time - 1)) + 1
}

// add a move animation at the start of an animation slice
func addAnimation(futur []move, m move) (newFutur []move) {
	newFutur = append(futur, m)
	for i := len(futur) - 1; i >= 0; i-- {
		newFutur[i+1] = newFutur[i]
	}
	newFutur[0] = m
	return
}
