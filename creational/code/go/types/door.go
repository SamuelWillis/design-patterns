package types

// IDoor Interface.
type IDoor interface {
	Open() bool
	IMapSite
}

// Door bridges the void between two rooms.
type Door struct {
	Room1 IRoom
	Room2 IRoom
	open bool
	mapSite
}

// Open to tell if the void is accessible.
func (door Door) Open() bool {
	return door.open
}
