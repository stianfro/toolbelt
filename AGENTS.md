# AGENTS.md

## Build/Test/Lint

- Build: go build ./...
- Format: go fmt ./...
- Vet: go vet ./...
- Test all: go test ./...
- Test single: go test ./... -run ^TestName$

## Formatting

- Imports ordered: stdlib, blank, external, blank, internal
- Use go fmt exclusively

## Error Handling

- Return errors via fmt.Errorf("%w", err) with context
- Check and return early; avoid panic

## Naming

- Exported: PascalCase; internal: camelCase
- No underscores; acronyms uppercase (HTTP, XML)

## Cobra CLI

- Commands in cmd/; root in root.go via Cobra

## Git

- Always use conventional commits, also for PR titles
