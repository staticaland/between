# between

A CLI tool that replaces text between two markers in files.

## Installation

### Homebrew (macOS/Linux)

```bash
brew install --cask staticaland/between/between
```

### From Source

```bash
go build -o between
```

Or install directly:

```bash
go install
```

## Usage

### Basic Example

Replace content between default HTML-style markers:

```bash
between replace -f index.html --content "Updated content"
```

This finds `<!-- BEGIN -->` and `<!-- END -->` markers and replaces everything between them.

### Custom Markers

Use custom markers for different file types:

```bash
# YAML configuration
between replace -f config.yaml \
  --start-marker "# BEGIN CONFIG" \
  --end-marker "# END CONFIG" \
  --content "database_url: postgres://localhost/mydb"

# Markdown documentation
between replace -f docs.md \
  --start-marker "## Generated Content" \
  --end-marker "## End Generated" \
  --content "This section is automatically updated"
```

### Real-world Examples

Update API documentation:

```bash
between replace -f README.md \
  --start-marker "<!-- API_ENDPOINTS -->" \
  --end-marker "<!-- /API_ENDPOINTS -->" \
  --content "$(curl -s http://localhost:3000/api/docs)"
```

Update configuration from template:

```bash
between replace -f app.config \
  --start-marker "# AUTO-GENERATED" \
  --end-marker "# END AUTO-GENERATED" \
  --content "$(cat config-template.txt)"
```

## Command Reference

```sh
between replace [flags]

Flags:
  -f, --file string          File to process (required)
      --content string       New content to insert (required)
      --start-marker string  Start marker (default "<!-- BEGIN -->")
      --end-marker string    End marker (default "<!-- END -->")
  -h, --help                help for replace
```
