package sprite

import (
	"pokered/pkg/joypad"
	"pokered/pkg/store"
	"pokered/pkg/util"
)

// Movment status
const (
	Uninitialized byte = iota
	OK
	Delay
	Movement
)

// UpdateSprites update sprite data
func UpdateSprites() {
	for offset, s := range store.SpriteData {
		if s == nil || s.ID == 0 {
			continue
		}
		if s.ID == 1 {
			UpdatePlayerSprite()
			continue
		}
		UpdateNPCSprite(uint(offset))
	}
}

// UpdateSpriteImage update sprite image offset
func UpdateSpriteImage(offset uint) {
	s := store.SpriteData[offset]
	if s == nil {
		return
	}
	length := len(s.VRAM.Images)
	switch length {
	case 1:
		s.VRAM.Index = 0
	case 4, 10:
	default:
		return
	}

	animCounter := s.AnimationFrame >> 2

	// ref:
	switch animCounter + uint(s.Direction) {
	case 0, 3:
		s.VRAM.Index = 1
		if length == 4 {
			s.VRAM.Index = 0
		}
	case 1:
	case 2:

	case 4:
	case 5:
	case 6:
	case 7:

	case 8, 11:
		s.VRAM.Index = 6
		if length == 4 {
			s.VRAM.Index = 2
		}
	case 9, 10:
		s.VRAM.Index = 7
		if length == 4 {
			s.VRAM.Index = 2
		}

	case 12, 15:
		s.VRAM.Index = 8
		if length == 4 {
			s.VRAM.Index = 3
		}
	case 13, 14:
		s.VRAM.Index = 9
		if length == 4 {
			s.VRAM.Index = 3
		}
	}
}

// DisableSprite hide sprite
func DisableSprite(offset uint) {
	s := store.SpriteData[offset]
	s.VRAM.Index = -1
}

// MoveSprite forcely move sprite by movement data
// set wNPCMovementDirections
func MoveSprite(offset uint, movement []byte) {
	copy(NPCMovementDirections, movement)
	util.SetBit(store.D730, 0)
	joypad.JoyIgnore = joypad.ByteToInput(0xff)
}
