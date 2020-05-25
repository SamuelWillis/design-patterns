package types

import (
	"strings"
)

type IRoom interface {
	GetSide(direction string) IMapSite
	SetSide(direction string, mapSite IMapSite)
	IMapSite
}

type roomSides struct {
	north IMapSite
	east IMapSite
	south IMapSite
	west IMapSite
}

type Room struct {
	RoomNo int
	sides roomSides
	mapSite
}

// GetSide to access what is on a side of the room.
func (room Room) GetSide(direction string) IMapSite {
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
func (room *Room) SetSide(direction string, mapSite IMapSite) {
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
