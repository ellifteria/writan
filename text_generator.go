package writan

type TextGenerator struct{}

func makeTextGenerator() TextGenerator {
	return TextGenerator{}
}

func (g TextGenerator) generate(node *Node) string {
	return node.nodeValue
}

func (g TextGenerator) getNodeType() string {
	return ANY_NODE_TYPE
}
