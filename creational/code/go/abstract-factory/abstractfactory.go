package abstractfactory

import (
	"github.com/SamuelWillis/design-patterns/creational/types"
)

type iMazeFactory interface {
	MakeMaze() types.Maze
	MakeWall() types.Wall
	MakeRoom(roomNo int) types.Room
	MakeDoor(room1 types.Room, room2 types.Room) types.Door
}

// MazeFactory for simple maze generation
type MazeFactory struct {}

// MakeMaze for easy maze making.
func (factory MazeFactory) MakeMaze() types.Maze {
	return types.Maze{}
}

// MakeWall for easy wall making
func (factory MazeFactory) MakeWall() types.Wall {
	return types.Wall{}
}

// Make room for easy room making.
func (factory MazeFactory) MakeRoom(roomNo int) types.Room {
	return types.Room{
		RoomNo: roomNo,
	}
}

// MakeDoor for easy door making.
func (factory MazeFactory) MakeDoor(room1 *types.Room, room2 *types.Room) types.Door {
	return types.Door{
		Room1: room1,
		Room2: room2,
	}
}
