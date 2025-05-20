package input

import (
	"sort"

	"github.com/hajimehoshi/ebiten/v2"
)

// Handler manages input state and provides methods to check for key events
type Handler struct {
	previousKeyStates map[ebiten.Key]bool
	combinationStates map[string]bool
}

// NewHandler creates a new input handler
func NewHandler() *Handler {
	return &Handler{
		previousKeyStates: make(map[ebiten.Key]bool),
		combinationStates: make(map[string]bool),
	}
}

// Update updates the input state for the current frame
func (h *Handler) Update() {
	// Update all tracked keys
	for key := range h.previousKeyStates {
		h.previousKeyStates[key] = ebiten.IsKeyPressed(key)
	}
}

// IsKeyJustPressed returns true if the key was just pressed this frame
func (h *Handler) IsKeyJustPressed(key ebiten.Key) bool {
	currentState := ebiten.IsKeyPressed(key)
	previousState := h.previousKeyStates[key]
	h.previousKeyStates[key] = currentState
	return currentState && !previousState
}

// IsKeyCombinationToggled returns true if the key combination was just pressed this frame,
// implementing a toggle behavior. The order of keys in the combination does not matter
// (e.g., Meta+D is the same as D+Meta).
func (h *Handler) IsKeyCombinationToggled(keys ...ebiten.Key) bool {
	// Create a sorted slice of key names to ensure consistent combination keys
	keyNames := make([]string, len(keys))
	for i, key := range keys {
		keyNames[i] = key.String()
	}
	sort.Strings(keyNames)

	// Create a unique key for this combination
	combinationKey := ""
	for _, name := range keyNames {
		combinationKey += name
	}

	// Check if all keys are currently pressed
	allPressed := true
	for _, key := range keys {
		if !ebiten.IsKeyPressed(key) {
			allPressed = false
			break
		}
	}

	// If not all keys are pressed, reset the combination state
	if !allPressed {
		h.combinationStates[combinationKey] = false
		return false
	}

	// If combination was already active, it's not a new toggle
	if h.combinationStates[combinationKey] {
		return false
	}

	// This is a new combination press, mark it as active
	h.combinationStates[combinationKey] = true
	return true
}
