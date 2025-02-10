package utils

func PtrToValue[T any](ptr *T) T {
	if ptr == nil {
		var defaultValue T

		return defaultValue
	}

	return *ptr
}

func ValueToPtr[T any](value T) *T {
	return &value
}
