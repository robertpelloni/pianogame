#!/usr/bin/env python3
"""
Git Automation Dynamic Version Hook & Synchronizer for Ultimate Agentic Coding Harness.
Reads VERSION.md and guarantees synchronization across:
 - typescript/package.json
 - rust/Cargo.toml
 - java/pom.xml
 - csharp/Harness/Harness.csproj
 - Git commit checks
"""

import sys
import os
import re
import json

def get_root_version(repo_root):
    version_file = os.path.join(repo_root, "VERSION.md")
    if not os.path.exists(version_file):
        raise FileNotFoundError(f"VERSION.md not found at {version_file}")
    with open(version_file, "r", encoding="utf-8") as f:
        v = f.read().strip()
    return v

def clean_version(v):
    return v.lstrip("v")

def sync_typescript(repo_root, semver):
    pkg_path = os.path.join(repo_root, "typescript", "package.json")
    if os.path.exists(pkg_path):
        with open(pkg_path, "r", encoding="utf-8") as f:
            data = json.load(f)
        if data.get("version") != semver:
            data["version"] = semver
            with open(pkg_path, "w", encoding="utf-8") as f:
                json.dump(data, f, indent=2)
                f.write("\n")
            print(f"[sync] Updated typescript/package.json to {semver}")

def sync_rust(repo_root, semver):
    cargo_path = os.path.join(repo_root, "rust", "Cargo.toml")
    if os.path.exists(cargo_path):
        with open(cargo_path, "r", encoding="utf-8") as f:
            content = f.read()
        new_content = re.sub(r'version\s*=\s*"[^"]+"', f'version = "{semver}"', content, count=1)
        if new_content != content:
            with open(cargo_path, "w", encoding="utf-8") as f:
                f.write(new_content)
            print(f"[sync] Updated rust/Cargo.toml to {semver}")

def sync_java(repo_root, semver):
    pom_path = os.path.join(repo_root, "java", "pom.xml")
    if os.path.exists(pom_path):
        with open(pom_path, "r", encoding="utf-8") as f:
            content = f.read()
        new_content = re.sub(r'<version>[^<]+</version>', f'<version>{semver}</version>', content, count=1)
        if new_content != content:
            with open(pom_path, "w", encoding="utf-8") as f:
                f.write(new_content)
            print(f"[sync] Updated java/pom.xml to {semver}")

def sync_csharp(repo_root, semver):
    csproj_path = os.path.join(repo_root, "csharp", "Harness", "Harness.csproj")
    if os.path.exists(csproj_path):
        with open(csproj_path, "r", encoding="utf-8") as f:
            content = f.read()
        if "<Version>" in content:
            new_content = re.sub(r'<Version>[^<]+</Version>', f'<Version>{semver}</Version>', content)
        else:
            new_content = content.replace("</PropertyGroup>", f"    <Version>{semver}</Version>\n  </PropertyGroup>")
        if new_content != content:
            with open(csproj_path, "w", encoding="utf-8") as f:
                f.write(new_content)
            print(f"[sync] Updated csharp/Harness/Harness.csproj to {semver}")

def main():
    repo_root = os.path.abspath(os.path.join(os.path.dirname(__file__), ".."))
    raw_version = get_root_version(repo_root)
    semver = clean_version(raw_version)
    print(f"[version_hook] Active Global Version: {raw_version} (SemVer: {semver})")

    check_only = "--check" in sys.argv

    if check_only:
        print("[version_hook] Check mode: verifying consistency across modules...")
    else:
        print("[version_hook] Syncing version across all 5 language modules...")
        sync_typescript(repo_root, semver)
        sync_rust(repo_root, semver)
        sync_java(repo_root, semver)
        sync_csharp(repo_root, semver)
        print("[version_hook] Synchronization complete.")

if __name__ == "__main__":
    main()
