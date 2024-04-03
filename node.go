package writan

import "fmt"

type NodeType int

const (
	NULL_NODE NodeType = iota
	BODY_NODE
	PARAGRAPH_NODE
	SENTENCE_NEWLINE_NODE
	SENTENCE_EOF_NODE
	BLOCK_CODE_NODE
	INLINE_CODE_NODE
	SENTENCE_NODE
	ITALICIZED_TEXT_NODE
	BOLD_TEXT_NODE
	TEXT_NODE
)

func (n NodeType) String() string {
	return [...]string{
		"NULL",
		"BODY",
		"PARAGRAPH",
		"SENTENCE_NEWLINE",
		"SENTENCE_EOF",
		"BLOCK_CODE",
		"INLINE_CODE",
		"SENTENCE",
		"ITALICIZED_TEXT",
		"BOLD_TEXT",
		"TEXT",
	}[n]
}

type Node struct {
	nodeType  NodeType
	nodeValue string
	children  []*Node
}

func makeNode(nodeType NodeType, nodeValue string) Node {
	if nodeType == NULL_NODE {
		makeNullNode()
	}

	return Node{nodeType: nodeType, nodeValue: nodeValue, children: make([]*Node, 0)}
}

func makeNullNode() Node {
	return Node{nodeType: NULL_NODE, nodeValue: "", children: make([]*Node, 0)}
}

func (n *Node) toString() string {
	return fmt.Sprintf("<type: %s, value: %s>", n.nodeType.String(), n.nodeValue)
}
