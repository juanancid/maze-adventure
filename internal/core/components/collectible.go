package components

type CollectibleKind int

const (
	CollectibleScore CollectibleKind = iota
)

type Collectible struct {
	Kind  CollectibleKind
	Value int
}
