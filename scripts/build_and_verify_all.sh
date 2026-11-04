#!/usr/bin/env bash
set -e

echo "=== 1. Checking Version Consistency ==="
VERSION=$(cat VERSION.md | tr -d '[:space:]')
echo "Active Version: ${VERSION}"

echo "=== 2. Verifying Rust Harness ==="
(cd rust && cargo check && cargo test)

echo "=== 3. Verifying Go Harness ==="
(cd go && go test ./... && go build ./...)

echo "=== 4. Verifying C# Harness ==="
(cd csharp/Harness && dotnet build)

echo "=== 5. Verifying Java Harness ==="
(cd java && mvn test-compile)

echo "=== 6. Verifying TypeScript Harness ==="
(cd typescript && npx tsc --noEmit)

echo "=== All builds and verifications succeeded! ==="
