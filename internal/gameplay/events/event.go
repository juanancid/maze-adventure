package events

// Event is the interface that all events must implement.
type Event interface {
	isEvent()
}

// LevelCompletedEvent indicates that a level has been successfully completed.
type LevelCompletedEvent struct{}

// isEvent implements the Event interface explicitly.
func (LevelCompletedEvent) isEvent() {}

// MissionAccomplished indicates that the game is over.
type MissionAccomplished struct{}

// isEvent implements the Event interface explicitly.
func (MissionAccomplished) isEvent() {}
