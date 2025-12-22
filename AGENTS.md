# Repository Guidelines

## Project Structure & Module Organization
Tyr365AdminCli is a Cobra-based Go CLI. Entry point `main.go` loads `config.InitViper` before delegating to `cmd/root.go`. Command implementations live in subfolders under `cmd/` (e.g., `cmd/archiver`, `cmd/azure`, `cmd/graphCommands`, `cmd/sp`, `cmd/teamGov`, `cmd/teamToolbox`), while reusable logic sits in helper packages like `AzureHelper`, `graphHelper`, `spHelper`, `m365Archiver`, `TeamToolBoxHelper`, `SaveToFile`, and `logger`. Shared data models remain in `structs/`, and OpenAPI descriptions are stored under `OpenApiSpecs/` for reference when working with Microsoft Graph endpoints.

## Build, Test, and Development Commands
Use `go build ./...` to ensure every package (commands plus helpers) compiles. Run `go run . --help` during development to exercise the CLI with the loaded configuration. Execute `go test ./...` for package tests; target subtrees with commands like `go test ./graphHelper`. Format code automatically via `go fmt ./...` and run `go vet ./...` before submitting.

## Coding Style & Naming Conventions
Follow idiomatic Go: tab indentation, `camelCase` for private identifiers, `PascalCase` for exported ones, and keep files scoped by package responsibility. Stick to Cobra conventions—each command file should expose a `New...Cmd()` factory and register with `rootCmd`. Prefer dependency injection via interfaces when wiring Azure or Graph helpers; keep configuration keys in `config.json` lowercase with dot-separated segments (e.g., `graph.clientId`).

## Testing Guidelines
Place tests alongside source files using the `_test.go` suffix and Go’s `testing` package. Focus on validating command option wiring, helper integrations with mocked Graph clients, and table-driven edge cases. Aim for at least 70% coverage on new packages and add regression tests when touching existing helpers. Run targeted suites with `go test ./cmd/... -run TestTeamGov` when iterating quickly.

## Commit & Pull Request Guidelines
Existing history is terse (`fix`, `updates`); move toward descriptive, imperative subject lines (`feat: add graph consent command`). Group related changes per commit, include configuration or schema updates explicitly, and reference issue IDs when available. Pull requests should summarize the user-facing impact, list new commands or flags, note config schema changes, and attach sample CLI output when behavior changes.

## Configuration & Secrets
`config/viperConfig.go` searches for `config.json` under `/root/condigurationFolder/` or `~/condigurationFolder/`. Keep environment-specific credentials out of version control; document required keys in the PR description and provide sanitized samples under `config/` if needed.
