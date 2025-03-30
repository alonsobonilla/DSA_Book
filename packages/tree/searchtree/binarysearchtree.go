package searchtree

type BinarySearchTree struct {
	root *binaryNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{
		root: nil,
	}
}

func (t *BinarySearchTree) MakeEmpty() {
	t.root = nil
}

func (t BinarySearchTree) IsEmpty() bool {
	return t.root == nil
}

func (t BinarySearchTree) Find(element Comparable) Comparable {
	return t.elementAt(t.find(element, t.root))
}

func (t BinarySearchTree) elementAt(node *binaryNode) Comparable {
	if node == nil {
		return nil
	}
	return node.element
}

func (t BinarySearchTree) find(element Comparable, node *binaryNode) *binaryNode {
	if node == nil {
		return nil
	}

	if element.compareTo(element) < 0 {
		return t.find(element, node.Left)
	} else if element.compareTo(element) > 0 {
		return t.find(element, node.Right)
	}

	return node
}
