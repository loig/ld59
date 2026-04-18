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
	"log"
	"math/rand"
)

func (l *level) getNew(levelNumber int) {
	l.playerProgress = 0
	l.moveDistance = 0
	l.moveRecord = l.moveRecord[:0]

	l.generateSignal()

	l.playerX, l.playerY = l.generateGrid()
}

func (l *level) generateSignal() {
	if len(l.signal) == 0 {
		l.signal = make([]signalElement, 3)
	}

	for pos := 0; pos < len(l.signal); pos++ {
		l.signal[pos] = signalElement(rand.Intn(signalElementNum))
	}
}

func (l *level) generateGrid() (startX, startY int) {
	if len(l.grid) == 0 {
		l.grid = make([][]signalElement, 0)
	}

	gridSize := rand.Intn(10) + 5
	log.Print(gridSize)

	startX = gridSize / 2
	startY = gridSize / 2

	for y := 0; y < gridSize; y++ {
		if len(l.grid) < y+1 {
			l.grid = append(l.grid, make([]signalElement, 0))
		}
		for x := 0; x < gridSize; x++ {
			element := signalElement(rand.Intn(signalElementNum))
			if x == startX && y == startY {
				element = signalElementNone
			}
			if len(l.grid[y]) < x+1 {
				l.grid[y] = append(l.grid[y], element)
			} else {
				l.grid[y][x] = element
			}
		}
		l.grid[y] = l.grid[y][:gridSize]
	}

	l.grid = l.grid[:gridSize]

	return
}
