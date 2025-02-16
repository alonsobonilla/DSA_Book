package lists

type linkedListItr struct {
	current *listNode
}

func newLinkedListItr(node *listNode) linkedListItr {
	return linkedListItr{
		current: node,
	}
}

func (itr linkedListItr) IsPastEnd() bool {
	return itr.current == nil
}

func (itr linkedListItr) Retrieve() any {
	if itr.IsPastEnd() {
		return nil
	}
	return itr.current.element
}

func (itr linkedListItr) Advance() {
	if !itr.IsPastEnd() {
		itr.current = itr.current.next
	}
}
