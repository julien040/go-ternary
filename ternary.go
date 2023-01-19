package ternary

// Replace the missing ternary operator in Go
//
// Usage:
//
//	ternary.If(condition bool, elementReturnedIfTrue, elementReturnedIfFalse)
//
// Example:
//
//	ternary.If(true, "foo", "bar") // returns "foo"
//	ternary.If(false, "foo", "bar") // returns "bar"
func If[T any](condition bool, a, b T) T {
	if condition {
		return a
	}
	return b
}
