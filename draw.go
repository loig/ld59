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

import "github.com/hajimehoshi/ebiten/v2"

func (g *game) Draw(screen *ebiten.Image) {

	switch g.state {
	case gameStateTitle:
		drawTitle(screen)
	case gameStateHowTo:
		drawHowTo(screen)
		g.playerTrail.draw(screen)
		freelyDrawPlayer(globalTutoPlayerX, globalTutoPlayerY, 1, screen)
	case gameStateLevelExposition:
		g.level.drawExposition(screen)
		drawNumberAt(globalLevelNumX, globalLevelNumY, g.levelNumber, screen)
	case gameStateLevelResolution:
		g.level.draw(g.playerTrail, screen)
		drawNumberAt(globalLevelNumX, globalLevelNumY, g.levelNumber, screen)
	case gameStateGameOver:
		drawGameOver(screen)
		drawNumberAt(globalEndLevelNumX, globalEndLevelNumY, g.levelNumber, screen)
	}
}

func drawTitle(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(titleImage, op)
}

func drawHowTo(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(howtoImage, op)
}

func drawGameOver(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(gameoverImage, op)
}
