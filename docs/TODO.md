# TODO

- Refactor `State_Playing.cpp` to load song settings (like LeadIn and LeadOut timing) from user settings or a configuration file rather than hardcoding them. (Done)
- Standardize the version display logic throughout the application so it always points to `src/version.h` without duplicating hard-coded version strings (e.g., in UI or installer). (Done)
- Address any remaining "TODO" comments in code where practical, especially concerning memory management (e.g., in `TextWriter.cpp` deletion on shutdown). (Partially Done)
- Replace generic "PianoGameVersionString" string concatenations with a centralized MACRO to avoid static initializers in C++ headers. (Done)
- Implement comprehensive fallback mechanism for Unicode string conversion (addressing `TODO: This isn't Unicode!` across string_util.h and Midi.cpp) without truncating characters to ASCII or crashing the OS X Carbon renderer.
- Implement OS-specific stack trace extraction for `PianoGameError.h` to make debugging easier for users. (Done)
- Add descriptive labels and tooltips to the OpenGL UI. Currently the UI is completely raw. Ensure every interactive component handles mouse-over descriptions natively. (Done)
