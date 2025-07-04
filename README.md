# Toolbelt

Toolbelt consolidates various scripts into a single CLI application built with Go and Cobra.

## Usage

Install Go 1.22 or newer and run:

```bash
go install github.com/stianfro/toolbelt@latest
```

Run the CLI with:

```bash
toolbelt
```

## Commands

### organize

Moves receipts from the `unorganized` directory into the monthly folder
matching the date prefix on each file.

Releases are published when tagging the repository. Prebuilt binaries for common platforms are available on the release page.
