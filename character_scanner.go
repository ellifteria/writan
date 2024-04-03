package writan

type CharScanner struct {
	validTokens map[string]string
}

func makeCharScanner() CharScanner {
	return CharScanner{
		map[string]string{
			"\n": NEWLINE_TOKEN,
		},
	}
}

func (s CharScanner) fromString(plainMarkdown string) Token {
	char := string(plainMarkdown[0])

	scannerToken := s.validTokens[char]
	if scannerToken == "" {
		return makeNullToken()
	}

	return makeToken(scannerToken, char)
}
