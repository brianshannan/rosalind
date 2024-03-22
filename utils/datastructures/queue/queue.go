package queue

type Queue[T any] struct {
	elems []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(elem T) {
	q.elems = append(q.elems, elem)
}

func (q *Queue[T]) Dequeue() T {
	elem := q.elems[0]
	q.elems = q.elems[1:]
	return elem
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elems) == 0
}
