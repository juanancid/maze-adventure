package events

// Event is the interface that all events must implement.
type Event interface {
	isEvent()
}

// LevelCompletedEvent indicates that a level has been successfully completed.
type LevelCompletedEvent struct{}

// isEvent implements the Event interface explicitly.
func (LevelCompletedEvent) isEvent() {}

// GameOverEvent indicates that the game is over.
type GameOverEvent struct{}

// isEvent implements the Event interface explicitly.
func (GameOverEvent) isEvent() {}
