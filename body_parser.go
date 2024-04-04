package writan

type BodyParser struct {
	paragraphParser *ParagraphParser
	newlineParser   *SingleTokenParser
}

func makeBodyParser(paragraphParser *ParagraphParser, newlineParser *SingleTokenParser) BodyParser {
	return BodyParser{paragraphParser: paragraphParser, newlineParser: newlineParser}
}

func (p BodyParser) match(token Token) (Node, *Token) {
	if token.tokenType == EOF_TOKEN || token.tokenType == NULL_TOKEN {
		return makeNullNode(), nil
	}

	currTokenPtr := &token

	bodyNode := makeNode(BODY_NODE, "")

	for currTokenPtr != nil && currTokenPtr.tokenType != EOF_TOKEN {
		var childNode Node
		var tmpCurrTokenPtr *Token
		childNode, tmpCurrTokenPtr = p.newlineParser.match(*currTokenPtr)
		if childNode.nodeType != NULL_NODE {
			bodyNode.children = append(bodyNode.children, &childNode)
		} else {
			childNode, tmpCurrTokenPtr = p.paragraphParser.match(*currTokenPtr)
			bodyNode.children = append(bodyNode.children, &childNode)
		}
		currTokenPtr = tmpCurrTokenPtr
	}

	return bodyNode, nil
}

func (p BodyParser) Match(token Token) Node {
	node, _ := p.match(token)
	return node
}
