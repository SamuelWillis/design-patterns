package types

import (
	"fmt"
)

// IMapSite interface.
type IMapSite interface {
	Enter()
}

type mapSite struct {}

func (mapSite mapSite) Enter() {
	fmt.Println("You enter the room")
}
