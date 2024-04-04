package writan

type SingleTokenGenerator struct {
	nodeType          string
	replacementString string
}

func makeSingleTokenGenerator(nodeType string, replacementString string) SingleTokenGenerator {
	return SingleTokenGenerator{nodeType: nodeType, replacementString: replacementString}
}

func (g SingleTokenGenerator) generate(node *Node) string {
	return g.replacementString
}

func (g SingleTokenGenerator) getNodeType() string {
	return g.nodeType
}
