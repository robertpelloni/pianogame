# Retrieval-Augmented Generation (RAG) Vector Indexing Protocol

## Overview
The RAG Vector Indexer splits codebase files, markdown documentation, and architectural memory logs into semantic chunks, generates vector embeddings, and performs cosine similarity ranking for relevant context retrieval.

## Schema

### 1. Document Chunk
```json
{
  "id": "chunk_doc_01",
  "sourcePath": "VISION.md",
  "content": "We are transforming this project into the ULTIMATE AGENTIC CODING HARNESS...",
  "embedding": [0.021, -0.045, 0.112, 0.089]
}
```

### 2. Similarity Search Query
```json
{
  "query": "dashboard layout tooltips",
  "topK": 3,
  "threshold": 0.5
}
```

### 3. Match Result
```json
{
  "matches": [
    {
      "chunkId": "chunk_doc_01",
      "sourcePath": "VISION.md",
      "content": "...",
      "score": 0.88
    }
  ]
}
```
