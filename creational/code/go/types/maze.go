package types

type IMaze interface {
	AddRoom(room IRoom)
	GetRooms() []IRoom
}

// Maze representation.
type Maze struct {
	rooms []IRoom
	mapSite
}

// AddRoom to add a room to the maze.
func (maze *Maze) AddRoom(room IRoom) {
	maze.rooms = append(maze.rooms, room)
}

// GetRooms to access the rooms in a maze.
func (maze Maze) GetRooms() []IRoom {
	return maze.rooms
}
