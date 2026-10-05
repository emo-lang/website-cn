// Package emocode renders Emo source snippets as syntax-highlighted HTML
// for the marketing pages. The tokenizer is deliberately small: it colors
// comments, strings, numbers, keywords, types, and call sites — enough for
// page-sized samples without pulling in a full grammar.
package emocode

import (
	"html"
	"strings"
	"unicode/utf8"
)

var keywords = map[string]bool{
	"class": true, "interface": true, "enum": true, "def": true,
	"const": true, "var": true, "if": true, "else": true, "case": true,
	"when": true, "return": true, "raise": true, "do": true, "receive": true,
	"halt": true, "self": true, "foreign": true, "require": true,
	"package": true, "deps": true, "targets": true, "init": true,
	"true": true, "false": true, "emo": true,
}

// Highlight tokenizes src and returns HTML with tokens wrapped in
// <span class="tok-*">. All source text is HTML-escaped per token.
func Highlight(src string) string {
	var out strings.Builder
	for _, line := range strings.Split(src, "\n") {
		highlightLine(line, &out)
		out.WriteByte('\n')
	}
	return strings.TrimRight(out.String(), "\n")
}

func highlightLine(line string, out *strings.Builder) {
	i := 0
	n := len(line)
	emit := func(cls, text string) {
		if cls == "" {
			out.WriteString(html.EscapeString(text))
			return
		}
		out.WriteString(`<span class="` + cls + `">` + html.EscapeString(text) + `</span>`)
	}

	for i < n {
		c := line[i]

		// comments: // to end of line
		if c == '/' && i+1 < n && line[i+1] == '/' {
			emit("tok-c", line[i:])
			return
		}

		// comments: # to end of line
		if c == '#' {
			emit("tok-c", line[i:])
			return
		}

		// strings and char literals
		if c == '"' || c == '\'' {
			j := i + 1
			for j < n {
				if line[j] == '\\' && j+1 < n {
					j += 2
					continue
				}
				if line[j] == c {
					j++
					break
				}
				j++
			}
			emit("tok-s", line[i:j])
			i = j
			continue
		}

		// numbers
		if isDigit(c) {
			j := i
			for j < n && (isDigit(line[j]) || line[j] == '.') {
				j++
			}
			emit("tok-n", line[i:j])
			i = j
			continue
		}

		// identifiers / keywords / types / call sites
		if isIdentStart(c) {
			j := i
			for j < n && isIdentPart(line[j]) {
				j++
			}
			if j < n && line[j] == '?' {
				j++
			}
			word := line[i:j]
			switch {
			case keywords[word]:
				emit("tok-k", word)
			case word[0] >= 'A' && word[0] <= 'Z':
				emit("tok-t", word)
			case j < n && line[j] == '(':
				emit("tok-f", word)
			default:
				emit("", word)
			}
			i = j
			continue
		}

		// operators that deserve emphasis
		if c == '<' && i+1 < n && line[i+1] == '-' {
			emit("tok-o", "<-")
			i += 2
			continue
		}
		if c == '-' && i+1 < n && line[i+1] == '>' {
			emit("tok-o", "->")
			i += 2
			continue
		}

		// multi-byte UTF-8: emit the whole rune, not one byte at a time
		if c >= utf8.RuneSelf {
			r, size := utf8.DecodeRuneInString(line[i:])
			emit("", string(r))
			i += size
			continue
		}

		emit("", string(c))
		i++
	}
}

func isDigit(c byte) bool { return c >= '0' && c <= '9' }

func isIdentStart(c byte) bool {
	return c == '_' || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isIdentPart(c byte) bool { return isIdentStart(c) || isDigit(c) }
