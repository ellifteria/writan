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
		chars := plainMarkdown[index : index+2]
		for _, scanner := range s.scanners {
			if !scanner.fromString(chars).isNull() {
				return makeToken(s.textTokenType, text)
			}
		}
		text += string(chars[0])
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
