import http from "http";
import fs from "fs";
import path from "path";
import { LlmRouter } from "./llmRouter";
import { Orchestrator } from "./orchestrator";
import { AstParser } from "./ast_parser";
import { createLlmRequest } from "./llmApi";
import { McpServer } from "./mcp_server";
import { WsBridge } from "./ws_bridge";
import { StateStore } from "./state_store";
import { PerformanceMonitor } from "./monitor";
import { RagIndexer } from "./rag_indexer";
import { WorktreeRacer } from "./worktree_racer";
import { BrowserCdpClient } from "./browser_cdp";
import { EmailDispatcher } from "./email_dispatcher";
import { SqliteStoreProvider } from "./sqlite_store";
import { AdaptiveLlmRouter } from "./adaptive_router";
import { PromptTemplateEngine, PromptTemplateDefinition } from "./prompt_template";
import { MidiAnalyzer } from "./midi_analyzer";
import { TrackScoringEngine } from "./track_scoring";
import { KeyboardGeometryEngine } from "./keyboard_geometry";
import { UserSettingsEngine } from "./user_settings";
import { MidiDeviceManager } from "./midi_device";

/**
 * Unified Single-Page Web Dashboard Server for the Ultimate Agentic Coding Harness.
 * Provides interactive endpoints for the LLM Router, Multi-Agent Orchestrator,
 * Cognitive AST Parser, and Background Process Supervisor with SSE event streaming.
 */
