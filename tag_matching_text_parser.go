package writan

type TagMatchingTextParser struct {
	baseParser          *BaseParser
	openingTagTokenType string
	closingTagTokenType string
	nodeType            string
}

func makeTagMatchingTextParser(opener string, closer string, nodeType string) TagMatchingTextParser {
	return TagMatchingTextParser{nil, opener, closer, nodeType}
}

func (p TagMatchingTextParser) match(token Token) (Node, *Token) {
	if !(token.matchesOpenClose(p.openingTagTokenType, p.closingTagTokenType)) {
		return makeNullNode(), nil
	}

	if token.next.tokenType == p.closingTagTokenType {
		return makeNode(p.nodeType, ""), token.next.next
	}

	interiorTokenPtr := token.next

	for (*interiorTokenPtr).next.tokenType != p.closingTagTokenType {
		interiorTokenPtr = interiorTokenPtr.next
	}

	nextToken := interiorTokenPtr.next.next
	interiorTokenPtr.next = nil

	interiorNode := makeNode(p.nodeType, "")
	interiorTokenPtr = token.next

	for interiorTokenPtr != nil {
		var childNode Node
		childNode, interiorTokenPtr = p.baseParser.match(*interiorTokenPtr)
		interiorNode.children = append(interiorNode.children, &childNode)
	}

	return interiorNode, nextToken
}
