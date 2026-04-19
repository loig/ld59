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

const (
	globalScreenWidth  = 400
	globalScreenHeight = 400

	// level drawing
	globalGridX        = 13
	globalGridY        = 62
	globalCellSize     = 64
	globalLevelNumX    = 110
	globalLevelNumY    = 10
	globalEndLevelNumX = 270
	globalEndLevelNumY = 125
	globalTutoPlayerX  = 90
	globalTutoPlayerY  = 235

	// game characteristics
	globalLevelSizeX = 5
	globalLevelSizeY = 5

	// times
	globalFramesBeforeFirstSymbol = 40
	globalFramesBeforeLevel       = 10
	globalOpacityFrames           = 10
	globalLevelOpacityFrames      = 20
	globalEndLevelWaitFrames      = 30

	// balancing
	globalMaxFramePerLevel = 600
	globalMaxFrameDecrease = 10
	globalFrameStep        = 2
)

// balancing

var globalSignalLenghtUpdate = [6]int{
	3, 5, 7, 15, 30, 50,
}
var globalFramesPerSignalElement = [7]int{
	60, 50, 40, 35, 30, 25, 20,
}