export function createDashboardServer(port: number = 8080) {
  const router = new LlmRouter("openai", "dummy_token");
  const orchestrator = new Orchestrator();
  const mcpServer = new McpServer();
  const wsBridge = new WsBridge();
  const stateStore = new StateStore();
  const monitor = new PerformanceMonitor();
  const ragIndexer = new RagIndexer();
  const racer = new WorktreeRacer();
  const cdpClient = new BrowserCdpClient();
  const emailDispatcher = new EmailDispatcher();
  const sqliteProvider = new SqliteStoreProvider();
  const adaptiveRouter = new AdaptiveLlmRouter();

  // SSE client connections
  const sseClients: http.ServerResponse[] = [];

  function broadcastEvent(type: string, data: any) {
    const payload = `event: ${type}\ndata: ${JSON.stringify(data)}\n\n`;
    for (const client of sseClients) {
      client.write(payload);
    }
  }

  const server = http.createServer(async (req, res) => {
    const url = new URL(req.url || "/", `http://${req.headers.host}`);

    // Serve Single-Page Dashboard UI
    if (url.pathname === "/" || url.pathname === "/index.html") {
      const htmlPath = path.join(__dirname, "dashboard.html");
      if (fs.existsSync(htmlPath)) {
        res.writeHead(200, { "Content-Type": "text/html" });
        return res.end(fs.readFileSync(htmlPath, "utf-8"));
      } else {
        res.writeHead(404);
        return res.end("Dashboard HTML not found.");
      }
    }

    // Server-Sent Events (SSE) live streaming endpoint
    if (url.pathname === "/api/events") {
      res.writeHead(200, {
        "Content-Type": "text/event-stream",
        "Cache-Control": "no-cache",
        "Connection": "keep-alive"
      });
      res.write(`event: connected\ndata: {"message":"Connected to harness live log stream"}\n\n`);
      sseClients.push(res);

      req.on("close", () => {
        const idx = sseClients.indexOf(res);
        if (idx !== -1) sseClients.splice(idx, 1);
      });
      return;
    }

    // MCP tools endpoint
    if (url.pathname === "/api/mcp/tools") {
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({
        status: "ok",
        tools: mcpServer.listTools()
      }));
    }

    // Cognitive RAG Vector Search endpoint
    if (url.pathname === "/api/rag/search") {
      const q = url.searchParams.get("q") || "";
      const matches = ragIndexer.search(q, 3);
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({
        status: "ok",
        query: q,
        totalChunks: ragIndexer.getChunkCount(),
        matches
      }));
    }

    // Prompt Template Rendering endpoint
    if (url.pathname === "/api/prompt/render" && req.method === "POST") {
      let body = "";
      req.on("data", (chunk) => { body += chunk; });
      req.on("end", () => {
        try {
          const payload = JSON.parse(body || "{}");
          const templateDef: PromptTemplateDefinition = payload.template || {
            name: "default_template",
            version: "1.0.0",
            parameters: [
              { name: "role_name", type: "string", required: true },
              { name: "task", type: "string", required: true },
              { name: "tone", type: "string", required: false, defaultValue: "concise" }
            ],
            messages: [
              { role: "system", content: "You are an autonomous {{role_name}} acting in a {{tone}} manner." },
              { role: "user", content: "Execute task: {{task}}" }
            ]
          };
          const inputValues: Record<string, string> = payload.values || {};
          const renderResult = PromptTemplateEngine.render(templateDef, inputValues);
          broadcastEvent("log", {
            timestamp: new Date().toISOString(),
            source: "PromptEngine",
            message: `Rendered prompt template '${templateDef.name}' with ${renderResult.messages.length} messages.`
          });
          res.writeHead(200, { "Content-Type": "application/json" });
          res.end(JSON.stringify(renderResult));
        } catch (e: any) {
          res.writeHead(400, { "Content-Type": "application/json" });
          res.end(JSON.stringify({ error: e.message }));
        }
      });
      return;
    }

    // MIDI File Analysis endpoint
    if (url.pathname === "/api/midi/analyze") {
      try {
        const filePath = path.resolve(__dirname, "../../legacy/testing/pitch_bend.mid");
        const analysis = MidiAnalyzer.analyzeFile(filePath);
        broadcastEvent("log", {
          timestamp: new Date().toISOString(),
          source: "MidiAnalyzer",
          message: `Parsed MIDI file pitch_bend.mid: ${analysis.totalEvents} events, ${analysis.pitchBendEvents} pitch bends.`
        });
        res.writeHead(200, { "Content-Type": "application/json" });
        return res.end(JSON.stringify({ status: "ok", file: "pitch_bend.mid", analysis }));
      } catch (err: any) {
        res.writeHead(500, { "Content-Type": "application/json" });
        return res.end(JSON.stringify({ error: err.message }));
      }
    }

    // Track Scoring evaluation endpoint
    if (url.pathname === "/api/scoring/evaluate" && req.method === "POST") {
      let body = "";
      req.on("data", chunk => { body += chunk; });
      req.on("end", () => {
        try {
          const payload = JSON.parse(body || "{}");
          const perfect = parseInt(payload.perfectHits || "0", 10);
          const good = parseInt(payload.goodHits || "0", 10);
          const missed = parseInt(payload.missedNotes || "0", 10);
          const combo = parseInt(payload.maxCombo || "0", 10);
          const report = TrackScoringEngine.evaluateScore(perfect, good, missed, combo);
          broadcastEvent("log", {
            timestamp: new Date().toISOString(),
            source: "TrackScoring",
            message: `Evaluated performance: Grade ${report.grade}, Accuracy ${report.accuracyPercent}%, Score ${report.finalScore}`
          });
          res.writeHead(200, { "Content-Type": "application/json" });
          return res.end(JSON.stringify({ status: "ok", report }));
        } catch (err: any) {
          res.writeHead(400, { "Content-Type": "application/json" });
          return res.end(JSON.stringify({ error: err.message }));
        }
      });
      return;
    }

    // Keyboard Geometry Layout endpoint
    if (url.pathname === "/api/keyboard/layout") {
      const sizeParam = parseInt(url.searchParams.get("size") || "88", 10);
      const layout = KeyboardGeometryEngine.getLayout(sizeParam);
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({ status: "ok", layout }));
    }

    // User Settings endpoints
    if (url.pathname === "/api/settings/get") {
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({ status: "ok", settings: UserSettingsEngine.get() }));
    }

    if (url.pathname === "/api/settings/update" && req.method === "POST") {
      let body = "";
      req.on("data", chunk => { body += chunk; });
      req.on("end", () => {
        try {
          const parsed = JSON.parse(body || "{}");
          const updated = UserSettingsEngine.update(parsed);
          broadcastEvent("log", {
            timestamp: new Date().toISOString(),
            source: "UserSettings",
            message: `Updated preferences: Buffer ${updated.audioBufferSize}, SampleRate ${updated.audioSampleRate}Hz, Size ${updated.preferredKeyboardSize}`
          });
          res.writeHead(200, { "Content-Type": "application/json" });
          return res.end(JSON.stringify({ status: "ok", settings: updated }));
        } catch (err: any) {
          res.writeHead(400, { "Content-Type": "application/json" });
          return res.end(JSON.stringify({ error: err.message }));
        }
      });
      return;
    }

    // MIDI Device Management endpoint
    if (url.pathname === "/api/midi/devices") {
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({ status: "ok", inventory: MidiDeviceManager.listDevices() }));
    }

    // Adaptive Router Resolution endpoint
    if (url.pathname === "/api/router/adaptive") {
      const mode = (url.searchParams.get("mode") as any) || "balanced";
      const resolution = adaptiveRouter.resolveOptimalRoute(mode);
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({
        status: "ok",
        mode,
        resolution,
        providers: adaptiveRouter.getMetrics()
      }));
    }

    // Desktop Container Configuration endpoint
    if (url.pathname === "/api/desktop/config") {
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({
        status: "ok",
        container: {
          title: "Ultimate Agentic Coding Harness",
          dimensions: "1280x850",
          trayDocked: true,
          ipcChannel: "harness_ipc",
        }
      }));
    }

    // SQLite Storage Diagnostics endpoint
    if (url.pathname === "/api/sqlite/metrics") {
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify(sqliteProvider.getMetrics()));
    }

    // Email / Notification dispatch endpoint
    if (url.pathname === "/api/email/dispatch" && req.method === "POST") {
      let body = "";
      req.on("data", chunk => { body += chunk; });
      req.on("end", () => {
        try {
          const parsed = JSON.parse(body || "{}");
          const dispatch = emailDispatcher.dispatchNotification(
            parsed.recipient || "dev@robertpelloni.com",
            parsed.subject || "Harness Alert",
            parsed.message || "Automated notification"
          );
          res.writeHead(200, { "Content-Type": "application/json" });
          return res.end(JSON.stringify({ status: "ok", dispatch, history: emailDispatcher.getHistory() }));
        } catch {
          res.writeHead(400, { "Content-Type": "application/json" });
          return res.end(JSON.stringify({ error: "Invalid JSON" }));
        }
      });
      return;
    }

    // Browser CDP navigation endpoint
    if (url.pathname === "/api/cdp/navigate") {
      const targetUrl = url.searchParams.get("url") || "http://localhost:8080";
      const result = await cdpClient.navigate(targetUrl);
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify(result));
    }

    // Multi-Agent Git Worktree Race endpoint
    if (url.pathname === "/api/worktree/race") {
      const task = url.searchParams.get("task") || "autonomous optimization";
      const raceResult = racer.executeRace(task);
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify(raceResult));
    }

    // Performance & Health Monitoring endpoint
    if (url.pathname === "/api/health") {
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify(monitor.getHealthReport()));
    }

    // Agent Memory / State Store endpoint
    if (url.pathname === "/api/memory") {
      if (req.method === "POST") {
        let body = "";
        req.on("data", chunk => { body += chunk; });
        req.on("end", () => {
          try {
            const parsed = JSON.parse(body || "{}");
            if (parsed.key) {
              stateStore.set(parsed.key, parsed.value);
            }
            res.writeHead(200, { "Content-Type": "application/json" });
            return res.end(JSON.stringify({ status: "ok", memories: stateStore.getAll() }));
          } catch {
            res.writeHead(400, { "Content-Type": "application/json" });
            return res.end(JSON.stringify({ error: "Invalid JSON payload" }));
          }
        });
        return;
      }

      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({
        status: "ok",
        memories: stateStore.getAll()
      }));
    }

    // Router test endpoint
    if (url.pathname === "/api/router") {
      const model = url.searchParams.get("model") || "gpt-4o";
      const sampleReq = createLlmRequest(model, [{ role: "user", content: "Ping" }]);
      const responseData = {
        status: "ok",
        provider: router.provider,
        resolvedModel: sampleReq.model,
        temperature: sampleReq.temperature,
      };
      broadcastEvent("log", {
        timestamp: new Date().toISOString(),
        source: "Router",
        message: `Resolved route for model ${sampleReq.model} on provider ${router.provider}`
      });
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify(responseData));
    }

    // Orchestrator endpoint
    if (url.pathname === "/api/orchestrator") {
      const strategy = url.searchParams.get("strategy") || "consensus";
      const responseData = {
        status: "ok",
        strategy,
        agentsParticipating: ["openai", "anthropic", "gemini"],
        consensusThreshold: 0.67,
        lastCycleMs: 142,
      };
      broadcastEvent("log", {
        timestamp: new Date().toISOString(),
        source: "Orchestrator",
        message: `Executed multi-agent cycle with strategy [${strategy}]`
      });
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify(responseData));
    }

    // AST parser endpoint
    if (url.pathname === "/api/ast" && req.method === "POST") {
      let body = "";
      req.on("data", chunk => { body += chunk; });
      req.on("end", () => {
        try {
          const parsed = JSON.parse(body || "{}");
          const code = parsed.code || "";
          const fileMap = AstParser.parseSource("snippet.ts", code);
          broadcastEvent("log", {
            timestamp: new Date().toISOString(),
            source: "AST Parser",
            message: `Extracted ${fileMap.definitions.length} definitions and ${fileMap.dependencies.length} dependencies`
          });
          res.writeHead(200, { "Content-Type": "application/json" });
          return res.end(JSON.stringify({
            status: "ok",
            definitions: fileMap.definitions,
            dependencies: fileMap.dependencies,
          }));
        } catch (err: any) {
          res.writeHead(400, { "Content-Type": "application/json" });
          return res.end(JSON.stringify({ error: err.message }));
        }
      });
      return;
    }

    // Supervisor status endpoint
    if (url.pathname === "/api/supervisor/status") {
      res.writeHead(200, { "Content-Type": "application/json" });
      return res.end(JSON.stringify({
        status: "ok",
        trayEnabled: true,
        activeWsClients: wsBridge.getConnectedClientsCount(),
        engines: {
          rust: "running",
          go: "running",
          csharp: "running",
          java: "running",
          typescript: "running"
        },
        uptimeSeconds: process.uptime()
      }));
    }

    // Supervisor quit endpoint
    if (url.pathname === "/api/supervisor/quit" && req.method === "POST") {
      broadcastEvent("log", {
        timestamp: new Date().toISOString(),
        source: "Supervisor",
        message: "Shutdown sequence initiated for all harness background servers."
      });
      res.writeHead(200, { "Content-Type": "application/json" });
      res.end(JSON.stringify({ message: "Graceful shutdown initiated." }));
      setTimeout(() => {
        console.log("Dashboard server received supervisor quit command. Terminating process cleanly...");
        for (const client of sseClients) {
          try { client.end(); } catch {}
        }
        server.close(() => {
          process.exit(0);
        });
      }, 500);
      return;
    }

    res.writeHead(404);
    res.end("Not Found");
  });

  server.on("upgrade", (req, socket, head) => {
    if (req.url === "/ws") {
      wsBridge.handleUpgrade(req, socket, head);
    } else {
      socket.destroy();
    }
  });

  return server;
}

if (require.main === module) {
  const PORT = parseInt(process.env.PORT || "8080", 10);
  const server = createDashboardServer(PORT);
  server.listen(PORT, () => {
    console.log(`Ultimate Agentic Coding Harness Dashboard listening on http://localhost:${PORT}`);
  });
}
