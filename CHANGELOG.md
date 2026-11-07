# CHANGELOG

See `VERSION.md` for the current active version string.

## [v5.85.0]

### Added
* Defined External Email & Notification Dispatch Protocol in `schema/email_protocol.md`.
* Implemented native Email & Notification Dispatcher in TypeScript (`typescript/src/email_dispatcher.ts`) and Go (`go/llmapi/email_dispatcher.go`) supporting inbound task message parsing and outbound system notifications.
* Integrated Email & Notification Dispatch card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.85.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.84.0]

### Added
* Defined Chrome DevTools Protocol (CDP) & Headless Browser Automation Protocol in `schema/browser_cdp_protocol.md`.
* Implemented native Headless Browser CDP Client in TypeScript (`typescript/src/browser_cdp.ts`) and Go (`go/llmapi/browser_cdp.go`) providing page navigation, DOM inspection, and script evaluation.
* Integrated Browser Automation & CDP card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.84.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).

## [v5.83.0]

### Added
* Defined Multi-Agent Git Worktree Racing Protocol in `schema/worktree_racing.md`.
* Implemented Git Worktree Racing Orchestrator in TypeScript (`typescript/src/worktree_racer.ts`) and Go (`go/llmapi/worktree_racer.go`) to spin up isolated branch worktrees for competing models, run automated tests, and determine winning solutions.
* Integrated Git Worktree Racing card into the Unified Single-Page Web Dashboard (`typescript/src/dashboard.html`, `dashboard.ts`).

### Changed
* Updated `VERSION.md` to `v5.83.0` and synchronized package manifests across TypeScript, Rust, Java, and C#.
* Synchronized documentation files (`ROADMAP.md`, `TODO.md`, `HANDOFF.md`, `MEMORY.md`).
