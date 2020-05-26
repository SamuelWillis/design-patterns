package types

// IWall interface.
type IWall interface {
	IMapSite
}

// Wall representation.
type Wall struct {
	mapSite
}
