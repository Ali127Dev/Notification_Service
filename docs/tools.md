# 🚀 Development Tools Setup

Follow this document to prepare your local development environment.

---

## 📦 Prerequisites

Before running any task, ensure you have these essential tools installed.

---

### 🛠️ 1. Install Go

Required for building, running, formatting, and testing the project.

- **Recommended version:** Go 1.22+
- **Download:** https://go.dev/dl

```bash
# Check installation
go version
```

---

### ⚙️ 2. Install Task (Taskfile Runner)

Task is required to execute all tasks defined in `Taskfile.yaml`.

install:

```bash
go install github.com/go-task/task/v3/cmd/task@latest
```

Check:

```bash
task --version
```

---

## 🔥 Development Tools (Used by Taskfile)

Below is a list of all tools required by specific Taskfile commands.

---

### 🔄 3. Air — Live Reload for Go

Install:

```bash
go install github.com/air-verse/air@latest
```

Check:

```bash
air -v
```

---

### 🐘 4. Goose — Database Migration Tool

Install:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Check:

```bash
goose -version
```

---

### 🧹 6. golangci-lint — Code Linting

Install:

```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Check:

```bash
golangci-lint --version
```

---

### 🐳 Docker & Docker Compose

Useful for running databases and services locally.

```bash
docker --version
docker compose version
```

---

### ✅ Environment Ready

Your development environment is now fully prepared.
