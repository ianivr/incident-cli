# incident-cli

A command-line tool written in Go for creating, listing, filtering and closing
incidents. Incidents are persisted to a local JSON file.

## Domain

An incident has an ID, title, description, severity
(low / medium / high / critical), status (open / in_progress / closed)
and timestamps.

Commands: `create`, `list` (filters: `--status`, `--severity`), `show <id>`, `close <id>`.

Out of scope: databases, HTTP, concurrency.

## Project structure

```
cmd/incident/        # entrypoint (main.go)
internal/incident/   # domain types and validation
internal/store/      # Store interface and JSON implementation
internal/cli/        # cobra commands
```

## Conventions

- Standard library + `github.com/spf13/cobra` only. Ask before adding dependencies.
- Never ignore errors. Wrap them with context: `fmt.Errorf("...: %w", err)`.
  Use sentinel errors (`errors.New`) for expected cases such as `ErrNotFound`.
- No `panic` outside `main`.
- Keep interfaces small and define them where they are consumed.
- Every package has a doc comment (`// Package x ...` or `// Command x ...`).
- Table-driven tests with `t.Run`, next to the code they test.
  Use `t.TempDir()` for tests that touch the filesystem.
- Code, identifiers and comments in English.

## Commands

```bash
go run ./cmd/incident <command>   # run
go test ./... -cover              # tests with coverage
golangci-lint run                 # linters
golangci-lint fmt                 # formatting
```

A change is done when tests pass and `golangci-lint run` reports no issues.
