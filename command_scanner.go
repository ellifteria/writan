package writan

type CmdScanner struct {
	cmdToken byte
	cmdChars map[byte]TokenType
}

func makeCmdScanner() CmdScanner {
	return CmdScanner{
		cmdToken: '@',
		cmdChars: map[byte]TokenType{
			'c': CODE_INLINE_TOKEN,
			'C': CODE_BLOCK_TOKEN,
			'>': QUOTE_BLOCK_TOKEN,
			'*': BOLD_TOKEN,
			'_': ITALICS_TOKEN,
			'/': CLOSE_TOKEN,
			'@': ESCAPE_TOKEN,
		},
	}
}

func (s CmdScanner) fromString(plainMarkdown string) Token {
	initialChar := plainMarkdown[0]
	if initialChar != s.cmdToken || len(plainMarkdown) == 1 {
		return makeNullToken()
	}
	secondChar := plainMarkdown[1]
	return makeToken(s.cmdChars[secondChar], plainMarkdown[0:2])
}
