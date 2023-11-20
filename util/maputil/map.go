package maputil

func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, len(m))
	var idx int
	for key := range m {
		keys[idx] = key
		idx++
	}
	return keys
}

func Values[K comparable, V any](m map[K]V) []V {
	values := make([]V, len(m))
	var idx int
	for _, value := range m {
		values[idx] = value
		idx++
	}
	return values
}
