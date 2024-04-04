package writan

type SentenceParser struct {
	baseParser *BaseParser
}

func makeSentenceParser() SentenceParser {
	return SentenceParser{}
}

func (p SentenceParser) match(token Token) (Node, *Token) {
	if token.tokenType == EOF_TOKEN || token.tokenType == NULL_TOKEN {
		return makeNullNode(), nil
	}

	currTokenPtr := &token

	sentenceNode := makeNode(SENTENCE_NODE, "")

	for currTokenPtr != nil && currTokenPtr.tokenType != EOF_TOKEN && currTokenPtr.tokenType != NEWLINE_TOKEN {
		var childNode Node
		childNode, currTokenPtr = p.baseParser.match(*currTokenPtr)
		sentenceNode.children = append(sentenceNode.children, &childNode)
	}

	return sentenceNode, currTokenPtr
}
