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

func KeysFiltered[K comparable, V any](m map[K]V, fn func(k K) bool) []K {
	keys := make([]K, 0)
	for key := range m {
		if fn(key) {
			keys = append(keys, key)
		}
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

func ValuesFiltered[K comparable, V any](m map[K]V, fn func(v V) bool) []V {
	values := make([]V, 0)
	for _, value := range m {
		if fn(value) {
			values = append(values, value)
		}
	}
	return values
}
