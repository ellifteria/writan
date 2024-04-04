package writan

import (
	"errors"
	"log"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Tokenizer struct {
	scanners []Scanner
}

func MakeDefaultTokenizer() Tokenizer {
	return Tokenizer{[]Scanner{makeCmdScanner(), makeCharScanner(), makeTextScanner()}}
}

func (t Tokenizer) Tokenize(plainMarkdown string) Token {
	if plainMarkdown == "" {
		return makeEOFToken()
	}

	token, err := t.scanOneToken(plainMarkdown)
	check(err)
	remainingPlainMarkdown := strings.TrimPrefix(plainMarkdown, token.tokenValue)
	nextToken := t.Tokenize(remainingPlainMarkdown)
	token.next = &nextToken

	return token
}

func (t Tokenizer) scanOneToken(plainMarkdown string) (Token, error) {
	for _, scanner := range t.scanners {
		token := scanner.fromString(plainMarkdown)
		if !token.isNull() {
			return token, nil
		}
	}

	return makeNullToken(), errors.New("token not found")
}
