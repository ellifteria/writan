package writan

import "strings"

type SingleTokenParser struct {
	tokenToParse string
	nodeToSet    string
	baseParser   *BaseParser
}

func makeSingleTokenParser(tokenToParse string, nodeToSet string) SingleTokenParser {
	return SingleTokenParser{tokenToParse: tokenToParse, nodeToSet: nodeToSet}
}

func (p SingleTokenParser) match(token Token) (Node, *Token) {
	if token.tokenType != p.tokenToParse {
		return makeNullNode(), nil
	}

	tokenValue := token.tokenValue
	value := strings.Join(strings.Split(tokenValue[1:len(tokenValue)-1], " ")[1:], " ")

	nextToken := token.next
	textNode := makeNode(p.nodeToSet, value)

	return textNode, nextToken
}
