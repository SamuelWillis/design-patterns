package main

import (
	"fmt"
	"github.com/SamuelWillis/design-patterns/creational/abstract-factory"
)

// This feels like it's essentially a factory method.
// Code is a little clunky as it stands but it is a first pass.
func main() {
	fmt.Println("Starting Creational Patterns Go Code")

	factory := abstractfactory.MazeFactory{}

	maze := factory.MakeMaze()

	room1 := factory.MakeRoom(1)
	room2 := factory.MakeRoom(2)

	door := factory.MakeDoor(&room1, &room2)

	maze.AddRoom(&room1)
	maze.AddRoom(&room2)

	room1.SetSide("North", factory.MakeWall())
	room1.SetSide("East", door)
	room1.SetSide("South", factory.MakeWall())
	room1.SetSide("West", factory.MakeWall())

	room2.SetSide("North", factory.MakeWall())
	room2.SetSide("East", factory.MakeWall())
	room2.SetSide("South", factory.MakeWall())
	room2.SetSide("West", door)


	fmt.Printf("Your Maze looks like this: %#v", maze);
}
