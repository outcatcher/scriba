package repo

type convertable[T any] interface {
	toEntity() *T
}

// convertEntitySlice convert slice of repo elements to entity elements.
// See tests for example usage.
func convertEntitySlice[F convertable[T], T any](from []F) []T {
	result := make([]T, len(from))

	for i, val := range from {
		result[i] = *val.toEntity()
	}

	return result
}
