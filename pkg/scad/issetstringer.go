package scad

// stringer is an interface for types that can tell you if they are set and what their SCADString value is.
type isSetStringer interface {
	isSet() bool
	SCADString() string
}
