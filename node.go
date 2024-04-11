package writan

import "fmt"

type NodeType int

const (
	NULL_NODE        = "NULL_NODE"
	BODY_NODE        = "BODY_NODE"
	PARAGRAPH_NODE   = "PARAGRAPH_NODE"
	CODE_BLOCK_NODE  = "CODE_BLOCK_NODE"
	CODE_INLINE_NODE = "CODE_INLINE_NODE"
	QUOTE_BLOCK_NODE = "QUOTE_BLOCK_NODE"
	SENTENCE_NODE    = "SENTENCE_NODE"
	ITALIC_NODE      = "ITALIC_NODE"
	BOLD_NODE        = "BOLD_NODE"
	TEXT_NODE        = "TEXT_NODE"
	NEWLINE_NODE     = "NEWLINE_NODE"
	AT_NODE          = "AT_NODE"
	LINK_NODE        = "LINK_NODE"
	IMAGE_NODE       = "IMAGE_NODE"
	ANY_NODE_TYPE    = "__any__"
)

var VALID_NODES_TYPES = []string{
	NULL_NODE,
	BODY_NODE,
	PARAGRAPH_NODE,
	CODE_BLOCK_NODE,
	CODE_INLINE_NODE,
	QUOTE_BLOCK_NODE,
	SENTENCE_NODE,
	ITALIC_NODE,
	BOLD_NODE,
	TEXT_NODE,
	NEWLINE_NODE,
	LINK_NODE,
	AT_NODE,
	IMAGE_NODE,
}

type Node struct {
	nodeType  string
	nodeValue string
	children  []*Node
}

func makeNode(nodeType string, nodeValue string) Node {
	if nodeType == NULL_NODE {
		makeNullNode()
	}

	return Node{nodeType: nodeType, nodeValue: nodeValue, children: make([]*Node, 0)}
}

func makeNullNode() Node {
	return Node{nodeType: NULL_NODE, nodeValue: "", children: make([]*Node, 0)}
}

func (n *Node) ToString() string {
	start := fmt.Sprintf("<type: %s, value: \"%s\" children: [", n.nodeType, n.nodeValue)
	middle := ""
	for _, child := range n.children {
		middle += child.ToString()
	}
	end := "]>"

	return start + middle + end
}
