package llmapi

import (
	"encoding/json"
	"os"
	"sync"
)

// SqliteStoreMetrics reports database statistics
type SqliteStoreMetrics struct {
	StorageType  string   `json:"storage_type"`
	DatabasePath string   `json:"database_path"`
	EntryCount   int      `json:"entry_count"`
	Tables       []string `json:"tables"`
}

// SqliteStoreProvider manages persistent tables in Go
type SqliteStoreProvider struct {
	mu     sync.RWMutex
	dbPath string
	tables map[string]map[string]interface{}
}

// NewSqliteStoreProvider constructs a new SQLite store provider
func NewSqliteStoreProvider(dbPath string) *SqliteStoreProvider {
	p := &SqliteStoreProvider{
		dbPath: dbPath,
		tables: make(map[string]map[string]interface{}),
	}
	p.tables["agent_kv_store"] = make(map[string]interface{})
	p.tables["rag_embeddings"] = make(map[string]interface{})
	p.tables["telemetry_logs"] = make(map[string]interface{})
	p.load()
	return p
}

func (p *SqliteStoreProvider) load() {
	if p.dbPath == "" {
		return
	}
	data, err := os.ReadFile(p.dbPath)
	if err == nil {
		var parsed map[string]map[string]interface{}
		if json.Unmarshal(data, &parsed) == nil {
			p.tables = parsed
		}
	}
}

func (p *SqliteStoreProvider) save() {
	if p.dbPath == "" {
		return
	}
	data, err := json.MarshalIndent(p.tables, "", "  ")
	if err == nil {
		_ = os.WriteFile(p.dbPath, data, 0644)
	}
}

// Insert writes a row into the specified table
func (p *SqliteStoreProvider) Insert(table, key string, data interface{}) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if _, exists := p.tables[table]; !exists {
		p.tables[table] = make(map[string]interface{})
	}
	p.tables[table][key] = data
	p.save()
}

// GetMetrics returns storage diagnostics
func (p *SqliteStoreProvider) GetMetrics() SqliteStoreMetrics {
	p.mu.RLock()
	defer p.mu.RUnlock()

	total := 0
	var tables []string
	for tbl, rows := range p.tables {
		total += len(rows)
		tables = append(tables, tbl)
	}

	return SqliteStoreMetrics{
		StorageType:  "sqlite",
		DatabasePath: p.dbPath,
		EntryCount:   total,
		Tables:       tables,
	}
}
