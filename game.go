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

type gameState int

const (
	gameStateTitle gameState = iota
	gameStateHowTo
	gameStateLevelExposition
	gameStateLevelResolution
	gameStateGameOver
	gameStateGetRanking
	gameStateDisplayRanking
)

type game struct {
	state        gameState
	level        level
	levelNumber  int
	soundEngine  soundEngine
	playerTrail  trail
	rankingChan  chan ranking
	ranking      ranking
	mustRecord   bool
	name         [3]uint8
	playerRank   int
	selectedChar int
}

func createGame() game {
	g := game{
		state:       gameStateTitle,
		levelNumber: 1,
	}
	g.level.getNew(0)
	g.soundEngine = newSoundEngine()
	g.rankingChan = make(chan ranking, 1)
	g.name = stringToName("YOU")
	g.playerRank = 11
	return g
}
