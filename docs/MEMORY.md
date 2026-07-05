# Memory

- The project relies on legacy OS APIs (Mac OS X Carbon and Win32).
- OpenGL is used for the raw drawing layer.
- `libmidi` submodule/dependency provides the MIDI file parsing logic.
- Avoid defining raw `static const std::wstring` in `.h` files since it creates a new copy per translation unit, resulting in binary bloat and static initialization issues.
- Do not blindly delete TODOs in legacy code unless you are actually fixing the underlying technical debt (like Unicode conversion limits).
- The user has provided an AI agent prompt that heavily references typical modern stack paradigms ("Dashboard UI with explicit interactive forms, clear labels, distinct descriptions, and detailed tooltips"). While this is a C++ game, the intent is clearly to ensure any new UI features in OpenGL are fully self-documenting and user-friendly. We document these goals in TODO.md.
