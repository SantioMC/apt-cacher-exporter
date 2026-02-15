package aptcache

func count[T comparable](slice []T, predicate func(T) bool) int {
	count := 0

	for _, item := range slice {
		if predicate(item) {
			count++
		}
	}

	return count
}

func reduce[T comparable](initial int, slice []T, mapper func(T) int) int {
	result := initial

	for _, item := range slice {
		result += mapper(item)
	}

	return result
}
