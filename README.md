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

### Homebrew

You can also install the CLI using Homebrew:

```bash
brew tap stianfro/toolbelt
brew install toolbelt
```

## Commands

### organize

Moves receipts from the `unorganized` directory into the monthly folder
matching the date prefix on each file.

### joincsv

Combines all CSV files in a directory into one file. Usage:

```bash
toolbelt joincsv <directory> <output.csv>
```

### joinpdf

Combines all PDF files in a directory into one file. Usage:

```bash
toolbelt joinpdf <directory> <output.pdf>
```

Releases are created automatically from conventional commits using release-please.
Prebuilt binaries for common platforms are available on the release page.
