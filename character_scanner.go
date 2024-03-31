package writan

type CharScanner struct {
	validTokens map[string]TokenType
}

func makeCharScanner() CharScanner {
	return CharScanner{
		map[string]TokenType{
			"`":  QUOTE_TOKEN,
			"[":  OPEN_BRACKET_TOKEN,
			"]":  CLOSE_BRACKET_TOKEN,
			"{":  OPEN_BRACE_TOKEN,
			"}":  CLOSE_BRACE_TOKEN,
			"(":  OPEN_PAREN_TOKEN,
			")":  CLOSE_PAREN_TOKEN,
			"<":  OPEN_ARROW_TOKEN,
			">":  CLOSE_ARROW_TOKEN,
			"_":  UNDERSCORE_TOKEN,
			"*":  STAR_TOKEN,
			"\n": NEWLINE_TOKEN,
			"!":  EXCLAMATION_TOKEN,
		},
	}
}

func (s CharScanner) fromString(plainMarkdown string) Token {
	char := string(plainMarkdown[0])
	return makeToken(s.validTokens[char], char)
}
