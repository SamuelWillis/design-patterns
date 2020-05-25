package types

type IDoor interface {
	IMapSite
}

// Door bridges the void between two rooms.
type Door struct {
	Room1 IRoom
	Room2 IRoom
	isOpen bool
	mapSite
}

// IsOpen to tell if the void is accessible.
func (door Door) IsOpen() bool {
	return door.isOpen
}
