# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.5.0] - 2025-06-04

### Added
- Introduced a **hazard system** with deadly cells that damage the player on contact.
- Added a **player health system** with visible units displayed in the HUD.
- Added a new **Game Over screen** shown when the player loses all health.
- Created a **damage feedback system** with repositioning to avoid repeated hits.
- Embedded all **game images and sounds** for easier distribution and predictable boot behavior.
- Added **sound effects** for collecting items, completing a level, and taking damage.
- Added **debug system** with toggles for development diagnostics.

### Changed
- Refactored maze generation to support typed cells (`Regular`, `Deadly`, `Freezing`).
- Improved maze and level configuration validation with better error reporting.
- Split HUD renderer into focused modules (e.g. Health, Memory, Level).
- Updated the damage logic to reposition the player to the center of the current cell after hit.
- Enhanced visuals for the HUD and hazards, improving readability and style.
- Renamed `EndState` to `VictoryState` for clarity.

### Maintenance
- Updated to the latest **Ebitengine** and **Golang** versions.
- Added build and license badges to the README.
- Cleaned up sound playback logic and image embedding system for consistency.

## [0.4.0] - 2025-04-24

### Added
- New **Boot Screen** introducing narrative tone and Picatoste, the exploration drone.
- New **Victory Screen** shown after completing all available levels.
- Added a **collectible system**: memory fragments now spawn in each level and are tracked in the HUD.
- Introduced a **HUD bar** displaying current level and collected memory fragments.
- Added a **GameSession and GameState** system to manage score, session-wide data, and player progress.

### Changed
- Updated game resolution to **480x270**, providing more screen space and a modern-retro aspect ratio.
- Replaced legacy YAML level definitions with **code-driven level configuration**, allowing more flexibility.
- Renamed game states to **BootScreen**, **PlayingState**, and **VictoryScreen** for better clarity.
- Improved maze visuals and color palette for **better contrast and readability**.
- Refined game loop logic: completing a level now properly advances or ends the session.

## [0.3.0] - 2025-04-04

### Added
- Game state management system explicitly (Menu, Playing, Game Over).
- Basic main menu explicitly prompting the player to start.
- Game Over screen explicitly after completing the available levels.

### Changed
- Refined ECS package structure explicitly, separating clearly entities, components, systems, queries, and events.
- Decoupled events explicitly from the World into their dedicated event bus.
- Explicitly refactored updaters (Movement, InputControl, MazeCollisionSystem) to clearly separate responsibilities and improve maintainability.
- Improved collision detection logic explicitly, enhancing clarity and accuracy.
- Explicitly clarified query methods (`Query` and `QueryComponents`) for easier component access.

## [0.2.0] - 2025-03-14

### Added
- Basic level generation system with support for multiple levels.
- An event-driven mechanism to signal level completion.
- Visual rendering for the exit door.

### Changed
- Extracted game logic into its own package for better modularity.
- Moved systems (now called updaters) and renderers from the World to the Game.
- Renamed "systems" to "updaters" for clarity.
- Refactored Game code for improved readability.
- Renamed maze disposition field to "layout" to clarify its role.
- Updated sprites to refine visuals.
- Refactored levels to a data-driven design for easier configuration.
- Consolidated maze handling to support a single maze per level.

## [0.1.0] - 2025-02-26

### Added
- Initial repository setup with the basic game architecture.
- Integration of [Ebitengine](https://ebiten.org/) as the graphics framework.
- Implementation of an ECS (Entity Component System) architectural pattern.
- On-the-fly maze generation.
- Smooth player movement.
- Basic collision handling between the player and maze walls.

[Unreleased]: https://github.com/juanancid/maze-adventure/compare/v0.5.0...HEAD
[0.5.0]: https://github.com/juanancid/maze-adventure/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/juanancid/maze-adventure/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/juanancid/maze-adventure/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/juanancid/maze-adventure/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/juanancid/maze-adventure/releases/tag/v0.1.0
