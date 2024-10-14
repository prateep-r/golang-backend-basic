package pointer_func

func ToPointer[T any](t T) *T {
	return &t
}

func ToValue[T any](t *T, defaultValue T) T {
	if t == nil {
		return defaultValue
	}
	return *t
}
