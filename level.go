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

type signalElement int

const (
	signalElementNum  = 6
	signalElementMax  = signalElementNum - 1
	signalElementNone = signalElementNum + 1
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
	grid                             [globalLevelSizeY][globalLevelSizeX]signalElement
	playerX, playerY, playerProgress int
	cursorX, cursorY                 int
	//moveDistance                     int
	moveRecord []playerState
}
