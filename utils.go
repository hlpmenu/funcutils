package funcutils

func RunFuncWithArgs(fn func(interface{}, ...interface{}), i interface{}, args ...interface{}) {
	fn(i, args)
}

func RunFunc(fn func()) {
	fn()
}
func RunFuncError(fn func() error) error {
	return fn()
}
