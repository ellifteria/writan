package writan

type TextParser struct{}

func (p TextParser) match(token Token) Node {
	if !(token.isAOr([]TokenType{
		QUOTE_TOKEN,
		OPEN_BRACKET_TOKEN,
		CLOSE_BRACKET_TOKEN,
		OPEN_BRACE_TOKEN,
		CLOSE_BRACE_TOKEN,
		OPEN_PAREN_TOKEN,
		CLOSE_PAREN_TOKEN,
		OPEN_ARROW_TOKEN,
		CLOSE_ARROW_TOKEN,
		UNDERSCORE_TOKEN,
		STAR_TOKEN,
		NEWLINE_TOKEN,
		EXCLAMATION_TOKEN,
		TEXT_TOKEN,
	})) {
		return makeNullNode()
	}

	return makeNode(TEXT_NODE, token.tokenValue, 1)
}
