package writan

import "fmt"

const (
	NULL_TOKEN        = "NULL_TOKEN"
	EOF_TOKEN         = "EOF_TOKEN"
	TEXT_TOKEN        = "TEXT_TOKEN"
	CODE_INLINE_TOKEN = "CODE_INLINE_TOKEN"
	CODE_BLOCK_TOKEN  = "CODE_BLOCK_TOKEN"
	QUOTE_BLOCK_TOKEN = "QUOTE_BLOCK_TOKEN"
	BOLD_TOKEN        = "BOLD_TOKEN"
	ITALICS_TOKEN     = "ITALICS_TOKEN"
	NEWLINE_TOKEN     = "NEWLINE_TOKEN"
	AT_TOKEN          = "AT_TOKEN"
	LINK_TOKEN        = "LINK_TOKEN"
)

var VALID_TOKEN_TYPES []string = []string{
	NULL_TOKEN,
	EOF_TOKEN,
	TEXT_TOKEN,
	CODE_INLINE_TOKEN,
	CODE_BLOCK_TOKEN,
	QUOTE_BLOCK_TOKEN,
	BOLD_TOKEN,
	ITALICS_TOKEN,
	NEWLINE_TOKEN,
	AT_TOKEN,
	LINK_TOKEN,
}

type Token struct {
	tokenType  string
	tokenValue string
	next       *Token
}

func makeToken(tokenType string, tokenValue string) Token {
	if tokenType == NULL_TOKEN || tokenValue == "" {
		return makeNullToken()
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

func (t Token) ToString() string {
	var nextTokenString string

	if t.next == nil {
		nextTokenString = "none"
	} else {
		nextToken := *t.next
		nextTokenString = nextToken.ToString()
	}
	return fmt.Sprintf("<type: %s, value: \"%s\" next: %s>", t.tokenType, t.tokenValue, nextTokenString)
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

func (t Token) matchesOpenClose(openingTokenType string, closingTokenType string) bool {
	if t.tokenType != openingTokenType {
		return false
	}
	currToken := t

	for currToken.next != nil {
		if currToken.next.tokenType == closingTokenType {
			return true
		}

		currToken = *currToken.next
	}

	return false
}
