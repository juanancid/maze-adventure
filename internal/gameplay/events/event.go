package events

// Event is the interface that all events must implement.
type Event interface {
	isEvent()
}

// CollectiblePicked indicates that a collectible has been picked up.
type CollectiblePicked struct {
	Value int
}

// isEvent implements the Event interface explicitly.
func (CollectiblePicked) isEvent() {}

// LevelCompletedEvent indicates that a level has been successfully completed.
type LevelCompletedEvent struct{}

// isEvent implements the Event interface explicitly.
func (LevelCompletedEvent) isEvent() {}

// GameComplete indicates that the game has been won.
type GameComplete struct{}

// isEvent implements the Event interface explicitly.
func (GameComplete) isEvent() {}
