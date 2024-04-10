package writan

import "strings"

type CmdScanner struct {
	cmdToken byte
	endToken byte
	cmdChars map[string]string
}

func makeCmdScanner() CmdScanner {
	return CmdScanner{
		cmdToken: '@',
		endToken: ';',
		cmdChars: map[string]string{
			"c": CODE_INLINE_TOKEN,
			"C": CODE_BLOCK_TOKEN,
			">": QUOTE_BLOCK_TOKEN,
			"*": BOLD_TOKEN,
			"_": ITALICS_TOKEN,
			"@": AT_TOKEN,
			"a": LINK_TOKEN,
		},
	}
}

func (s CmdScanner) fromString(plainMarkdown string) Token {
	initialChar := plainMarkdown[0]
	if initialChar != s.cmdToken || len(plainMarkdown) == 1 {
		return makeNullToken()
	}

	endIndex := 1

	for plainMarkdown[endIndex] != s.endToken {
		endIndex += 1
	}

	commandString := strings.Split(plainMarkdown[1:endIndex], " ")[0]

	commandToken := s.cmdChars[commandString]
	if commandToken == "" {
		return makeNullToken()
	}

	totalString := plainMarkdown[0 : endIndex+1]

	return makeToken(commandToken, totalString)
}
