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

func (l level) drawExposition(screen *ebiten.Image) {

	for pos, element := range l.signal {
		element.draw(pos, 0, screen)
	}

}

func updateExposition() (done bool) {
	return inpututil.IsKeyJustPressed(ebiten.KeyEnter) || inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}
