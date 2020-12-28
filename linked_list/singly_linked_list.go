package linked_list

import "log"

func (s *SinglyLinkedList) Insert (val int) {
	// Initialize a new Node
	node := &Node{}
	node.value = val

	// For first node, next will still be null, thankfully we have the length value to help
	if s.length < 1 {
		s.head = node // Since its the first, this will be the head node now
		s.length++
		return // No further action required
	}

	ptr := s.head
	// For a list with existing nodes
	for i := 0; i < s.length; i++ {
		// If next is nil, then this is the last node
		if ptr.next == nil {
			ptr.next = node
			s.length++
			return
		}
		ptr = ptr.next
	}
}

// Search function will find the item and return its position
func (s *SinglyLinkedList) SearchListByValue (val int) *int {
	var position int
	if s.length == 0 {
		return nil
	}
	ptr := s.head

	for ptr.next != nil {
		if ptr.value == val {
			break
		}
		ptr = ptr.next
		position++
	}
	return &position
}

// Search function will find the item and return its position
func (s *SinglyLinkedList) SearchListByPosition (position int) *int {
	if s.length == 0 {
		return nil
	}
	ptr := s.head

	for i := 0; i < s.length; i++ {
		if i == position {
			break
		}
		ptr = ptr.next
	}
	return &ptr.value
}

func (s *SinglyLinkedList) PrintNodes () {
	if s.length == 0 {
		log.Println("No nodes in list")
	}
	ptr := s.head
	for ptr.next != nil {
		log.Println(ptr)
		ptr = ptr.next
	}
	log.Println(ptr)
}

func (s *SinglyLinkedList) DeleteNodeByValue (val int) {
	prev := &Node{}
	if s.length == 0 {
		log.Println("No nodes in list")
	}
	ptr := s.head
	for ptr.next != nil {
		if ptr.value == val {
			prev.next = ptr.next
			s.length--
			break
		}
		prev = ptr
		ptr = ptr.next
	}
}

func (s *SinglyLinkedList) DeleteNodeByPosition (position int) {
	prev := &Node{}
	if s.length == 0 {
		log.Println("No nodes in list")
	}

	if position - 1 >= s.length {
		log.Println("Provided position does not exist")
	}
	ptr := s.head

	for i := 0; i < s.length; i++ {
		if i == position - 1 {
			prev.next = ptr.next
			break
		}
		prev = ptr
		ptr = ptr.next
	}
}

