package tui

import "strings"


// WrapText is a public wrapper around wrapText for use by external packages.
func WrapText(text string, maxWidth int) []string {
	return wrapText(text, maxWidth)
}
// wrapText wraps text to fit within maxWidth terminal cells using word boundaries.
// It breaks at spaces, falling back to mid-character breaks when a single word
// exceeds maxWidth. Existing newlines in the text are preserved.
func wrapText(text string, maxWidth int) []string {
	if maxWidth < 1 {
		return []string{""}
	}
	if text == "" {
		return []string{""}
	}

	var result []string
	for _, paragraph := range strings.Split(text, "\n") {
		result = append(result, wrapParagraph(paragraph, maxWidth)...)
	}
	return result
}

// wrapParagraph wraps a single paragraph (no newlines) to maxWidth.
func wrapParagraph(text string, maxWidth int) []string {
	if text == "" {
		return []string{""}
	}
	text = strings.ReplaceAll(text, "\t", " ")

	var lines []string
	var buf strings.Builder
	lineWidth := 0
	i := 0
	textLen := len(text)
	lastWasSpace := false

	for i < textLen {
		if text[i] == ' ' {
			lastWasSpace = true
			spaces := 1
			for i+spaces < textLen && text[i+spaces] == ' ' {
				spaces++
			}
			for s := 0; s < spaces; s++ {
				if lineWidth >= maxWidth {
					lines = append(lines, buf.String())
					buf.Reset()
					lineWidth = 0
				}
				buf.WriteByte(' ')
				lineWidth++
			}
			i += spaces
		} else {
			wordStart := i
			for i < textLen && text[i] != ' ' {
				i++
			}
			word := text[wordStart:i]
			ww := stringWidth(word)

			if ww > maxWidth && lineWidth > 0 {
				lines = append(lines, buf.String())
				buf.Reset()
				lineWidth = 0
				lastWasSpace = false
			}
			if ww > maxWidth {
				for _, r := range word {
					rw := RuneWidth(r)
					if lineWidth+rw > maxWidth && lineWidth > 0 {
						lines = append(lines, buf.String())
						buf.Reset()
						lineWidth = 0
					}
					buf.WriteRune(r)
					lineWidth += rw
				}
			} else if lineWidth == 0 {
				buf.WriteString(word)
				lineWidth = ww
			} else if lastWasSpace {
				if lineWidth+ww <= maxWidth {
					buf.WriteString(word)
					lineWidth += ww
				} else {
					lines = append(lines, buf.String())
					buf.Reset()
					buf.WriteString(word)
					lineWidth = ww
				}
			} else if lineWidth+1+ww <= maxWidth {
				buf.WriteByte(' ')
				buf.WriteString(word)
				lineWidth += 1 + ww
			} else {
				lines = append(lines, buf.String())
				buf.Reset()
				buf.WriteString(word)
				lineWidth = ww
			}
			lastWasSpace = false
		}
	}

	lines = append(lines, buf.String())
	for i := 0; i < len(lines)-1; i++ {
		lines[i] = strings.TrimRight(lines[i], " ")
	}
	return lines
}
