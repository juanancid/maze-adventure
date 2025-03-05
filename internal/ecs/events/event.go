package events

// Event is a marker interface for events.
type Event interface{}

// LevelCompletedEvent is emitted when the player reaches the maze exit.
type LevelCompletedEvent struct{}
