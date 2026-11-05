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
