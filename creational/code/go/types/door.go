package types

// Door bridges the void between two rooms.
type Door struct {
	Room1 *Room
	Room2 *Room
	isOpen bool
	mapSite
}

// IsOpen to tell if the void is accessible.
func (door Door) IsOpen() bool {
	return door.isOpen
}
