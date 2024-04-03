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

// func (t Token) isPresent() bool {
// 	return !t.isNull()
// }

// func (t Token) length() int {
// 	if t.isNull() {
// 		return 0
// 	}

// 	return len(t.tokenValue)
// }

func (t Token) toString() string {
	var nextTokenString string

	if t.next == nil {
		nextTokenString = "none"
	} else {
		nextToken := *t.next
		nextTokenString = nextToken.toString()
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

// func (t Token) matchesTypes(tokenTypes []string) bool {
// 	if len(tokenTypes) == 0 {
// 		return true
// 	}

// 	firstTokenType := tokenTypes[0]
// 	if t.tokenType != firstTokenType {
// 		return false
// 	}

// 	if len(tokenTypes) == 1 {
// 		return true
// 	}

// 	if t.next == nil {
// 		return false
// 	}

// 	remainingTokenTypes := tokenTypes[1:]

// 	return t.next.matchesTypes(remainingTokenTypes)
// }

// func (t Token) isA(tokenType string) bool {
// 	return t.tokenType == tokenType
// // }

// func (t Token) isAOr(tokenTypes []string) bool {
// 	for _, tokenType := range tokenTypes {
// 		if t.isA(tokenType) {
// 			return true
// 		}
// 	}

// 	return false
// }

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

// func (t Token) copy() Token {
// 	if t.next == nil {
// 		return makeToken(t.tokenType, t.tokenValue)
// 	}

// 	token := makeToken(t.tokenType, t.tokenValue)

// 	nextToken := t.next.copy()
// 	token.next = &nextToken

// 	return token
// }

func (t Token) copyUntil(untilTokenType string) (Token, *Token) {
	if t.tokenType == untilTokenType {
		token := makeNullToken()
		nextTokenPointer := token.next

		if t.next == nil {
			childOfFinalToken := makeNullToken()
			nextTokenPointer = &childOfFinalToken
		}

		return token, nextTokenPointer
	}

	token := makeToken(t.tokenType, t.tokenValue)
	var nextTokenPointer *Token

	if t.next == nil {
		childOfFinalToken := makeNullToken()
		nextTokenPointer = &childOfFinalToken
	} else {
		if t.next.tokenType != untilTokenType {
			var nextToken Token
			nextToken, nextTokenPointer = t.next.copyUntil(untilTokenType)
			token.next = &nextToken
		} else {
			nextTokenPointer = t.next.next
		}
	}

	return token, nextTokenPointer
}
