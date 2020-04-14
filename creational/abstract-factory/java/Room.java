import java.util.*;

public class Room extends MapSite {
  private int roomNo;
  private HashMap<Direction, MapSite> sides;

  public Room(int roomNo) {
    this.roomNo = roomNo;
  }

  public Room setSide(Direction direction, MapSite mapSite) {
    this.sides.put(direction, mapSite);
    return this;
  }

  public MapSite getSide(Direction direction) {
    return this.sides.get(direction);
  }
}
