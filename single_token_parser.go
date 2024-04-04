package writan

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

	nextToken := token.next
	textNode := makeNode(p.nodeToSet, token.tokenValue)

	return textNode, nextToken
}
