# Ideas

- Port the OpenGL/OS layer to SDL2. This would completely remove the need for `src/os.h` platform-specific forks, Carbon APIs, and Win32 HWND creations.
- Integrate FreeType for text rendering instead of relying on WGL/AGL native OS font APIs, fixing the Unicode rendering issues simultaneously across all platforms.
- Create an ImGui layer for the "Dashboard UI" to satisfy user requirements for "interactive forms" and "tooltips" natively within the OpenGL context.
