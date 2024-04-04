package writan

import (
	"errors"
	"fmt"
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

func (t Tokenizer) Test(plainMarkdown string) {
	token := t.Tokenize(plainMarkdown)
	fmt.Println(token.toString())
	fmt.Println(token.valuesToString())

	if token.matchesOpenClose(BOLD_TOKEN, BOLD_TOKEN) {
		insideToken, nextToken := token.next.copyUntil(BOLD_TOKEN)
		fmt.Println(insideToken.toString())
		fmt.Println(nextToken.toString())
		if insideToken.matchesOpenClose(ITALICS_TOKEN, ITALICS_TOKEN) {
			insideInsideToken, nextNextToken := insideToken.next.copyUntil(ITALICS_TOKEN)
			fmt.Println(insideInsideToken.toString())
			if nextNextToken != nil {
				fmt.Println(nextNextToken.toString())
			}
		}
	}
	fmt.Println(t.Tokenize("@*@_hello@_@*").matchesOpenClose(BOLD_TOKEN, BOLD_TOKEN))
}

func (t Tokenizer) TestBTP(plainMarkdown string) {
	token := t.Tokenize(plainMarkdown)
	fmt.Println(token.toString())
	// baseParser := makeBaseParser()
	// btp := makeTagMatchingTextParser(BOLD_TOKEN, BOLD_TOKEN, BOLD_TEXT_NODE)
	// itp := makeTagMatchingTextParser(ITALICS_TOKEN, ITALICS_TOKEN, ITALICIZED_TEXT_NODE)
	// cip := makeTagMatchingTextParser(CODE_INLINE_TOKEN, CODE_INLINE_TOKEN, CODE_INLINE_NODE)
	// cbp := makeTagMatchingTextParser(CODE_BLOCK_TOKEN, CODE_BLOCK_TOKEN, CODE_BLOCK_NODE)
	// qbp := makeTagMatchingTextParser(QUOTE_BLOCK_TOKEN, QUOTE_BLOCK_TOKEN, QUOTE_BLOCK_NODE)
	// tp := makeTextParser()
	// btp.baseParser = &baseParser
	// itp.baseParser = &baseParser
	// tp.baseParser = &baseParser
	// cbp.baseParser = &baseParser
	// cip.baseParser = &baseParser
	// qbp.baseParser = &baseParser
	// baseParser.parsers = []Parser{&cbp, &cip, &qbp, &btp, &itp, &tp}

	_, parser := MakeParser()

	node, _ := parser.match(token)
	fmt.Println(node.toString())
}
