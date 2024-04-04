package writan

type Parser interface {
	match(token Token) (Node, *Token)
}

func MakeParser() (*BaseParser, *BodyParser) {
	baseParser := makeBaseParser()

	boldTextParser := makeTagMatchingTextParser(BOLD_TOKEN, BOLD_TOKEN, BOLD_TEXT_NODE)
	italicsTextParser := makeTagMatchingTextParser(ITALICS_TOKEN, ITALICS_TOKEN, ITALICIZED_TEXT_NODE)
	codeInlineParser := makeTagMatchingTextParser(CODE_INLINE_TOKEN, CODE_INLINE_TOKEN, CODE_INLINE_NODE)
	codeBlockParser := makeTagMatchingTextParser(CODE_BLOCK_TOKEN, CODE_BLOCK_TOKEN, CODE_BLOCK_NODE)
	quoteBlockParser := makeTagMatchingTextParser(QUOTE_BLOCK_TOKEN, QUOTE_BLOCK_TOKEN, QUOTE_BLOCK_NODE)
	textParser := makeTextParser()

	newlineParser := makeSingleTokenParser(NEWLINE_TOKEN, NEWLINE_NODE)
	atParser := makeSingleTokenParser(AT_TOKEN, AT_NODE)

	sentenceParser := makeSentenceParser()
	paragraphParser := makeParagraphParser(&sentenceParser)
	bodyParser := makeBodyParser(&paragraphParser, &newlineParser)

	sentenceParser.baseParser = &baseParser

	boldTextParser.baseParser = &baseParser
	italicsTextParser.baseParser = &baseParser
	textParser.baseParser = &baseParser
	codeBlockParser.baseParser = &baseParser
	codeInlineParser.baseParser = &baseParser
	quoteBlockParser.baseParser = &baseParser
	atParser.baseParser = &baseParser
	newlineParser.baseParser = &baseParser

	baseParser.parsers = []Parser{
		&newlineParser,
		&codeBlockParser,
		&codeInlineParser,
		&quoteBlockParser,
		&boldTextParser,
		&italicsTextParser,
		&atParser,
		&textParser,
	}

	return &baseParser, &bodyParser
}
