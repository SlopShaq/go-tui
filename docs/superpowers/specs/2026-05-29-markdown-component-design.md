# Markdown Component Design

- Date: 2026-05-29
- Status: Approved (design); revised after technical review; pending implementation plan
- Issue: [#62 markdown support](https://github.com/grindlemire/go-tui/issues/62)

> Revision note: a technical review verified against the renderer found that text
> rendering is gated on `e.text != ""` in three separate places (normal render,
> clipped/scroll render, and layout measurement), that no-wrap multiline text
> collapses to a single line, that OSC 8 must thread through the cell-diff/Flush
> path rather than `bufferRowToANSI` alone, that cached component elements need
> `PropsUpdater`, and that gsx cannot type-discriminate expressions. Full scope
> (lists, blockquotes, setext, links via OSC 8) is retained by decision; the spec
> below incorporates the correctness fixes.

## Summary

Add a `Markdown` widget that renders a markdown string into the go-tui widget
tree, in the spirit of `glow`. The widget is a pure content renderer: it returns
a block tree and leaves scrolling to the caller. It is usable from Go directly
(`tui.NewMarkdown(...)`) and from `.gsx` via a self-closing `<markdown>` tag.

Rendering markdown correctly requires inline mixed styling (a bold word inside an
otherwise-plain sentence that still wraps at word boundaries). The current
`Element` holds a single `text` string with a single `textStyle`, so this design
introduces a reusable rich-text primitive (`TextSpan` + `WithRichText`) as a
first-class public framework feature. Markdown is its first consumer.

## Goals

- Render headings, bold, italic, inline code, fenced code blocks, pipe tables,
  ordered and unordered lists (including nesting), blockquotes, and links.
- Provide a public, documented, tested rich-text primitive that other widgets
  can reuse (inline code, styled log views, future syntax highlighting).
- Keep the project's zero-external-dependency policy intact.
- Support both static `string` content and reactive `*State[string]` content.

## Non-goals (v1)

- Code-block syntax highlighting. The rich-text primitive makes it possible
  later, but per-language tokenizers are a separate effort.
- Images, HTML passthrough, footnotes, task lists, definition lists. Table column
  alignment markers (`:--`, `:-:`, `--:`) are parsed but not applied in v1; all
  columns render left-aligned.
- Writing literal markdown as `.gsx` children (would require a raw-text lexer
  mode). Content arrives through an expression attribute only.
- The widget owning scroll state or key bindings. Callers wrap it in a
  `scrollable` container, consistent with existing examples.

## Architecture

The work divides into five layers, each independently testable:

```
Layer 1  Rich-text primitive            (tui package: TextSpan, wrapping, render, measure)
Layer 2  OSC 8 hyperlink support        (tui package: Cell, buffer Diff, Flush, escape)
Layer 3  Markdown parser                (internal/markdown: zero-dep, recursive block tree)
Layer 4  Markdown component             (tui package: markdown.go, markdown_options.go, theme)
Layer 5  gsx integration                (tuigen generator, analyzer, LSP schema, testdata)
```

Layers 1 and 2 are general framework features with no markdown knowledge. Layer 3
has no dependency on the `tui` package. Layer 4 composes 1, 2, and 3. Layer 5 is
codegen plumbing.

## Layer 1: Rich-text primitive

### Types and API

```go
// TextSpan is a run of text sharing one style. A zero-value Style means the
// span inherits the element's textStyle; set fields override it.
type TextSpan struct {
    Text  string
    Style Style
    Link  string // optional OSC 8 hyperlink target (see Layer 2)
}

// WithRichText sets styled, multi-segment text on an element. When set it takes
// precedence over WithText. Wrapping and alignment behave as for plain text.
func WithRichText(spans ...TextSpan) Option

// Mutation + accessor, mirroring SetText/Text:
func (e *Element) SetRichText(spans ...TextSpan)
func (e *Element) RichText() []TextSpan
```

`Element` gains a `richText []TextSpan` field alongside `text`. When `richText`
is non-empty it is the source of truth for text rendering and measurement.

Precedence and clearing semantics (so the two fields never disagree):

- `SetRichText`/`WithRichText` clears `text`.
- `SetText`/`WithText` clears `richText`.
- If both are somehow set, `richText` wins.

### Style-merge semantics

For each span, the effective per-cell style is the element's inherited
`textStyle` with the span's set fields layered on top:

- Attribute bits (bold, italic, underline, etc.) are OR-ed in.
- A non-default `Fg`/`Bg` on the span overrides the inherited color; a default
  color leaves the inherited value untouched.
- Background from the element/inherited context still applies, matching the
  existing plain-text background handling in `renderTextContent`.

### Wrapping

Add `wrapSpans(spans []TextSpan, maxWidth int) [][]TextSpan` in `text_wrap.go`,
mirroring `wrapParagraph`'s word-packing algorithm:

- Break at word boundaries; fall back to mid-word breaks when a single word
  exceeds `maxWidth` (same rule as today).
- Each emitted word carries the style of the span it came from, so a multi-word
  bold run stays bold across a line break.
- Adjacent same-style segments on a line are merged to keep lines compact.
- Existing newlines split paragraphs, as in `wrapText`.

### Rendering (three call sites, not one)

Text rendering currently lives in three places, each gated on `e.text != ""` and
each with its own copy of the wrap/align/per-cell logic:

1. `renderElement` -> `renderTextContent` (`element_render.go:137`), normal pass.
2. `renderClippedElement` (`element_render.go:246`), the scroll/clip pass, with a
   fully duplicated text loop. **This is why callers wrapping markdown in a
   scrollable container would otherwise see nothing** unless rich text is wired
   in here too.
3. `element_layout.go:89`, the intrinsic-measurement path.

All three gates must also fire on `len(e.richText) > 0`, and all three must
render/measure spans. The existing per-cell loops (used today by the text-gradient
path) already walk runes calling `buf.SetRune(x, y, r, style)`; the rich-text
branch reuses that loop, supplying each segment's merged style instead of a
gradient color. Per-line alignment, truncation, and scroll offset operate on line
counts and carry over unchanged.

Because we are adding a second branch to three near-identical copies, the
implementation extracts the shared text layout (wrap -> lines -> per-line
align -> per-cell emit) into one helper consumed by all three sites. This is a
targeted refactor of code we are already editing, not a speculative one.

### Measurement

Auto-width rich text must report intrinsic size from its spans. Today
`element_layout.go:89-91` measures plain text as `stringWidth(e.text)` wide and a
hard-coded `1` row. The rich-text path computes width from the concatenated span
text and height from `wrapSpans` at the resolved content width (matching how the
wrapped plain-text height is computed at `element_layout.go:208`). A short spike
confirms plain-text sizing does not regress; this is the primary technical risk.

## Layer 2: OSC 8 hyperlinks

Links render as styled text and, on capable terminals, as real clickable
hyperlinks via the OSC 8 escape sequence. On other terminals the text is still
shown styled and the URL is inert.

This is the largest framework change in the design because the live render path
is `Render -> buf.Diff() -> Terminal.Flush([]CellChange)` (`render.go:8`,
`terminal_ansi.go:78`), driven by per-cell diffing, not by `bufferRowToANSI`
(which only serves the standalone/`PrintAbove` path). Link state must therefore
travel with the cell through the whole pipeline:

- `Cell` gains a link target (interned id into a per-buffer URL table, so cell
  copies/comparisons stay cheap; the table maps id -> URL).
- `Cell` equality / diff must include the link id, so a link appearing or
  changing produces a `CellChange`. This touches `cell.go`, `buffer.go` Diff, and
  the `CellChange` type.
- `ANSITerminal.Flush` tracks the current link id across emitted cells and opens
  `OSC 8 ; ; URL ST` when entering a run and closes with `OSC 8 ; ; ST` when
  leaving it (or when a non-contiguous jump occurs). The standalone
  `bufferRowToANSI` path (`render_element.go`) gets the same run handling.
- Capability gating reuses `caps.go`; when unsupported, `Flush` omits the escape
  and applies style only.
- `MockTerminal` records link runs so golden tests can assert them.

This layer is general framework functionality; markdown links are its first
consumer via `TextSpan.Link`. Risk: Flush is the hottest path in the renderer, so
the link-run state machine must add no work when no cell carries a link.

## Layer 3: Markdown parser (`internal/markdown`)

Zero-dependency, line-oriented, and independent of the `tui` package so it can be
unit-tested on its own data types.

### Block tree

Because blockquotes and nested lists contain other blocks, the parser produces a
recursive block tree rather than a flat list:

```go
type Block struct {
    Kind     BlockKind   // Heading, Paragraph, CodeFence, Table, List, ListItem, Blockquote
    Level    int         // heading level, or list-nesting depth
    Ordered  bool        // list ordering
    Lang     string      // code fence language tag (stored, unused in v1)
    Inline   []Inline      // inline content for leaf blocks (heading, paragraph)
    Rows     [][]TableCell // table rows; first row is the header
    Children []Block       // nested blocks (blockquote contents, list items)
    Lines    []string      // raw lines for code fences
}

// TableCell holds one cell's inline content (named to avoid colliding with the
// buffer Cell type in the tui package).
type TableCell struct {
    Inline []Inline
}
```

(Concrete field set is finalized in the plan; the shape above is the contract.)

### Inline scanner

Turns a string into `[]Inline` (which Layer 4 converts to `[]TextSpan`) by
toggling state on markers:

- `**text**` / `__text__` -> bold
- `*text*` / `_text_` -> italic
- `` `code` `` -> inline code; suppresses other markers inside
- `[text](url)` -> link with text + target

### Block constructs and disambiguation

- ATX headings: `#`..`######`.
- Setext headings: a paragraph line followed by `===` (h1) or `---` (h2). The
  `---` form is a setext underline only when the preceding line is non-blank
  paragraph text; otherwise `---` is a horizontal rule or, inside a table, the
  separator row.
- Fenced code blocks: ```` ``` ```` delimited; contents are literal, not scanned
  for inline markers.
- Pipe tables: header row, separator row (`---`/`:--:` etc.), body rows.
- Lists: ordered (`1.`) and unordered (`-`, `*`, `+`); nesting tracked by
  indentation, producing nested `List`/`ListItem` blocks.
- Blockquotes: `>` prefix; contents are parsed recursively as blocks.
- Anything unrecognized degrades to a Paragraph of plain text.

## Layer 4: Markdown component

### Files

- `markdown.go` -- `Markdown` struct and `Render`.
- `markdown_options.go` -- option funcs.
- `markdown_theme.go` -- `MarkdownTheme` and `DefaultMarkdownTheme()`.

### Struct and interfaces

```go
type Markdown struct {
    source string
    state  *State[string]   // optional reactive source
    width  int              // 0 = fill available width
    theme  MarkdownTheme
    cache  parseCache       // last source -> parsed block tree
}

var (
    _ Component    = (*Markdown)(nil)
    _ AppBinder    = (*Markdown)(nil)
    _ PropsUpdater = (*Markdown)(nil)
)
```

Component elements are cached via `mount`/`MountPersistent`, which calls
`UpdateProps` on the cached instance when present (`mount.go:69`). So `Markdown`
must implement `PropsUpdater`: `UpdateProps(fresh)` copies the new source, width,
and theme from the freshly-constructed instance, otherwise `<markdown source={expr}/>`
renders stale content when `expr` changes.

`BindApp` is implemented unconditionally (interfaces are type-level in Go, so it
cannot be "only when state-backed"); it binds the `*State[string]` when one is
present and is a no-op otherwise.

`Render` resolves the current source (state takes precedence when set), parses it
(or returns the cached tree when the source is unchanged), and walks the block
tree into a `flex-col` `*Element` root.

### Block-to-element mapping

| Block        | Element                                                                 |
|--------------|-------------------------------------------------------------------------|
| Heading      | `WithRichText(spans)` styled by `theme.Heading[level]`                   |
| Paragraph    | `WithRichText(spans)` styled by `theme.Paragraph`                        |
| Inline code  | a `TextSpan` styled by `theme.CodeSpan`                                  |
| Link         | a `TextSpan` styled by `theme.Link` with `Link` set                     |
| Code fence   | bordered/`bg` `flex-col`, **one child element per line** (`WithWrap(false)`, see note) |
| Table        | existing `<table>/<tr>/<th>/<td>` element tree; each cell is rich text   |
| List         | `flex-col`; each item prefixed `theme.BulletMarker` or `"N. "`, indented per depth |
| Blockquote   | `flex-row`: a styled bar column + indented content `flex-col` (see note), recursive |

Code fences must not use a single multiline `WithText`: the no-wrap text path
collapses to one line (`element_render.go:263`) and measures as height 1
(`element_layout.go:91`). Each source line becomes its own child element inside a
bordered/`bg` column. (Once multiline no-wrap text is supported framework-wide,
this can collapse to a single element; that is future work, not v1.)

Blockquotes cannot use a `BorderStyle` for the left bar: borders draw a full box
(`DrawBox`). The bar is a 1-cell-wide child column filled with the bar glyph
(`theme.BlockquoteBar`, e.g. `│`) styled by `theme.BlockquoteBarStyle`, placed in
a `flex-row` beside the indented, recursively-rendered content.

### Options

```go
func WithMarkdownSource(s string) MarkdownOption
func WithMarkdownState(s *State[string]) MarkdownOption
func WithMarkdownWidth(w int) MarkdownOption
func WithMarkdownTheme(t MarkdownTheme) MarkdownOption
```

### Caching

A single-entry cache keyed on the source string avoids re-parsing identical
content every frame. State-backed content invalidates the cache when the string
changes. The precise invalidation hook is finalized in the plan.

## MarkdownTheme

A flat struct of `Style` fields plus a few non-style extras, with a
`DefaultMarkdownTheme()` constructor:

```go
type MarkdownTheme struct {
    Heading    [6]Style    // per-level heading styles
    Paragraph  Style
    Bold       Style
    Italic     Style
    CodeSpan   Style       // inline `code`
    Link       Style

    CodeBlockText   Style
    CodeBlockBg     Color
    CodeBlockBorder BorderStyle // full box around the code element; this one is a real border

    // Tables use the existing table layout, which draws a 1-char column gap, not
    // grid lines. v1 styles the header and an optional separator row rather than
    // a grid border. A fully bordered table is future work.
    TableHeader        Style
    TableSeparator     bool  // draw a "---" separator row under the header
    TableSeparatorChar rune  // e.g. '-'

    // Blockquotes render a 1-wide glyph column, not a BorderStyle (borders are
    // full boxes).
    BlockquoteBar      rune  // e.g. '│'
    BlockquoteBarStyle Style
    BlockquoteText     Style

    BulletMarker string      // e.g. "• "
}
```

## Layer 5: gsx integration

Expose a self-closing `<markdown>` tag fed by an expression attribute.

- `internal/tuigen/generator_element.go`: add `markdown` to `isComponentElement`,
  map it to `tui.NewMarkdown` in `componentConstructor`, and add an attribute map.
  The generator cannot type-check arbitrary Go expressions, so the source kind is
  expressed by **distinct attributes** rather than inferred:
  `source={stringExpr}` emits `WithMarkdownSource`, `state={stateExpr}` emits
  `WithMarkdownState`. There is no `value` attribute. Plus `width` and `theme`.
- `internal/tuigen/analyzer.go`: add `markdown` to `knownTags`, mark it
  self-closing/void (children rejected).
- `internal/lsp/schema/schema.go`: add the element definition with attribute
  documentation.
- `cmd/tui/testdata/`: add a `markdown.gsx` fixture and its expected
  `markdown_gsx.go`.

Usage:

```gsx
<markdown source={readme} width={80} />
<markdown state={mdState} />
```

## Testing strategy

- Parser (`internal/markdown`): table-driven tests asserting the block tree and
  inline output for each construct and for the disambiguation rules (setext vs
  rule vs table separator, nested lists, blockquote recursion).
- `wrapSpans`: table-driven wrapping cases, including style runs that straddle a
  line break and mid-word breaks.
- Rendering: `MockTerminal` golden tests asserting cell styles (bold/italic runs,
  code-span background, link styling) and layout for tables, code blocks,
  blockquotes, and nested lists.
- OSC 8: assert the emitted ANSI wraps link runs when capable and omits the
  escape when not.
- Codegen: `testdata` golden comparison for the `<markdown>` tag.

## Open questions resolved during planning

These are implementation-level and are answered by reading the relevant code
while writing the plan. They do not change the public design above.

1. The exact signature/placement of the shared text-layout helper extracted from
   the three current text sites, and confirming the spike shows no plain-text
   render or sizing regression. (Primary risk.)
2. Parse-cache invalidation hook for state-backed content.
3. Confirming `<td>`/`<th>` intrinsic sizing works when fed `richText` (uses the
   same measurement path as #1).
4. Code-block overflow: fixed to content width, wrap disabled, horizontal scroll
   left to the caller's container.
5. The per-buffer URL intern-table lifecycle (allocation, reset on resize) for
   OSC 8 link ids, and how `CellChange` carries the link id.
6. Whether `BlockquoteBar`/`BulletMarker` defaults need width-aware handling for
   wide glyphs.

## Future work

- Code-block syntax highlighting built on the rich-text primitive.
- Additional inline/block constructs (images, task lists, footnotes).
- Optional self-scrolling convenience wrapper if demand appears.
