package queue

import "errors"

const max = 10

type Queue struct {
	Channel []int
}

func (q *Queue) Enqueue(num int) error {
	if len(q.Channel) == max {
		return errors.New("OverFlow")
	}
	q.Channel = append([]int{num}, q.Channel...)
	return nil
}

func (q *Queue) Dequeue() (int, error) {
	if len(q.Channel) == 0 {
		return 0, errors.New("UnderFlow")
	}
	out := q.Channel[0]
	q.Channel = q.Channel[1:]
	return out, nil
}

func (q *Queue) IsEmpty() bool {
	if len(q.Channel) == 0 {
		return true
	}
	return false
}
