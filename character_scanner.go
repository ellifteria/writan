package writan

type CharScanner struct {
	validTokens map[string]TokenType
}

func makeCharScanner() CharScanner {
	return CharScanner{
		map[string]TokenType{
			"\n": NEWLINE_TOKEN,
		},
	}
}

func (s CharScanner) fromString(plainMarkdown string) Token {
	char := string(plainMarkdown[0])
	return makeToken(s.validTokens[char], char)
}
