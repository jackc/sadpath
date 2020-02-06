// sadpath provides a simple way to centralize error handling and reduce boilerplate.
package sadpath

type failedCheck struct {
	err error
}

// Handle catches errors thrown by Check. It must be called with defer.
func Handle(f func(error)) {
	a := recover()
	if a != nil {
		if fc, ok := a.(failedCheck); ok {
			f(fc.err)
		} else {
			panic(a)
		}
	}
}

// Check throws err when err is not nil. It does nothing when err is nil.
func Check(err error) {
	if err != nil {
		panic(failedCheck{err: err})
	}
}
