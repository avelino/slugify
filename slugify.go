package slugify

import (
	"fmt"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

var SKIP = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk,
	unicode.Lm,
}

var SAFE = []*unicode.RangeTable{
	unicode.Letter,
	unicode.Number,
}

// Slugify a string. The result will only contain lowercase letters,
// digits and dashes. It will not begin or end with a dash, and it
// will not contain runs of multiple dashes.
//
// It is NOT forced into being ASCII, but may contain any Unicode
// characters, with the above restrictions.
func Slugify(text string) string {
	buf := make([]rune, 0, len(text))
	dash := false
	for _, r := range norm.NFKD.String(text) {
		switch {
		case unicode.IsOneOf(SAFE, r):
			buf = append(buf, unicode.ToLower(r))
			dash = true
		case unicode.IsOneOf(SKIP, r):
		case dash:
			buf = append(buf, '-')
			dash = false
		}
	}
	if i := len(buf) - 1; i >= 0 && buf[i] == '-' {
		buf = buf[:i]
	}
	return string(buf)
}

// Slugifyf slugfy a formated string
func Slugifyf(format string, a ...interface{}) string {
	return Slugify(fmt.Sprintf(format, a...))
}
