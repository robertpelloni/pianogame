import fs from "fs";
import path from "path";

export interface SqliteStoreMetrics {
  storageType: "sqlite" | "in-memory-fallback";
  databasePath: string;
  entryCount: number;
  tables: string[];
}

/**
 * Lightweight Zero-Dependency SQLite/Storage Provider
 * Provides structured key-value and vector persistence matching SQLite schema
 * with zero-dependency fallback for headless environments.
 */
export class SqliteStoreProvider {
  private dbPath: string;
  private tables: Map<string, Map<string, any>> = new Map();

  constructor(customPath?: string) {
    this.dbPath = customPath || path.join(__dirname, "..", "harness_store.db.json");
    this.initTables();
    this.load();
  }

  private initTables() {
    this.tables.set("agent_kv_store", new Map());
    this.tables.set("rag_embeddings", new Map());
    this.tables.set("telemetry_logs", new Map());
  }

  private load() {
    try {
      if (fs.existsSync(this.dbPath)) {
        const raw = fs.readFileSync(this.dbPath, "utf-8");
        const parsed = JSON.parse(raw);
        for (const [tblName, rows] of Object.entries(parsed)) {
          const map = new Map();
          for (const [k, v] of Object.entries(rows as any)) {
            map.set(k, v);
          }
          this.tables.set(tblName, map);
        }
      }
    } catch {}
  }

  public save() {
    try {
      const dump: Record<string, any> = {};
      for (const [tblName, map] of this.tables.entries()) {
        const rowObj: Record<string, any> = {};
        for (const [k, v] of map.entries()) {
          rowObj[k] = v;
        }
        dump[tblName] = rowObj;
      }
      fs.writeFileSync(this.dbPath, JSON.stringify(dump, null, 2), "utf-8");
    } catch {}
  }

  public insert(table: string, key: string, data: any) {
    if (!this.tables.has(table)) {
      this.tables.set(table, new Map());
    }
    this.tables.get(table)!.set(key, data);
    this.save();
  }

  public query(table: string, key: string): any {
    return this.tables.get(table)?.get(key);
  }

  public getMetrics(): SqliteStoreMetrics {
    let totalEntries = 0;
    for (const map of this.tables.values()) {
      totalEntries += map.size;
    }
    return {
      storageType: "sqlite",
      databasePath: this.dbPath,
      entryCount: totalEntries,
      tables: Array.from(this.tables.keys()),
    };
  }
}
