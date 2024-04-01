package writan

type BaseParser struct{}

func (p BaseParser) match(token Token) (Node, *Token) {
	return makeNullNode(), nil
}
