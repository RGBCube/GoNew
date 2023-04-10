package new

func New[T any](thing T) *T {
	return &thing
}
