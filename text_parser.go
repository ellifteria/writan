package writan

type TextParser struct {
	baseParser *BaseParser
}

func makeTextParser() TextParser {
	return TextParser{}
}

func (p TextParser) match(token Token) (Node, *Token) {
	nextToken := token.next

	textNode := makeNode(TEXT_NODE, token.tokenValue)

	return textNode, nextToken
}
