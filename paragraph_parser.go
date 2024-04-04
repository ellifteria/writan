package writan

type ParagraphParser struct {
	sentenceParser *SentenceParser
}

func makeParagraphParser(sentenceParser *SentenceParser) ParagraphParser {
	return ParagraphParser{sentenceParser: sentenceParser}
}

func (p ParagraphParser) match(token Token) (Node, *Token) {
	if token.tokenType == EOF_TOKEN || token.tokenType == NULL_TOKEN {
		return makeNullNode(), nil
	}

	currTokenPtr := &token

	paragraphNode := makeNode(PARAGRAPH_NODE, "")

	for currTokenPtr != nil && currTokenPtr.tokenType != EOF_TOKEN {
		if currTokenPtr.tokenType == NEWLINE_TOKEN {
			if currTokenPtr.next != nil && currTokenPtr.next.tokenType == NEWLINE_TOKEN {
				return paragraphNode, currTokenPtr.next.next
			}
			currTokenPtr = currTokenPtr.next
		} else {
			var childNode Node
			childNode, currTokenPtr = p.sentenceParser.match(*currTokenPtr)
			paragraphNode.children = append(paragraphNode.children, &childNode)
		}
	}

	return paragraphNode, currTokenPtr
}
