# Go ternary

[![Go Reference](https://pkg.go.dev/badge/github.com/julien040/go-ternary.svg)](https://pkg.go.dev/github.com/julien040/go-ternary)

The missing ternary operator for Go.

It's a blazing fast, 7 lines of code package that will save you a lot of lines of code over readability.

Honestly, don't take this package too seriously. It's just a joke to keep me busy one evening.

## Installation

```bash
go get github.com/julien040/go-ternary
```

## Usage/Examples

```go
import (
    "github.com/julien040/go-ternary"
)

ternary.If(true, "foo", "bar") // "foo"
ternary.If(false, "foo", "bar") // "bar"

// Use it for pluralization
ternary.If(len(slice) != 1, "objects", "object") // "objects"

ternary.If(true, 5.4, 3.2) // 5.4

// Do not use it like this, because both funcs will be evaluated regardless of condition!
ternary.If(true, trueFunc(), falseFunc())

// Instead, use the Iff function...
ternary.Iff(true, trueFunc, falseFunc)

// or with closures if the funcs require arguments
ternary.Iff(true, func() string { return truefunc(...)}, , func() string { return falsefunc(...)})
```

## FAQ

### Which type can I use?

Thanks to generics, you can use any type you want.

### What is the minimum Go version?

Because of generics, you need at least Go 1.18.

## Contributing

Contributions are always welcome!

I accept any contribution, from code to documentation, bug reports, and feature requests.

However, I wonder if there is much to contribute to this project since it's pretty simple.

## License

[MIT](https://choosealicense.com/licenses/mit/)

## Authors

- [@julien040](https://www.github.com/julien040)
