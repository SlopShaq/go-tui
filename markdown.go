package tui

import (
	"fmt"
	"strings"

	"github.com/grindlemire/go-tui/internal/markdown"
)

// Markdown renders a markdown string into the widget tree. It is a pure content
// renderer: it owns no scroll state or key bindings. Wrap it in a scrollable
// container to scroll long documents. Construct with NewMarkdown.
type Markdown struct {
	source string
	state  *State[string] // optional reactive source; takes precedence over source
	width  int            // 0 = fill available width
	theme  MarkdownTheme

	// single-entry parse cache keyed on the resolved source string
	lastSource string
	cached     []markdown.Block
	parsed     bool
}

var (
	_ Component    = (*Markdown)(nil)
	_ AppBinder    = (*Markdown)(nil)
	_ PropsUpdater = (*Markdown)(nil)
)

// NewMarkdown creates a Markdown component.
func NewMarkdown(opts ...MarkdownOption) *Markdown {
	m := &Markdown{
		theme: DefaultMarkdownTheme(),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// BindApp binds the reactive source (if any) to the app. It is a no-op when the
// component has only a static source.
func (m *Markdown) BindApp(app *App) {
	if m.state != nil {
		m.state.BindApp(app)
	}
}

// UpdateProps copies new props from a freshly-constructed instance when this
// cached instance is re-rendered. The parse cache is intentionally preserved;
// Render re-parses when the resolved source string changes.
func (m *Markdown) UpdateProps(fresh Component) {
	f, ok := fresh.(*Markdown)
	if !ok {
		return
	}
	m.source = f.source
	m.state = f.state
	m.width = f.width
	m.theme = f.theme
}

// resolveSource returns the current markdown text (state wins when present).
func (m *Markdown) resolveSource() string {
	if m.state != nil {
		return m.state.Get()
	}
	return m.source
}

// ensureParsed (re)parses when the resolved source changed since last parse.
func (m *Markdown) ensureParsed() {
	src := m.resolveSource()
	if m.parsed && src == m.lastSource {
		return
	}
	m.cached = markdown.Parse(src)
	m.lastSource = src
	m.parsed = true
}

// Render parses the current source and walks the block tree into a flex-col root.
func (m *Markdown) Render(app *App) *Element {
	m.ensureParsed()

	opts := []Option{WithDirection(Column)}
	if m.width > 0 {
		opts = append(opts, WithWidth(m.width))
	}
	root := New(opts...)

	for _, b := range m.cached {
		if el := m.renderBlock(b, m.width, m.theme.Paragraph); el != nil {
			root.AddChild(el)
		}
	}
	return root
}

// renderBlock dispatches one block to its renderer. contentWidth is the width
// available to this block (0 = auto/unknown). textStyle is the base style for
// paragraph and list-item text in this context (e.g. italic inside a blockquote);
// headings, code, and tables use their own theme styles.
func (m *Markdown) renderBlock(b markdown.Block, contentWidth int, textStyle Style) *Element {
	switch b.Kind {
	case markdown.KindHeading:
		return m.renderHeading(b)
	case markdown.KindParagraph:
		return m.renderParagraph(b, textStyle)
	case markdown.KindCodeFence:
		return m.renderCodeFence(b)
	case markdown.KindList:
		return m.renderList(b, 0, contentWidth, textStyle)
	case markdown.KindBlockquote:
		return m.renderBlockquote(b, contentWidth)
	case markdown.KindTable:
		return m.renderTable(b)
	default:
		// Unknown leaf: render its inline text as a paragraph so nothing is
		// silently dropped.
		return m.renderParagraph(b, textStyle)
	}
}

func (m *Markdown) renderHeading(b markdown.Block) *Element {
	level := b.Level
	if level < 1 {
		level = 1
	}
	if level > 6 {
		level = 6
	}
	return New(
		WithTextStyle(m.theme.Heading[level-1]),
		WithRichText(m.inlineToSpans(b.Inline)...),
		WithMarginTRBL(0, 0, 1, 0), // a blank line after every heading
	)
}

func (m *Markdown) renderParagraph(b markdown.Block, textStyle Style) *Element {
	return New(
		WithTextStyle(textStyle),
		WithRichText(m.inlineToSpans(b.Inline)...),
	)
}

// renderCodeFence renders a fenced code block as a bordered/background column
// with one child element per source line. A single multiline WithText would
// collapse to one line under no-wrap and measure as height 1, so each line is
// its own element; blank lines render a space to keep height 1.
func (m *Markdown) renderCodeFence(b markdown.Block) *Element {
	opts := []Option{WithDirection(Column)}
	if m.theme.CodeBlockBorder != BorderNone {
		opts = append(opts, WithBorder(m.theme.CodeBlockBorder))
	}
	if !m.theme.CodeBlockBg.IsDefault() {
		opts = append(opts, WithBackground(NewStyle().Background(m.theme.CodeBlockBg)))
	}
	box := New(opts...)

	for _, line := range b.Lines {
		text := line
		if text == "" {
			text = " " // keep blank lines from collapsing to height 0
		}
		box.AddChild(New(
			WithText(text),
			WithWrap(false),
			WithTextStyle(m.theme.CodeBlockText),
		))
	}
	return box
}

// renderList renders a list and its items at the given nesting depth.
// contentWidth is the width available to the list (0 = auto/unknown). textStyle is
// the base style for item text (e.g. italic inside a blockquote).
func (m *Markdown) renderList(list markdown.Block, depth, contentWidth int, textStyle Style) *Element {
	col := New(WithDirection(Column))
	for i, item := range list.Children {
		marker := m.theme.BulletMarker
		if list.Ordered {
			marker = fmt.Sprintf("%d. ", i+1)
		}
		col.AddChild(m.renderListItem(item, marker, depth, contentWidth, textStyle))
	}
	return col
}

// renderListItem renders one item: an indented "marker + inline text" row,
// followed by any nested list rendered at depth+1.
func (m *Markdown) renderListItem(item markdown.Block, marker string, depth, contentWidth int, textStyle Style) *Element {
	itemCol := New(WithDirection(Column))

	markerText := strings.Repeat("  ", depth) + marker
	spans := m.inlineToSpans(item.Inline)

	rowOpts := []Option{WithDisplay(DisplayFlex), WithDirection(Row)}
	content := New(WithTextStyle(textStyle), WithRichText(spans...))
	if contentWidth > 0 {
		// Constrain content width so it wraps, and size the row to the wrapped
		// height: a Row sizes its height from children's intrinsic height, which
		// is 1 for rich text, so without an explicit height wrapped lines clip.
		cw := contentWidth - stringWidth(markerText)
		if cw < 1 {
			cw = 1
		}
		content = New(WithDirection(Column), WithWidth(cw))
		content.AddChild(New(WithTextStyle(textStyle), WithRichText(spans...)))
		h := content.HeightForWidth(cw)
		if h < 1 {
			h = 1
		}
		rowOpts = append(rowOpts, WithHeight(h))
	}

	row := New(rowOpts...)
	row.AddChild(New(WithTextStyle(textStyle), WithText(markerText), WithWrap(false)))
	row.AddChild(content)
	itemCol.AddChild(row)

	for _, child := range item.Children {
		if child.Kind == markdown.KindList {
			itemCol.AddChild(m.renderList(child, depth+1, contentWidth, textStyle))
		}
	}
	return itemCol
}

// renderBlockquote renders a recursive blockquote: a 1-wide glyph bar column
// beside the indented, recursively-rendered content. The bar's height matches
// the content height (measured at the available width; at auto width it assumes
// no wrapping).
func (m *Markdown) renderBlockquote(b markdown.Block, contentWidth int) *Element {
	childWidth := 0
	if contentWidth > 0 {
		childWidth = contentWidth - 2 // bar (1) + gap (1)
		if childWidth < 1 {
			childWidth = 1
		}
	}

	// Constrain the content width (when known) so paragraphs wrap to it; this
	// also makes the measured height below match what actually renders.
	contentOpts := []Option{WithDirection(Column)}
	if childWidth > 0 {
		contentOpts = append(contentOpts, WithWidth(childWidth))
	}
	content := New(contentOpts...)
	for _, child := range b.Children {
		if el := m.renderBlock(child, childWidth, m.theme.BlockquoteText); el != nil {
			content.AddChild(el)
		}
	}

	// Measure content height to size the bar.
	height := 0
	if childWidth > 0 {
		height = content.HeightForWidth(childWidth)
	} else {
		_, height = content.IntrinsicSize()
	}
	if height < 1 {
		height = 1
	}

	bar := New(WithDirection(Column), WithWidth(1))
	for i := 0; i < height; i++ {
		bar.AddChild(New(
			WithText(string(m.theme.BlockquoteBar)),
			WithTextStyle(m.theme.BlockquoteBarStyle),
		))
	}

	row := New(WithDisplay(DisplayFlex), WithDirection(Row), WithGap(1))
	row.AddChild(bar)
	row.AddChild(content)
	return row
}

// tableGrid holds the box-drawing runes for a full table grid in a given style:
// horizontal/vertical lines plus the nine junctions (top/middle/bottom row, each
// left/mid/right).
type tableGrid struct {
	h, v       rune
	tl, tm, tr rune
	ml, mm, mr rune
	bl, bm, br rune
}

func tableGridFor(b BorderStyle) tableGrid {
	switch b {
	case BorderDouble:
		return tableGrid{'═', '║', '╔', '╦', '╗', '╠', '╬', '╣', '╚', '╩', '╝'}
	case BorderThick:
		return tableGrid{'━', '┃', '┏', '┳', '┓', '┣', '╋', '┫', '┗', '┻', '┛'}
	case BorderSingle:
		return tableGrid{'─', '│', '┌', '┬', '┐', '├', '┼', '┤', '└', '┴', '┘'}
	default: // rounded (and the default theme)
		return tableGrid{'─', '│', '╭', '┬', '╮', '├', '┼', '┤', '╰', '┴', '╯'}
	}
}

// gridRule builds one horizontal rule spanning all columns (the top border, the
// header separator, or the bottom border), padding each column by one space on
// each side to match the content rows.
func gridRule(widths []int, h, left, mid, right rune) string {
	var sb strings.Builder
	sb.WriteRune(left)
	for i, w := range widths {
		sb.WriteString(strings.Repeat(string(h), w+2))
		if i == len(widths)-1 {
			sb.WriteRune(right)
		} else {
			sb.WriteRune(mid)
		}
	}
	return sb.String()
}

// renderTable renders a pipe table as a full grid: an outer box plus column
// separators and a rule under the header, in the theme.TableBorder style. Each
// cell keeps its inline (rich-text) styling. Column widths fit the widest cell.
func (m *Markdown) renderTable(b markdown.Block) *Element {
	if len(b.Rows) == 0 {
		return New(WithDirection(Column))
	}

	cols := 0
	for _, row := range b.Rows {
		if len(row) > cols {
			cols = len(row)
		}
	}
	if cols == 0 {
		return New(WithDirection(Column))
	}

	widths := make([]int, cols)
	for _, row := range b.Rows {
		for c := 0; c < cols && c < len(row); c++ {
			w := 0
			for _, in := range row[c].Inline {
				w += stringWidth(in.Text)
			}
			if w > widths[c] {
				widths[c] = w
			}
		}
	}
	for c := range widths {
		if widths[c] < 1 {
			widths[c] = 1
		}
	}

	g := tableGridFor(m.theme.TableBorder)
	gridStyle := NewStyle()
	rule := func(left, mid, right rune) *Element {
		return New(WithText(gridRule(widths, g.h, left, mid, right)), WithWrap(false), WithTextStyle(gridStyle))
	}

	table := New(WithDirection(Column))
	table.AddChild(rule(g.tl, g.tm, g.tr))
	for r, row := range b.Rows {
		table.AddChild(m.renderTableRow(row, widths, g, r == 0))
		// A rule between every row (after the header and between body rows).
		if r < len(b.Rows)-1 {
			table.AddChild(rule(g.ml, g.mm, g.mr))
		}
	}
	table.AddChild(rule(g.bl, g.bm, g.br))
	return table
}

// renderTableRow builds one content row: vertical separators around each
// fixed-width cell, with one space of padding on each side of the content.
func (m *Markdown) renderTableRow(cells []markdown.TableCell, widths []int, g tableGrid, header bool) *Element {
	gridStyle := NewStyle()
	bar := func(text string) *Element {
		return New(WithText(text), WithWrap(false), WithTextStyle(gridStyle))
	}

	row := New(WithDisplay(DisplayFlex), WithDirection(Row), WithHeight(1))
	row.AddChild(bar(string(g.v) + " "))
	for c := range widths {
		var spans []TextSpan
		if c < len(cells) {
			spans = m.inlineToSpans(cells[c].Inline)
		}
		cellOpts := []Option{WithWidth(widths[c]), WithWrap(false), WithRichText(spans...)}
		if header {
			cellOpts = append(cellOpts, WithTextStyle(m.theme.TableHeader))
		}
		row.AddChild(New(cellOpts...))
		if c < len(widths)-1 {
			row.AddChild(bar(" " + string(g.v) + " "))
		}
	}
	row.AddChild(bar(" " + string(g.v)))
	return row
}

// inlineToSpans converts parser inline runs into themed TextSpans. The element's
// textStyle supplies the base; each span layers only its inline-specific style.
func (m *Markdown) inlineToSpans(inls []markdown.Inline) []TextSpan {
	spans := make([]TextSpan, 0, len(inls))
	for _, in := range inls {
		st := Style{}
		if in.Bold {
			st = mergeSpanStyle(st, m.theme.Bold)
		}
		if in.Italic {
			st = mergeSpanStyle(st, m.theme.Italic)
		}
		if in.Code {
			st = mergeSpanStyle(st, m.theme.CodeSpan)
		}
		if in.Link != "" {
			st = mergeSpanStyle(st, m.theme.Link)
		}
		spans = append(spans, TextSpan{Text: in.Text, Style: st, Link: in.Link})
	}
	return spans
}
