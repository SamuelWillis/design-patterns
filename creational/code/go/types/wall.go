package types

type IWall interface {
	IMapSite
}

// Wall representation.
type Wall struct {
	mapSite
}
