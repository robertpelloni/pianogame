# MEMORY

## Architectural Observations
* **Legacy Base:** The original codebase is a C++ application named "PianoGame" (later renamed Synthesia) for macOS and Windows, utilizing OpenGL and MIDI APIs.
* **Paradigm Shift:** The current directive requires a massive pivot from a C++ piano teaching game to a multi-language agentic coding harness.
* **Target Architectures:** Rust, Go, C#, Java, TypeScript. The C++ legacy will likely need to be archived or heavily compartmentalized as the project shifts focus.

## Design Preferences
* Strict documentation governance (`VISION.md`, `MEMORY.md`, etc.).
* Continuous autonomous execution: commit and push after every major feature/step.
* 100% feature parity across the 5 target languages.

## Codebase Traits
* Currently contains legacy C++ code (`src/` folder, `libmidi`, `CompatibilitySystem`).
* `Synthesia.xcodeproj`, `PianoGame.sln`, `PianoGame.vcproj` present for legacy builds.
