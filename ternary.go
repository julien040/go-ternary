package ternary

import (
	"errors"
	"slices"
	"strconv"
)

// If is the replacement ternary operator in Go.
//
// Usage:
//
//	ternary.If(condition, elementReturnedIfTrue, elementReturnedIfFalse)
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

// IfFunc implements a given function using the chosen element as parameter.
//
// Usage:
//
//	ternary.If(condition, elementPassedIfTrue, elementPassedIfFalse, function func(element))
//
// Example:
//
//	ternary.If(true, "foo", "bar", func(e any) {
//		return fmt.Sprintf("%v", e)
//	}) // returns "foo"
func IfFunc[T, R any](condition bool, a, b T, f func(T) R) R {
	if condition {
		return f(a)
	}
	return f(b)
}

// Third implements a third option alongside true and false.
// An error is returned if the string value does not match options.
//
// Usage:
//
//	ternary.Third(value, thirdOption, elementPassedIfTrue, elementPassedIfFalse, elementIfThird, optionalErrorMessage)
//
// Example:
//
//	ternary.Third("true", "strict", "foo", "bar", "baz", errMsg) // returns "foo"
//	ternary.Third("false", "strict", "foo", "bar", "baz", errMsg) // returns "bar"
//	ternary.Third("strict", "strict", "foo", "bar", "baz", errMsg) // returns "baz"
func Third[T any](value, third string, a, b, c T, err string) (T, error) {
	var none T
	switch value {
	case "true":
		return a, nil
	case "false":
		return b, nil
	case third:
		return c, nil
	}
	return none, errString(err, "value does not match options")
}

func ThirdExt[T any](value string, options [3][]string, a, b, c T, err string) (T, error) {
	var none T
	if options[2] == nil {
		return none, errors.New("third option is nil")
	}

	for i := range options {
		if i == 2 {
			break
		}

		boolValue := strconv.FormatBool(i == 0)
		if options[i] == nil {
			options[i] = append(options[i], boolValue)
		} else if len(options[i]) == 1 && options[i][0] == "" {
			options[i][0] = boolValue
		} else if slices.Contains(options[i], "") {
			return none, errString(err, "options has empty values")
		}
	}

	if slices.Contains(options[0], value) {
		return a, nil
	} else if slices.Contains(options[1], value) {
		return b, nil
	} else if slices.Contains(options[2], value) {
		return c, nil
	}
	return none, errString(err, "value does not match options")
}

func errString(s, d string) error {
	err := If(s == "", d, s)
	return errors.New(err)
}
