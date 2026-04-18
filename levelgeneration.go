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
	//l.moveDistance = 0
	l.moveRecord = l.moveRecord[:0]

	l.generateSignal(levelNumber)

	l.playerX, l.playerY = l.generateGrid()

	l.frames = getFrames(l.frames, levelNumber)
	l.framesLeft = l.frames

	l.signalElementFrames = globalFramesPerSignalElement[len(l.signal)-1]
	l.signalElementFramesLeft = l.signalElementFrames
	l.signalPosition = 0
	l.expositionDone = false
	l.signalX = (globalScreenWidth - globalCellSize) / 2
	l.signalY = (globalScreenHeight - globalCellSize) / 2
}

func (l *level) generateSignal(levelNumber int) {

	signalLen := 1
	for _, num := range globalSignalLenghtUpdate {
		if levelNumber < num {
			break
		}
		signalLen++
	}

	if len(l.signal) == 0 {
		l.signal = make([]signalElement, signalLen)
	}

	for pos := 0; pos < signalLen; pos++ {
		newSignalElement := signalElement(rand.Intn(signalElementNum))
		if pos < len(l.signal) {
			l.signal[pos] = newSignalElement
		} else {
			l.signal = append(l.signal, newSignalElement)
		}
	}

	l.signal = l.signal[:signalLen]
}

func (l *level) generateGrid() (startX, startY int) {

	// Step 0: erase grid
	for y := 0; y < globalLevelSizeY; y++ {
		for x := 0; x < globalLevelSizeX; x++ {
			l.grid[y][x] = signalElementNone
		}
	}

	// Step 1: generate solution
	startX = rand.Intn(globalLevelSizeX)
	startY = rand.Intn(globalLevelSizeY)

	currentX, currentY := startX, startY

	// record data for avoiding duplicate solutions later on
	forbidenSignalElementPerX := make(map[int][]signalElement)
	forbidenSignalElementPerY := make(map[int][]signalElement)

	// number of possible moves from a position while staying in the grid
	// including the position itself, counted twice, for simplicity and for
	// slightly higher chance that steps of a solution are close in the grid
	numMoves := globalLevelSizeX + globalLevelSizeY

	// this is not robust: avoid long signals (anything shorter than globalLevelSizeX + globalLevelSizeY should be safe)
	for _, element := range l.signal {
		next := rand.Intn(numMoves)
		for try := 0; try < numMoves; try++ {
			// check if the position corresponding to next is suitable
			newX := currentX
			newY := currentY
			if next < globalLevelSizeX {
				newX = next
			} else {
				newY = next - globalLevelSizeX
			}
			if (newX != startX || newY != startY) &&
				(newX != currentX || newY != currentY) &&
				l.grid[newY][newX] == signalElementNone {
				// this is a suitable position
				// record info in order to prevent branching at this point in the solution from the later generated remaining of the grid
				if currentX != startX || currentY != startY {
					forbidenSignalElementPerX[newX] = append(forbidenSignalElementPerX[newX], l.grid[currentY][currentX])
					forbidenSignalElementPerY[newY] = append(forbidenSignalElementPerY[newY], l.grid[currentY][currentX])
				}
				// setup the solution
				l.grid[newY][newX] = element
				currentX = newX
				currentY = newY
				break
			}
			next = (next + 1) % numMoves
		}
	}

	// Step 2: generate remaining of the grid, while avoiding multiple solutions
	for y := 0; y < globalLevelSizeY; y++ {
		for x := 0; x < globalLevelSizeX; x++ {
			if l.grid[y][x] == signalElementNone && (x != startX || y != startY) {
				// need to add an element
				allowed := make([]bool, signalElementNum)
				for i := 0; i < len(allowed); i++ {
					allowed[i] = true
				}
				numAllowed := signalElementNum

				// avoid branching into the existing solution
				for _, s := range forbidenSignalElementPerX[x] {
					if allowed[s] {
						allowed[s] = false
						numAllowed--
					}
				}
				for _, s := range forbidenSignalElementPerY[y] {
					if allowed[s] {
						allowed[s] = false
						numAllowed--
					}
				}

				// avoid creating a new solution from scratch
				lastInSignal := l.signal[len(l.signal)-1]
				if len(l.signal) > 1 {
					beforeLastInSignal := l.signal[len(l.signal)-2]
					if l.isReachable(lastInSignal, x, y) {
						if allowed[beforeLastInSignal] {
							numAllowed--
							allowed[beforeLastInSignal] = false
						}
					}
					if l.isReachable(beforeLastInSignal, x, y) {
						if allowed[lastInSignal] {
							numAllowed--
							allowed[lastInSignal] = false
						}
					}
				} else {
					if x == startX || y == startY {
						if allowed[lastInSignal] {
							numAllowed--
							allowed[lastInSignal] = false
						}
					}
				}

				if numAllowed <= 0 {
					log.Print("Oups, pas assez d'éléments possibles pour cette position", x, y)
					l.grid[y][x] = signalElementNone
					continue
				}

				// choose element to add
				numElement := rand.Intn(numAllowed)
				pos := 0
				for numElement > 0 || !allowed[pos] {
					if allowed[pos] {
						numElement--
					}
					pos++
				}
				l.grid[y][x] = signalElement(pos)

			}
		}
	}

	return
}

// check if signal element is directly reachable from position (x, y)
func (l level) isReachable(s signalElement, x, y int) bool {
	for yy := 0; yy < globalLevelSizeY; yy++ {
		if yy != y {
			if l.grid[yy][x] == s {
				return true
			}
		}
	}
	for xx := 0; xx < globalLevelSizeX; xx++ {
		if xx != x {
			if l.grid[y][xx] == s {
				return true
			}
		}
	}
	return false
}
