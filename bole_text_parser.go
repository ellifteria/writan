package writan

type BoldTextParser struct{}

func (p BoldTextParser) match(token Token) Node {
	if !(token.matchesTypes([]TokenType{
		STAR_TOKEN,
		STAR_TOKEN,
		TEXT_TOKEN,
		STAR_TOKEN,
		STAR_TOKEN,
	})) {
		return makeNullNode()
	}

	return makeNode(TEXT_NODE, token.tokenValue, 1)
}
