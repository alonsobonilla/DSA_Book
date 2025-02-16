package lists

type listNode struct {
	element any
	next    *listNode
}

func newListNode(x any, next *listNode) *listNode {
	return &listNode{
		element: x,
		next:    next,
	}
}
