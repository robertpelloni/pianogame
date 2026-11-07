# Multi-Agent Git Worktree Racing Protocol

## Overview
Git Worktree Racing is an advanced orchestration strategy where multiple distinct AI models (e.g. GPT-4o, Claude 3.5 Sonnet, Gemini 1.5 Pro) receive the same software engineering issue and concurrently develop candidate solutions in isolated Git worktrees.

## Lifecycle Stages

1. **Worktree Provisioning (`git worktree add`):**
   - An isolated temporary branch and directory is provisioned for each participating model:
     - `candidate/model_gpt4o` -> `.worktrees/gpt4o`
     - `candidate/model_claude` -> `.worktrees/claude`
     - `candidate/model_gemini` -> `.worktrees/gemini`

2. **Concurrent Execution:**
   - Each model synthesizes diffs and runs the project's automated build and unit test suite in its isolated worktree.

3. **Benchmarking & Scoring:**
   - Candidate solutions are scored on:
     - Test pass rate (100% required)
     - Execution latency (speed to complete)
     - Diff cleanliness & cyclomatic complexity
     - Token cost efficiency

4. **Auto-Merge & Cleanup:**
   - The winning solution branch is merged into the active feature branch.
   - Worktrees are pruned via `git worktree prune`.
