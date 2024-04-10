package writan

type Generator interface {
	getNodeType() string
	generate(node *Node) string
}

func MakeGenerator() *BaseGenerator {
	baseGenerator := makeBaseGenerator()

	textGenerator := makeTextGenerator()

	boldGenerator := makeTagMatchingGenerator(BOLD_NODE, "<strong>", "</strong>")
	italicsGenerator := makeTagMatchingGenerator(ITALIC_NODE, "<em>", "</em>")
	codeInlineGenerator := makeTagMatchingGenerator(CODE_INLINE_NODE, "<code>", "</code>")
	codeBlockGenerator := makeTagMatchingGenerator(CODE_BLOCK_NODE, "<pre><code>", "</code></pre>")
	quoteBlockGenerator := makeTagMatchingGenerator(QUOTE_BLOCK_NODE, "<blockquote>", "</blockquote>")
	paragraphGenerator := makeTagMatchingGenerator(PARAGRAPH_NODE, "<p>", "</p>")
	linkGenerator := makeLinkGenerator()

	bodyGenerator := makeBodyGenerator()
	sentenceGenerator := makeSentenceGenerator()

	atCharacterGenerator := makeSingleTokenGenerator(AT_NODE, "@")

	boldGenerator.baseGenerator = &baseGenerator
	codeInlineGenerator.baseGenerator = &baseGenerator
	italicsGenerator.baseGenerator = &baseGenerator
	paragraphGenerator.baseGenerator = &baseGenerator
	bodyGenerator.baseGenerator = &baseGenerator
	sentenceGenerator.baseGenerator = &baseGenerator
	codeBlockGenerator.baseGenerator = &baseGenerator
	quoteBlockGenerator.baseGenerator = &baseGenerator
	linkGenerator.baseGenerator = &baseGenerator

	baseGenerator.generators = []Generator{
		codeInlineGenerator,
		codeBlockGenerator,
		quoteBlockGenerator,
		linkGenerator,
		sentenceGenerator,
		bodyGenerator,
		paragraphGenerator,
		italicsGenerator,
		boldGenerator,
		atCharacterGenerator,
		textGenerator,
	}

	return &baseGenerator
}
