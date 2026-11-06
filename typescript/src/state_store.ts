import fs from "fs";
import path from "path";

/**
 * Lightweight Persistent Key-Value State Store and Memory Manager
 * for the Ultimate Agentic Coding Harness.
 */
export class StateStore {
  private storePath: string;
  private memory: Map<string, any> = new Map();

  constructor(customPath?: string) {
    this.storePath = customPath || path.join(__dirname, "..", "agent_memory.json");
    this.load();
  }

  private load() {
    try {
      if (fs.existsSync(this.storePath)) {
        const raw = fs.readFileSync(this.storePath, "utf-8");
        const parsed = JSON.parse(raw);
        for (const [k, v] of Object.entries(parsed)) {
          this.memory.set(k, v);
        }
      }
    } catch {
      // In-memory fallback
    }
  }

  public save() {
    try {
      const obj: Record<string, any> = {};
      for (const [k, v] of this.memory.entries()) {
        obj[k] = v;
      }
      fs.writeFileSync(this.storePath, JSON.stringify(obj, null, 2), "utf-8");
    } catch {
      // Ignore write failures in read-only setups
    }
  }

  public get(key: string): any {
    return this.memory.get(key);
  }

  public set(key: string, value: any): void {
    this.memory.set(key, value);
    this.save();
  }

  public getAll(): Record<string, any> {
    const obj: Record<string, any> = {};
    for (const [k, v] of this.memory.entries()) {
      obj[k] = v;
    }
    return obj;
  }
}
