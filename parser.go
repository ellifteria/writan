package writan

type Parser interface {
	match(token Token) (Node, *Token)
}

func MakeParser() *BodyParser {
	baseParser := makeBaseParser()

	boldParser := makeTagMatchingParser(BOLD_TOKEN, BOLD_TOKEN, BOLD_NODE)
	italicsParser := makeTagMatchingParser(ITALICS_TOKEN, ITALICS_TOKEN, ITALIC_NODE)
	codeInlineParser := makeTagMatchingParser(CODE_INLINE_TOKEN, CODE_INLINE_TOKEN, CODE_INLINE_NODE)
	codeBlockParser := makeTagMatchingParser(CODE_BLOCK_TOKEN, CODE_BLOCK_TOKEN, CODE_BLOCK_NODE)
	quoteBlockParser := makeTagMatchingParser(QUOTE_BLOCK_TOKEN, QUOTE_BLOCK_TOKEN, QUOTE_BLOCK_NODE)
	linkParser := makeTagMatchingParser(LINK_TOKEN, LINK_TOKEN, LINK_NODE)

	textParser := makeTextParser()

	newlineParser := makeSingleTokenParser(NEWLINE_TOKEN, NEWLINE_NODE)
	atParser := makeSingleTokenParser(AT_TOKEN, AT_NODE)
	imageParser := makeSingleTokenParser(IMAGE_TOKEN, IMAGE_NODE)

	sentenceParser := makeSentenceParser()
	paragraphParser := makeParagraphParser(&sentenceParser)
	bodyParser := makeBodyParser(&paragraphParser, &newlineParser)

	sentenceParser.baseParser = &baseParser

	boldParser.baseParser = &baseParser
	italicsParser.baseParser = &baseParser
	textParser.baseParser = &baseParser
	codeBlockParser.baseParser = &baseParser
	codeInlineParser.baseParser = &baseParser
	quoteBlockParser.baseParser = &baseParser
	linkParser.baseParser = &baseParser
	imageParser.baseParser = &baseParser

	atParser.baseParser = &baseParser
	newlineParser.baseParser = &baseParser

	baseParser.parsers = []Parser{
		&newlineParser,
		&codeBlockParser,
		&codeInlineParser,
		&quoteBlockParser,
		&boldParser,
		&italicsParser,
		&atParser,
		&linkParser,
		&imageParser,
		&textParser,
	}

	return &bodyParser
}
