package writan

type CmdScanner struct {
	cmdToken byte
	cmdChars map[byte]string
}

func makeCmdScanner() CmdScanner {
	return CmdScanner{
		cmdToken: '@',
		cmdChars: map[byte]string{
			'c': CODE_INLINE_TOKEN,
			'C': CODE_BLOCK_TOKEN,
			'>': QUOTE_BLOCK_TOKEN,
			'*': BOLD_TOKEN,
			'_': ITALICS_TOKEN,
			'@': AT_TOKEN,
		},
	}
}

func (s CmdScanner) fromString(plainMarkdown string) Token {
	initialChar := plainMarkdown[0]
	if initialChar != s.cmdToken || len(plainMarkdown) == 1 {
		return makeNullToken()
	}

	secondChar := plainMarkdown[1]

	scannerToken := s.cmdChars[secondChar]
	if scannerToken == "" {
		return makeNullToken()
	}

	return makeToken(scannerToken, plainMarkdown[0:2])
}
