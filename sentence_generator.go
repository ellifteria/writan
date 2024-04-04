package writan

type SentenceGenerator struct {
	baseGenerator *BaseGenerator
}

func makeSentenceGenerator() SentenceGenerator {
	return SentenceGenerator{}
}

func (g SentenceGenerator) generate(node *Node) string {
	sentence := ""

	for _, child := range node.children {
		sentence += g.baseGenerator.generate(child)
	}

	return sentence
}

func (g SentenceGenerator) getNodeType() string {
	return SENTENCE_NODE
}
