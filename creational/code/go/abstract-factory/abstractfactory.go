package abstractfactory

import (
	"github.com/SamuelWillis/design-patterns/creational/types"
)

type IMazeFactory interface {
	MakeMaze() types.IMaze
	MakeWall() types.IWall
	MakeRoom(roomNo int) types.IRoom
	MakeDoor(room1 *types.IRoom, room2 *types.IRoom) types.IDoor
}

type DefaultMazeFactory struct {}

// MakeMaze for easy maze making.
func (factory DefaultMazeFactory) MakeMaze() types.IMaze {
	return &types.Maze{}
}

// MakeWall for easy wall making
func (factory DefaultMazeFactory) MakeWall() types.IWall {
	return types.Wall{}
}

// MakeRoom for easy room making.
func (factory DefaultMazeFactory) MakeRoom(roomNo int) types.IRoom {
	return &types.Room{
		RoomNo: roomNo,
	}
}

// MakeDoor for easy door making.
func (factory DefaultMazeFactory) MakeDoor(room1 types.IRoom, room2 types.IRoom) types.IDoor {
	return &types.Door{
		Room1: room1,
		Room2: room2,
	}
}

func (factory DefaultMazeFactory) CreateMaze() types.IMaze {
	maze := factory.MakeMaze()

	room1 := factory.MakeRoom(1)
	room2 := factory.MakeRoom(2)

	door := factory.MakeDoor(room1, room2)

	maze.AddRoom(room1)
	maze.AddRoom(room2)

	room1.SetSide("North", factory.MakeWall())
	room1.SetSide("East", door)
	room1.SetSide("South", factory.MakeWall())
	room1.SetSide("West", factory.MakeWall())

	room2.SetSide("North", factory.MakeWall())
	room2.SetSide("East", factory.MakeWall())
	room2.SetSide("South", factory.MakeWall())
	room2.SetSide("West", door)

	return maze
}
