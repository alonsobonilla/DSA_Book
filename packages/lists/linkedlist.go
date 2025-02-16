package lists

type LinkedList struct {
	header *listNode
}

func NewLinkedList() LinkedList {
	return LinkedList{
		header: newListNode(nil, nil),
	}
}

func (l LinkedList) IsEmpty() bool {
	return l.header.next == nil
}

func (l LinkedList) MakeEmpty() {
	l.header.next = nil
}

func (l LinkedList) Zeroth() linkedListItr {
	return newLinkedListItr(l.header)
}

func (l LinkedList) First() linkedListItr {
	return newLinkedListItr(l.header.next)
}

func (l LinkedList) Find(element any) linkedListItr {

	itr := l.header

	for itr != nil && itr.element != element {
		itr = itr.next
	}

	return newLinkedListItr(itr)
}

func (l LinkedList) Remove(element any) {
	p := l.FindPrevious(element)

	if p.current.next != nil {
		p.current.next = p.current.next.next
	}
}

func (l LinkedList) FindPrevious(element any) linkedListItr {
	itr := l.header

	for itr.next != nil && itr.next.element != element {
		itr = itr.next
	}
	return newLinkedListItr(itr)
}

// Insert after p;
// element the item to insert;
// p the position prior to the newly inserted item.
func (l LinkedList) Insert(element any, p linkedListItr) {
	if p.current != nil {
		p.current.next = newListNode(element, p.current.next)
	}
}
