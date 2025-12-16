# devtool

An **extremely stupid and simple** Go-based CLI utility whose only job is to **manage other development tools**.

`devtool` is intentionally minimal. It creates and manages a **single, separate Go module file named `go.tool.mod`** whose purpose is **tooling only** — nothing else.

It is designed to be used as a **`go tool`**, and is meant to become the **main entry point for managing all project tools**.

---

## 🧠 Philosophy

This tool is:

* 🧱 **Stupid** — no magic, no abstractions
* ✂️ **Simple** — does one thing and does it plainly
* 🧰 **Opinionated** — tooling should be isolated
* 🧑‍💻 **Explicit** — tools are dependencies too

If you are looking for a flexible framework or a smart dependency manager — this is **not** it.

`devtool` exists to make tool management boring and obvious.

---

## 🎯 Problem

Go projects often depend on tools such as:

* linters (`golangci-lint`)
* generators (`mockgen`, `stringer`)
* formatters (`goimports`)
* analyzers and CI helpers

The usual solutions are awkward.

### ❌ `tools.go` pattern

```go
//go:build tools
package tools

import _ "github.com/golangci/golangci-lint/cmd/golangci-lint"
```

Problems:

* Pollutes the **main** module dependency graph
* Confusing to humans
* Mixes runtime and tooling concerns
* Exists only to trick `go mod`

---

### ⚠️ Note on Go 1.24 `go get -tool`

Go 1.24 introduces the official `go get -tool` mechanism from the Go team. While it is a step forward, in practice it can still messy:

* Tool dependencies are still resolved within the **main module graph**
* Tool requirements can influence versions of shared dependencies
* Upgrading a tool may unexpectedly affect application dependencies
* The boundary between *runtime* and *tooling* remains blurred

For larger projects and teams, this can reintroduce the very problems tool management is meant to solve.

`devtool` deliberately avoids this by fully isolating tools into `go.tool.mod`, making conflicts structurally impossible.

---

## ✅ Solution

`devtool` creates a **dedicated tooling module** using a **separate file**:

```
go.tool.mod
```

This file:

* Contains **only** tool dependencies
* Is never imported by application code
* Is committed to the repository
* Is treated as first-class configuration

No fake packages. No build tags. No hacks.

---

## 🚀 What `devtool` Does

Very little — on purpose:

* 📄 Creates `go.tool.mod`
* ➕ Adds tool dependencies to it
* ▶️ Runs / executes managed tools
* 🧼 Keeps your main `go.mod` clean

That’s it.

---

## 📦 Installation & Usage

### As a project tool: (Go 1.24+)

Added as a tool to the current module:

```bash
go get -tool github.com/ErkebulanMukhamedkali/devtool/cmd/devtool
```

This is the **intended** way to use `devtool`.

---

### Globally as a binary

```bash
go install github.com/ErkebulanMukhamedkali/devtool/cmd/devtool@latest
```

---

## 📖 Typical Workflow

### 1️⃣ Initialize tooling

```bash
go tool devtool init
```

Creates:

```
go.tool.mod
go.tool.sum
```

---

### 2️⃣ Add a tool

```bash
go tool devtool add golang.org/x/vuln/cmd/govulncheck
```

The dependency goes **only** into `go.tool.mod`.

---

### 3️⃣ Use tools

```bash
go tool devtool run govulncheck
```

---

## 📁 Project Layout Example

```
my-project/
├── go.mod            # application dependencies only
├── go.sum
├── go.tool.mod       # tooling dependencies only
├── go.tool.sum
├── cmd/
├── internal/
└── Makefile
```

---

## 👥 Team Rules (Recommended)

* `go.mod` → **project only**
* `go.tool.mod` → **tools only**
* Commit both `go.tool.mod` and `go.tool.sum`
* Never add tools to the main module
* Use `devtool` as the **single interface** for tooling

---

## 🛠 Development

Requirements:

* Go 1.24+

---

## 🧑‍💻 Author

Created by **Erkebulan Mukhamedkali**.

If you like stupidly simple tools — ⭐ the repository.
