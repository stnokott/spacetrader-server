package mocks

func ptr[T any](x T) *T {
	return &x
}
