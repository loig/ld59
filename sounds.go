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
	"bytes"
	_ "embed"
	"io"
	"log"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

//go:embed assets/ok.wav
var okBytes []byte

//go:embed assets/lost.wav
var lostBytes []byte

//go:embed assets/won.wav
var wonBytes []byte

//go:embed assets/jump.wav
var jumpBytes []byte

//go:embed assets/miss.wav
var missBytes []byte

//go:embed assets/note1.wav
var note1Bytes []byte

//go:embed assets/note2.wav
var note2Bytes []byte

//go:embed assets/note3.wav
var note3Bytes []byte

//go:embed assets/note4.wav
var note4Bytes []byte

//go:embed assets/note5.wav
var note5Bytes []byte

//go:embed assets/note6.wav
var note6Bytes []byte

//go:embed assets/music.wav
var musicBytes []byte

// A sound engine is responsible for playing sounds
// so that the same sound is not played twice at the same
// frame.
type soundEngine struct {
	audioContext *audio.Context
	nextSounds   [numSounds]bool
	sounds       [numSounds][]byte
	mute         bool
	musicPlayer  *audio.Player
}

// The list of existing sounds.
const (
	soundOk int = iota
	soundLost
	soundWon
	soundJump
	soundMiss
	soundNote1
	soundNote2
	soundNote3
	soundNote4
	soundNote5
	soundNote6
	numSounds
)

// Toggle sound
func (s *soundEngine) toggleSound() {
	s.mute = !s.mute
}

// Initialisation of the sound engine (sound decoding).
func newSoundEngine() (engine soundEngine) {

	var err error
	var sound *wav.Stream
	engine.audioContext = audio.NewContext(44100)

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(okBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundOk], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(lostBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundLost], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(wonBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundWon], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(jumpBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundJump], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(missBytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundMiss], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(note1Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundNote1], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(note2Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundNote2], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(note3Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundNote3], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(note4Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundNote4], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(note5Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundNote5], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	sound, err = wav.DecodeWithSampleRate(engine.audioContext.SampleRate(), bytes.NewReader(note6Bytes))
	if err != nil {
		log.Panic("Audio problem: ", err)
	}
	engine.sounds[soundNote6], err = io.ReadAll(sound)
	if err != nil {
		log.Panic("Audio problem: ", err)
	}

	// music

	wavMusic, err := wav.DecodeF32(bytes.NewReader(musicBytes))
	if err != nil {
		log.Fatal(err)
	}

	introSec := 0.8
	loopSec := 3.2 //3.025
	introDuration := int64(introSec * 8 * 44100)
	loopDuration := int64((loopSec - introSec) * 8 * 44100)
	s := audio.NewInfiniteLoopWithIntroF32(wavMusic, introDuration, loopDuration)

	engine.musicPlayer, err = engine.audioContext.NewPlayerF32(s)
	if err != nil {
		log.Fatal(err)
	}

	engine.musicPlayer.SetVolume(0.4)

	return
}

// Play all sounds that have been registered for
// playing in e.nextSounds.
func (e *soundEngine) playNow() {
	e.musicPlayer.Play()
	for soundID, play := range e.nextSounds {
		if play && soundID != soundJump {
			e.playSound(soundID)
			e.nextSounds[soundID] = false
		}
	}
}

// Play one sound by generating a new player.
func (e soundEngine) playSound(ID int) {
	if !e.mute {
		soundPlayer := e.audioContext.NewPlayerFromBytes(e.sounds[ID])
		volume := 0.7
		soundPlayer.SetVolume(volume)
		switch ID {
		case soundNote1, soundNote2, soundNote3, soundNote4:
			soundPlayer.SetVolume(volume + 0.6)
		case soundNote5, soundNote6:
			soundPlayer.SetVolume(volume + 0.3)
		default:
			soundPlayer.SetVolume(volume - 0.5)
		}
		soundPlayer.Play()
	}
}

// Ask to play a sound
func (e *soundEngine) register(ID int) {
	e.nextSounds[ID] = true
}
