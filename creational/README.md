# Creational Designs

Creational designs abstract the instantiation process.
This allows a system to be independent of how its objects are created, composed, and represented.

A class creational pattern uses inheritance to vary the class that is instantiated, whereas an object creational pattern will delegate instantiation to another object.

These patterns become more important as systems evolve to depend more on object composition than class inheritance.
As this happens, emphasis shifts away from hard coding a fixed set of behaviours towards defining a smaller set of fundamental behaviours that can be composed into any number of more complex ones.
Thus creating objects with particular behaviours requires more than simply instantiating a class.

There are two recurring themes in these patterns:

1. They all encapsulate knowledge about which concrete classes the system uses.
1. They hide how instances of these classes are created and put together.

All the system comes to know of these classes is their interfaces as defined by abstract classes.

As a consequence, there is a tremendous amount of flexibility in *what* gets created, *who* creates it, *how* it gets created, and *when*.
They also allow the system to be configured with "product" objects that vary wildly in structure and functionality.
Configuration can be static (specified at compile-time) or dynamic (or run-time).

## Patterns

### [Abstract Factory](/creational/abstract-factory)

## Example problem
Each pattern will be illustrated through the common example of building a maze for a computer game.
This maze and game will vary slightly from pattern to pattern.
The main focus will be the creation of the maze, however.

The classes in the maze are as follows:

### MazeGame
Creates the maze.

### Maze
A **maze** is defined as a set of **rooms**. A room knows it's neighbours; possible neighbors are another room, a wall, or a door to another room.

A maze can also find a room given a room number.

### Mapsite
This is the common abstraction for each maze component. It defines one operation: `Enter`.
It's meaning depends on what you're entering. If you enter a room then your location changes.
If you enter a door then one of two things happen: if the door is open you go into the next room, otherwise you hurt your nose.

### Room
Each room has 4 sides, `north`, `south`, `east`, and `west` as well as a number to identify it.

The rooms define the key relationships between components in the maze. It maintains references to other `MapSite` objects.

### Wall
A wall...

### Door
A door that, when open, allows you to pass into the next room.
