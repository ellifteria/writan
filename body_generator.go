package writan

type BodyGenerator struct {
	baseGenerator *BaseGenerator
}

func makeBodyGenerator() BodyGenerator {
	return BodyGenerator{}
}

func (g BodyGenerator) generate(node *Node) string {
	body := ""

	for _, child := range node.children {
		body += g.baseGenerator.generate(child)
	}

	return body
}

func (g BodyGenerator) getNodeType() string {
	return BODY_NODE
}
