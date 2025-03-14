# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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

[Unreleased]: https://github.com/juanancid/maze-adventure/compare/v0.2.0...HEAD
[0.2.0]: https://github.com/juanancid/maze-adventure/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/juanancid/maze-adventure/releases/tag/v0.1.0
