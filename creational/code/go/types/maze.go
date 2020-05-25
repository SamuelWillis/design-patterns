package types

// Maze representation.
type Maze struct {
	rooms []*Room
	mapSite
}

// AddRoom to add a room to the maze.
func (maze *Maze) AddRoom(room *Room) {
	maze.rooms = append(maze.rooms, room)
}

// GetRooms to access the rooms in a maze.
func (maze Maze) GetRooms() []*Room {
	return maze.rooms
}
