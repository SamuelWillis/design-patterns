package types

import (
	"strings"
)

type roomSides struct {
	north iMapSite
	east iMapSite
	south iMapSite
	west iMapSite
}

// Room representation.
type Room struct {
	RoomNo int
	sides roomSides
	mapSite
}

// GetSide to access what is on a side of the room.
func (room Room) GetSide(direction string) iMapSite {
	lower := strings.ToLower(direction)
	switch lower {
	case "north":
		return room.sides.north
	case "east":
		return room.sides.east
	case "south":
		return room.sides.south
	case "west":
		return room.sides.west
	default:
		return room.sides.north
	}
}

// SetSide to set what is on a side of the room.
func (room *Room) SetSide(direction string, mapSite iMapSite) {
	lower := strings.ToLower(direction)
	switch lower {
	case "north":
		room.sides.north = mapSite
	case "east":
		room.sides.east = mapSite
	case "south":
		room.sides.south = mapSite
	case "west":
		room.sides.west = mapSite
	default:
		room.sides.north = mapSite
	}
}
