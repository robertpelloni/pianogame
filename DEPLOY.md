# DEPLOY

## Deployment & Build Instructions

### Prerequisites
* Rust 1.80+ / Cargo
* Go 1.22+
* .NET SDK 8.0+ / 10.0+
* Java 21+ / Apache Maven 3.9+
* Node.js v20+ / v22+, npm, and ts-node

### Automated Multi-Language Verification
Run the unified script from repository root:
```bash
./scripts/build_and_verify_all.sh
```

To run all language engines sequentially:
```bash
./scripts/run_all.sh
```

### Unified Single-Page Mission Control Dashboard (Umbrella v6)
Start the unified WebUI dashboard server on port 8080:
```bash
cd typescript
npm install
npx tsx src/dashboard.ts
```
Open `http://localhost:8080` in your browser. All services, including the LLM Router, Adaptive Balancer, Multi-Agent Orchestrator, AST Parser, Prompt Template Studio, MCP Tool Registry, Cognitive RAG, Browser CDP, and Process Supervisor, are consolidated onto this single control surface.

### Cross-Engine System Tray Supervisor
To launch the background process supervisor and native desktop tray integration:
```bash
./scripts/start_tray.sh
```
This initializes signal trapping (`SIGINT`, `SIGTERM`), OS desktop notifications, browser launcher capabilities, and one-click graceful shutdown (`QuitAll`).

### Dynamic Manifest & Version Synchronization
To verify or synchronize SemVer versions across `package.json`, `Cargo.toml`, `pom.xml`, and `Harness.csproj`:
```bash
# Check version consistency across all 5 engines
python3 scripts/version_hook.py --check

# Synchronize manifests to match VERSION.md
python3 scripts/version_hook.py
```

### Individual Harness Execution

#### Rust Harness
```bash
cd rust
cargo run
```

#### Go Harness
```bash
cd go
go run main.go
```

#### C# Harness
```bash
cd csharp/Harness
dotnet run
```

#### Java Harness
```bash
cd java
mvn exec:java -Dexec.mainClass="com.harness.Main"
```

#### TypeScript Harness
```bash
cd typescript
npm install
npx ts-node src/index.ts
```
