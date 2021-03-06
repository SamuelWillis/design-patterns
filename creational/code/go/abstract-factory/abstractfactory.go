package abstractfactory

import (
	"math/rand"
	"time"
	"github.com/SamuelWillis/design-patterns/creational/types"
)

// IMazeFactory interface.
type IMazeFactory interface {
	MakeMaze() types.IMaze
	MakeWall() types.IWall
	MakeRoom(roomNo int) types.IRoom
	MakeDoor(room1 types.IRoom, room2 types.IRoom) types.IDoor
}

// DefaultMazeFactory to create default mazes.
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
	return types.Door{
		Room1: room1,
		Room2: room2,
	}
}

// CreateMaze function.
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

// EnchantedMazeFactory to create mazes with enchated rooms.
type EnchantedMazeFactory struct {}

// MakeMaze for easy maze making.
func (factory EnchantedMazeFactory) MakeMaze() types.IMaze {
	return &types.Maze{}
}

// MakeWall for easy wall making
func (factory EnchantedMazeFactory) MakeWall() types.IWall {
	return types.Wall{}
}

// MakeRoom for easy room making.
func (factory EnchantedMazeFactory) MakeRoom(roomNo int) types.IRoom {
	rand.Seed(time.Now().UTC().UnixNano())

	if (rand.Intn(2) == 1) {
		return &types.EnchantedRoom{
			RoomNo: roomNo,
		}
	}
	return &types.Room{
		RoomNo: roomNo,
	}
}

// MakeDoor for easy door making.
func (factory EnchantedMazeFactory) MakeDoor(room1 types.IRoom, room2 types.IRoom) types.IDoor {
	return &types.Door{
		Room1: room1,
		Room2: room2,
	}
}

// CreateMaze implementation for EnchantedMazeFactory.
func (factory EnchantedMazeFactory) CreateMaze() types.IMaze {
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
