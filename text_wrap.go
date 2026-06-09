package tui

import (
	"strings"
	"unicode"
)

// isDefaultEmojiPresentation returns true for BMP characters that are
// Emoji_Presentation=Yes per Unicode — they render as 2-cell emoji even
// without a following VS16. Characters with Emoji=Yes but NOT
// Emoji_Presentation=Yes render as 1 cell (text presentation) and need
// VS16 to display as 2-cell emoji.
func isDefaultEmojiPresentation(r rune) bool {
	// Specific ranges and points from Unicode Emoji v15.0 Emoji_Presentation=Yes
	switch {
	case r == 0x231A || r == 0x231B: // ⌚⌛ watch, hourglass
		return true
	case r == 0x23E9 || r == 0x23EA || r == 0x23EB || r == 0x23EC: // ⏩⏪⏫⏬ fast-forward etc
		return true
	case r == 0x23F0 || r == 0x23F3: // ⏰⏳ alarm clock, hourglass
		return true
	case r == 0x25FD || r == 0x25FE: // ◽◾ medium squares
		return true
	case r >= 0x2614 && r <= 0x2615: // ☔☕
		return true
	case r >= 0x2648 && r <= 0x2653: // ♈-♓ zodiac
		return true
	case r == 0x267F: // ♿ wheelchair
		return true
	case r == 0x2693: // ⚓ anchor
		return true
	case r == 0x26A1: // ⚡ high voltage
		return true
	case r >= 0x26AA && r <= 0x26AB: // ⚪⚫ white/black circle
		return true
	case r >= 0x26BD && r <= 0x26BE: // ⚽⚾ soccer/baseball
		return true
	case r >= 0x26C4 && r <= 0x26C5: // ⛄⛅ snowman, sun behind cloud
		return true
	case r == 0x26CE: // ⛎ Ophiuchus
		return true
	case r == 0x26D4: // ⛔ no entry
		return true
	case r == 0x26EA: // ⛪ church
		return true
	case r >= 0x26F2 && r <= 0x26F3: // ⛲⛳ fountain, golf
		return true
	case r == 0x26F5: // ⛵ sailboat
		return true
	case r == 0x26FA: // ⛺ tent
		return true
	case r == 0x26FD: // ⛽ fuel pump
		return true
	case r == 0x2702: // ✂ scissors
		return true
	case r == 0x2705: // ✅ check mark
		return true
	case r >= 0x2708 && r <= 0x270D: // ✈-✍ airplane to writing hand
		return true
	case r == 0x270F: // ✏ pencil
		return true
	case r == 0x2712: // ✒ black nib
		return true
		return true
	case r == 0x2716: // ✖ multiplication
		return true
	case r == 0x271D: // ✝ latin cross
		return true
	case r == 0x2721: // ✡ star of David
		return true
	case r == 0x2728: // ✨ sparkles
		return true
	case r >= 0x2733 && r <= 0x2734: // ✳✴ asterisk
		return true
	case r == 0x2744: // ❄ snowflake
		return true
	case r == 0x2747: // ❇ sparkle
		return true
	case r == 0x274C: // ❌ cross mark
		return true
	case r == 0x274E: // ❎ cross mark
		return true
	case r >= 0x2753 && r <= 0x2755: // ❓❔❕ question/exclamation
		return true
	case r == 0x2757: // ❗ exclamation
		return true
	case r >= 0x2763 && r <= 0x2764: // ❣❤ heart
		return true
	case r >= 0x2795 && r <= 0x2797: // ➕➖➗ math
		return true
	case r == 0x27A1: // ➡ right arrow
		return true
	case r == 0x27B0: // ➰ curly loop
		return true
	case r == 0x27BF: // ➿ double curly loop
		return true
	case r >= 0x2B05 && r <= 0x2B07: // ⬅⬆⬇ arrows
		return true
	case r == 0x2B1B || r == 0x2B1C: // ⬛⬜ large squares
		return true
	case r == 0x2B50: // ⭐ star
		return true
	case r == 0x2B55: // ⭕ hollow red circle
		return true
	case r == 0x3030: // 〰 wavy dash
		return true
	case r == 0x303D: // 〽 part alternation mark
		return true
	case r == 0x3297: // ㊗ congratulations
		return true
	case r == 0x3299: // ㊙ secret
		return true
	}
	return false
}

// visualWidth returns the visual cell width of a rune, returning 0 for
// zero-width characters (variation selectors, combining marks, ZWJ, etc.)
// and 1 for BMP emoji that are text-presentation by default.
func visualWidth(r rune) int {
	if r == 0xFE0E || r == 0xFE0F || r == 0x200D || unicode.In(r, unicode.Mn, unicode.Me, unicode.Cf) {
		return 0
	}
	// For BMP emoji-wide characters, only return 2 for Emoji_Presentation=Yes.
	// Characters that need VS16 to be emoji render as 1 cell (text presentation).
	if r >= 0x2190 && r <= 0x2BFF && !isDefaultEmojiPresentation(r) {
		return 1
	}
	return RuneWidth(r)
}

// visualStringWidth returns the total visual cell width of a string.
func visualStringWidth(s string) int {
	w := 0
	for _, r := range s {
		w += visualWidth(r)
	}
	return w
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

// WrapText is a public wrapper around wrapText for use by external packages.
func WrapText(text string, maxWidth int) []string {
	return wrapText(text, maxWidth)
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
			ww := visualStringWidth(word)

			if ww > maxWidth && lineWidth > 0 {
				lines = append(lines, buf.String())
				buf.Reset()
				lineWidth = 0
				lastWasSpace = false
			}
			if ww > maxWidth {
				for _, r := range word {
					rw := visualWidth(r)
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
