package writan

type BaseGenerator struct {
	generators []Generator
}

func makeBaseGenerator() BaseGenerator {
	return BaseGenerator{}
}

func (g BaseGenerator) generate(node *Node) string {
	generated := ""
	for _, generator := range g.generators {
		if generator.getNodeType() == node.nodeType || generator.getNodeType() == ANY_NODE_TYPE {
			generated += generator.generate(node)
			break
		}
	}
	return generated
}

func (g BaseGenerator) Generate(node *Node) string {
	return g.generate(node)
}

func (g BaseGenerator) getNodeType() string {
	return "__none__"
}
