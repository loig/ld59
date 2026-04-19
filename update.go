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
			if len(g.rankingChan) > 0 {
				<-g.rankingChan
			}
			go getRank(g.levelNumber, g.rankingChan)
		}
	case gameStateGameOver:
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			g.soundEngine.register(soundOk)
			g.state = gameStateGetRanking
		}
	case gameStateGetRanking, gameStateDisplayRanking:
		if g.updateRanking() {
			g.soundEngine.register(soundOk)
			g.levelNumber = 1
			g.level.getNew(0)
			g.state = gameStateTitle
			g.mustRecord = false
			g.playerRank = 11
			g.selectedChar = 0
		}
	}

	return nil
}

func (g *game) updateRanking() bool {

	switch g.state {
	case gameStateGetRanking:

		select {
		case r := <-g.rankingChan:
			log.Print(r)
			g.ranking = r
			g.state = gameStateDisplayRanking
			g.setupRanking()
		default:
			return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
		}

	case gameStateDisplayRanking:
		if g.mustRecord {
			return g.updateEnterName()
		} else {
			return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
		}
	}

	return false
}

func (g *game) setupRanking() {

	curLevel := g.levelNumber
	curName := nameToString(g.name)
	changed := false

	for pos := 0; pos < len(g.ranking.top10); pos++ {
		if curLevel > g.ranking.top10[pos] {
			if !changed {
				g.playerRank = pos + 1
				changed = true
			}
			tmp := curLevel
			tmpName := curName
			curLevel = g.ranking.top10[pos]
			curName = g.ranking.top10Names[pos]
			g.ranking.top10[pos] = tmp
			g.ranking.top10Names[pos] = tmpName
		}
	}

	if curLevel < g.levelNumber {
		g.mustRecord = true
	}

}

func (g *game) updateEnterName() bool {

	switch {
	case inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft):
		g.selectedChar++
		if g.selectedChar >= 3 {
			go sendRank(nameToString(g.name), g.levelNumber)
		}
		return g.selectedChar >= 3
	case inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight):
		g.selectedChar--
		if g.selectedChar < 0 {
			g.selectedChar = 0
		}
	case isWheelUpJustUsed():
		g.name[g.selectedChar] = (g.name[g.selectedChar] + 1) % 26
	case isWheelDownJustUsed():
		g.name[g.selectedChar] = (g.name[g.selectedChar] + 25) % 26
	}

	return false
}
