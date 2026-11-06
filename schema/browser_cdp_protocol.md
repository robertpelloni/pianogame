# Chrome DevTools Protocol (CDP) & Headless Browser Automation Protocol

## Overview
The Browser Automation & CDP Protocol enables the AI Harness to connect to external Chrome/Chromium/Edge instances over WebSocket, inspect the DOM tree, execute JavaScript in the page context, evaluate accessibility trees, take screenshots, and intercept console logs.

## Core Commands

### 1. `Page.navigate`
Navigates the current frame to a specified URL:
```json
{
  "id": 1,
  "method": "Page.navigate",
  "params": {
    "url": "http://localhost:8080"
  }
}
```

### 2. `Runtime.evaluate`
Evaluates an expression in the global context:
```json
{
  "id": 2,
  "method": "Runtime.evaluate",
  "params": {
    "expression": "document.title",
    "returnByValue": true
  }
}
```

### 3. `Page.captureScreenshot`
Captures a base64 encoded PNG screenshot:
```json
{
  "id": 3,
  "method": "Page.captureScreenshot",
  "params": {
    "format": "png"
  }
}
```
