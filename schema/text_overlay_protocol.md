# Rich Text Overlay & Head-Up Display (HUD) Protocol (v1.0.0)

## Overview
Connects legacy C++ OpenGL typography (`TextWriter.h`, `TextWriter.cpp`) with the modern Agentic Coding Harness. Standardizes text overlay placement, bounding box estimation, line alignment (left, center, right), multi-color text spans, and shadow rendering for in-game HUDs and terminal dashboards.

## HUD Span Schema
```json
{
  "x": 640,
  "y": 480,
  "alignment": "center",
  "fontSize": 14,
  "fontFamily": "Trebuchet MS",
  "dropShadow": true,
  "spans": [
    {
      "text": "Score: ",
      "color": "#94a3b8"
    },
    {
      "text": "98,400",
      "color": "#10b981"
    }
  ],
  "estimatedWidth": 112,
  "estimatedHeight": 18
}
```

## Protocol Operations
* `CalculateLayout(options)`: Computes bounding boxes, character widths, and anchor offsets based on alignment.
* `FormatHudMetric(label, value, color)`: Formats standard two-span HUD metric badges.
