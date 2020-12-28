package linked_list

import "log"

type Node struct {
	value 	int
	next 	*Node
}

type LinkedList struct {
	head 	*Node
	length 	int
}

type ILinkedList interface {
	Insert (int)
	PrintNodes ()
	SearchListByValue (val int) *int // Find and return position
	SearchListByPosition (position int) *int // Find and return position
	DeleteNodeByValue (val int)
	DeleteNodeByPosition (position int)
}

type SinglyLinkedList struct {
	LinkedList
}

func LinkedListExample () {
	sk := &SinglyLinkedList{}
	for i := 10; i < 15; i++ {
		sk.Insert(i)
	}

	sk.PrintNodes()
	entryPosition := sk.SearchListByValue(12)
	if entryPosition != nil {
		log.Println(*entryPosition, "entry at position")
	}
	itemAtPosition := sk.SearchListByPosition(3)
	if itemAtPosition != nil {
		log.Println(*itemAtPosition, "item at position")
	}

	sk.DeleteNodeByValue(13)
	sk.PrintNodes()

	sk.DeleteNodeByPosition(3)
	sk.PrintNodes()
}
