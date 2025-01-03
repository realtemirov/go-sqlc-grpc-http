package utils

func HasArray[T comparable](target T, source []T) bool {
	for _, v := range source {
		if v == target {
			return true
		}
	}

	return false
}

func ArrayIndexOrDefault[T any](source []T, index int) T {
	if index < 0 {
		index = 0
	}
	if len(source) > index {
		return source[index]
	}

	var empty T
	return empty
}
