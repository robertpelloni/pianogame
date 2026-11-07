# SQLite Vector & Persistent Storage Protocol

## Overview
Defines SQLite database schemas for durable key-value caching, RAG vector embeddings, session execution logs, and cognitive codebase AST maps.

## Table Definitions

### 1. `agent_kv_store`
```sql
CREATE TABLE IF NOT EXISTS agent_kv_store (
    key TEXT PRIMARY KEY,
    val TEXT NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 2. `rag_embeddings`
```sql
CREATE TABLE IF NOT EXISTS rag_embeddings (
    chunk_id TEXT PRIMARY KEY,
    source_path TEXT NOT NULL,
    content TEXT NOT NULL,
    embedding_json TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 3. `telemetry_logs`
```sql
CREATE TABLE IF NOT EXISTS telemetry_logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    source TEXT NOT NULL,
    message TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
