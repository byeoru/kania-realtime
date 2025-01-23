package util

func Map[T any, R any](input []T, mapper func(T) R) []R {
	result := make([]R, len(input))
	for i, v := range input {
		result[i] = mapper(v)
	}
	return result
}
