package ds
import "errors"

type Queue struct {
    list LinkedList
}

func (q *Queue) Push(value string) { // add tail node
	q.list.Insert(value)
}
func (q *Queue) Pop() (string, error) { // remove the head
	if q.list.IsEmpty() {
		return "", errors.New("list is empty")
	}
	queueHead := q.list.Head.data
	q.list.RemoveAt(0)
	return queueHead, nil
}
