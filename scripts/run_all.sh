#!/usr/bin/env bash
set -e

echo "=== Ultimate Agentic Coding Harness: Multi-Language Demonstration ==="

echo "--- 1. Rust Harness ---"
(cd rust && cargo run --quiet)

echo "--- 2. Go Harness ---"
(cd go && go run main.go)

echo "--- 3. C# Harness ---"
(cd csharp/Harness && dotnet run --no-build)

echo "--- 4. Java Harness ---"
(cd java && mvn exec:java -q -Dexec.mainClass="com.harness.Main")

echo "--- 5. TypeScript Harness ---"
(cd typescript && npx ts-node src/index.ts)

echo "=== All 5 Harness Engines Executed Successfully! ==="
