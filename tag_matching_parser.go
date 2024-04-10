package writan

import (
	"strings"
)

type TagMatchingParser struct {
	baseParser          *BaseParser
	openingTagTokenType string
	closingTagTokenType string
	nodeType            string
}

func makeTagMatchingParser(opener string, closer string, nodeType string) TagMatchingParser {
	return TagMatchingParser{nil, opener, closer, nodeType}
}

func (p TagMatchingParser) match(token Token) (Node, *Token) {
	if !(token.matchesOpenClose(p.openingTagTokenType, p.closingTagTokenType)) {
		return makeNullNode(), nil
	}

	tokenValue := token.tokenValue
	value := strings.Join(strings.Split(tokenValue[1:len(tokenValue)-1], " ")[1:], " ")

	if token.next.tokenType == p.closingTagTokenType {
		return makeNode(p.nodeType, value), token.next.next
	}

	interiorTokenPtr := token.next

	for (*interiorTokenPtr).next.tokenType != p.closingTagTokenType {
		interiorTokenPtr = interiorTokenPtr.next
	}

	nextToken := interiorTokenPtr.next.next
	interiorTokenPtr.next = nil

	interiorNode := makeNode(p.nodeType, value)
	interiorTokenPtr = token.next

	for interiorTokenPtr != nil {
		var childNode Node
		childNode, interiorTokenPtr = p.baseParser.match(*interiorTokenPtr)
		interiorNode.children = append(interiorNode.children, &childNode)
	}

	return interiorNode, nextToken
}
