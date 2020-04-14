import java.util.*;

class Maze {
  private ArrayList<Room> rooms;

  public Maze() {
    this.rooms = new ArrayList<Room>();
  }

  public Maze addRoom(Room room) {
    this.rooms.add(room);

    return this;
  }
}
