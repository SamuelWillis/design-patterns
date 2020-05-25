package types

import (
	"fmt"
)

type mapSite struct {}

func (mapSite mapSite) Enter() {
	fmt.Println("You enter the room")
}
