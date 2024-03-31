package writan

import "fmt"

type TokenType int

const (
	NULL_TOKEN TokenType = iota
	QUOTE_TOKEN
	OPEN_BRACKET_TOKEN
	CLOSE_BRACKET_TOKEN
	OPEN_BRACE_TOKEN
	CLOSE_BRACE_TOKEN
	OPEN_PAREN_TOKEN
	CLOSE_PAREN_TOKEN
	OPEN_ARROW_TOKEN
	CLOSE_ARROW_TOKEN
	UNDERSCORE_TOKEN
	STAR_TOKEN
	NEWLINE_TOKEN
	EXCLAMATION_TOKEN
	TEXT_TOKEN
	EOF_TOKEN
)

func (t TokenType) String() string {
	return [...]string{
		"NULL",
		"QUOTE",
		"OPEN_BRACKET",
		"CLOSE_BRACKET",
		"OPEN_BRACE",
		"CLOSE_BRACE",
		"OPEN_PAREN",
		"CLOSE_PAREN",
		"OPEN_ARROW",
		"CLOSE_ARROW",
		"UNDERSCORE",
		"STAR",
		"NEWLINE",
		"EXCLAMATION",
		"TEXT",
		"EOF",
	}[t]
}

type Token struct {
	tokenType  TokenType
	tokenValue string
	next       *Token
}

func makeToken(tokenType TokenType, tokenValue string) Token {
	if tokenType == NULL_TOKEN || tokenValue == "" {
		makeNullToken()
	}

	return Token{tokenType: tokenType, tokenValue: tokenValue, next: nil}
}

func makeNullToken() Token {
	return Token{tokenType: NULL_TOKEN, tokenValue: "", next: nil}
}

func makeEOFToken() Token {
	return Token{tokenType: EOF_TOKEN, tokenValue: "", next: nil}
}

func (t Token) isNull() bool {
	return t.tokenType == NULL_TOKEN
}

func (t Token) isPresent() bool {
	return !t.isNull()
}

func (t Token) length() int {
	if t.isNull() {
		return 0
	}

	return len(t.tokenValue)
}

func (t Token) toString() string {
	var nextTokenString string

	if t.next == nil {
		nextTokenString = "none"
	} else {
		nextToken := *t.next
		nextTokenString = nextToken.toString()
	}
	return fmt.Sprintf("<type: %s, value: %s next: %s>", t.tokenType.String(), t.tokenValue, nextTokenString)
}

func (t Token) valuesToString() string {
	var nextTokenString string

	if t.next == nil {
		nextTokenString = ""
	} else {
		nextToken := *t.next
		nextTokenString = nextToken.valuesToString()
	}
	return fmt.Sprintf("%s%s", t.tokenValue, nextTokenString)
}

func (t Token) matchesTypes(tokenTypes []TokenType) bool {
	if len(tokenTypes) == 0 {
		return true
	}

	firstTokenType := tokenTypes[0]
	if t.tokenType != firstTokenType {
		return false
	}

	if len(tokenTypes) == 1 {
		return true
	}

	if t.next == nil {
		return false
	}

	remainingTokenTypes := tokenTypes[1:]

	return t.next.matchesTypes(remainingTokenTypes)
}

func (t Token) isA(tokenType TokenType) bool {
	return t.tokenType == tokenType
}

func (t Token) isAOr(tokenTypes []TokenType) bool {
	for _, tokenType := range tokenTypes {
		if t.isA(tokenType) {
			return true
		}
	}

	return false
}
