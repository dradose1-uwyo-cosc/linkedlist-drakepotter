package ds
import "fmt"
import "errors"

type LinkedList struct {
    Head *Node
    Tail *Node
    Size int
}

func (l *LinkedList) Insert(value string) { // insert at the tail
	newNode := &Node{value, nil}
	if l.IsEmpty() {
		l.Head = newNode
		l.Tail = newNode
		l.Size++
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
		l.Size++
	}
}

func (l *LinkedList) InsertAt(position int, value string) error { //inserts at a position, returns true if success or false if not, like if pos doesn't exist
	if position < 0 || position > l.Size {
		return errors.New("index out of bounds")
	} else if position == 0 {
		newNode := &Node{value, l.Head} //make a new node pointing to current head
		l.Head = newNode //new node now is new head
		if l.IsEmpty() {
			l.Tail = newNode //if this is first node, it's also the tail
		}
		l.Size++ //increase the size
		return nil
	} else if position == l.Size {
		l.Insert(value)
		return nil
	} else {
		current := l.Head
		for i := 0; i < position; i++ {
			if i == position - 1 {
				newNode := &Node{value, current.Next}
				current.Next = newNode
				l.Size++
			}
			current = current.Next
		}
		return nil
	}
}

func (l *LinkedList) Remove(value string) error { // remove first occurrence of the value
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	if l.Head.data == value {
		l.Head = l.Head.Next
		l.Size--
		if l.IsEmpty() {
			l.Tail = nil
		}
		return nil
	}
	current := l.Head
	for i := 0; i < l.Size - 1; i++ {
		if current.Next.data == value {
			if current.Next == l.Tail {
				l.Tail = current
			}
			current.Next = current.Next.Next
			l.Size--
			return nil
		}
		current = current.Next
	}
	return errors.New("value not found")
}

func (l *LinkedList) RemoveAll(value string) error { //remove all occurrences of a value
	if l.IsEmpty() {
		return errors.New("list is empty")
	}
	for l.Head != nil && l.Head.data == value { //keep removing the head while it matches the value
		l.Head = l.Head.Next
		l.Size--
		if l.Size == 0 {
			l.Tail = nil
			return nil
		}
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.data == value {
			if current.Next == l.Tail {
				l.Tail = current
			}
			current.Next = current.Next.Next
			l.Size--
		} else {
			current = current.Next
		}
	}
	return nil
}

func (l *LinkedList) RemoveAt(pos int) error { // remove at a position, if index exists
	if pos < 0 || pos >= l.Size {
		return errors.New("invalid position")
	}
	if pos == 0 {
		l.Head = l.Head.Next
		l.Size--
		if l.IsEmpty() {
			l.Tail = nil
		}
		return nil
	} else {
		current := l.Head
		for i:= 0; i < pos - 1; i++ {
			current = current.Next
		}
		if current.Next == l.Tail {
			l.Tail = current
		}
		current.Next = current.Next.Next
		l.Size--
		return nil
	}
}

func (l *LinkedList) IsEmpty() bool { // checks if the linked list is empty
	if l.Size == 0 {
		return true
	}
	return false
}

func (l *LinkedList) GetSize() int { // get size of ll
	return l.Size
}

func (l *LinkedList) Reverse() { //reverse the list
	if !l.IsEmpty() && l.Head.Next != nil {
		current := l.Head
		var previous *Node = nil
		l.Tail = l.Head //set the head as the new tail
		for current != nil {
			next := current.Next //get next node
			current.Next = previous //flip the pointer
			previous = current //move forward in the list
			current = next
		}
		l.Head = previous //set the tail as the new head
	}

}
func (l *LinkedList) PrintList() { //print the list 
	if l.IsEmpty() {
		fmt.Println("List is empty")
	} else {
		current := l.Head
		for current != nil {
			if current.Next == nil {
				fmt.Print(current.data)
			} else {
				fmt.Print(current.data + " -> ")
			}
			current = current.Next
		}
	}
	fmt.Println() //add a new line below the list
}
