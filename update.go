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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *game) Update() error {

	g.soundEngine.playNow()

	mouseX, mouseY := ebiten.CursorPosition()

	switch g.state {
	case gameStateTitle:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.soundEngine.register(soundOk)
			g.state = gameStateHowTo
			g.playerTrail.init(globalTutoPlayerX, globalTutoPlayerY)
		}
	case gameStateHowTo:
		g.playerTrail.update(globalTutoPlayerX, globalTutoPlayerY)
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.soundEngine.register(soundOk)
			g.state = gameStateLevelExposition
		}
	case gameStateLevelExposition:
		done, playNote, note := g.level.updateExposition()
		if playNote {
			g.soundEngine.register(soundNote1 + note)
		}
		if done {
			g.state = gameStateLevelResolution
			g.playerTrail.init(
				globalGridX+g.level.playerX*globalCellSize,
				globalGridY+g.level.playerY*globalCellSize,
			)
		}
	case gameStateLevelResolution:
		finished, dead, playMove, playMiss, playNote, note, x, y := g.level.update(mouseX, mouseY)
		g.playerTrail.update(x, y)
		if playMove {
			g.soundEngine.register(soundJump)
		}
		if playMiss {
			g.soundEngine.register(soundMiss)
		}
		if playNote {
			g.soundEngine.register(soundNote1 + note)
		}
		if finished {
			g.soundEngine.register(soundWon)
			g.level.getNew(g.levelNumber)
			g.levelNumber++
			g.state = gameStateLevelExposition
		}
		if dead {
			g.soundEngine.register(soundLost)
			g.state = gameStateGameOver
		}
	case gameStateGameOver:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.soundEngine.register(soundOk)
			g.levelNumber = 1
			g.level.getNew(0)
			g.state = gameStateTitle
		}
	}

	return nil
}
