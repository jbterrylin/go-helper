package arrayhelper

// Entry represents a key-value pair.
type Entry[T any] struct {
	Index int
	Value T
}

// Iterator represents a slice iterator.
type Iterator[T any] struct {
	slice []T
	index int
}

// NewIterator creates a new iterator for a slice.
func NewIterator[T any](slice []T) *Iterator[T] {
	return &Iterator[T]{slice: slice, index: -1}
}

// Next advances the iterator and returns the next entry.
func (it *Iterator[T]) Next() (Entry[T], bool) {
	if it.index+1 < len(it.slice) {
		it.index++
		return Entry[T]{Index: it.index, Value: it.slice[it.index]}, true
	}
	return Entry[T]{}, false
}
