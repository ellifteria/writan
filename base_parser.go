package writan

type BaseParser struct {
	parsers []Parser
}

func makeBaseParser() BaseParser {
	return BaseParser{}
}

func (p BaseParser) match(token Token) (Node, *Token) {
	var node Node
	var nextTokenPtr *Token
	for _, parser := range p.parsers {
		node, nextTokenPtr = parser.match(token)
		if node.nodeType != NULL_NODE {
			return node, nextTokenPtr
		}
	}

	return makeNullNode(), nil
}
