# Desktop Webview Container & Native IPC Protocol

## Overview
Defines specifications for embedding the Unified Single-Page Web Dashboard within a native OS desktop webview container (e.g. Tauri, Wry, Webview), managing window chrome, tray docking, and IPC message routing.

## Configuration Schema

```json
{
  "app": {
    "title": "Ultimate Agentic Coding Harness",
    "version": "v5.87.0",
    "width": 1280,
    "height": 850,
    "resizable": true,
    "frameless": false,
    "trayDocked": true
  },
  "webview": {
    "url": "http://localhost:8080",
    "devTools": true,
    "ipc": {
      "channel": "harness_ipc",
      "allowedOrigins": ["http://localhost:8080"]
    }
  }
}
```
