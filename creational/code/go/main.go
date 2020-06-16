package main

import (
	"fmt"
	"github.com/SamuelWillis/design-patterns/creational/abstract-factory"
)

// This feels like it's essentially a factory method.
// Code is a little clunky as it stands but it is a first pass.
func main() {
	fmt.Println("Starting Creational Patterns Go Code")

	maze := abstractfactory.DefaultMazeFactory{}.CreateMaze()

	fmt.Printf("Your Maze looks like this: %#v\n", maze);

	enchantedMaze := abstractfactory.EnchantedMazeFactory{}.CreateMaze()

	fmt.Printf("Your EnchantedMaze looks like this: %#v\n", enchantedMaze);
}

type mazeGame struct {
	factory abstractfactory.IMazeFactory
}

