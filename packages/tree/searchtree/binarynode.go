package searchtree

type binaryNode struct {
	element Comparable
	Left    *binaryNode
	Right   *binaryNode
}

func NewBinaryNode(element Comparable, left *binaryNode, right *binaryNode) *binaryNode {
	return &binaryNode{
		element: element,
		Left:    left,
		Right:   right,
	}
}

type Comparable interface {
	compareTo(element any) int
}
