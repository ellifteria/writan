package writan

import "strings"

type ImageGenerator struct {
	baseGenerator *BaseGenerator
	nodeType      string
}

func makeImageGenerator() ImageGenerator {
	return ImageGenerator{nodeType: IMAGE_NODE}
}

func (g ImageGenerator) generate(node *Node) string {
	unpackedNodeValue := strings.Split(node.nodeValue, " ")
	src := unpackedNodeValue[0]
	alt := unpackedNodeValue[1]
	start := "<img src=\"" + src + "\" alt=\"" + alt + "\">"

	middle := ""

	for _, child := range node.children {
		middle += g.baseGenerator.generate(child)
	}

	end := "</img>"

	return start + middle + end
}

func (g ImageGenerator) getNodeType() string {
	return g.nodeType
}
