package writan

type CharScanner struct {
	validTokens map[byte]string
}

func makeCharScanner() CharScanner {
	return CharScanner{
		map[byte]string{
			'\n': NEWLINE_TOKEN,
		},
	}
}

func (s CharScanner) fromString(plainMarkdown string) Token {
	char := plainMarkdown[0]

	scannerToken := s.validTokens[char]
	if scannerToken == "" {
		return makeNullToken()
	}

	return makeToken(scannerToken, string(char))
}
