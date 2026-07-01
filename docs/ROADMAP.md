# Roadmap

- Better configuration and settings management (replacing hard-coded configuration).
- Expanded cross-platform support (specifically modernizing away from Carbon to Cocoa/Metal for Mac, or a unified framework like SDL2/GLFW).
- Refactoring the OpenGL and OS interaction layers to make it easier to add new UI elements (like tooltips and dashboard descriptors as requested).
- Implement proper Unicode string conversions universally, so MIDI paths and metadata from Japan/Asia render and load correctly instead of being truncated by naïve `std::string` casting.
