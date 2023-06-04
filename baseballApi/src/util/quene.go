package util

type queue[T any] struct {
	data []T
}

// Queue ...
type Queue[T any] interface {
	IsEmpty() bool
	Push(T)
	Pop() (T, bool)
}

// NewQueue ...
func NewQueue[T any]() Queue[T] {
	return &queue[T]{}
}

// Push ...
func (q *queue[T]) Push(n T) {
	q.data = append(q.data, n)
}

// Pop ...
func (q *queue[T]) Pop() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	dequeueData := q.data[0]
	q.data = q.data[1:]
	return dequeueData, true
}

// IsEmpty ...
func (q *queue[T]) IsEmpty() bool {
	return len(q.data) == 0
}
