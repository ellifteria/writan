package writan

import (
	"strings"
)

type TextScanner struct {
	CharScanner CharScanner
}

func makeTextScanner() TextScanner {
	return TextScanner{makeCharScanner()}
}

func (s TextScanner) fromString(plainMarkdown string) Token {
	text := ""
	for _, char := range strings.Split(plainMarkdown, "") {
		if !s.CharScanner.fromString(char).isNull() {
			break
		}
		text += char
	}
	return makeToken(TEXT_TOKEN, text)
}
