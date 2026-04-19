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
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ranking struct {
	rankInfo   int
	numPlayed  int
	numTop10Ok int
	top10      [8]int
	top10Names [8]string
	problem    bool
	errnum     int
}

func sendRank(name string, level int) {

	if name == "YOU" {
		name = nameToString([3]uint8{uint8(rand.Intn(26)), uint8(rand.Intn(26)), uint8(rand.Intn(26))})
	}

	response, err := http.PostForm("https://games.balotchka.fr/ld59/", url.Values{
		"uname": {name},
		"level": {fmt.Sprint(level)},
	})
	if err != nil {
		log.Print(err)
		return
	}

	defer response.Body.Close()

	return
}

func getRank(level int, c chan ranking) {
	response, err := http.PostForm("https://games.balotchka.fr/ld59/", url.Values{
		"uname": {"dummy"},
		"level": {fmt.Sprint(level)},
	})
	if err != nil {
		log.Print(err)
		c <- ranking{problem: true, errnum: 1}
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Print(err)
		c <- ranking{problem: true, errnum: 2}
		return
	}

	parts := strings.Split(string(body), ":")
	if len(parts) < 2 {
		c <- ranking{problem: true, errnum: 3}
	}

	pos, err := strconv.Atoi(parts[0])
	if err != nil {
		c <- ranking{problem: true, errnum: 4}
		return
	}

	var top10 [8]int
	var top10Names [8]string
	var numTop10Ok int

	for i := 1; i < len(parts)-1; i++ {
		if len(parts[i]) < 5 {
			c <- ranking{problem: true, errnum: 5}
			return
		}
		name := parts[i][1:4]
		level, err := strconv.Atoi(parts[i][5 : len(parts[i])-1])
		if err != nil {
			c <- ranking{problem: true, errnum: 6}
			return
		}
		top10[i-1] = level
		top10Names[i-1] = name
		numTop10Ok = i
	}

	numPos, err := strconv.Atoi(parts[len(parts)-1])
	if err != nil {
		c <- ranking{problem: true, errnum: 7}
		return
	}

	c <- ranking{
		rankInfo:   pos,
		numPlayed:  numPos,
		numTop10Ok: numTop10Ok,
		top10:      top10,
		top10Names: top10Names,
	}
	return
}
