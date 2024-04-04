package writan

type TagMatchingGenerator struct {
	baseGenerator *BaseGenerator
	nodeType      string
	openingTag    string
	closingTag    string
}

func makeTagMatchingGenerator(nodeType string, openingTag string, closingTag string) TagMatchingGenerator {
	return TagMatchingGenerator{nodeType: nodeType, openingTag: openingTag, closingTag: closingTag}
}

func (g TagMatchingGenerator) generate(node *Node) string {
	start := g.openingTag

	middle := ""

	for _, child := range node.children {
		middle += g.baseGenerator.generate(child)
	}

	end := g.closingTag

	return start + middle + end
}

func (g TagMatchingGenerator) getNodeType() string {
	return g.nodeType
}
