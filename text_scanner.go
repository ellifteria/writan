package writan

type TextScanner struct {
	scanners      []Scanner
	textTokenType string
}

func makeTextScanner() TextScanner {
	return TextScanner{[]Scanner{makeCmdScanner(), makeCharScanner()}, TEXT_TOKEN}
}

func (s TextScanner) fromString(plainMarkdown string) Token {
	text := ""
	for index := 0; index < len(plainMarkdown)-1; index++ {
		for _, scanner := range s.scanners {
			if !scanner.fromString(plainMarkdown[index:]).isNull() {
				return makeToken(s.textTokenType, text)
			}
		}
		text += string(plainMarkdown[index])
	}

	char := string(plainMarkdown[len(plainMarkdown)-1])
	for _, scanner := range s.scanners {
		if !scanner.fromString(char).isNull() {
			break
		}
	}
	text += string(char[0])

	return makeToken(s.textTokenType, text)
}
