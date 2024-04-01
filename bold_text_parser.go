package writan

type BoldTextParser struct {
	baseParser BaseParser
}

func (p BoldTextParser) match(token Token) (Node, *Token) {
	if !(token.matchesOpenClose(BOLD_TOKEN, BOLD_TOKEN)) {
		return makeNullNode(), nil
	}

	insideToken, nextToken := token.next.copyUntil(BOLD_TOKEN)
	interiorNode, _ := p.baseParser.match(insideToken)

	node := makeNode(BOLD_TEXT_NODE, "TEST", 1)
	node.children = append(node.children, &interiorNode)

	return node, nextToken
}
