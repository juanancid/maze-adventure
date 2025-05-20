package debug

import (
	"fmt"
	"runtime"
	"time"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/juanancid/maze-adventure/internal/engine/input"
)

var (
	fps         float64
	fpsCount    int
	fpsTime     = time.Now()      // Initialize to current time
	fpsUpdate   = time.Second / 2 // Update FPS every half second
	debugString string
)

// TrackFPS updates the FPS counter
func TrackFPS() {
	fpsCount++
	now := time.Now()

	if now.Sub(fpsTime) >= fpsUpdate {
		fps = float64(fpsCount) / now.Sub(fpsTime).Seconds()
		fpsCount = 0
		fpsTime = now

		// Update debug string with current stats
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		debugString = fmt.Sprintf(
			"FPS: %.1f\n"+
				"Memory: %.1f MB\n"+
				"GC: %d\n"+
				"Threads: %d\n"+
				"Goroutines: %d\n"+
				"Stack: %.1f MB\n"+
				"Heap: %.1f MB\n"+
				"Next GC: %.1f MB",
			fps,
			float64(m.Alloc)/1024/1024,
			m.NumGC,
			runtime.NumCPU(),
			runtime.NumGoroutine(),
			float64(m.StackInuse)/1024/1024,
			float64(m.HeapInuse)/1024/1024,
			float64(m.NextGC)/1024/1024)
	}
}

// GetDebugInfo returns the current debug information
func GetDebugInfo() string {
	return debugString
}

// System handles engine-level debug functionality
type System struct {
	showDebug    bool
	inputHandler *input.Handler
}

// NewSystem creates a new debug system
func NewSystem(inputHandler *input.Handler) *System {
	return &System{
		inputHandler: inputHandler,
	}
}

// Update handles debug-related updates
func (s *System) Update() {
	TrackFPS()
	s.inputHandler.Update()

	if s.inputHandler.IsKeyCombinationToggled(ebiten.KeyMeta, ebiten.KeyD) {
		s.showDebug = !s.showDebug
	}
}

// IsDebugEnabled returns whether debug mode is enabled
func (s *System) IsDebugEnabled() bool {
	return s.showDebug
}
