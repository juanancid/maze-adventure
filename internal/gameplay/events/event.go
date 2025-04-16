package events

// Event is the interface that all events must implement.
type Event interface {
	isEvent()
}

// LevelCompletedEvent indicates that a level has been successfully completed.
type LevelCompletedEvent struct{}

// isEvent implements the Event interface explicitly.
func (LevelCompletedEvent) isEvent() {}

// Victory indicates that the game has been won.
type Victory struct{}

// isEvent implements the Event interface explicitly.
func (Victory) isEvent() {}
