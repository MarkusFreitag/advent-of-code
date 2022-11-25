package util

func Keys[T any](m map[T]any) []T {
	keys := make([]T, len(m))
	var idx int
	for key := range m {
		keys[idx] = key
		idx++
	}
	return keys
}

func Values[T any](m map[any]T) []T {
	values := make([]T, len(m))
	var idx int
	for _, value := range m {
		keys[idx] = value
		idx++
	}
	return values
}
