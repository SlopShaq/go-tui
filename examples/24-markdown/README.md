# 24 - Markdown

Renders a markdown document with the `<markdown>` element inside a scrollable
container.

```bash
go run ../../cmd/tui generate markdown.gsx
go run .
```

## What it shows

- The `<markdown>` tag fed by a `source` expression, with an explicit `width` so
  list and blockquote content wraps.
- A `MarkdownTheme` default styling headings, emphasis, inline code, code blocks,
  tables, lists, and blockquotes.
- The component wrapped in an `overflow-y-scroll` container, since `Markdown`
  owns no scroll state of its own.

The sample document in `main.go` exercises every supported construct:

- ATX headings (`#`..`######`) and single-line setext headings (`===`, `---`)
- Bold, italic, combined bold-italic, inline code, and links (OSC 8 on capable
  terminals), with both `*`/`_` and `**`/`__` markers
- Edge cases: a delimiter with no closer stays literal (`see **docs`, `3 * 4`)
- Fenced code blocks, including a preserved blank line
- Pipe tables with inline formatting in cells
- Unordered lists (`-`, `*`, `+`), ordered lists, and nesting
- Blockquotes, including a nested quote and a quote containing a list

## Controls

- `j` / `k` or arrow keys: scroll
- `PageUp` / `PageDown`: scroll a page
- mouse wheel: scroll
- `q` / `Esc`: quit
