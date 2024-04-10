package writan

type LinkGenerator struct {
	baseGenerator *BaseGenerator
	nodeType      string
}

func makeLinkGenerator() LinkGenerator {
	return LinkGenerator{nodeType: LINK_NODE}
}

func (g LinkGenerator) generate(node *Node) string {
	start := "<a href=\"" + node.nodeValue + "\">"

	middle := ""

	for _, child := range node.children {
		middle += g.baseGenerator.generate(child)
	}

	end := "</a>"

	return start + middle + end
}

func (g LinkGenerator) getNodeType() string {
	return g.nodeType
}
