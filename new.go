package new

// New allocates the thing on the heap and returns a non-nil pointer to it.
func New[T any](thing T) *T {
	return &thing
}
