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
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

const numPart = 7

type particle struct {
	x, y int
}

type trail [numPart]particle

func (p *particle) update(x, y int) {
	p.x = x
	p.y = y
}

func (t *trail) update(x, y int) {

	xshift := rand.Intn(3) - 1
	yshift := rand.Intn(3) - 2
	x += xshift
	y += yshift

	for i := 0; i < len(t); i++ {
		tmpX, tmpY := t[i].x, t[i].y
		t[i].x = x + xshift
		t[i].y = y + yshift
		x = tmpX
		y = tmpY
	}
}

func (t *trail) init(x, y int) {
	for i := 0; i < len(t); i++ {
		t[i].x = x
		t[i].y = y
	}
}

func (t trail) draw(screen *ebiten.Image) {
	for num := len(t) - 1; num >= 0; num-- {
		part := t[num]
		opacity := float32(numPart-num) / float32(numPart)
		freelyDrawPlayer(part.x, part.y, opacity, screen)
	}
}
