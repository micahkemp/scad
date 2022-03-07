package scad

// ChildWrapper is a function that wraps an interface within another
// interface.
type ChildWrapper func(interface{}) interface{}

// Wrap applies a set of ChildWrapper functions create a new interface
// from the input interface.
func Wrap(i interface{}, wrappers ...ChildWrapper) interface{} {
	for _, wrapper := range wrappers {
		i = wrapper(i)
	}

	return i
}
